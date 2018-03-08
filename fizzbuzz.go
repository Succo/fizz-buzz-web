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

func fizzbuzz(q query, w io.Writer) error {
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
				return nil
			}
			i1 = q.i1
			i2 = q.i2
		case i1 == 0:
			w.Write(fizz)
			i1 = q.i1
		case i2 == 0:
			w.Write(buzz)
			i2 = q.i2
		default:
			w.Write([]byte(strconv.Itoa(i)))
		}
	}
	w.Write([]byte("]"))
	return nil
}
