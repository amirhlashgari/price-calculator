package main

import (
	"fmt"

	"github.com/amirhlashgari/price-calculator/filemanager"
	"github.com/amirhlashgari/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// NOTE: with using interfaces we now can accept both reading from file and accepting from command line

		// cmd := cmdmanager.New()
		// priceJob1 := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		// priceJob1.Process()

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob2 := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob2.Process()
	}
}
