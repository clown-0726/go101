package app

import "fmt"

/*
合并两个有序数组
最简单的思路就是暴力合并，也就是把两个数组合并在一起之后进行排序。

效率比较高的解法是采用双指针：
- 从两个数组中去分别取出数据进行比较并将比较小的值放到第三个数组中
- 最后处理剩下没有遍历完的数组，直接合并到第三个数组中
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	var ans = make([]int, m+n)
	p1, p2 := 0, 0
	p := 0
	for p1 < m && p2 < n {
		if nums1[p1] < nums2[p2] {
			ans[p] = nums1[p1]
			p1++
		} else {
			ans[p] = nums2[p2]
			p2++
		}
		p++
	}

	if p1 < m {
		for i := p1; i < m; i++ {
			ans[p] = nums1[i]
			p++
		}
	}

	if p2 < n {
		for i := p2; i < n; i++ {
			ans[p] = nums2[i]
			p++
		}
	}

	copy(nums1, ans)

	// 校验结果
	fmt.Println(nums1)
}

func MainMerge() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
}
