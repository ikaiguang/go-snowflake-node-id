# 雪花算法 节点ID

`snowflake`雪花算法，依赖节点id(nodeID:节点ID，machineID:机器编码)

当业务服务需要扩容时，节点ID必须全局**唯一**，然后才能生产**唯一ID**

本工具主要功能：提供全局唯一的**节点ID**

## 数据存储

数据存储组件：数据库

数据库操作使用了[GORM](https://gorm.io/)，支持一下的数据库：

> 使用示例，请参考：[./node-id/node-id_test.go](node-id/node-id_test.go)

- [x] `MySQL`
- [x] `PostgreSQL`
- [x] `SQLite`
- [x] `SQL Server`

```mysql

-- 创建数据库
CREATE DATABASE `srv_snowflake` DEFAULT CHARSET utf8mb4;
USE `srv_snowflake`;

DROP TABLE IF EXISTS snowflake_worker_node;
CREATE TABLE snowflake_worker_node
(
    id                     BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'id',
    node_uuid              VARCHAR(255) DEFAULT '' NOT NULL COMMENT '唯一ID；instance_id+node_id',
    instance_launch_time   DATETIME(3)             NOT NULL COMMENT '实例启动时间',
    instance_extend_time   DATETIME(3)             NOT NULL COMMENT '实例续期时间',
    instance_id            VARCHAR(255) DEFAULT '' NOT NULL COMMENT '实例ID',
    snowflake_node_id      INT          DEFAULT 0  NOT NULL COMMENT '雪花算法节点id',
    instance_name          VARCHAR(255) DEFAULT '' NOT NULL COMMENT '实例名称',
    instance_endpoint_list JSON                    NOT NULL COMMENT '实例端点数组',
    instance_metadata      JSON COMMENT '实例元数据',
    created_time           DATETIME(3)             NOT NULL COMMENT '创建时间',
    PRIMARY KEY (id),
    UNIQUE KEY node_uuid (node_uuid),
    KEY instance_id (instance_id),
    KEY instance_extend_time (instance_extend_time),
    KEY snowflake_node_id (snowflake_node_id)
) ENGINE InnoDB
  DEFAULT CHARSET utf8mb4
    COMMENT '雪花算法节点ID';

```

**时序图**
![雪花算法ID节点颁发流程图](./statics/雪花算法ID节点颁发流程图@开广.drawio.png)