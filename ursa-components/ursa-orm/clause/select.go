package clause

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 14:33
* @version: 1.0
* @description: select模块
*********************************************************/

// Select 表示SELECT语句
type Select struct {
	Distinct bool     // DISTINCT修饰符
	Columns  []Column // SELECT查询的字段
}

// Name 返回语句字符串表示
func (s Select) Name() string {
	return "SELECT"
}

// Build 构建SELECT语句
// 若Columns无数据，则表示SELECT *
// 否则 SELECT `field1`, `field2`
func (s Select) Build(builder Builder) {
	if len(s.Columns) > 0 {
		if s.Distinct {
			builder.WriteString("DISTINCT ")
		}
		for idx, column := range s.Columns {
			if idx > 0 {
				builder.WriteByte(',')
			}
			builder.WriteQuoted(column)
		}
	} else {
		builder.WriteByte('*')
	}
}
