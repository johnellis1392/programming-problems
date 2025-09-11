namespace ProgrammingProblems.LeetCode;

public static class StrStrSolution {
  public static int StrStr(string haystack, string needle) {
    if (needle.Length == 0 || needle.Length > haystack.Length)
      return -1;
    int n = haystack.Length, m = needle.Length;
    for (var i = 0; i < n - m + 1; i++)
      if (haystack.Substring(i, m) == needle)
        return i;
    return -1;
  }
}
