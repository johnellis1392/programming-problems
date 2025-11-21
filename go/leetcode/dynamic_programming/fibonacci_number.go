package dynamicprogramming

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	a := 0
	b := 1
	for i := 2; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}
