package ursa_orm

import "database/sql"

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 11:06
* @version: 1.0
* @description:
*********************************************************/

type DB struct {
	*Config
	dbConn *sql.DB
}

func (db *DB) DBConn() *sql.DB {
	return db.dbConn
}

// Open 建立与数据库的连接
func Open(driverName, dataSourceName string) (*DB, error) {
	dbConn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	db := &DB{dbConn: dbConn}
	return db, nil
}
