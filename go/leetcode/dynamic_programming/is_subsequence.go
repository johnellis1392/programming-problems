package dynamicprogramming

func isSubsequence(s, t string) bool {
	si := 0
	ti := 0
	for si < len(s) && ti < len(t) {
		if s[si] == t[ti] {
			si++
			ti++
		} else {
			ti++
		}
	}
	return si == len(s)
}
