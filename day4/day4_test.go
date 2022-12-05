package main

import "testing"


func TestCheckContainment(t *testing.T) {
	var tests = []struct {
		a int
		b int
		c int
		d int
		expected int
	}{
		{20, 67, 3, 21, 0}, // spot check for alphabetic comparators "20" < "3"
		{1, 2, 3, 4, 0}, // no overlap
		{3, 4, 1, 2, 0}, // no overlap
		{1, 2, 4, 5, 0}, // no overlap, gap
		{4, 5, 1, 2, 0}, // no overlap, gap
		{1, 2, 2, 4, 0}, // partial overlap
		{2, 4, 1, 2, 0}, // partial overlap
		{1, 1, 1, 2, 1}, // single area, contained
		{1, 2, 1, 1, 1}, // single area, contained
		{1, 2, 1, 2, 1}, // 100% overlap
		{1, 3, 1, 2, 1}, // Contained, shared first
		{1, 2, 1, 3, 1}, // Contained, shared first
		{2, 3, 1, 3, 1}, // Contained, shared last
		{1, 3, 2, 3, 1}, // Contained, shared last
		{1, 4, 2, 3, 1}, // Contained, fully
		{2, 3, 1, 4, 1}, // Contained, fully
		{2, 4, 6, 8, 0}, // supplied check
		{2, 3, 4, 5, 0}, // supplied check
		{5, 7, 7, 9, 0}, // supplied check
		{2, 8, 3, 7, 1}, // supplied check
		{6, 6, 4, 6, 1}, // supplied check
		{2, 6, 4, 8, 0}, // supplied check
	}

	for _, test := range tests {
		if output := checkContainment(test.a, test.b, test.c, test.d); output != test.expected {
			t.Errorf("Test Failed: %v-%v,%v-%v inputted, %v expected, received: %v", test.a, test.b, test.c, test.d, test.expected, output)
		}
	}

}
