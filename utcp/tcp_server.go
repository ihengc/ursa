package utcp

import "net"

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/6 21:32
 * @description:
 ***************************************************************/

type TCPServer struct {
	ln      net.Listener
	running bool
	codec   ICodec
	route   IRoute
}

func (tcpServer *TCPServer) Start() {
	tcpServer.running = true
	for tcpServer.running {
		conn, err := tcpServer.ln.Accept()
		if err != nil {
			continue
		}
		tcpServer.handle(conn)
	}
}

func (tcpServer *TCPServer) handle(conn net.Conn) {
	tcpConn := NewTCPConnection(1, conn)

	reqChannel := tcpConn.GetReqChannel()
	respChannel := tcpConn.GetRespChannel()
	for {
		req, ok := <-reqChannel
		if !ok {
			break
		}
		handler := tcpServer.route.GetHandler(req.GetId())
		resp := NewResponse()
		handler(req, resp)
		respChannel <- resp
	}
}
