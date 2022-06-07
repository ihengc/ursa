package utcp

import "time"

/********************************************************
* @author: Ihc
* @date: 2022/6/7 0007 11:11
* @version: 1.0
* @description:
*********************************************************/

// DefaultConfigure 服务默认配置
var DefaultConfigure *Configure

// Configure 服务配置信息
type Configure struct {
	Host              string        // Host 主机地址
	Port              int           // Port 端口号
	Router            IRouter       // Route 路由
	Codec             ICodec        // Codec 编解码器
	HeartbeatInterval time.Duration // HeartbeatInterval 心跳间隔
}

func init() {
	DefaultConfigure = &Configure{
		Host:              "localhost",
		Port:              9090,
		Router:            NewRouter(),
		Codec:             NewJsonCodec(),
		HeartbeatInterval: 1 * time.Second,
	}
}
