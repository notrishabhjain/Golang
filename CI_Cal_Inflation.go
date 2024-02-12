package main

import (
	"fmt"
	"math"
)

func compoundInterestInflation(principal float64, rate float64, time float64, inflation float64) float64 {
	var interest = (principal * math.Pow(1+rate/100, time))
	fmt.Printf("Compound Interest (without adjusted inflation): %.2f\n", interest)
	return (interest / math.Pow(1+inflation/100, time))
}

func compoundInterestInflationFunc() {
	const inflation float64 = 5.5
	var principal, rate, time float64

	fmt.Println("Enter principal amount:")
	fmt.Scan(&principal)

	fmt.Println("Enter rate of interest:")
	fmt.Scan(&rate)

	fmt.Println("Enter time (in years):")
	fmt.Scan(&time)

	interestInflation := compoundInterestInflation(principal, rate, time, inflation)
	fmt.Printf("Compound Interest (with adjusted inflation): %.2f\n", interestInflation)
}
