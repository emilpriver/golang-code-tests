package main

import (
	"encoding/csv"
	"os"
)

/**
* Read and parse the CSV file.
 */
func ReadAndParseCsvFile() [][]string {
	csvFile, err := os.Open("./customers.csv")

	if err != nil {
		panic("Error opening CSV file")
	}

	// We dont need to wait for Go to close the file
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	csvRecords, err := csvReader.ReadAll()

	if err != nil {
		panic("Error reading CSV content")
	}

	return csvRecords
}
