package utcp

import (
	"context"
	"net"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/6 21:32
 * @description:
 ***************************************************************/

type TCPServer struct {
	ln      net.Listener
	running bool
	codec   ICodec
	route   IRouter
	ctx     context.Context
}

func (tcpServer *TCPServer) handle(conn net.Conn) {

}

func (tcpServer *TCPServer) Start() {
	tcpServer.running = true
	for tcpServer.running {
		conn, err := tcpServer.ln.Accept()
		if err != nil {
			continue
		}
		go tcpServer.handle(conn)
	}
}

func (tcpServer *TCPServer) LocalAddr() string {
	return tcpServer.ln.Addr().String()
}

func (tcpServer *TCPServer) Stop() {
	tcpServer.running = false
}
