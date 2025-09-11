namespace ProgrammingProblems.LeetCode;

public static class ValidParentheses {
  public static bool IsValid(string s) {
    var stack = new Stack<char>();
    foreach (var c in s)
      if (c == '(' || c == '[' || c == '{')
        stack.Push(c);
      else if (stack.Count == 0)
        return false;
      else if (
        (c == ')' && stack.Peek() == '(') ||
        (c == ']' && stack.Peek() == '[') ||
        (c == '}' && stack.Peek() == '{')
      )
        stack.Pop();
      else
        return false;

    return stack.Count == 0;
  }
}
