package huawei

import (
	"github.com/wangyide/golearn/huawei/V01"
	"github.com/wangyide/golearn/huawei/V02"
	"github.com/wangyide/golearn/huawei/base"
)

func ShowA() {
	vb := base.Vlan{}
	v1 := V01.Vlan{}
	v2 := V02.Vlan{}
	vb.ShowA()
	vb.ShowB()
	v1.ShowA()
	v1.ShowB()
	v2.ShowA()
	v2.ShowB()
}
