package utcp

/********************************************************
* @author: Ihc
* @date: 2022/6/6 0006 17:12
* @version: 1.0
* @description:
*********************************************************/

type IRequest interface {
	GetId() int
}

type Request struct {
	Id int
}
