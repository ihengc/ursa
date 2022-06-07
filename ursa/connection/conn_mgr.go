package connection

import "sync"

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/7 20:43
 * @description:
 ***************************************************************/

var tcpConnMgr *TcpConnMgr

type TcpConnMgr struct {
	lock   *sync.Mutex
	counts int
}

func (t *TcpConnMgr) Counts() int {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.counts
}

func init() {
	tcpConnMgr = &TcpConnMgr{}
}

func GetTcpConnMgr() *TcpConnMgr {
	return tcpConnMgr
}
