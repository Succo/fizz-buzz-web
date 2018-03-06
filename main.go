package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", fizzBuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	string1 := r.FormValue("string1")
	if string1 == "" {
		string1 = "fizz"
	}
	string2 := r.FormValue("string2")
	if string2 == "" {
		string2 = "buzz"
	}
	int1, err := strconv.Atoi(r.FormValue("int1"))
	if err != nil {
		int1 = 3
	}
	int2, err := strconv.Atoi(r.FormValue("int2"))
	if err != nil {
		int2 = 5
	}
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		limit = 100
	}
	fmt.Fprintf(w, "str1: %s, str2 %s, int1 %d, int2 %d, limit %d", html.EscapeString(string1), html.EscapeString(string2), int1, int2, limit)
}
