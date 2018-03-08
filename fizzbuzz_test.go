package main

import (
	"bytes"
	"testing"
)

var testOutput = "[1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz]"

var testQuery = query{
	str1:  "fizz",
	str2:  "buzz",
	i1:    3,
	i2:    5,
	limit: 15,
}

func TestFizzBuzz(t *testing.T) {
	var b bytes.Buffer
	fizzbuzz(testQuery, &b)
	out := b.String()
	if out != testOutput {
		t.Errorf("Invalid output was\n %s\n expected\n %s", out, testOutput)
		t.Fail()
	}
}
