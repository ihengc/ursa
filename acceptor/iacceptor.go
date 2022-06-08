package acceptor

import "net"

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/7 22:44
 * @description: 监听模块接口
IAcceptor接口应该被用于前台服务(与客户端之间交互的服务)
后台服务之间的交互应该使用RPC接口
 ***************************************************************/

// IAcceptor 监听模块接口
type IAcceptor interface {
	ListenAndServe()               // 启动监听
	Stop()                         // 停止监听
	GetConnChannel() chan net.Conn // 获取链接队列
	LocalAddr() string             // 获取监听地址和端口号
}
