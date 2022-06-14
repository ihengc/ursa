package Lepus

import "database/sql"

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 17:54
* @version: 1.0
* @description:
*********************************************************/

type LepusDB struct {
	dbConn   *sql.DB
	Stmt     *Stmt
	executor *Executor
}

func Open(driverName, dataSourceName string) *LepusDB {
	dbConn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil
	}
	lepusDB := &LepusDB{dbConn: dbConn}
	lepusDB.Stmt = &Stmt{
		db: lepusDB,
	}
	return lepusDB
}
