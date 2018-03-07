package main

import "testing"

var testOutput = output{
	"1",
	"2",
	"fizz",
	"4",
	"buzz",
	"fizz",
	"7",
	"8",
	"fizz",
	"buzz",
	"11",
	"fizz",
	"13",
	"14",
	"fizzbuzz",
}

func TestFizzBuzz(t *testing.T) {
	out := fizzbuzz("fizz", "buzz", 3, 5, 15)
	if len(out) != len(testOutput) {
		t.Errorf("Invalid length for output was %d expected %d", len(out), len(testOutput))
		t.Fail()
	}
	for i, v := range testOutput {
		if v != out[i] {
			t.Errorf("Invalid value at index %d for output was %s expected %s", i, out[i], v)
			t.Fail()
		}
	}
}
