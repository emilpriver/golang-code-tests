package main

import (
	"fmt"
)

func main() {
	records := ReadAndParseCsvFile()
	fmt.Println(records)
}
