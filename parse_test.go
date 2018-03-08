package main

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	var b bytes.Buffer

	// Test loading default values
	url := "0.0.0.0/fizzbuzz/"
	r, err := http.NewRequest("GET", url, &b)
	if err != nil {
		t.Fatalf("Invalid request creation for %s\n", url)
	}
	q, err := parseRequest(r)
	if err != nil || !reflect.DeepEqual(q, query{str1: "fizz", str2: "buzz", i1: 3, i2: 5, limit: 100}) {
		t.Fatalf("Invalid query for %s\n", url)
	}

	// Test loading values from the url
	url = "0.0.0.0/fizzbuzz?int1=1&int2=2&limit=3&string1=test1&string2=test2"
	r, err = http.NewRequest("GET", url, &b)
	if err != nil {
		t.Fatalf("Invalid request creation for %s\n", url)
	}
	q, err = parseRequest(r)
	if err != nil || !reflect.DeepEqual(q, query{str1: "test1", str2: "test2", i1: 1, i2: 2, limit: 3}) {
		t.Fatalf("Invalid query for %s\n", url)
	}

	// Test error in case of invalid value
	url = "0.0.0.0/fizzbuzz/?int1=d"
	r, err = http.NewRequest("GET", url, &b)
	if err != nil {
		t.Fatalf("Invalid request creation for %s\n", url)
	}
	q, err = parseRequest(r)
	if err == nil {
		t.Fatalf("Did not fail parsing i1 invalid number for %s\n", url)
	}

	url = "0.0.0.0/fizzbuzz/?int2=d"
	r, err = http.NewRequest("GET", url, &b)
	if err != nil {
		t.Fatalf("Invalid request creation for %s\n", url)
	}
	q, err = parseRequest(r)
	if err == nil {
		t.Fatalf("Did not fail parsing i2 invalid number for %s\n", url)
	}
}
