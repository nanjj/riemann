package riemann

import (
	"context"
	"math"
	"math/rand"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestCoprime(t *testing.T) {
	tcs := []struct {
		m  uint64
		n  uint64
		cp bool
	}{
		{1, 2, true},
		{1, 100, true},
		{2, 100, false},
		{6, 27, false},
	}
	for _, tc := range tcs {
		cp := Coprime(tc.m, tc.n)
		if cp != tc.cp {
			t.Fatal(tc, cp)
		}
	}
}

func TestCoprimeProbability(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	max := 1000 * math.MaxInt16
	const num = 2
	var cp [num]int
	g, _ := errgroup.WithContext(context.Background())
	for l := 0; l < num; l++ {
		l := l
		g.Go(func() error {
			for i := 0; i < max/num; i++ {
				m := rand.Uint64()
				n := rand.Uint64()
				if Coprime(m, n) {
					cp[l] = cp[l] + 1
				}
			}
			return nil
		})
	}
	g.Wait()
	hit := 0
	for n := 0; n < num; n++ {
		hit += cp[n]
	}
	t.Log(float64(hit) / float64(max))
}
