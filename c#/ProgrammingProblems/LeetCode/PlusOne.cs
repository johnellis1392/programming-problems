namespace ProgrammingProblems.LeetCode;

public static class PlusOneSolution {
  public static int[] PlusOne(int[] digits) {
    var l = new List<int>();
    int c = 1, i = digits.Length - 1;
    while (i >= 0) {
      l.Insert(0, (digits[i] + c) % 10);
      c = (digits[i] + c) / 10;
      i--;
    }

    if (c > 0) l.Insert(0, c);
    return l.ToArray();
  }

  // public static int[] PlusOne2(int[] digits) {
  //   for (var i = digits.Length - 1; i >= 0; i--)
  //     if (digits[i] == 9) {
  //       digits[i] = 0;
  //     }
  //     else {
  //       digits[i]++;
  //       return digits;
  //     }
  //
  //   if (digits[0] > 0) return digits;
  //   var a = new int[digits.Length + 1];
  //   a[0] = 1;
  //   Array.Copy(digits, 0, a, 1, digits.Length);
  //   return a;
  // }
}
