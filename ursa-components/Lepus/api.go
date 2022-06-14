package Lepus

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 17:56
* @version: 1.0
* @description:
*********************************************************/

func (l *LepusDB) Create(value interface{}) {
	l.Stmt.Dest = value
	l.executor.Execute()
}
