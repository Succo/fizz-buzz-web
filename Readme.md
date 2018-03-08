# Exercise: Write a simple fizz-buzz REST server.
 
The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by “fizz”, all multiples of 5 by “buzz”, and all multiples of 15 by “fizzbuzz”. The output would look like this:

“1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,fizz,...”
 
More specifically :

Expose a REST API endpoint that accepts five parameters : two strings (say, string1 and string2), and three integers (say, int1, int2 and limit), and returns a JSON

 - It must return a list of strings with numbers from 1 to limit, where:
 - All multiples of int1 are replaced by string1,
 - All multiples of int2 are replaced by string2,
 - All multiples of int1 and int2 are replaced by string1string2

# Note

I've added multiples fizzbuzz implementations to benchmark performances.
Benchmark can be ran with 

```
go test -bench=.
```

On my computer the results are
```
oos: linux
goarch: amd64
pkg: github.com/Succo/fizz-buzz-web
BenchmarkAllFizzBuzz/naive-4         	  500000	      2889 ns/op
BenchmarkAllFizzBuzz/count_down-4    	  500000	      2433 ns/op
BenchmarkAllFizzBuzz/memorize_modulo-4         	  500000	      2780 ns/op
BenchmarkAllFizzBuzz/updated_var-4             	  500000	      2849 ns/op
PASS
ok  	github.com/Succo/fizz-buzz-web	5.613s
```
