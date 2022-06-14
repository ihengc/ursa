package ursa_orm

import "reflect"

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 10:55
* @version: 1.0
* @description: api
*********************************************************/

// Create 创建数据
func (db *DB) Create(value interface{}) *DB {
	if db.CreateBatchSize > 0 {
		return db.CreateBatches(value, db.CreateBatchSize)
	}
	// 获取sql执行对象
	processor := db.executor.Create()
	// 执行sql
	return processor.Execute(db)
}

// CreateBatches 批量创建
func (db *DB) CreateBatches(value interface{}, batchSize int) *DB {
	// 解析玩家传入的value数据类型
	// 若value是一个指针,则需要获取其真实的值
	// 若value不是指针,则直接获取其值
	refValue := reflect.Indirect(reflect.ValueOf(value))

	switch refValue.Kind() {
	case reflect.Slice, reflect.Array:
		// 创建数据可能出错
		for i := 0; i < refValue.Len(); i += batchSize {
			endIdx := i + batchSize
			if endIdx > refValue.Len() {
				endIdx = refValue.Len()
			}
			// 获取切片或数据中的值
			db.Statement.Dest = refValue.Slice(i, endIdx).Interface()
			// 使用切片中的值执行sql
			processor := db.executor.Create()
			processor.Execute(db)
		}
		return db
	default:
		processor := db.executor.Create()
		return processor.Execute(db)
	}
}
