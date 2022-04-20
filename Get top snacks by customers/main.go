package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Customer struct {
	Name  string
	Candy string
	Eaten int
}

type CustomerTopSnack struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int    `json:"totalSnacks"`
}

type Snack struct {
	Candy string
	Eaten int
}

/**
 * Read and parse the CSV file.
 */
func readAndParseCsvFile() [][]string {
	csvFile, err := os.Open("./customers.csv")

	if err != nil {
		panic("Error opening CSV file")
	}

	// Close the csvFile when program is closed
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	csvRecords, err := csvReader.ReadAll()

	if err != nil {
		panic("Error reading CSV content")
	}

	return csvRecords
}

func convertRecordsIntoStructs(records [][]string) []Customer {
	var customers []Customer

	for index, row := range records {
		// We dont want the CSV headers. Headers have the first row which index is 0. So lets skip first row.
		if index > 0 {
			for _, line := range row {
				/**
				 * If we look at the CSV file do we know that
				 * Name = index 0
				 * Candy = index 1
				 * Eaten = index 2
				 */
				splittedString := strings.Split(line, ";")

				// We will compair and add values to a total snack. so convert to int before added to customers
				eatenAsInt, err := strconv.Atoi(splittedString[2])

				if err != nil {
					panic("Couldn't convert string to int")
				}

				customer := Customer{
					Name:  splittedString[0],
					Candy: splittedString[1],
					Eaten: eatenAsInt,
				}

				customers = append(customers, customer)
			}
		}
	}

	return customers
}

/**
 * Check if value exists in a list, return bool depending on result.
 */
func contains(list []string, value string) bool {
	for _, arrayValue := range list {
		if arrayValue == value {
			return true
		}
	}

	return false
}

/**
 * Take the customers which earlier on is readed from the file and get the customers top snacks
 */
func customersTopSnack(customers []Customer) []CustomerTopSnack {
	var customerTopSnacks []CustomerTopSnack

	/**
	 * Create a map which have the name of the customer as key
	 * and have snack struct which includes Eaten and Candy
	 * where candy is the name of the Snacka and eaten is how
	 * times the customer have eaten the snack
	 */
	groupedCustomerSnacks := make(map[string][]Snack)

	for _, customer := range customers {
		s := Snack{
			Candy: customer.Candy,
			Eaten: customer.Eaten,
		}

		groupedCustomerSnacks[customer.Name] = append(groupedCustomerSnacks[customer.Name], s)
	}

	// loop thrue the new grouped customer snack
	for customerName := range groupedCustomerSnacks {
		// Check if the snack name exists
		if _, exists := groupedCustomerSnacks[customerName]; exists {
			snackTotals := make(map[string]int)

			// Loop through all snacks eaten by customer
			for _, snack := range groupedCustomerSnacks[customerName] {
				// If the snack exists append the eaten value to the already added eaten value
				if _, exists := snackTotals[snack.Candy]; exists {
					snackTotals[snack.Candy] = snackTotals[snack.Candy] + snack.Eaten
				} else {
					// Add a array item with candy name as key and eaten as value
					snackTotals[snack.Candy] = snack.Eaten
				}
			}

			// Create default value and name which is used to check against and then add to top snacks
			highestSnackValue := 0
			highestSnackName := ""

			// name is the key and value is the eaten
			for name, value := range snackTotals {
				// if array items eaten is more then the current heightSnackValue, replace it
				if value > highestSnackValue {
					highestSnackValue = value
					highestSnackName = name
				}
			}

			/**
			 * Create a new customer top snack stat with customer name, heighestValue and name of the array
			 * and the return this information.
			 */
			customTopSnack := CustomerTopSnack{
				Name:           customerName,
				FavouriteSnack: highestSnackName,
				TotalSnacks:    highestSnackValue,
			}

			customerTopSnacks = append(customerTopSnacks, customTopSnack)
		}
	}

	return customerTopSnacks
}

/**
 * Sort the top snacks information by highest total snacks
 * Requires go 1.8
 */
func sortTopSnacksByHighest(snacks []CustomerTopSnack) []CustomerTopSnack {
	sort.Slice(snacks, func(i, j int) bool {
		return snacks[i].TotalSnacks > snacks[j].TotalSnacks
	})

	return snacks
}

func CustomerSnacksSorted() []CustomerTopSnack {
	records := readAndParseCsvFile()

	customers := convertRecordsIntoStructs(records)

	stats := customersTopSnack(customers)

	sortedStats := sortTopSnacksByHighest(stats)

	return sortedStats
}

func main() {
	/*
	 * Generate json data out of the array struc that have been working with which is indented and
	 * formated
	 */
	jsonOutput, err := json.MarshalIndent(CustomerSnacksSorted(), "", "  ")

	if err != nil {
		panic("Error converting array struct to json")
	}

	log.Println(string(jsonOutput))
}
