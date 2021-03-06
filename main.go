package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	str1  = "fizz"
	str2  = "buzz"
	i1    = 3
	i2    = 5
	limit = 100
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/fizzbuzz", fizzBuzzHandler)
	mux.HandleFunc("/fizzbuzz/", fizzBuzzHandler)

	c := loadConfig()

	srv := &http.Server{
		Addr:         c.addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}

	log.Printf("Starting to serve on %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

// parseRequest reads parameters from a request to create a valid query
// It replaces them with default values when missing but fails on invalid values
func parseRequest(r *http.Request) (q query, err error) {
	q.str1 = r.FormValue("string1")
	if q.str1 == "" {
		q.str1 = str1
	}
	q.str2 = r.FormValue("string2")
	if q.str2 == "" {
		q.str2 = str2
	}

	int1 := r.FormValue("int1")
	if int1 == "" {
		q.i1 = i1
	} else {
		q.i1, err = strconv.Atoi(int1)
		if err != nil {
			return q, fmt.Errorf("Unable to parse int1 value %s not an integer", int1)
		}
	}

	int2 := r.FormValue("int2")
	if int2 == "" {
		q.i2 = i2
	} else {
		q.i2, err = strconv.Atoi(r.FormValue("int2"))
		if err != nil {
			return q, fmt.Errorf("Unable to parse int2 value %s not an integer", int2)
		}
	}

	l := r.FormValue("limit")
	if l == "" {
		q.limit = limit
	} else {
		q.limit, err = strconv.Atoi(r.FormValue("limit"))
		if err != nil {
			return q, fmt.Errorf("Unable to parse limit value %s not an integer", l)
		}
	}
	return q, nil
}

// fizzBuzzHandler is the handler for all request to the service
func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	q, err := parseRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	log.Printf("Serving query: %v", q)
	err = FizzBuzzCountDown(q, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
