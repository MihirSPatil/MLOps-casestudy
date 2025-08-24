package sieve

import (
	"math"
)

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	// Create the boolean slice.
	primes := make([]bool, limit+1)

	// FIX 1: Initialize all numbers from 2 to limit as potentially prime (true).
	for i := 2; i <= limit; i++ {
		primes[i] = true
	}

	// Calculate the limit for the outer loop.
	loop_limit := int(math.Sqrt(float64(limit)))

	// FIX 3: The outer loop should only go up to the square root of the limit.
	for i := 2; i <= loop_limit; i++ {
		// Only check for multiples if 'i' is still considered a prime.
		if primes[i] {
			// FIX 4: The inner loop must mark multiples all the way up to 'limit'.
			// We start from i*i for optimization.
			for num := i * i; num <= limit; num += i {
				// FIX 2: Mark the multiple 'num' as not prime, not the prime 'i'.
				primes[num] = false
			}
		}
	}

	// Collect the results.
	var num_list []int
	for i := 2; i <= limit; i++ {
		if primes[i] {
			num_list = append(num_list, i)
		}
	}

	return num_list
}