package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/10 0010 09:29
* @version: 1.0
* @description: uuid生成模块
参照RFC 4122
*********************************************************/

// 1582.09.15 00:00:00 - 1970.01.01 00:00:00
// UUID的时间戳起始时间为1582.09.15 00:00:00,时间单位为100ns
// epochStart 为1582.09.15 00:00:00到1970.01.01 00:00:00的时间差值(单位为ns)
const epochStart = 122192928000000000

// epochFunc 获取当前时间,时间起点为1582.09.15 00:00:00
type epochFunc func() time.Time

// hardwareAddrFunc 获取硬件地址函数
type hardwareAddrFunc func() (net.HardwareAddr, error)

// IGenerator 生成UUID接口
type IGenerator interface {
	NewUUID1() (UUID, error) // 基于时间的UUID生成版本
	// TODO 实现UUID生成2-5
}

var generator = newUUIDGenerator()

// uuidGenerator RFC4122
type uuidGenerator struct {
	storageLock       sync.Mutex       // 用于同步时钟序列
	clockSequenceOnce sync.Once        // 保证时间序列只生成一次
	hardwareAddrOnce  sync.Once        // 保证硬件地址只被获取一次
	hardwareAddrFunc  hardwareAddrFunc // 用于获取硬件地址
	epochFunc         epochFunc        // 用于获取当前时间戳
	rand              io.Reader        // 伪随机接口
	lastGenerateTime  uint64           // 上次生成时间戳
	clockSequence     uint16           // 时间序列2 字节
	hardwareAddr      [6]byte          // 硬件地址6 字节
}

// getHardwareAddr 获取硬件地址
// 获取硬件地址是随机的还是真的网卡MAC基于 hardwareAddrFunc
// 的实现方式，在创建生成器后，生成器存活的周期内，硬件地址只
// 会被获取一次（在第一次生成UUID时硬件地址将会被获取到，保存
// 在生成器中，下一次再次生成UUID将会直接使用已经保存的硬件地
// 址。
func (ug *uuidGenerator) getHardwareAddr() ([]byte, error) {
	var err error
	ug.hardwareAddrOnce.Do(func() {
		if hardwareAddr, err := ug.hardwareAddrFunc(); err == nil {
			copy(ug.hardwareAddr[:], hardwareAddr)
			return
		}

		if _, err := io.ReadFull(ug.rand, ug.hardwareAddr[:]); err != nil {
			return
		}
		ug.hardwareAddr[0] |= 0x01
	})
	if err != nil {
		return []byte{}, err
	}
	return ug.hardwareAddr[:], nil
}

// getEpoch 获取从1582.09.15 00:00:00到现在的时间戳（单位ns)
func (ug *uuidGenerator) getEpoch() uint64 {
	return epochStart + uint64(ug.epochFunc().UnixNano()/100)
}

// getClockSequence 获取时钟序列，当前时间戳
// 在第一次生成UUID时，生成时钟序列。之后每次生成UUID时，若系
// 统时钟发生回退，则时钟序列会自增；否则保持不变
func (ug *uuidGenerator) getClockSequence() (uint64, uint16, error) {
	var err error
	ug.clockSequenceOnce.Do(func() {
		buffer := make([]byte, 2)
		if _, err := io.ReadFull(ug.rand, buffer); err != nil {
			return
		}
		ug.clockSequence = binary.BigEndian.Uint16(buffer)
	})
	if err != nil {
		return 0, 0, err
	}
	ug.storageLock.Lock()
	defer ug.storageLock.Unlock()

	now := ug.getEpoch()
	if ug.lastGenerateTime >= now {
		ug.clockSequence++
	}
	ug.lastGenerateTime = now
	return now, ug.clockSequence, nil
}

// newUUIDGenerator 创建uuid生成器
func newUUIDGenerator() *uuidGenerator {
	return &uuidGenerator{
		epochFunc:        time.Now,
		hardwareAddrFunc: defaultHardwareAddrFunc,
		rand:             rand.Reader,
	}
}

// defaultHardwareAddrFunc 获取硬件地址
// net.HardwareAddr 是一个字节数组
func defaultHardwareAddrFunc() (net.HardwareAddr, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return []byte{}, err
	}
	for _, iface := range ifaces {
		if len(iface.HardwareAddr) >= 6 {
			return iface.HardwareAddr, nil
		}
	}
	return []byte{}, fmt.Errorf("uuid no hardware address found")
}

// NewUUID1 UUID版本1生成接口
func (ug *uuidGenerator) NewUUID1() (UUID, error) {
	uuid := UUID{}
	timestamp, clockSeq, err := ug.getClockSequence() // 获取时间戳,时钟序列
	if err != nil {
		return NilUUID, err
	}

	binary.BigEndian.PutUint32(uuid[0:], uint32(timestamp))     // 时间戳低位
	binary.BigEndian.PutUint16(uuid[4:], uint16(timestamp>>32)) // 时间戳中位
	binary.BigEndian.PutUint16(uuid[6:], uint16(timestamp>>48)) // 时间戳低位
	binary.BigEndian.PutUint16(uuid[8:], clockSeq)              // 时钟序列

	hardwareAddr, err := ug.getHardwareAddr() // 获取硬件地址
	if err != nil {
		return NilUUID, err
	}
	copy(uuid[10:], hardwareAddr)

	uuid.SetVersion(V1)
	uuid.SetVariant(VariantRFC4122)

	return uuid, nil
}

// 对包外暴露生成接口

// NewUUID1 UUID生成(版本1)
func NewUUID1() (UUID, error) {
	return generator.NewUUID1()
}
