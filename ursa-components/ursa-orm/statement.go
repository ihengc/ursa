package ursa_orm

import (
	"strings"
	"ursa/ursa-components/ursa-orm/clause"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 11:22
* @version: 1.0
* @description: statement模块

1。构建sql
	执行的语句由sql语句和参数构成

*********************************************************/

type IStatement interface {
	WriteByte(c byte) error
	WriteString(s string) (int, error)
	AddVar()
	AddClause()
	AddCondition()
}

type Statement struct {
	SQL  strings.Builder // sql语句
	Vars []interface{}   // 语句参数
}

// WriteString 构建SQL
func (stmt *Statement) WriteString(s string) (int, error) {
	return stmt.SQL.WriteString(s)
}

// WriteByte 构建SQL
func (stmt *Statement) WriteByte(c byte) error {
	return stmt.SQL.WriteByte(c)
}

// QuoteTo 给语句增加引号
func (stmt *Statement) QuoteTo(writer clause.Writer, field interface{}) {

}
