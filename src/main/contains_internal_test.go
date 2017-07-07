package main

import "testing"

var tests = []struct {
	s string
	res bool
}{
	{"google.com", false},
	{"http://google.com", true},
	{"https://google.com", true},
}

func TestContains(t *testing.T) {
	for _, test := range tests {
		res := contains(test.s, Schemes)
		if res != test.res {
			t.Error("For string", test.s, "the result should be", test.res, "but got", res, "instead")
		}
	}
}