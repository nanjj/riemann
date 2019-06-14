package towards

var (
	cache = map[int][]int{}
)

func Towards1(n int) (a []int) {
	for n > 1 {
		a = append(a, n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}
	a = append(a, n)
	return
}
