package uuid

import "bytes"

/********************************************************
* @author: Ihc
* @date: 2022/6/9 0009 14:01
* @version: 1.0
* @description: uuid
uuid占16字节
uuid形式:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx(32位16进制数)
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

// Equal 若u1和u2相等返回true；否则返回false
func Equal(u1 UUID, u2 UUID) bool {
	return bytes.Equal(u1[:], u2[:])
}
