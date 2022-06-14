package schema

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 10:35
* @version: 1.0
* @description: naming模块

*********************************************************/

// IName 命名接口
type IName interface {
	TableName(table string) string // 生成新的表名
}

// NamingStrategy 命名策略
type NamingStrategy struct {
	TablePrefix string // 表名前缀，每个表名都会加上此前缀
}

// TableName 根据指定的字符串生成表名
func (n NamingStrategy) TableName(table string) string {
	return n.TablePrefix + table
}

// IndexName 生成索引名称
func (n NamingStrategy) IndexName(table, column string) string {
	return "idx" + table + column
}
