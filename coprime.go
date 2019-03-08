package riemann

func Coprime(m, n uint64) bool {
	for {
		if m < n {
			m, n = n, m
		}
		if n == 0 {
			if m == 1 {
				return true
			} else {
				return false
			}
		}
		m %= n
	}
}
