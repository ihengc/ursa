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
	AddConn(conn IConnection)
	DelConnByConnUId(connUId int)
	Counts() int
}

var tcpConnMgr *TCPConnectionManager

// TCPConnectionManager 管理tcp链接
type TCPConnectionManager struct {
	lock        sync.Mutex          // lock 添加和删除可能会在多线程环境中,需要在多线程环境中同步数据
	counts      int                 // counts 表示已经保存的链接数
	connections map[int]IConnection // connections 用于保存链接
}

// AddConn 添加链接
func (tcpConnMgr *TCPConnectionManager) AddConn(conn IConnection) {
	tcpConnMgr.lock.Lock()
	defer tcpConnMgr.lock.Unlock()
	if _, ok := tcpConnMgr.connections[conn.GetConnUId()]; !ok {
		tcpConnMgr.connections[conn.GetConnUId()] = conn
		tcpConnMgr.counts++
	}
}

// DelConnByConnUId 通过ConnUId删除链接
func (tcpConnMgr *TCPConnectionManager) DelConnByConnUId(connUId int) {
	tcpConnMgr.lock.Lock()
	defer tcpConnMgr.lock.Unlock()
	if _, ok := tcpConnMgr.connections[connUId]; ok {
		delete(tcpConnMgr.connections, connUId)
		tcpConnMgr.counts--
	}
}

// Counts 返回链接数
func (tcpConnMgr *TCPConnectionManager) Counts() int {
	tcpConnMgr.lock.Lock()
	defer tcpConnMgr.lock.Unlock()
	return tcpConnMgr.counts
}

// init 直接初始化
func init() {
	if tcpConnMgr == nil {
		tcpConnMgr = &TCPConnectionManager{
			connections: make(map[int]IConnection),
		}
	}
}

// GetTCPConnectionMgr 获取tcp链接管理对象(单例)
func GetTCPConnectionMgr() *TCPConnectionManager {
	return tcpConnMgr
}
