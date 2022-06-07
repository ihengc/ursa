package utcp

import (
	"net"
	"ursa/utcp/internal"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/6 21:32
 * @description:
 ***************************************************************/

// TCPServer tcp服务
type TCPServer struct {
	ln           net.Listener    // ln 监听套接字
	isRunning    bool            // isRunning 表示服务是否正在运行
	limitLoad    int             // limitLoad 负载上限
	waitingQueue internal.IQueue // waitingQueue 排队队列
}

// ListenAndServe 启动监听
func (tcpServer *TCPServer) ListenAndServe() {
	tcpServer.isRunning = true
	for tcpServer.isRunning {
		// 1.若当前服务已经负载
	}
}
