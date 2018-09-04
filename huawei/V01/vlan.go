package V01

import (
	"fmt"

	"github.com/wangyide/golearn/huawei/base"
)

type Vlan struct {
	base.Vlan
}

func (c *Vlan) ShowA() {
	c.FieldB = "v01"
	fmt.Printf("%v\n", c)
}
