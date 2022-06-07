package utcp

/********************************************************
* @author: Ihc
* @date: 2022/6/6 0006 17:23
* @version: 1.0
* @description:
*********************************************************/

type ICodec interface {
	Decode(data []byte) interface{}
	Encode(data interface{}) []byte
}

type JsonCodec struct {
}

func (jsonCodec *JsonCodec) Decode(data []byte) interface{} {
	//TODO implement me
	panic("implement me")
}

func (jsonCodec *JsonCodec) Encode(data interface{}) []byte {
	//TODO implement me
	panic("implement me")
}

func NewJsonCodec() *JsonCodec {
	jsonCodec := new(JsonCodec)
	return jsonCodec
}
