package utcp

import (
	"net"
	"ursa/utcp/internal"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/6 21:32
 * @description:
状态值:
	1.服务的运行状态
		1) 运行
		2) 关闭
	2.服务的压力状态
		1) 流畅
		2) 拥挤
		3) 负载
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

// IsLoad 服务是否负载
// 如何判断服务是否负载?
// 1.根据压测值,设置在线人数,当总的在线人数到达预设值,则标记服务为负载
// 2.增加限流方案,被限流则表示已经负载
func (tcpServer *TCPServer) IsLoad() bool {
	return false
}
