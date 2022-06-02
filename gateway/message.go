package gateway

/********************************************************
* @author: Ihc
* @date: 2022/6/2 0002 17:19
* @version: 1.0
* @description:
request message length: 4 + 4 + 8 + 1 + 1 + 1 + len(data) = 19 + len(data)
response message length: 4 = 4 + 8 + 1 + 1 + 1 + 4 + len(data) = 23 + len(data)
*********************************************************/

const (
	RequestMessageHeaderSize  = 19
	ResponseMessageHeaderSize = 23
)

type RequestMessage struct {
	id            int32
	sequenceId    int32
	sendTimestamp int64
	isCompressed  bool
	isEncrypted   bool
	serviceType   byte
	data          []byte
}

type ResponseMessage struct {
	id            int32
	sequenceId    int32
	sendTimestamp int64
	isCompressed  bool
	isEncrypted   bool
	serviceType   byte
	errorCode     int32
	data          []byte
}
