using ProgrammingProblems.LeetCode;

namespace Tests.LeetCode;

public class StrStrTest {
  [Test]
  public void Test1() {
    const string haystack = "sadbutsad";
    const string needle = "sad";
    const int expected = 0;
    Assert.That(
      StrStrSolution.StrStr(haystack, needle),
      Is.EqualTo(expected)
    );
  }

  [Test]
  public void Test2() {
    const string haystack = "leetcode";
    const string needle = "leeto";
    const int expected = -1;
    Assert.That(
      StrStrSolution.StrStr(haystack, needle),
      Is.EqualTo(expected)
    );
  }

  [Test]
  public void Test3() {
    const string haystack = "a";
    const string needle = "a";
    const int expected = 0;
    Assert.That(
      StrStrSolution.StrStr(haystack, needle),
      Is.EqualTo(expected)
    );
  }

  [Test]
  public void Test4() {
    const string haystack = "ba";
    const string needle = "a";
    const int expected = 1;
    Assert.That(
      StrStrSolution.StrStr(haystack, needle),
      Is.EqualTo(expected)
    );
  }
}
