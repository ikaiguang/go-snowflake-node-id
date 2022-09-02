// Package assemblers
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package assemblers

import (
	timeutil "github.com/ikaiguang/go-srv-kit/kit/time"

	apiv1 "github.com/ikaiguang/go-snowflake-node-id/api"
	entities "github.com/ikaiguang/go-snowflake-node-id/node-id/entity"
)

var _ = timeutil.YmdHms

// AssembleListSnowflakeWorkerNode assemble listSnowflakeWorkerNode
func AssembleListSnowflakeWorkerNode(dataModels []*entities.SnowflakeWorkerNode) []*apiv1.SnowflakeWorkerNode {
	var newDataModels = make([]*apiv1.SnowflakeWorkerNode, len(dataModels))
	for index := range dataModels {
		newDataModels[index] = AssembleSnowflakeWorkerNode(dataModels[index])
	}
	return newDataModels
}

// AssembleSnowflakeWorkerNode assemble SnowflakeWorkerNode
func AssembleSnowflakeWorkerNode(dataModel *entities.SnowflakeWorkerNode) *apiv1.SnowflakeWorkerNode {
	newDataModel := &apiv1.SnowflakeWorkerNode{
		Id:                   dataModel.Id,                                         // id
		InstanceLaunchTime:   dataModel.InstanceLaunchTime.Format(timeutil.YmdHms), // 实例启动时间
		InstanceExtendTime:   dataModel.InstanceExtendTime.Format(timeutil.YmdHms), // 实例续期时间
		InstanceId:           dataModel.InstanceId,                                 // 实例ID
		NodeId:               dataModel.SnowflakeNodeId,                            // 雪花算法节点id
		InstanceName:         dataModel.InstanceName,                               // 实例名称
		InstanceEndpointList: dataModel.InstanceEndpointList,                       // 实例端点数组
		CreatedTime:          dataModel.CreatedTime.Format(timeutil.YmdHms),        // 创建时间
		InstanceMetadata:     dataModel.InstanceMetadata,                           // 实例元数据
	}
	return newDataModel
}
