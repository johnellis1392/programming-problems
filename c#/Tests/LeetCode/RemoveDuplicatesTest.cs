using ProgrammingProblems.LeetCode;

namespace Tests.LeetCode;

public class RemoveDuplicatesTest {
  [Test]
  public void Test1() {
    int[] input = [1, 2, 3];
    int[] expected = [1, 2, 3];
    const int k = 3;
    var actual = RemoveDuplicatesSolution.RemoveDuplicates(input);
    Assert.Multiple(() => {
      Assert.That(actual, Is.EqualTo(k));
      Assert.That(input[..k], Is.EqualTo(expected));
    });
  }

  [Test]
  public void Test2() {
    int[] input = [1, 1, 2];
    int[] expected = [1, 2];
    const int k = 2;
    var actual = RemoveDuplicatesSolution.RemoveDuplicates(input);
    Assert.Multiple(() => {
      Assert.That(actual, Is.EqualTo(k));
      Assert.That(input[..k], Is.EqualTo(expected));
    });
  }

  [Test]
  public void Test3() {
    int[] input = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
    int[] expected = [0, 1, 2, 3, 4];
    const int k = 5;
    var actual = RemoveDuplicatesSolution.RemoveDuplicates(input);
    Assert.Multiple(() => {
      Assert.That(actual, Is.EqualTo(k));
      Assert.That(input[..k], Is.EqualTo(expected));
    });
  }
}
