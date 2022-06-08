package codec

import "ursa/connection/packet"

/********************************************************
* @author: Ihc
* @date: 2022/6/8 0008 16:21
* @version: 1.0
* @description: 编解码器
*********************************************************/

// IPacketDecoder 数据解码接口
type IPacketDecoder interface {
	GetHeaderSize() int                           // 获取数据包头大小
	Decode(data []byte) ([]*packet.Packet, error) // 解析数据生成数据包
}

// IPacketEncoder 数据编码接口
type IPacketEncoder interface {
	Encode(packetType packet.PacketType, data []byte) ([]byte, error) // 生成数据包
}
