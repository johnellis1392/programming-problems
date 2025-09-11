namespace ProgrammingProblems.LeetCode;

public static class TwoSumSolution {
  public static int[] TwoSum(int[] nums, int target) {
    var m = new Dictionary<int, int>();
    for (var i = 0; i < nums.Length; i++) {
      if (m.TryGetValue(nums[i], out var v)) return new[] { v, i };

      m[target - nums[i]] = i;
    }

    return [];
  }
}
