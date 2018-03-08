package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
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

	log.Fatal(srv.ListenAndServe())
}

func parseRequest(r *http.Request) (q query) {
	q.str1 = r.FormValue("string1")
	if q.str1 == "" {
		q.str1 = "fizz"
	}
	q.str2 = r.FormValue("string2")
	if q.str2 == "" {
		q.str2 = "buzz"
	}
	var err error
	q.i1, err = strconv.Atoi(r.FormValue("int1"))
	if err != nil {
		q.i1 = 3
	}
	q.i2, err = strconv.Atoi(r.FormValue("int2"))
	if err != nil {
		q.i2 = 5
	}
	q.limit, err = strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		q.limit = 100
	}
	return q
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	q := parseRequest(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := FizzBuzzCountDown(q, w)

	log.Printf("Serving query: %v", q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
