package dialect

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 10:13
* @version: 1.0
* @description:
*********************************************************/

// IDialect 数据库类型
type IDialect interface {
	Name() string // 数据库名称
}
