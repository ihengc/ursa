package ursa_orm

import (
	"strings"
	"testing"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 14:09
* @version: 1.0
* @description: statement模块测试
*********************************************************/

var stmt *Statement

func init() {
	stmt = &Statement{SQL: strings.Builder{}, Vars: make([]interface{}, 0)}
}

func TestStatement_WriteString(t *testing.T) {
	stmt.WriteString("SELECT")
	t.Log(stmt.SQL.String())
}
