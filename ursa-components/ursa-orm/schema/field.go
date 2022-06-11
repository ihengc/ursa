package schema

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/11 11:54
 * @description: field模块

field数据结构:
	1.字段名称
	2.约束
		2.1 唯一约束
			唯一约束是不考虑NULL。
			复合唯一约束
		2.2 主键约束
			单一主键
			复合主键
		2.3 非空约束
		2.4 外键约束
		2.5 检查约束
		2.6 默认值约束
		2.7 自增约束
			自增约束要求
				数据类型为整型
				不能用于多个自增列
				主键列或唯一键列
	3.数据类型
		数据库与golang数据类型不一致如何解决？
	4.注释
	5.默认值

Golang数据类型与数据库字段数据类型的映射：
	MySQL:
		整型:
			tinyint		1字节	uint8/int8
			smallint 	2字节	uint16/int16
			mediumint	3字节
			int			4		uint32/int32
			bigint		8		uint64/int64
		浮点型:
			float 		4		float32
			double		8		float64
			decimal	    decimal(m,d)
		时间日期类型：
			date		3字节	YYYY-MM-DD
			time		3字节	HH:MM:SS
			year		1字节
			datetime	8字节
			timestamp	4字节
		字符串类型:
			char		0-255
			varchar		0-65535

 ***************************************************************/

// FieldConstraint 字段约束
type FieldConstraint struct {
	PrimaryKey    bool // 是否是主键
	Unique        bool // 是否唯一
	AutoIncrement bool // 是否自增
	NotNull       bool // 非空
}
