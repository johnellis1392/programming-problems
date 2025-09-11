namespace ProgrammingProblems.LeetCode;

public static class RemoveElementClass {
  // public static int RemoveElement2(int[] nums, int val) {
  //   int k = nums.Length;
  //   int i = 0;
  //   while (i < k) {
  //     if (nums[i] == val) {
  //       int temp = nums[i];
  //       for (int j = i; j < k-1; j++)
  //         nums[j] = nums[j+1];
  //       nums[k-1] = temp;
  //       k--;
  //     } else {
  //       i++;
  //     }
  //   }
  //   return k;
  // }

  public static int RemoveElement(int[] nums, int val) {
    var k = nums.Length;
    var i = 0;
    while (i < k)
      if (nums[i] == val) {
        (nums[i], nums[k - 1]) = (nums[k - 1], nums[i]);
        k--;
      }
      else {
        i++;
      }

    return k;
  }
}