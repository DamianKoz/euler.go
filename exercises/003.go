package exercises

import (
	"fmt"
	"time"
)

// Exercise:
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
// -----------
// Notes:
// Prime Number: A whole number, that cannot be made by multiplying other whole numbers
// Prime Factorization: Find the prime numbers which add up to make the original number

func Solve003() int {
	prime_factors := []int{}
	original_num := 600851475143
	divider := 2
	start := time.Now()

	for original_num > 1 {
		if original_num%divider == 0 {
			prime_factors = append(prime_factors, divider)
			original_num /= divider
		}
		divider += 1
	}
	end := time.Now()
	fmt.Printf("THIS PROGRAMM RAN %v SECONDS\n", end.Sub(start))
	return prime_factors[len(prime_factors)-1]
}
