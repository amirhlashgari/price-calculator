package main

import (
	"fmt"

	"github.com/amirhlashgari/price-calculator/filemanager"
	"github.com/amirhlashgari/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		// NOTE: with using interfaces we now can accept both reading from file and accepting from command line

		// cmd := cmdmanager.New()
		// priceJob1 := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		// priceJob1.Process()
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])
		// if err != nil {
		// 	fmt.Println("Could not process job", err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println("Error:", err)
			}
		case <-doneChans[index]:
			fmt.Println("Done")
		}
	}
}
