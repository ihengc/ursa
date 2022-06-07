package utcp

/********************************************************
* @author: Ihc
* @date: 2022/6/7 0007 11:17
* @version: 1.0
* @description:
*********************************************************/

// PacketType 数据包的类型
type PacketType byte

const (
	Heartbeat PacketType = iota + 1 // 心跳数据包
)

type IPacket interface {
}
