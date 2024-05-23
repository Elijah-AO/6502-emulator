package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(findKthLargest(nums, 4))
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 1. 先排序
// 2. 返回倒数第k个元素
// 3. 时间复杂度O(nlogn) 空间复杂度O(1)

// 作者：carlsun-2
// 链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/215-shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-carlsu/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, k int) int {
	pivot := partition(nums, left, right)
	if pivot == k {
		return nums[pivot]
	} else if pivot < k {
		return quickSelect(nums, pivot+1, right, k)
	} else {
		return quickSelect(nums, left, pivot-1, k)
	}
}
