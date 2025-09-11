namespace ProgrammingProblems.LeetCode;

public static class RemoveDuplicatesSolution {
  public static int RemoveDuplicates(int[] nums) {
    var k = nums.Length;
    var i = 0;
    while (i < k - 1)
      if (nums[i] == nums[i + 1]) {
        var temp = nums[i];
        for (var j = i; j < k - 1; j++) nums[j] = nums[j + 1];
        nums[k - 1] = temp;
        k--;
      }
      else {
        i++;
      }

    return k;
  }
}
