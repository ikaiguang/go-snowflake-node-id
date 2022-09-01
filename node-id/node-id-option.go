package nodeid

import (
	"gorm.io/gorm"
)

// options 配置可选项
type options struct {
	maxNodeID int64
	dbConn    *gorm.DB
}

// Option is config option.
type Option func(*options)

// WithMaxNodeID ...
func WithMaxNodeID(maxNodeID int64) Option {
	return func(o *options) {
		o.maxNodeID = maxNodeID
	}
}

// WithDBConn ...
func WithDBConn(dbConn *gorm.DB) Option {
	return func(o *options) {
		o.dbConn = dbConn
	}
}
