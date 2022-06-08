package packet

/********************************************************
* @author: Ihc
* @date: 2022/6/8 0008 16:23
* @version: 1.0
* @description: 数据包
*********************************************************/

// PacketType 将数据包进行分类，方便处理
type PacketType byte

const (
	Handshake PacketType = iota + 1 // 在握手时发送此类型数据报
	Heartbeat                       // 心跳检测包
	Data                            // 用于业务的数据包
)

// Packet 数据包
type Packet struct {
	Type PacketType // 数据包的类型
	Len  int        // 数据包的长度
	Data []byte     // 数据包的数据
}
