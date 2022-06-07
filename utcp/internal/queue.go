package internal

import "net"

/********************************************************
* @author: Ihc
* @date: 2022/6/7 0007 17:55
* @version: 1.0
* @description:
*********************************************************/

type IQueue[T any] interface {
	IsEmpty() bool
	Get() net.Conn
	Len() int
}

type Queue struct {
}
