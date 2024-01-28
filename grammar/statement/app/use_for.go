package app

import "fmt"

func UseFor() {
	// for 的条件里不需要括号
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for 的条件里可以省略初始条件，结束条件，递增表达式
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
