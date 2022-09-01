package nodeid

import (
	"gorm.io/gorm"
	"time"
)

// options 配置可选项
type options struct {
	dbConn *gorm.DB

	// maxNodeID 节点ID；默认 _maxNodeId = 1023
	maxNodeID int64
	// idleDuration 闲置的时间
	idleDuration time.Duration
}

// Option is config option.
type Option func(*options)

// WithMaxNodeID ...
func WithMaxNodeID(maxNodeID int64) Option {
	return func(o *options) {
		o.maxNodeID = maxNodeID
	}
}

// WithIdleDuration ...
func WithIdleDuration(idleDuration time.Duration) Option {
	return func(o *options) {
		o.idleDuration = idleDuration
	}
}

// WithDBConn ...
func WithDBConn(dbConn *gorm.DB) Option {
	return func(o *options) {
		o.dbConn = dbConn
	}
}
