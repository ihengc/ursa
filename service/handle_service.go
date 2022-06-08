package service

import "net"

/********************************************************
* @author: Ihc
* @date: 2022/6/8 0008 13:48
* @version: 1.0
* @description:
*********************************************************/

// HandleService 负责连接的处理
type HandleService struct {
}

// Handle 处理成功建立的连接
func (h *HandleService) Handle(conn net.Conn) {

}
