package app

import "fmt"

func Main4LearnPoint() {
	str := "A"
	pointer := &str            // 数据的指针
	anotherString := *&str     // 数据指针的值，也就是 value
	fmt.Println(str)           // A
	fmt.Println(pointer)       // 0xc000096230
	fmt.Println(anotherString) // A

	fmt.Println("---------")

	str = "B"
	fmt.Println(str)           // B
	fmt.Println(&str)          // 0xc000096230
	fmt.Println(pointer)       // 0xc000096230
	fmt.Println(anotherString) // A
}
