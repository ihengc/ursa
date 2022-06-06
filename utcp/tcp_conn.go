package utcp

import (
	"net"
	"time"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/6 0006 16:54
* @version: 1.0
* @description:
*********************************************************/

type TCPConnection struct {
	uid         int64
	running     bool
	conn        net.Conn
	reqChannel  chan []byte
	respChannel chan []byte
	readTimeout time.Duration
}

// I/O intensive
func (tcpConn *TCPConnection) read() {
	for tcpConn.running {
		// read header by header size
		// decode header
		// read body by body size
		// put the data into work channel
	}
}

func (tcpConn *TCPConnection) GetReqChannel() chan []byte {
	return tcpConn.reqChannel
}

func (tcpConn *TCPConnection) Close() {
	tcpConn.running = false
	close(tcpConn.reqChannel)
	close(tcpConn.respChannel)
	tcpConn.conn.Close()
}

func NewTCPConnection(uid int64, conn net.Conn) *TCPConnection {
	tcpConn := new(TCPConnection)
	tcpConn.uid = uid
	tcpConn.conn = conn
	return tcpConn
}
