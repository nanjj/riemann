package ipart

import "testing"

func TestIntegerPartition(t *testing.T) {
	tcs := []struct {
		n     int
		total int
	}{ // from https://proofwiki.org/wiki/Integer_Partition/Table
		{1, 1}, // 1
		{2, 2}, // 2, 1+1
		{3, 3}, // 3, 2+1, 1+1+1
		{4, 5}, // 4, 3+1, 2+2, 2+1+1, 1+1+1+1
		{5, 7}, // 5, 4+1, 3+2, 3+1+1, 2+2+1, 2+1+1+1, 1+1+1+1+1
		{6, 11},
		{7, 15},
		{8, 22},
		{9, 30},
		{10, 42},
		{11, 56},
		{12, 77},
		{13, 101},
		{14, 135},
		{15, 176},
		{16, 231},
		{17, 297},
		{18, 385},
		{19, 490},
		{20, 627},
		{21, 792},
		{22, 1002},
		{23, 1255},
		{24, 1575},
		{25, 1958},
		{26, 2436},
		{27, 3010},
		{28, 3718},
		{29, 4565},
		{30, 5604},
	}
	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			if q := p(tc.n, tc.n); q != tc.total {
				t.Fatal(tc.n, q)
			}
		})
	}
}
