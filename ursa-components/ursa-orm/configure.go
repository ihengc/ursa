package ursa_orm

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 10:19
* @version: 1.0
* @description:
*********************************************************/

// Option 获取数据库连接时的配置选项接口
type Option interface {
	Apply(config *Config) error // 应用config配置
}

// Config orm配置
type Config struct {
	CreateBatchSize int // 批量新建
}
