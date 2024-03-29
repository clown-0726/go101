package app

/*
去除数组重复元素

通过快慢指针实现，整体思路是快指针遍历整个数组，慢指针标识最后更改的位置
当快指针发现任意一个元素与目标不一致时（需要留下的），就将这个元素往慢指针的位置更新，慢指针 +1
*/

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func MainRemoveElement() {
	nums := []int{3, 2, 2, 3}
	val := 3

	removeElement(nums, val)
}
