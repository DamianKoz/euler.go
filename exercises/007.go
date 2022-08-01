package exercises

import (
	"fmt"
	"time"
)

func Solve007() int {
	primes := []int{}
	upper_bound := 10001

	for i := 2; len(primes) < upper_bound; i++ {
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
	end := time.Now()
	fmt.Printf("END %v ", end)
	return primes[len(primes)-1]
}
