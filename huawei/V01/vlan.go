package V01

import (
	"fmt"

	"github.com/wangyide/golearn/huawei/base"
)

type Vlan struct {
	base.Vlan
	FieldC bool
}

func (c *Vlan) ShowA() {
	c.FieldB = "v01"
	c.Vlan.FieldA = 10
	fmt.Printf("%v\n", c)
}
