package main

import "testing"

type testpair struct {
	a     int
	multi int
	sum   int
}

var tests = []testpair{
	{3, 3, 1},
	{9, 3, 1},
	{10, 3, 0},
	{0, 3, 1},
}

func TestMultiples(t *testing.T) {
	for _, pair := range tests {
		v := multiples(pair.a, pair.multi)
		if v != pair.sum {
			t.Error(
				"A", pair.a,
				"Multi", pair.multi,
				"Sum", pair.sum,
				"got", v,
			)
		}
	}
}
