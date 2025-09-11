using ProgrammingProblems.LeetCode;

namespace Tests.LeetCode;

public class LengthOfLastWordTest {
  [Test]
  public void Test1() {
    Assert.That(
      LengthOfLastWordSolution.LengthOfLastWord("Hello World"),
      Is.EqualTo(5)
    );
  }

  [Test]
  public void Test2() {
    Assert.That(
      LengthOfLastWordSolution.LengthOfLastWord("   fly me   to   the moon  "),
      Is.EqualTo(4)
    );
  }

  [Test]
  public void Test3() {
    Assert.That(
      LengthOfLastWordSolution.LengthOfLastWord("luffy is still joyboy"),
      Is.EqualTo(6)
    );
  }
}