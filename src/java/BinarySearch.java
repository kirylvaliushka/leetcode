//Given an array of integers nums which is sorted in ascending order, and an
//integer target, write a function to search target in nums. If target exists, then
//return its index. Otherwise, return -1.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
//
//
//Input: nums = [-1,0,3,5,9,12], target = 9
//Output: 4
//Explanation: 9 exists in nums and its index is 4
//
//
// Example 2:
//
//
//Input: nums = [-1,0,3,5,9,12], target = 2
//Output: -1
//Explanation: 2 does not exist in nums so return -1
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ < nums[i], target < 10⁴
// All the integers in nums are unique.
// nums is sorted in ascending order.
//
// Related Topics Array Binary Search 👍 3659 👎 92

import java.util.Arrays;

public class BinarySearch {

  public static void main(String[] args) {
    System.out.println(new BinarySearch().search(new int[]{-1,0,3,5,9,12}, 2));
  }

  public int search(int[] nums, int target) {
    if (nums.length == 0) {
      return -1;
    } else if (nums.length == 1 && nums[0] != target) {
      return -1;
    } else {
      int index = nums.length / 2;
      int value = nums[index];
      if (value == target) {
        return index;
      } else if (value > target) {
        int[] leftSubNums = Arrays.copyOfRange(nums, 0, index);
        return search(leftSubNums, target);
      } else {
        int[] rightSubNums = Arrays.copyOfRange(nums, index +1, nums.length);
        int foundIndexOfRightSubNums = search(rightSubNums, target);
        if(foundIndexOfRightSubNums == -1){
          return -1;
        } else {
          return nums.length - rightSubNums.length + foundIndexOfRightSubNums;
        }
      }
    }
  }

}
