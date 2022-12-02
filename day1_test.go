package main

import "testing"

func TestFindMax(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	max, _ := findMax(slice)
	if max != 4 {
		t.Error("Expected max to equal 4")
	}
}

func TestTableFindMax(t *testing.T) {
	var tests = []struct {
		input []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{-5, 0, 3, 4}, 4},
	}

	for _, test := range tests {
		if output, _ := findMax(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}

}