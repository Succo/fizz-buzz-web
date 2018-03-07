package main

import "strconv"

type output []string

func fizzbuzz(string1, string2 string, int1, int2, limit int) output {
	res := make(output, limit)
	i1 := int1
	i2 := int2
	fizzbuzz := string1 + string2
	for i := 1; i <= limit; i++ {
		i1 = i1 - 1
		i2 = i2 - 1
		switch {
		case i1 == 0 && i2 == 0:
			res[i-1] = fizzbuzz
			i1 = int1
			i2 = int2
		case i1 == 0:
			res[i-1] = string1
			i1 = int1
		case i2 == 0:
			res[i-1] = string2
			i2 = int2
		default:
			res[i-1] = strconv.Itoa(i)
		}
	}
	return res
}
