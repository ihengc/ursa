package utcp

import (
	"context"
	"fmt"
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
	router  IRouter
}

// handleConnection 链接处理
func (tcpServer *TCPServer) handleConnection(ctx context.Context, conn net.Conn) {
	tcpConn := NewTCPConnection(ctx, 1, conn, tcpServer.codec)
	tcpConnMgr := GetTCPConnectionMgr()
	tcpConnMgr.AddConn(tcpConn)
	tcpConn.Start()
}

// Start 启动TCP服务
func (tcpServer *TCPServer) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	tcpServer.running = true
	for tcpServer.running {
		conn, err := tcpServer.ln.Accept()
		if err != nil {
			continue
		}
		go tcpServer.handleConnection(ctx, conn)
	}
	defer cancel() // 关闭协程
}

// LocalAddr 获取监听的地址和端口号
func (tcpServer *TCPServer) LocalAddr() string {
	return tcpServer.ln.Addr().String()
}

func (tcpServer *TCPServer) Stop() {
	tcpServer.running = false
}

// NewTCPServerWithConf 创建TCP服务
func NewTCPServerWithConf(configure *Configure) *TCPServer {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", configure.Host, configure.Port))
	if err != nil {
		panic(err)
	}
	tcpServer := &TCPServer{
		ln:     ln,
		router: configure.Router,
		codec:  configure.Codec,
	}
	return tcpServer
}
