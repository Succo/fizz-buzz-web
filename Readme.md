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
goos: linux
goarch: amd64
pkg: github.com/Succo/fizz-buzz-web
BenchmarkAllFizzBuzz/naive_i1:3_i2:5_limit:15-4         	  300000	      4848 ns/op
BenchmarkAllFizzBuzz/count_down_i1:3_i2:5_limit:15-4    	  300000	      4358 ns/op
BenchmarkAllFizzBuzz/memorize_modulo_i1:3_i2:5_limit:15-4         	  300000	      4913 ns/op
BenchmarkAllFizzBuzz/updated_var_i1:3_i2:5_limit:15-4             	  300000	      4871 ns/op
BenchmarkAllFizzBuzz/naive_i1:10_i2:3_limit:100-4                 	   50000	     32278 ns/op
BenchmarkAllFizzBuzz/count_down_i1:10_i2:3_limit:100-4            	   50000	     28137 ns/op
BenchmarkAllFizzBuzz/memorize_modulo_i1:10_i2:3_limit:100-4       	   50000	     33056 ns/op
BenchmarkAllFizzBuzz/updated_var_i1:10_i2:3_limit:100-4           	   50000	     32756 ns/op
BenchmarkAllFizzBuzz/naive_i1:1000_i2:1000_limit:500-4            	   10000	    183031 ns/op
BenchmarkAllFizzBuzz/count_down_i1:1000_i2:1000_limit:500-4       	   10000	    178977 ns/op
BenchmarkAllFizzBuzz/memorize_modulo_i1:1000_i2:1000_limit:500-4  	   10000	    186772 ns/op
BenchmarkAllFizzBuzz/updated_var_i1:1000_i2:1000_limit:500-4      	   10000	    187417 ns/op
PASS
ok  	github.com/Succo/fizz-buzz-web	21.279s
```

There is multiples query types, some are made to have a lot of fizz and buzz remplacement, whereas other are more focused on printing numbers (which is why i1 and i2 are superior to the limit).

Reordered the results are (the column names are i1 i2 limit)

Implementation  | 3 5 15 | 10 3 100 | 1000 1000 500 |
---			    | ---	 | ---	    | ---		    |
Naive 		    | 4848   | 32278    | 183031        |
Memorize modulo | 4913   | 33056	| 186772		|
Updated var		| 4871	 | 32756	| 187417		|
Count down		| 4358	 | 28137	| 178977		|

So it seems that using the "count down" method to avoid all modulo operation if the fatest on those small benchmark.
