package acceptor

import "net"

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/7 22:29
 * @description: 连接监听,接入
 ***************************************************************/

// TcpAcceptor 负责tcp链接的监听和接入
type TcpAcceptor struct {
	address     string        // 服务地址和端口号
	running     bool          // 表示监听服务是否运行
	ln          net.Listener  // 监听套接字
	connChannel chan net.Conn // 用于传递已经建立成功的链接
}

// ListenAndServe 启动监听
func (t *TcpAcceptor) ListenAndServe() {
	ln, err := net.Listen("tcp", t.address)
	if err != nil {
		// TODO record error
	}
	t.ln = ln
	t.running = true
	for t.running {
		conn, err := ln.Accept()
		if err != nil {
			// TODO record error
			continue
		}
		t.connChannel <- conn
	}
	defer t.Stop()
}

// Stop 停止监听
func (t *TcpAcceptor) Stop() {
	t.running = false
	t.ln.Close()
}

// LocalAddr 返回服务地址
func (t *TcpAcceptor) LocalAddr() string {
	return t.address
}

// GetConnChannel 返回链接队列
func (t *TcpAcceptor) GetConnChannel() chan net.Conn {
	return t.connChannel
}

// NewTcpAcceptor 创建 TcpAcceptor 对象,用于监听指定的tcp端口
func NewTcpAcceptor(address string) *TcpAcceptor {
	return &TcpAcceptor{
		address:     address,
		connChannel: make(chan net.Conn),
	}
}
