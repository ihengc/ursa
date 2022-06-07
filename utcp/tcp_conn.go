package utcp

import (
	"context"
	"net"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/6 0006 16:54
* @version: 1.0
* @description:
*********************************************************/

type IConnection interface {
	Start()
	Close()
	GetConnUId() int
}

// TCPConnection 表示一个成功建立的tcp链接
type TCPConnection struct {
	ctx     context.Context // ctx 用于关闭读写协程
	connUId int             // connUId 链接的唯一标识
	conn    net.Conn        // conn 已连接套接字
	codec   ICodec          // codec 编解码器;用于解析来自对端的数据或编码返回的数据
	running bool            // running 表示此链接的状态
}

// read 读取来自对端的数据
func (tcpConn *TCPConnection) read() {
	// 1.数据包在这里被解析成,协议头和协议体
	// 2.调用业务函数
	// 3.将结果
}

// write 写出本端的数据
func (tcpConn *TCPConnection) write() {

}

// Start 接收和发送数据
func (tcpConn *TCPConnection) Start() {
	// 读处理业务与写异步
	go tcpConn.read()
	go tcpConn.write()
}

// GetConnUID 获取链接的唯一标识
func (tcpConn *TCPConnection) GetConnUId() int {
	return tcpConn.connUId
}

// Close 关闭此链接
func (tcpConn *TCPConnection) Close() {
	tcpConn.running = false
	tcpConn.conn.Close()
}

// NewTCPConnection 创建tcp链接
func NewTCPConnection(ctx context.Context, connUId int, conn net.Conn, codec ICodec) *TCPConnection {
	tcpConn := &TCPConnection{
		ctx:     ctx,
		codec:   codec,
		connUId: connUId,
		conn:    conn,
	}
	return tcpConn
}
