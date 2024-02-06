package app

import "fmt"

func SliceBasic() {
	// 切片是数组的视图
	// 切片都是前闭后开的
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])
}

func SliceUpdate() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	s1 := arr[:4]
	updateSlice(s1)
	fmt.Println(s1)
}

func updateSlice(s []int) {
	s[0] = 1000
}

func ReSliceExtendingSlice() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	// ReSlice
	s2 := arr[:]
	s2 = arr[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	// Extend Slice
	// 可以向后扩展，不可以向前扩展
	s3 := arr[2:6]
	s4 := arr[3:5]
	// fmt.Println(s3[4]) // index out of range
	fmt.Println(s3, s4)
}

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func SliceOps() {
	// 创建 slice 并赋值
	var s []int // Zero value for slice is nil
	// 在向切片中加入元素的时候，切片的最大长度 cap 会根据元素多少进行翻倍扩张
	for i := 0; i < 30; i++ {
		printSlice(s)
		s = append(s, i)
	}
	fmt.Println(s)

	// Make slice
	// 建立一个已知大小的 slice
	s2 := make([]int, 16)
	s22 := new([]int)         // new 返回的是数组的指针
	s3 := make([]int, 16, 32) // 但是其后面的数组要申请 32 的长度
	fmt.Println(s2, s22)
	fmt.Println(s3)

	// Copy slice
	sCopy := make([]int, 16, 32)
	copy(sCopy, s)
	fmt.Println(sCopy)

	// 删除一个元素
	s4 := append(s[:3], s[4:]...)
	fmt.Println(s4)

	// 从头尾删除元素
	// 从头部删除
	front := s[0]
	s = s[1:]

	// 从尾部删除
	end := s[len(s)-1]
	s = s[:len(s)-2]
	fmt.Println(front, end)
}
