package ursa_orm

import "time"

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 10:19
* @version: 1.0
* @description:
*********************************************************/

type Model struct {
	CreateAt time.Time // 新建时间
	UpdateAt time.Time // 更新时间
}
