package ursa_orm

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 10:19
* @version: 1.0
* @description:
*********************************************************/

// Config orm配置
type Config struct {
	MaxOpenConns int // 与数据库建立连接的最大数目
	MaxIdleConns int // 数据库连接池中最大空置的连接数
}
