using ProgrammingProblems.LeetCode;

namespace Tests.LeetCode;

public class TwoSumTest {
  [Test]
  public void Test1() {
    int[] input = [2, 7, 11, 15];
    const int target = 9;
    int[] expected = [0, 1];
    Assert.That(
      TwoSumSolution.TwoSum(input, target),
      Is.EqualTo(expected)
    );
  }

  [Test]
  public void Test2() {
    int[] input = [3, 2, 4];
    const int target = 6;
    int[] expected = [1, 2];
    Assert.That(
      TwoSumSolution.TwoSum(input, target),
      Is.EqualTo(expected)
    );
  }

  [Test]
  public void Test3() {
    int[] input = [3, 3];
    const int target = 6;
    int[] expected = [0, 1];
    Assert.That(
      TwoSumSolution.TwoSum(input, target),
      Is.EqualTo(expected)
    );
  }
}
