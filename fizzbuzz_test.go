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

func testFizzBuzz(t *testing.T, f FizzBuzzFunction, name string) {
	var b bytes.Buffer
	err := f(testQuery, &b)
	if err != nil {
		t.Error(err)
	}
	out := b.String()
	if out != testOutput {
		t.Errorf("Invalid output for %s was\n %s\n expected\n %s", name, out, testOutput)
		t.Fail()
	}
}

func TestAllFizzBuzz(t *testing.T) {
	testFizzBuzz(t, FizzBuzzNaive, "naive")
	testFizzBuzz(t, FizzBuzzCountDown, "count down")
	testFizzBuzz(t, FizzBuzzMemorizeModulo, "memorize modulo")
	testFizzBuzz(t, FizzBuzzUpdatedVar, "memorize modulo")
}
