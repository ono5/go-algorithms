package main

import "testing"

func TestSum(t *testing.T) {
	input1 := 5
	input2 := 7
	expected := 11
	res := sum(input1, input2)

	if res != expected {
		t.Errorf(
			"Our sum functions doesn't work, %d+%d isn't %d.\n", input1, input2, res)
	}
}
