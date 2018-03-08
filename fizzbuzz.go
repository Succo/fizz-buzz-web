package main

import (
	"io"
	"strconv"
)

type query struct {
	str1  string
	str2  string
	i1    int
	i2    int
	limit int
}

type output []string

type FizzBuzzFunction func(query, io.Writer) error

func FizzBuzzNaive(q query, w io.Writer) error {
	fizz := []byte(q.str1)
	buzz := []byte(q.str2)
	fizzbuzz := []byte(q.str1 + q.str2)
	w.Write([]byte("["))
	for i := 1; i <= q.limit; i++ {
		if i > 1 {
			w.Write([]byte(", "))
		}
		switch {
		case i%q.i1 == 0 && i%q.i2 == 0:
			_, err := w.Write(fizzbuzz)
			if err != nil {
				return err
			}
		case i%q.i1 == 0:
			_, err := w.Write(fizz)
			if err != nil {
				return err
			}
		case i%q.i2 == 0:
			_, err := w.Write(buzz)
			if err != nil {
				return err
			}
		default:
			_, err := w.Write([]byte(strconv.Itoa(i)))
			if err != nil {
				return err
			}
		}
	}
	_, err := w.Write([]byte("]"))
	return err
}

func FizzBuzzMemorizeModulo(q query, w io.Writer) error {
	fizz := []byte(q.str1)
	buzz := []byte(q.str2)
	w.Write([]byte("["))
	for i := 1; i <= q.limit; i++ {
		if i > 1 {
			w.Write([]byte(", "))
		}
		i1 := i % q.i1
		i2 := i % q.i2
		if i1 == 0 {
			_, err := w.Write(fizz)
			if err != nil {
				return err
			}
		}
		if i2 == 0 {
			_, err := w.Write(buzz)
			if err != nil {
				return err
			}
		}
		if i1 != 0 && i2 != 0 {
			_, err := w.Write([]byte(strconv.Itoa(i)))
			if err != nil {
				return err
			}
		}
	}
	_, err := w.Write([]byte("]"))
	return err
}

func FizzBuzzUpdatedVar(q query, w io.Writer) error {
	fizz := []byte(q.str1)
	buzz := []byte(q.str2)
	w.Write([]byte("["))
	for i := 1; i <= q.limit; i++ {
		if i > 1 {
			w.Write([]byte(", "))
		}
		matched := false
		if i%q.i1 == 0 {
			_, err := w.Write(fizz)
			matched = true
			if err != nil {
				return err
			}
		}
		if i%q.i2 == 0 {
			_, err := w.Write(buzz)
			matched = true
			if err != nil {
				return err
			}
		}
		if !matched {
			_, err := w.Write([]byte(strconv.Itoa(i)))
			if err != nil {
				return err
			}
		}
	}
	_, err := w.Write([]byte("]"))
	return err
}

func FizzBuzzCountDown(q query, w io.Writer) error {
	i1 := q.i1
	i2 := q.i2
	fizz := []byte(q.str1)
	buzz := []byte(q.str2)
	fizzbuzz := []byte(q.str1 + q.str2)
	w.Write([]byte("["))
	for i := 1; i <= q.limit; i++ {
		if i > 1 {
			w.Write([]byte(", "))
		}
		i1 = i1 - 1
		i2 = i2 - 1
		switch {
		case i1 == 0 && i2 == 0:
			_, err := w.Write(fizzbuzz)
			if err != nil {
				return err
			}
			i1 = q.i1
			i2 = q.i2
		case i1 == 0:
			_, err := w.Write(fizz)
			if err != nil {
				return err
			}
			i1 = q.i1
		case i2 == 0:
			_, err := w.Write(buzz)
			if err != nil {
				return err
			}
			i2 = q.i2
		default:
			_, err := w.Write([]byte(strconv.Itoa(i)))
			if err != nil {
				return err
			}
		}
	}
	_, err := w.Write([]byte("]"))
	return err
}
