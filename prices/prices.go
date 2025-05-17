package prices

import (
	"fmt"

	"github.com/RevanthGovindan/go-learning/conversions"
	"github.com/RevanthGovindan/go-learning/iomanager"
)

type TaxIncludedPriceJob struct {
	IoManager         iomanager.IoManger
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(fm iomanager.IoManger, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{},
		TaxRate:     taxRate,
		IoManager:   fm,
	}
}

func (t *TaxIncludedPriceJob) Process() error {
	err := t.LoadData()
	if err != nil {
		return err
	}
	var result map[string]float64 = make(map[string]float64)

	for _, price := range t.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + t.TaxRate)
	}
	t.TaxIncludedPrices = result
	t.IoManager.WriteResult(t)
	return nil
}

func (t *TaxIncludedPriceJob) LoadData() error {
	lines, err := t.IoManager.ReaLines()
	if err != nil {
		return err
	}
	t.InputPrices, err = conversions.StringsToFloats(lines)
	if err != nil {
		return err
	}
	return nil
}
