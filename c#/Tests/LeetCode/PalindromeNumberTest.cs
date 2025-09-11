using ProgrammingProblems.LeetCode;

namespace Tests.LeetCode;

public class PalindromeNumberTest {
  [Test]
  public void Test1() {
    Assert.That(PalindromeNumber.IsPalindrome(5), Is.EqualTo(true));
  }

  [Test]
  public void Test2() {
    Assert.That(PalindromeNumber.IsPalindrome(121), Is.EqualTo(true));
  }

  [Test]
  public void Test3() {
    Assert.That(PalindromeNumber.IsPalindrome(-121), Is.EqualTo(false));
  }

  [Test]
  public void Test4() {
    Assert.That(PalindromeNumber.IsPalindrome(10), Is.EqualTo(false));
  }
}
