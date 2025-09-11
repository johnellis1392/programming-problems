namespace ProgrammingProblems.LeetCode;

public static class PalindromeNumber {
  public static bool IsPalindrome(int x) {
    if (x < 0) return false;
    var s = x.ToString();
    var n = s.Length;
    for (var i = 0; i < n / 2; i++)
      if (s[i] != s[n - i - 1])
        return false;
    return true;
  }
}
