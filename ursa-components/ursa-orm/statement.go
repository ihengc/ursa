package ursa_orm

import (
	"database/sql"
	"database/sql/driver"
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

type Statement struct {
	DB   *DB
	SQL  strings.Builder // sql语句
	Vars []interface{}   // 语句参数
	Dest interface{}
}

func (stmt *Statement) WriteString(s string) (int, error) {
	return stmt.SQL.WriteString(s)
}

func (stmt *Statement) WriteByte(c byte) error {
	return stmt.WriteByte(c)
}

func (stmt *Statement) WriteQuoted(field interface{}) {
	//TODO implement me
	panic("implement me")
}

func (stmt *Statement) AddVar(writer clause.Writer, vars ...interface{}) {
	for idx, v := range vars {
		if idx > 0 {
			writer.WriteByte(',')
		}

		switch v := v.(type) {
		case sql.NamedArg:
			stmt.Vars = append(stmt.Vars, v.Value)
		case clause.Column, clause.Table:
		case driver.Valuer:
		case []byte:
		case []interface{}:
		case *DB:
		default:

		}
	}
}
