namespace ProgrammingProblems.LeetCode;

public static class SearchInsertPositionSolution {
  public static int SearchInsert(int[] nums, int target) {
    int i = 0, j = nums.Length;
    while (i < j) {
      var mid = (i + j) / 2;
      if (nums[mid] == target) return mid;

      if (nums[mid] < target)
        i = mid + 1;
      else
        j = mid;
    }

    return i;
  }
}