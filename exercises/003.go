package exercises

import "fmt"

// Exercise:
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
// -----------
// Notes:
// Prime Number: A whole number, that cannot be made by multiplying other whole numbers
// Prime Factorization: Find the prime numbers which add up to make the original number

func Solve003() []int {
	result := []int{}
	original_num := 13195
	fmt.Print(result, original_num)
	primes := get_primes()

	return primes
}

func get_prime_factors(num int) []int {
	return []int{}
}

func get_primes() []int {
	primes := []int{}
	upper_bound := 8000
	for i := 2; i <= upper_bound; i++ {
		is_prime := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			primes = append(primes, i)
		}
	}
	return primes
}
