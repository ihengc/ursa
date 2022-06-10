package uuid

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/6/10 0010 14:46
* @version: 1.0
* @description: 测试UUID生成算法
*********************************************************/

func TestUUID1(t *testing.T) {
	v1, err := NewUUID1()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v1.String())
}
