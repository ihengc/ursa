package utcp

import "sync"

/********************************************************
* @author: Ihc
* @date: 2022/6/7 0007 10:28
* @version: 1.0
* @description:
*********************************************************/

// ConnectionManager 链接管理接口
type ConnectionManager interface {
	AddConn(connId int, conn IConnection)
	DelConn(connId int)
	Counts() int
}

var (
	tcpConnMgr *TCPConnectionManager
	lock       = sync.Mutex{}
)

// TCPConnectionManager 管理tcp链接
type TCPConnectionManager struct {
	connections map[int]IConnection
	counts      int
}

// AddConn 添加链接
func (tcpConnMgr *TCPConnectionManager) AddConn(conn IConnection) {
	if _, ok := tcpConnMgr.connections[conn.GetConnUId()]; !ok {
		tcpConnMgr.connections[conn.GetConnUId()] = conn
		tcpConnMgr.counts++
	}
}

// DelConnByConnUId 通过ConnUId删除链接
func (tcpConnMgr *TCPConnectionManager) DelConnByConnUId(connUId int) {
	if _, ok := tcpConnMgr.connections[connUId]; ok {
		delete(tcpConnMgr.connections, connUId)
		tcpConnMgr.counts--
	}
}

// Counts 返回链接数
func (tcpConnMgr *TCPConnectionManager) Counts() int {
	return tcpConnMgr.counts
}

// GetTCPConnectionMgr 获取tcp链接管理对象(单例)
func GetTCPConnectionMgr() *TCPConnectionManager {
	lock.Lock()
	if tcpConnMgr == nil {
		tcpConnMgr = new(TCPConnectionManager)
		tcpConnMgr.connections = make(map[int]IConnection)
	}
	lock.Unlock()
	return tcpConnMgr
}
