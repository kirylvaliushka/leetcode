package main

import (
	"fmt"
)

//Given a sorted array of distinct integers and a target value, return the
//index if the target is found. If not, return the index where it would be if it were
//inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
//
//
//Input: nums = [1,3,5,6], target = 5
//Output: 2
//
//
// Example 2:
//
//
//Input: nums = [1,3,5,6], target = 2
//Output: 1
//
//
// Example 3:
//
//
//Input: nums = [1,3,5,6], target = 7
//Output: 4
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums contains distinct values sorted in ascending order.
// -10⁴ <= target <= 10⁴
//
// Related Topics Array Binary Search 👍 7021 👎 386

func main() {
	nums := []int{1, 3, 5, 6}
	found := searchInsert(nums, 2)
	fmt.Println(found)
}

func searchInsert(nums []int, target int) int {
	var pivot int
	left := 0
	right := len(nums) - 1
	for left <= right {
		pivot = left + (right-left)/2
		if nums[pivot] == target {
			return pivot
		} else if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return left
}
