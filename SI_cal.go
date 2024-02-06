package main

import (
	"fmt"
)

func simpleInterest(principal float64, rate float64, time float64) float64 {
	return (principal * rate * time) / 100
}

func main() {
	var principal, rate, time float64

	fmt.Println("Enter principal amount:")
	fmt.Scan(&principal)

	fmt.Println("Enter rate of interest:")
	fmt.Scan(&rate)

	fmt.Println("Enter time (in years):")
	fmt.Scan(&time)

	interest := simpleInterest(principal, rate, time)

	fmt.Printf("Simple Interest: %.2f\n", interest)
}
