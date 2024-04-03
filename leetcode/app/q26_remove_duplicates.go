package app

import "fmt"

/*
去除数组重复元素

解题思路是快慢指针，快指针遍历整个元素，慢指针指向最后一个与快指针不同的元素位置
*/

func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

func MainRemoveDuplicates() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := removeDuplicates(nums)
	fmt.Println(res)
}
