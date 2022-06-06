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
