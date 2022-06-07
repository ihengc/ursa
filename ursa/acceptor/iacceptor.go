package acceptor

import "net"

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/7 22:44
 * @description: 定义接口
 ***************************************************************/

// IAcceptor 监听模块接口
type IAcceptor interface {
	ListenAndServe()               // 启动监听
	Stop()                         // 停止监听
	GetConnChannel() chan net.Conn // 获取链接队列
	LocalAddr() string             // 获取监听地址和端口号
}
