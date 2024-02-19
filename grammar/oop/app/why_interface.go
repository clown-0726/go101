package app

import "fmt"

type Laptop struct{}

func (Laptop) GetContent(param string) string {
	return "Get content from Laptop: " + param
}

type Desktop struct{}

func (Desktop) GetContent(param string) string {
	return "Get content from Desktop: " + param
}

type ComputerIntl interface {
	GetContent(string) string
}

func getComputer(cond int) ComputerIntl {
	if cond == 1 {
		return Desktop{}
	} else if cond == 2 {
		return Laptop{}
	} else {
		panic("Wrong cond...")
	}
}

func MainCalc() {
	cond := 1
	var c ComputerIntl = getComputer(cond)
	resStr := c.GetContent("http://abc.com/")
	fmt.Println(resStr)
}
