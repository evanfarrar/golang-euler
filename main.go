package main

import "fmt"

// problem 1
func multiples_of_three_and_five(limit int) (sum int) {
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// problem 2
func sum_of_even_fib(limit int) (sum uint64) {
	previous := 1
	for current := 2; current < limit; current += previous {
		if current%2 == 0 {
			sum += uintcurrent
		}
	}
	return sum
}

func main() {
	fmt.Printf("%d\n", sum_of_even_fib(4000000))
}
