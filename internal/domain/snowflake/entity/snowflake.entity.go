// Package entities
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package entities

import (
	"strconv"
	"time"
)

var _ = time.Time{}

// SnowflakeWorker ENGINE InnoDB CHARSET utf8mb4 COMMENT '雪花算法节点ID'
type SnowflakeWorker struct {
	Id                   uint64    `gorm:"column:id;primaryKey" json:"id"`                              // id
	NodeUuid             string    `gorm:"column:node_uuid" json:"node_uuid"`                           // 唯一ID；instance_id+node_id
	InstanceLaunchTime   time.Time `gorm:"column:instance_launch_time" json:"instance_launch_time"`     // 实例启动时间
	InstanceExtendTime   time.Time `gorm:"column:instance_extend_time" json:"instance_extend_time"`     // 实例续期时间
	InstanceId           string    `gorm:"column:instance_id" json:"instance_id"`                       // 实例ID
	SnowflakeNodeId      int64     `gorm:"column:snowflake_node_id" json:"snowflake_node_id"`           // 雪花算法节点id
	InstanceName         string    `gorm:"column:instance_name" json:"instance_name"`                   // 实例名称
	InstanceEndpointList string    `gorm:"column:instance_endpoint_list" json:"instance_endpoint_list"` // 实例端点数组
	InstanceMetadata     string    `gorm:"column:instance_metadata" json:"instance_metadata"`           // 实例元数据
	CreatedTime          time.Time `gorm:"column:created_time" json:"created_time"`                     // 创建时间
}

// TableName table name
func (s *SnowflakeWorker) TableName() string {
	return "snowflake_worker_node"
}

// GenNodeUUID uuid
func (s *SnowflakeWorker) GenNodeUUID() string {
	return s.InstanceId + ":" + strconv.FormatInt(s.SnowflakeNodeId, 10)
}

// InstanceMaxNodeID ...
type InstanceMaxNodeID struct {
	Id                 uint64    `gorm:"column:id;primaryKey" json:"id"`                          // id
	InstanceId         string    `gorm:"column:instance_id" json:"instance_id"`                   // 实例ID
	SnowflakeNodeId    int64     `gorm:"column:snowflake_node_id" json:"snowflake_node_id"`       // 雪花算法节点id
	InstanceExtendTime time.Time `gorm:"column:instance_extend_time" json:"instance_extend_time"` // 实例续期时间
}

// InstanceIdleNodeIDReq 获取闲置的NodeID
type InstanceIdleNodeIDReq struct {
	InstanceId            string
	MaxInstanceExtendTime time.Time
}

// InstanceMissingNodeIDReq 获取缺失的NodeID
type InstanceMissingNodeIDReq struct {
	InstanceId string
}
