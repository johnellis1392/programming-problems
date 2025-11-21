package dynamicprogramming

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}
	for r := 2; r <= numRows; r++ {
		row := make([]int, r)
		res[r-1] = row
		row[0] = 1
		row[r-1] = 1
		prev := res[r-2]
		for c := 1; c < r-1; c++ {
			row[c] = prev[c-1] + prev[c]
		}
	}
	return res
}
