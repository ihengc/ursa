package service

import (
	"net"
	"ursa/connection/codec"
	"ursa/connection/packet"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/8 0008 13:48
* @version: 1.0
* @description: 请求分发服务
*********************************************************/

// HandleService 负责连接的处理
type HandleService struct {
	localHandleChannel  chan []byte          // 存放需要本地处理的请求
	remoteHandleChannel chan []byte          // 存放需要其他服务处理的请求
	decoder             codec.IPacketDecoder // 数据包解码器
}

// Handle 处理成功建立的连接
// 读取数据，解析数据
func (h *HandleService) Handle(conn net.Conn) {
	for {
		data, err := h.readData(conn)
		if err != nil {
			continue
		}

		packets, err := h.decoder.Decode(data)
		if err != nil {
			// TODO log error
			return
		}

		for i := range packets {
			if err := h.handlePacket(packets[i]); err != nil {
				// TODO log error
				return
			}
		}
	}
}

func (h *HandleService) readData(conn net.Conn) ([]byte, error) {
	return nil, nil
}

// handlePacket 处理数据包
func (h *HandleService) handlePacket(p *packet.Packet) error {
	switch p.Type {
	case packet.Handshake:
	case packet.Heartbeat:
	case packet.Data:

	}
	return nil
}
