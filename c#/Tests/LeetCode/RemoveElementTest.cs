using ProgrammingProblems.LeetCode;

namespace Tests.LeetCode;

public class RemoveElementTest {
  [Test]
  public void Test1() {
    int[] input = [3, 2, 2, 3];
    const int val = 3;
    const int k = 2;
    int[] expected = [2, 2];
    Assert.Multiple(() => {
      Assert.That(
        RemoveElementClass.RemoveElement(input, val),
        Is.EqualTo(k)
      );
      Assert.That(
        input[..k],
        Is.EqualTo(expected)
      );
    });
  }

  [Test]
  public void Test2() {
    int[] input = [0, 1, 2, 2, 3, 0, 4, 2];
    const int val = 2;
    const int k = 5;
    int[] expected = [0, 1, 4, 0, 3];
    Assert.Multiple(() => {
      Assert.That(
        RemoveElementClass.RemoveElement(input, val),
        Is.EqualTo(k)
      );
      Assert.That(
        input[..k],
        Is.EqualTo(expected)
      );
    });
  }
}