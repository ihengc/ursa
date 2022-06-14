package Lepus

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/6/14 0014 18:01
* @version: 1.0
* @description:
*********************************************************/

var lepusDB *LepusDB

func init() {
	lepusDB = &LepusDB{}
}

type Player struct {
	ID   uint
	Name string
	Age  uint
}

func TestLepusDB_Create(t *testing.T) {
	players := make([]*Player, 0)
	players = append(players, &Player{ID: 1})
	players = append(players, &Player{ID: 2})
	players = append(players, &Player{ID: 3})
	lepusDB.Create(players[:])
}
