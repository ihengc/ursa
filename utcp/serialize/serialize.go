package serialize

import (
	"io"
	"ursa/utcp"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/7 0007 15:23
* @version: 1.0
* @description:
*********************************************************/

// ISerialization 序列化接口
type ISerialization interface {
	Serialize(response interface{}) []byte          // Serialize 序列化
	Deserialize(data []byte, request utcp.IRequest) // Deserialize 反序列化
}

// Serialization 序列化
type Serialization struct {
}

func (sl *Serialization) Serialize(response interface{}) []byte {
	//TODO implement me
	panic("implement me")
}

// Deserialize 反序列化生成 utcp.IRequest
func (sl *Serialization) Deserialize(reader io.Reader, request utcp.IRequest) {
	// 读取请求头
	rawHeader := make([]byte, request.GetRequestHeaderSize())
	_, err := io.ReadFull(reader, rawHeader)
	if err != nil {
		return
	}
	//TODO 解析请求头,生成请求头

	// 读取请求体
	rawBody := make([]byte, request.GetBodySize())
	_, err = io.ReadFull(reader, rawBody)
	if err != nil {
		return
	}
	//TODO 生成请求对象
}
