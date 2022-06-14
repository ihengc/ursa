package ursa_orm

import (
	"database/sql"
	"ursa/ursa-components/ursa-orm/dialect"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 11:06
* @version: 1.0
* @description:
*********************************************************/

type DB struct {
	*Config
	rawDB     *sql.DB
	executor  *Executor
	Statement *Statement
}

func (db *DB) getInstance() *DB {
	return db
}

// Open 打开数据库连接
func Open(dialect dialect.IDialect, opts ...Option) (*DB, error) {
	conf := &Config{}
	for _, opt := range opts {
		if opt != nil {
			if err := opt.Apply(conf); err != nil {
				return nil, err
			}
		}
	}
	db := &DB{}
	db.Statement = &Statement{
		DB: db,
	}
	return db, nil
}
