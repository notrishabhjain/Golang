package main

import (
  "fmt"
  "math"
)

func compoundInterest(principal float64, rate float64, time float64) float64 {
  return (principal * math.Pow(1 + rate/100,time))

}

func compoundInterestFunc() {
  var principal, rate, time float64

  fmt.Println("Enter principal amount:")
  fmt.Scan(&principal)

  fmt.Println("Enter rate of interest:")
  fmt.Scan(&rate)

  fmt.Println("Enter time (in years):")
  fmt.Scan(&time)

  interest := compoundInterest(principal, rate, time)

  fmt.Printf("Compound Interest: %.2f\n", interest)
}
