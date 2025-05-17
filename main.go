package main

import (
	"github.com/RevanthGovindan/go-learning/cmdmanager"
	"github.com/RevanthGovindan/go-learning/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		// var fm = filemanager.New("prices.txt", fmt.Sprintf("output_%d.json", index))
		var cmd = cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		err := priceJob.Process()
		if err != nil {
			panic(err)
		}
	}
}
