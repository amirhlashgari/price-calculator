package prices

import (
	"fmt"

	"github.com/amirhlashgari/price-calculator/conversion"
	"github.com/amirhlashgari/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"` // IMPORTANT: this way you can avoid and attribute to come in the json files
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

/*
IMPORTANT: we should pass the pointer to file in this receiver function
because if we don't pass pointer it would be a copy of the struct
and in this case the InputPrices in the struct won't change after reading from file
*/
func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool) {
	err := job.LoadData()
	if err != nil {
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	// filemanager.WriteJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
