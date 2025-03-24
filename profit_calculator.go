package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var tax_rate float64

	fmt.Print("Enter the Revenue Generated: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the Tax Rate: ")
	fmt.Scan(&tax_rate)

	var earnings_before_tax float64 = revenue - expenses
	var earnings_after_tax float64 = (revenue - expenses) * (1 - tax_rate)

	fmt.Println("Earnings Before Tax: ", earnings_before_tax)
	fmt.Println("Earnings After Tax: ", earnings_after_tax)
}