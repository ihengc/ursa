package utcp

import "time"

/********************************************************
* @author: Ihc
* @date: 2022/6/6 0006 17:12
* @version: 1.0
* @description:
*********************************************************/

// IRequestHeader 请求头接口
type IRequestHeader interface {
	GetRouteId() int            // GetRouteId 获取此数据报的路由Id
	GetSendTime() time.Duration // GetSendTime 获取此数据报的发送时间
	GetServerType() int         // GetServerType 获取服务类型
	GetRequestHeaderSize() int  // GetRequestHeaderSize 获取请求头大小
	GetBodySize() int           // GetBodySize 获取请求体大小
}

// IRequest 请求接口
type IRequest interface {
	IRequestHeader          // IRequestHeader 请求头
	GetRequestBody() []byte // GetRequestBody 获取请求体
}

// RequestHeader 请求头
type RequestHeader struct {
	routeId      int           // routeId 路由Id
	sequenceId   int           // sequenceId 序列号
	isCompressed bool          // isCompressed 是否压缩
	isEncrypted  bool          // isEncrypted是否加密
	serverType   int           // serverType 服务类型(分布式服务中使用)
	sendTime     time.Duration // sendTime 包的发送时间
	packetSize   int           // packetSize 请求的总大小
}

// Request 请求
type Request struct {
	requestHeader IRequestHeader
	body          []byte // body 请求体
}

func (req *Request) GetServerType() int {
	return req.requestHeader.GetServerType()
}

func (req *Request) GetRequestHeaderSize() int {
	return req.requestHeader.GetRequestHeaderSize()
}

func (req *Request) GetRouteId() int {
	return req.requestHeader.GetRouteId()
}

func (req *Request) GetSendTime() time.Duration {
	return req.requestHeader.GetSendTime()
}

func (req *Request) GetRequestBody() []byte {
	return req.body
}
