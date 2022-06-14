package clause

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 14:34
* @version: 1.0
* @description:
*********************************************************/

type Column struct {
	Raw   bool
	Table string
	Name  string
	Alias string
}
