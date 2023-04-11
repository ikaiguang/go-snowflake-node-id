package services

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	commonv1 "github.com/ikaiguang/go-snowflake-node-id/api/common/v1"
	srvs "github.com/ikaiguang/go-snowflake-node-id/internal/domain/ping/service"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	stdhttp "net/http"
)

// WebsocketService ...
type WebsocketService struct {
	wsSrv *srvs.WebsocketSrv
}

// NewWebsocketService ...
func NewWebsocketService(wsSrv *srvs.WebsocketSrv) *WebsocketService {
	return &WebsocketService{
		wsSrv: wsSrv,
	}
}

// TestWebsocket ...
func (s *WebsocketService) TestWebsocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != stdhttp.MethodGet {
		err := errorutil.InternalServer(commonv1.ERROR_METHOD_NOT_ALLOWED.String(), "ERROR_METHOD_NOT_ALLOWED")
		w.WriteHeader(stdhttp.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err := s.wsSrv.Wss(r.Context(), w, r)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	return
}
