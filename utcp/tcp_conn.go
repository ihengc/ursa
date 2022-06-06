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
	uid          int64
	running      bool
	conn         net.Conn
	workChannel  chan []byte
	writeChannel chan []byte
	readTimeout  time.Duration
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

// CPU intensive
func (tcpConn *TCPConnection) handle() {
	for tcpConn.running {
		for data := range tcpConn.workChannel {
			// decode
			panic(data)
		}
	}
}

func (tcpConn *TCPConnection) Handle() {
	tcpConn.running = true
}

func (tcpConn *TCPConnection) Close() {
	tcpConn.running = false
	close(tcpConn.workChannel)
	close(tcpConn.writeChannel)
	tcpConn.conn.Close()
}

func NewTCPConnection(uid int64, conn net.Conn) *TCPConnection {
	tcpConn := new(TCPConnection)
	tcpConn.uid = uid
	tcpConn.conn = conn
	return tcpConn
}
