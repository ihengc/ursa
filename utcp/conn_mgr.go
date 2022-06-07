package utcp

/********************************************************
* @author: Ihc
* @date: 2022/6/7 0007 10:28
* @version: 1.0
* @description:
*********************************************************/

type ConnectionManager interface {
	AddConn(connId int, conn IConnection)
	DelConn(connId int)
	Counts() int
}

var tcpConnMgr *TCPConnectionManager

type TCPConnectionManager struct {
	connections map[int]IConnection
	counts      int
}

func (tcpConnMgr *TCPConnectionManager) AddConn(connId int, conn IConnection) {
	if _, ok := tcpConnMgr.connections[connId]; !ok {
		tcpConnMgr.connections[connId] = conn
		tcpConnMgr.counts++
	}
}

func (tcpConnMgr *TCPConnectionManager) DelConn(connId int) {
	if _, ok := tcpConnMgr.connections[connId]; ok {
		delete(tcpConnMgr.connections, connId)
		tcpConnMgr.counts--
	}
}

func (tcpConnMgr *TCPConnectionManager) Counts() int {
	return tcpConnMgr.counts
}

func init() {
	tcpConnMgr = NewTCPConnectionManager()
}

func GetTCPConnectionMgr() *TCPConnectionManager {
	return tcpConnMgr
}

func NewTCPConnectionManager() *TCPConnectionManager {
	if tcpConnMgr == nil {
		tcpConnMgr = new(TCPConnectionManager)
		tcpConnMgr.connections = make(map[int]IConnection)
	}
	return tcpConnMgr
}
