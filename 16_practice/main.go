package main

import (
	"fmt"

	"example.com/practice/filemanager"
	"example.com/practice/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Error processing price job:", err)
		// }
	}

	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Printf("Error in job with tax rate %.0f%%: %v\n", taxRates[i]*100, err)
			}
		case <-doneChans[i]:
			fmt.Printf("Done!")
		}
	}

	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// }

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }

}
