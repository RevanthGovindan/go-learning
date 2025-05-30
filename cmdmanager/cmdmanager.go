package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() CMDManager {
	return CMDManager{}
}

func (f CMDManager) ReaLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with enter")
	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (f CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
