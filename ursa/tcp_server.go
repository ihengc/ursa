package ursa

import (
	context "context"
	"net"
	"ursa/ursa/connection"
	"ursa/ursa/internal/linkedlist"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/7 19:21
 * @description: tcp服务模块
 ***************************************************************/

// INetServer 网络服务模块接口
type INetServer interface {
	ListenAndServe()   // 启动监听
	Stop()             // 停止服务
	GetRunningStatus() // 获取服务运行状态
	LocalAddr() string // 获取监听的地址和端口号
}

// TcpServer tcp服务
type TcpServer struct {
	ln           net.Listener               // 监听套接字
	running      bool                       // 服务运行状态
	maxConnNum   int                        // 最大链接数
	waitingQueue *linkedlist.List[net.Conn] // 等待处理的链接
}

func (ts *TcpServer) handleConn(ctx context.Context, conn net.Conn) {

}

// ListenAndServe 启动监听
func (ts *TcpServer) ListenAndServe() {
	ts.running = true
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		conn net.Conn
		err  error
	)
	tcpConnMgr := connection.GetTcpConnMgr()
	for ts.running {
		if tcpConnMgr.Counts() >= ts.maxConnNum {
			ts.waitingQueue.PushBack(conn)
			continue
		}
		if ts.waitingQueue.Len() != 0 {
			conn = ts.waitingQueue.PopFront().Value
		} else {
			conn, err = ts.ln.Accept()
			if err != nil {
				continue
			}
		}
		go ts.handleConn(ctx, conn)
	}
}
