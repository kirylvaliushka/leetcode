package main

import "fmt"

//You are a product manager and currently leading a team to develop a new
//product. Unfortunately, the latest version of your product fails the quality check.
//Since each version is developed based on the previous version, all the versions
//after a bad version are also bad.
//
// Suppose you have n versions [1, 2, ..., n] and you want to find out the
//first bad one, which causes all the following ones to be bad.
//
// You are given an API bool isBadVersion(version) which returns whether
//version is bad. Implement a function to find the first bad version. You should
//minimize the number of calls to the API.
//
//
// Example 1:
//
//
//Input: n = 5, bad = 4
//Output: 4
//Explanation:
//call isBadVersion(3) -> false
//call isBadVersion(5) -> true
//call isBadVersion(4) -> true
//Then 4 is the first bad version.
//
//
// Example 2:
//
//
//Input: n = 1, bad = 1
//Output: 1
//
//
//
// Constraints:
//
//
// 1 <= bad <= n <= 2³¹ - 1
//
// Related Topics Binary Search Interactive 👍 4389 👎 1643

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func main() {
	found := firstBadVersion(5)
	fmt.Println(found)
}

func firstBadVersion(n int) int {
	var pivot int
	left := 1
	right := n - 1
	for left <= right {
		pivot = left + (right-left)/2
		if isBadVersion(pivot) {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	if version == 1 || version == 2 || version == 3 {
		return false
	}
	return true
}
