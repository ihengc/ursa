package ursa_orm

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 10:59
* @version: 1.0
* @description:
*********************************************************/

type Executor struct {
	processor *processor
}

func (e *Executor) Create() *processor {
	return e.processor
}

// processor 负责sql执行
type processor struct {
}

// Execute 执行sql
func (p *processor) Execute(db *DB) *DB {
	// 使用执行sql接口
	return nil
}
