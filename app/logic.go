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
