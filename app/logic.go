package main

import (
	"math"
	"math/rand"
	"time"
)

//50% Probability Solution
func FlipCoin() int {
	max := 2
	min := 0
	rand.Seed(time.Now().UTC().UnixNano())
	randomNum := min + rand.Intn(max-min)
	return randomNum
}

func GenerateRandomNumber() int {
	max := 21 //20
	min := 0
	rand.Seed(time.Now().UTC().UnixNano())
	randomNum := min + rand.Intn(max-min)
	return randomNum
}

//Prime Number Checker
func CheckPrimeNumber(num int) bool {
	prime := true
	if num < 2 {
		prime = false
	}
	square_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= square_root; i++ {
		if num%i == 0 {
			prime = false
		}
	}
	return prime
}

//fibonnacci
func Fibonacci(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func GenerateRandomNumberPokemon() int {
	max := 905 //20
	min := 1
	rand.Seed(time.Now().UTC().UnixNano())
	randomNum := min + rand.Intn(max-min)
	return randomNum
}
