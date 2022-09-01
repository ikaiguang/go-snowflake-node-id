package nodeid

import (
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
	mysqlutil "github.com/ikaiguang/go-srv-kit/data/mysql"
	"gorm.io/gorm"
)

// Interface ...
type Interface interface {
}

// NewMysqlDB ...
func NewMysqlDB(conf *confv1.Data_MySQL, opts ...gormutil.Option) (db *gorm.DB, err error) {
	return mysqlutil.NewMysqlDB(conf, opts...)
}
