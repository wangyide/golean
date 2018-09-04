package V02

import "fmt"

type Vlan struct {
	FieldA int
	FieldB string
}

func (c *Vlan) ShowA() {
	c.FieldB = "v02"
	fmt.Printf("%v\n", c)
}

func (c *Vlan) ShowB() {
	c.FieldA = 1
	c.FieldB = "v02"
	fmt.Printf("%v\n", c)
}
