package app

import "fmt"

// 使用常量
func UseConst() {
	const name = "abc"
	fmt.Println(name)
}

// 使用枚举
func UseEnums() {
	// 在这个例子中，iota 的初始值为 0，然后在每次使用时自动递增，每次递增 1。
	//因此，b 的值为 1 左移 10 位（即 1024），kb 的值为 1 左移 20 位（即 1024 的平方），以此类推。
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(b, kb, mb, gb, tb)
}
