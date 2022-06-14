package clause

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 14:04
* @version: 1.0
* @description: clause模块
*********************************************************/

// Writer 语句写入接口
type Writer interface {
	WriteString(s string) (int, error) // 写入语句
	WriteByte(c byte) error            // 写入符号
}

// Builder 语句构建接口
type Builder interface {
	Writer
	WriteQuoted(field interface{}) // 给字段增加引号
	AddVar(Writer, ...interface{}) // 给语句增加参数
}
