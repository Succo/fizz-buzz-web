package main

import (
	"bytes"
	"testing"
)

var testOutput = `["1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"]`

var testQuery = query{
	str1:  "fizz",
	str2:  "buzz",
	i1:    3,
	i2:    5,
	limit: 15,
}

var testQueries = map[string]query{
	"i1:3 i2:5 limit:15":        testQuery,
	"i1:10 i2:10 limit:100":     query{str1: "fizz", str2: "buzz", i1: 10, i2: 10, limit: 100},
	"i1:1000 i2:1000 limit:500": query{str1: "fizz", str2: "buzz", i1: 100, i2: 100, limit: 500},
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

func getFizzBuzzBenchmark(f FizzBuzzFunction, q query) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			err := f(q, &buf)
			if err != nil {
				b.Log(err)
			}
		}
	}
}

func TestAllFizzBuzz(t *testing.T) {
	testFizzBuzz(t, FizzBuzzNaive, "naive")
	testFizzBuzz(t, FizzBuzzCountDown, "count down")
	testFizzBuzz(t, FizzBuzzMemorizeModulo, "memorize modulo")
	testFizzBuzz(t, FizzBuzzUpdatedVar, "updated var")
}

func BenchmarkAllFizzBuzz(b *testing.B) {
	for name, q := range testQueries {
		b.Run("naive "+name, getFizzBuzzBenchmark(FizzBuzzNaive, q))
		b.Run("count down "+name, getFizzBuzzBenchmark(FizzBuzzCountDown, q))
		b.Run("memorize modulo "+name, getFizzBuzzBenchmark(FizzBuzzMemorizeModulo, q))
		b.Run("updated var "+name, getFizzBuzzBenchmark(FizzBuzzUpdatedVar, q))
	}
}
