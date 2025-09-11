namespace ProgrammingProblems.LeetCode;

public static class LengthOfLastWordSolution {
  public static int LengthOfLastWord(string s) {
    var v = s.TrimEnd().Split(' ', StringSplitOptions.RemoveEmptyEntries);
    return v.Length > 0 ? v[^1].Length : 0;
  }
}