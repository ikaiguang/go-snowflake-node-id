package nodeid

import (
	"context"
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	apiv1 "github.com/ikaiguang/go-snowflake-node-id/api"
)

// ExampleNewWorker ...
func ExampleNewWorker() {
	var workerHandler WorkerRepo

	conf := &confv1.Data_MySQL{
		Dsn:            "root:Mysql.123456@tcp(127.0.0.1:3306)/srv_snowflake?charset=utf8mb4&timeout=30s&parseTime=True&loc=Local",
		LoggerEnable:   true,
		LoggerColorful: true,
		LoggerLevel:    "DEBUG",
	}
	dbConn, err := NewMysqlDB(conf)
	if err != nil {
		panic(err)
	}
	workerHandler, err = NewWorker(
		WithDBConn(dbConn),
		WithMaxNodeID(5),
		WithIdleDuration(16*time.Second),
	)
	if err != nil {
		panic(err)
	}
	_, _ = workerHandler.GetNodeId(context.Background(), nil)
	_, _ = workerHandler.ExtendNodeId(context.Background(), nil)
}

// go test -v -count=1 ./node-id/ -run=Test_worker_GetNodeId
func Test_worker_GetNodeId(t *testing.T) {
	conf := &confv1.Data_MySQL{
		Dsn:            "root:Mysql.123456@tcp(127.0.0.1:3306)/srv_snowflake?charset=utf8mb4&timeout=30s&parseTime=True&loc=Local",
		LoggerEnable:   true,
		LoggerColorful: true,
		LoggerLevel:    "DEBUG",
	}
	dbConn, err := NewMysqlDB(conf)
	require.NoError(t, err)
	workerHandler, err := NewWorker(
		WithIdleDuration(_idleDuration),
		WithDBConn(dbConn),
		WithMaxNodeID(5),
	)
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		req *apiv1.GetNodeIdReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp *apiv1.SnowflakeWorkerNode
		wantErr  bool
	}{
		{
			name: "###获取nodeID",
			args: args{
				ctx: context.Background(),
				req: &apiv1.GetNodeIdReq{
					InstanceId:   "test-instance",
					InstanceName: "test-instance",
					Endpoints: []string{
						"http://127.0.0.1:8000?isSecure=false",
						"grpc://127.0.0.1:9000?isSecure=false",
					},
					Metadata: map[string]string{
						"test": "instance",
					},
				},
			},
			wantResp: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := workerHandler.GetNodeId(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNodeId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp != nil {
				t.Log("==> node id :", gotResp.NodeId)
			} else {
				t.Log("==> gotResp = nil")
			}
			//if !reflect.DeepEqual(gotResp, tt.wantResp) {
			//	t.Errorf("GetNodeId() gotResp = %v, want %v", gotResp, tt.wantResp)
			//}
		})
	}
}

// go test -v -count=1 ./node-id/ -run=Test_worker_ExtendNodeId
func Test_worker_ExtendNodeId(t *testing.T) {
	conf := &confv1.Data_MySQL{
		Dsn:            "root:Mysql.123456@tcp(127.0.0.1:3306)/srv_snowflake?charset=utf8mb4&timeout=30s&parseTime=True&loc=Local",
		LoggerEnable:   true,
		LoggerColorful: true,
		LoggerLevel:    "DEBUG",
	}
	dbConn, err := NewMysqlDB(conf)
	require.NoError(t, err)
	workerHandler, err := NewWorker(
		WithIdleDuration(_idleDuration),
		WithDBConn(dbConn),
		WithMaxNodeID(5),
	)
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		req *apiv1.ExtendNodeIdReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp *apiv1.SnowflakeWorkerNode
		wantErr  bool
	}{
		{
			name: "###续期nodeID",
			args: args{
				ctx: context.Background(),
				req: &apiv1.ExtendNodeIdReq{
					Id:         2,
					InstanceId: "test-instance",
					NodeId:     2,
				},
			},
			wantResp: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := workerHandler.ExtendNodeId(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNodeId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp != nil {
				t.Log("==> gotResp.Success :", gotResp.Success)
			} else {
				t.Log("==> gotResp = nil")
			}
			//if !reflect.DeepEqual(gotResp, tt.wantResp) {
			//	t.Errorf("GetNodeId() gotResp = %v, want %v", gotResp, tt.wantResp)
			//}
		})
	}
}
