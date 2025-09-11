using ProgrammingProblems.LeetCode;

namespace Tests.LeetCode;

public class ValidParenthesesTest {
  [Test]
  public void Test1() {
    const string input = "()";
    const bool expected = true;
    Assert.That(ValidParentheses.IsValid(input), Is.EqualTo(expected));
  }

  [Test]
  public void Test2() {
    const string input = "()[]{}";
    const bool expected = true;
    Assert.That(ValidParentheses.IsValid(input), Is.EqualTo(expected));
  }

  [Test]
  public void Test3() {
    const string input = "(]";
    const bool expected = false;
    Assert.That(ValidParentheses.IsValid(input), Is.EqualTo(expected));
  }

  [Test]
  public void Test4() {
    const string input = "]";
    const bool expected = false;
    Assert.That(ValidParentheses.IsValid(input), Is.EqualTo(expected));
  }
}
