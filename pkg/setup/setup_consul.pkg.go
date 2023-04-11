package setuppkg

import (
	"github.com/hashicorp/consul/api"
	pkgerrors "github.com/pkg/errors"
	stdlog "log"
	"sync"

	consulutil "github.com/ikaiguang/go-srv-kit/data/consul"
)

// GetConsulClient consul 客户端
func (s *engines) GetConsulClient() (*api.Client, error) {
	var err error
	s.consulClientMutex.Do(func() {
		s.consulClient, err = s.loadingConsulClient()
	})
	if err != nil {
		s.consulClientMutex = sync.Once{}
	}
	return s.consulClient, err
}

// SetConsulClient consul 客户端
//func (s *engines) SetConsulClient(cc *api.Client) {
//	if cc == nil {
//		return
//	}
//	var hasSet bool
//	s.consulClientMutex.Do(func() {
//		hasSet = true
//		s.consulClient = cc
//	})
//	if !hasSet {
//		s.consulClient = cc
//	}
//}

// loadingConsulClient consul 客户端
func (s *engines) loadingConsulClient() (*api.Client, error) {
	if s.Config.ConsulConfig() == nil {
		stdlog.Println("|*** 加载：Consul客户端：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载：Consul客户端：...")

	return consulutil.NewConsulClient(s.Config.ConsulConfig())
}
