package srvs

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	commonv1 "github.com/ikaiguang/go-snowflake-node-id/api/common/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	websocketutil "github.com/ikaiguang/go-srv-kit/kratos/websocket"
	stdhttp "net/http"
)

// WebsocketSrv ...
type WebsocketSrv struct {
	log *log.Helper
}

// NewWebsocketSrv ...
func NewWebsocketSrv(logger log.Logger) *WebsocketSrv {
	return &WebsocketSrv{
		log: log.NewHelper(log.With(logger, "module", "admin/domain/service/websocket")),
	}
}

// WsMessage ws
type WsMessage struct {
	Type    int
	Content []byte
}

// Wss ws
func (s *WebsocketSrv) Wss(ctx context.Context, w stdhttp.ResponseWriter, r *stdhttp.Request) (err error) {
	// 升级连接
	cc, err := websocketutil.UpgradeConn(w, r, w.Header())
	if err != nil {
		err = errorutil.InternalServer(commonv1.ERROR_CONNECTION.String(), "upgrade conn failed", err)
		s.log.WithContext(ctx).Error(err)
		return
	}
	defer func() { _ = cc.Close() }()

	// 处理消息
	err = s.ProcessWssMsg(ctx, cc)
	if err != nil {
		return err
	}
	return err
}

func (s *WebsocketSrv) ProcessWssMsg(ctx context.Context, cc *websocket.Conn) error {
	// 读消息
	for {
		messageType, messageBytes, wsErr := cc.ReadMessage()
		if wsErr != nil {
			if websocketutil.IsCloseError(wsErr) {
				s.log.WithContext(ctx).Infow("websocket close", wsErr.Error())
				break
			}
			err := errorutil.InternalServer(commonv1.ERROR_CONNECTION.String(), "ws读取信息失败", wsErr)
			s.log.WithContext(ctx).Error(err)
			return err
		}

		// 消息
		wsMessage := &WsMessage{
			Type:    messageType,
			Content: messageBytes,
		}
		//messageChan <- wsMessage

		// 处理
		needCloseConn, err := s.processWsMessage(ctx, wsMessage)
		if err != nil {
			err = errorutil.InternalServer(commonv1.ERROR_INTERNAL_SERVER.String(), "ws处理信息失败", err)
			s.log.WithContext(ctx).Error(err)
			return err
		}

		// 响应
		err = cc.WriteMessage(messageType, messageBytes)
		if err != nil {
			if websocketutil.IsCloseError(wsErr) {
				s.log.WithContext(ctx).Infow("websocket close", wsErr.Error())
				break
			}
			err = errorutil.InternalServer(commonv1.ERROR_INTERNAL_SERVER.String(), "JSON编码错误", err)
			s.log.WithContext(ctx).Error(err)
			return err
		}

		// 关闭
		if needCloseConn {
			return err
		}
	}
	return nil
}

// processWsMessage 处理ws-message
func (s *WebsocketSrv) processWsMessage(ctx context.Context, wsMessage *WsMessage) (needCloseConn bool, err error) {
	s.log.WithContext(ctx).Infow("ws-message type", wsMessage.Type)
	s.log.WithContext(ctx).Infow("ws-message message", string(wsMessage.Content))
	if string(wsMessage.Content) == "close" {
		needCloseConn = true
	}
	return needCloseConn, err
}
