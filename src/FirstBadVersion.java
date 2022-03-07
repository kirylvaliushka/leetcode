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
// Related Topics Binary Search Interactive 👍 4350 👎 1630

/* The isBadVersion API is defined in the parent class VersionControl.
      boolean isBadVersion(int version); */
public class FirstBadVersion {

  public static void main(String[] args) {
    System.out.println(
        new FirstBadVersion().firstBadVersion(5));
  }

  public int firstBadVersion(int n) {
    return firstBadVersion(1, n);
  }

  private int firstBadVersion(int startVersion, int endVersion) {
    if (startVersion == endVersion) {
      if (isBadVersion(startVersion)) {
        return startVersion;
      } else {
        return -1;
      }
    } else {
      int found;
      int step = (endVersion - startVersion) / 2;
      if (step == 0) {
        step = 1;
      }
      if (isBadVersion(startVersion + step)) {
        found = firstBadVersion(startVersion, endVersion - step);
      } else {
        found = firstBadVersion(startVersion + step, endVersion);
      }
      if (found == -1) {
        return startVersion + 1;
      }
      return found;
    }
  }

  private boolean isBadVersion(int version) {
    if (version == 3) {
      return false;
    }
    return true;
  }
}