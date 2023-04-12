package snowflakeutil

import (
	"gorm.io/gorm"
	"time"
)

// options 配置可选项
type options struct {
	dbConn *gorm.DB

	// maxNodeID 节点ID；默认 DefaultMaxNodeId = 1023
	maxNodeID int64
	// idleDuration 空闲ID时间：超过此时间不续期，节点ID变为空闲的ID；
	// 默认 DefaultIdleDuration = 16s
	idleDuration time.Duration
}

// Option is config option.
type Option func(*options)

// WithDBConn ...
func WithDBConn(dbConn *gorm.DB) Option {
	return func(o *options) {
		o.dbConn = dbConn
	}
}

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
