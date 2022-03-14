package main

import "fmt"

//Given an integer array nums sorted in non-decreasing order, return an array
//of the squares of each number sorted in non-decreasing order.
//
//
// Example 1:
//
//
//Input: nums = [-4,-1,0,3,10]
//Output: [0,1,9,16,100]
//Explanation: After squaring, the array becomes [16,1,0,9,100].
//After sorting, it becomes [0,1,9,16,100].
//
//
// Example 2:
//
//
//Input: nums = [-7,-3,2,3,11]
//Output: [4,9,9,49,121]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums is sorted in non-decreasing order.
//
//
//
//Follow up: Squaring each element and sorting the new array is very trivial,
//could you find an O(n) solution using a different approach? Related Topics Array
//Two Pointers Sorting 👍 4841 👎 147

func main() {
	orig := []int{-5, -3, -2, -1}
	result := sortedSquares(orig)
	fmt.Println(result)
}

func sortedSquares(nums []int) []int {
	origLength := len(nums)
	result := make([]int, origLength)
	left := 0
	right := origLength - 1

	for i := origLength - 1; i >= 0; i-- {
		if abs(nums[left]) > abs(nums[right]) {
			result[i] = nums[left] * nums[left]
			left++
		} else {
			result[i] = nums[right] * nums[right]
			right--
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
