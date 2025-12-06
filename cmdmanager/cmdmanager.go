package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() *CMDManager {
	return &CMDManager{}
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Enter prices")

	var prices []string
	for {
		fmt.Print("Price: ")

		var price string
		_, err := fmt.Scan(&price)
		if err != nil {
			return nil, err
		}

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
