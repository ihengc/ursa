package utcp

import (
	"fmt"
	"net"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/7 0007 11:53
* @version: 1.0
* @description:
*********************************************************/

var defaultTCPClient *TCPClient

type IClient interface {
	Send(message []byte)
}

// TCPClient 网络对端
type TCPClient struct {
	tcpConn net.Conn // tcpConn 已连接套接字
}

// Send 发送消息到对端
func (tcpCli *TCPClient) Send(data []byte) {
	tcpCli.tcpConn.Write(data)
}

// NewTCPClientWithConf 创建请求端
func NewTCPClientWithConf(configure *Configure) *TCPClient {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", configure.Host, configure.Port))
	if err != nil {
		panic(err)
	}
	return &TCPClient{
		tcpConn: conn,
	}
}

func GetDefaultTCPClient() *TCPClient {
	if defaultTCPClient == nil {
		defaultTCPClient = NewTCPClientWithConf(GetDefaultConfigure())
	}
	return defaultTCPClient
}
