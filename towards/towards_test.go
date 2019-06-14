package towards

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestTowardsOne(t *testing.T) {
	tcs := []struct {
		n    int
		want []int
	}{
		{1, []int{1}},
		{2, []int{2, 1}},
		{3, []int{3, 10, 5, 16, 8, 4, 2, 1}},
		{4, []int{4, 2, 1}},
		{5, []int{5, 16, 8, 4, 2, 1}},
		{6, []int{6, 3, 10, 5, 16, 8, 4, 2, 1}},
		{7, []int{7, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			trace := Towards1(tc.n)
			if !reflect.DeepEqual(tc.want, trace) {
				t.Log(trace)
				t.Log(tc.want)
				t.Fatal()
			}
		})
	}
}

func BenchmarkTowardsOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Int() + 1
		Towards1(n)
	}
}
