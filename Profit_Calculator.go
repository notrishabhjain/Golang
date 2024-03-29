package main

import "fmt"

func profitCalculatorFunc() {
	var revenue, expenses, taxRate float64

	fmt.Println("Enter the total revenue")
	fmt.Scan(&revenue)

	fmt.Println("Enter the total expenses")
	fmt.Scan(&expenses)

	fmt.Println("Enter the tax rate")
	fmt.Scan(&taxRate)

	earningBeforeTax := earningBeforeTaxCalculator(revenue, expenses)
	fmt.Printf("Earning Before Tax : %.2f\n", earningBeforeTax)

	profit := profitEarnedCalculator(earningBeforeTax, taxRate)
	fmt.Printf("Profit after Tax : %.2f\n", profit)

	ratio := fmt.Sprintf("Profit Ratio : %1f\n", earningBeforeTax/profit)
	fmt.Print(ratio)
}

func earningBeforeTaxCalculator(revenue, expenses float64) float64 {
	return revenue - expenses
}

func profitEarnedCalculator(earningBeforeTax, taxRate float64) (eBT float64) {
	eBT = earningBeforeTax * (1 - taxRate/100)
	return eBT
	// return
}
