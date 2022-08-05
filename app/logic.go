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
	if num == 2 {
		prime = true
		return prime
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
	var result int
	var tempA int
	var tempB int

	for i := 0; i < n; i++ {
		if i == 0 {
			tempA = i
			result = tempA
		} else if i == 1 {
			tempB = i
			result = tempB
		} else {
			result = tempA + tempB
			tempA = tempB
			tempB = result
		}
	}
	return result
}

func GenerateRandomNumberPokemon() int {
	max := 905 //20
	min := 1
	rand.Seed(time.Now().UTC().UnixNano())
	randomNum := min + rand.Intn(max-min)
	return randomNum
}
