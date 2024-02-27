package app

import "fmt"

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

type Plane struct {
	vendor string
	model  string
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

type Car struct {
	factory, model string
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func MainInterfaceArray() {
	interfaces := []IF{}

	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)

	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)

	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	fmt.Println(p.getName())
	interfaces = append(interfaces, p)

	// 遍历接口集合中的元素，并调用 getName 方法
	for _, f := range interfaces {
		fmt.Println(f.getName())
	}
}
