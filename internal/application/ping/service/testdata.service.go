package services

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	commonv1 "github.com/ikaiguang/go-snowflake-node-id/api/common/v1"
	srvs "github.com/ikaiguang/go-snowflake-node-id/internal/domain/ping/service"
	testdatav1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1/resources"
	testdataservicev1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1/services"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
)

// testdata .
type testdata struct {
	testdataservicev1.UnimplementedSrvTestdataServer

	log    *log.Helper
	wssSrv *srvs.WebsocketSrv
}

// NewTestdataService .
func NewTestdataService(logger log.Logger, wssSrv *srvs.WebsocketSrv) testdataservicev1.SrvTestdataServer {
	return &testdata{
		log:    log.NewHelper(log.With(logger, "module", "admin/application/service/testdata")),
		wssSrv: wssSrv,
	}
}

// Websocket websocket
func (s *testdata) Websocket(ctx context.Context, in *testdatav1.TestReq) (resp *testdatav1.TestResp, err error) {
	err = errorutil.NotImplemented(commonv1.ERROR_NOT_IMPLEMENTED.String(), "unimplemented")
	return &testdatav1.TestResp{}, err
}

// wss ws
func (s *testdata) wss(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
	return s.wssSrv.Wss(ctx, w, r)
}
