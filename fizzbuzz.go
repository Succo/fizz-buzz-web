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

// FizzBuzzFunction writes a json encoded list of string that correspond to a fizz buzz list
type FizzBuzzFunction func(query, io.Writer) error

// FizzBuzzNaive generates values with a switch case and all modulo in an explicit way
func FizzBuzzNaive(q query, w io.Writer) error {
	fizz := []byte(q.str1)
	buzz := []byte(q.str2)
	fizzbuzz := []byte(q.str1 + q.str2)
	_, err := w.Write([]byte("["))
	if err != nil {
		return err
	}
	for i := 1; i <= q.limit; i++ {
		if i > 1 {
			_, err := w.Write([]byte(", "))
			if err != nil {
				return err
			}
		}
		_, err := w.Write([]byte("\""))
		if err != nil {
			return err
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
		_, err = w.Write([]byte("\""))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("]"))
	return err
}

// FizzBuzzMemorizeModulo store modulo values and avoid an explicit check for fizzbuzz cases
func FizzBuzzMemorizeModulo(q query, w io.Writer) error {
	fizz := []byte(q.str1)
	buzz := []byte(q.str2)
	_, err := w.Write([]byte("["))
	if err != nil {
		return err
	}
	for i := 1; i <= q.limit; i++ {
		if i > 1 {
			_, err := w.Write([]byte(", "))
			if err != nil {
				return err
			}
		}
		_, err := w.Write([]byte("\""))
		if err != nil {
			return err
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
		_, err = w.Write([]byte("\""))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("]"))
	return err
}

// FizzBuzzUpdatedVar avoid having an explicit check for fizzbuzz case by noting fizz and buzz cases
func FizzBuzzUpdatedVar(q query, w io.Writer) error {
	fizz := []byte(q.str1)
	buzz := []byte(q.str2)
	_, err := w.Write([]byte("["))
	if err != nil {
		return err
	}
	for i := 1; i <= q.limit; i++ {
		if i > 1 {
			_, err := w.Write([]byte(", "))
			if err != nil {
				return err
			}
		}
		_, err := w.Write([]byte("\""))
		if err != nil {
			return err
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
		_, err = w.Write([]byte("\""))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("]"))
	return err
}

// FizzBuzzCountDown avoid modulo operation by keeping counters that match fizz and buzz integer
func FizzBuzzCountDown(q query, w io.Writer) error {
	i1 := q.i1
	i2 := q.i2
	fizz := []byte(q.str1)
	buzz := []byte(q.str2)
	fizzbuzz := []byte(q.str1 + q.str2)
	_, err := w.Write([]byte("["))
	if err != nil {
		return err
	}
	for i := 1; i <= q.limit; i++ {
		if i > 1 {
			_, err := w.Write([]byte(", "))
			if err != nil {
				return err
			}
		}
		_, err := w.Write([]byte("\""))
		if err != nil {
			return err
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
		_, err = w.Write([]byte("\""))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("]"))
	return err
}
