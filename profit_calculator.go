package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	{
		fmt.Println("Revenue:")
		fmt.Scan(&revenue)
	}
	{
		fmt.Println("expenses:")
		fmt.Scan(&expenses)
	}
	{
		fmt.Println("taxRate:")
		fmt.Scan(&taxRate)
	}

	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit
	fmt.Println("ebt:", ebt)
	fmt.Println("prfit: ", profit)
	fmt.Println("ratio:", ratio)
}
