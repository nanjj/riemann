package ipart

func p(n, m int) int {
	if m > n {
		m = n
	}
	if n == 1 || m == 1 {
		return 1
	} else if m >= n-1 {
		return 1 + p(n, m-1)
	} else { // m < n
		return p(n-m, m) + p(n, m-1)
	}
}
