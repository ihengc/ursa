package uuid

import (
	"bytes"
	"encoding/hex"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/9 0009 14:01
* @version: 1.0
* @description: uuid
RFC URL: https://www.ietf.org/rfc/rfc4122.txt

uuid占16字节
uuid形式:
	xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx(32位16进制数)
	4-2-2-2-6
	时间值低位(4)-时间值中位(2)-时间值高位和版本号(2)-时钟序列高位和变体值与时钟序列低位(2)-节点值(6)
uuid构成:
	00-05(6字节):节点值。
	06(1字节):时钟序列低位。
	07(1字节):
		2bit:变体值
		6bit:时钟序列高位
	08(1字节):时间值的高位。
	09(1字节):
		4bit:版本号
		4bit:时间值高位
	10-11(2字节):时间值中位。
	12-15(4字节):时间值低位

uuid占16 * 8 = 128bit
时间值:时间值占60位,单位是100纳秒。
节点值:由机器的MAC地址构成。若机器有多个MAC地址，则随机选取其中一个；若机器没有MAC地址，则采用随机数。
*********************************************************/

// Size UUID总长度
const Size = 16

// UUID 版本
const (
	V1 byte = iota + 1 // 基于时间的UUID，依赖当前时间戳以及机器mac地址
	v2                 // 分布式安全的UUID
	v3                 // 基于命名空间的UUID(MD5版)
	v4                 // 基于随机数的UUID
	v5                 // 基于命名空间的UUID（SHA1版)
)

// UUID 变体
const (
	VariantNCS byte = iota
	VariantRFC4122
	VariantMicrosoft
	VariantFuture
)

type UUID [Size]byte

var NilUUID = UUID{}

// SetVersion 设置UUID的版本号
func (u *UUID) SetVersion(v byte) {
	// 第7个字节的前4位存放UUID的版本,将前4位设置为0,然后设置前4位为新的版本号
	//   0000 0000
	// & 0000 1111
	// | 0001 0000
	u[6] = (u[6] & 0x0F) | (v << 4)
}

// Version 获取UUID的版本号
func (u UUID) Version() byte {
	return u[6] >> 4
}

// SetVariant 设置变体值
func (u *UUID) SetVariant(v byte) {
	switch v {
	case VariantNCS:
	case VariantRFC4122:
	case VariantMicrosoft:
	case VariantFuture:
	default:

	}
}

// Variant 获取变体值
func (u UUID) Variant() byte {
	switch {
	case (u[8] >> 7) == 0x00:
		return VariantNCS
	case (u[8] >> 6) == 0x02:
		return VariantRFC4122
	case (u[8] >> 5) == 0x06:
		return VariantMicrosoft
	case (u[8] >> 5) == 0x07:
		fallthrough
	default:
		return VariantFuture
	}
}

// String 将字节形式的uuid转换成
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx
func (u UUID) String() string {
	buffer := make([]byte, 36)

	hex.Encode(buffer[0:8], u[0:4])
	buffer[8] = '-'
	hex.Encode(buffer[9:13], u[4:6])
	buffer[13] = '-'
	hex.Encode(buffer[14:18], u[6:8])
	buffer[18] = '-'
	hex.Encode(buffer[19:23], u[8:10])
	buffer[23] = '-'
	hex.Encode(buffer[24:], u[10:])

	return string(buffer)
}

// Equal 若u1和u2相等返回true；否则返回false
func Equal(u1 UUID, u2 UUID) bool {
	return bytes.Equal(u1[:], u2[:])
}
