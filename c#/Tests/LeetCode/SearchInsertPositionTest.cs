using ProgrammingProblems.LeetCode;

namespace Tests.LeetCode;

public class SearchInsertPositionTest {
  [Test]
  public void Test1() {
    int[] input = [1, 3, 5, 6];
    const int target = 5;
    const int expected = 2;
    Assert.That(
      SearchInsertPositionSolution.SearchInsert(input, target),
      Is.EqualTo(expected)
    );
  }

  [Test]
  public void Test2() {
    int[] input = [1, 3, 5, 6];
    const int target = 7;
    const int expected = 4;
    Assert.That(
      SearchInsertPositionSolution.SearchInsert(input, target),
      Is.EqualTo(expected)
    );
  }
}