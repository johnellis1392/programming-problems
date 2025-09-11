using ProgrammingProblems.LeetCode;

namespace Tests.LeetCode;

public class PlusOneTest {
  [Test]
  public void Test1() {
    int[] expected = [1, 2, 4];
    Assert.That(
      PlusOneSolution.PlusOne([1, 2, 3]),
      Is.EqualTo(expected)
    );
  }

  [Test]
  public void Test2() {
    int[] expected = [4, 3, 2, 2];
    Assert.That(
      PlusOneSolution.PlusOne([4, 3, 2, 1]),
      Is.EqualTo(expected)
    );
  }

  [Test]
  public void Test3() {
    int[] expected = [1, 0];
    Assert.That(
      PlusOneSolution.PlusOne([9]),
      Is.EqualTo(expected)
    );
  }
}
