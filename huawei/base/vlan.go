package base

import (
	"fmt"
)

type Vlan struct {
	FieldA int
	FieldB string
}

func (c *Vlan) ShowA() {
	c.FieldB = "base"
	fmt.Printf("%v\n", c)
}
func (c *Vlan) ShowB() {
	c.FieldA = 1
	c.FieldB = "base"
	fmt.Printf("%v\n", c)
}
