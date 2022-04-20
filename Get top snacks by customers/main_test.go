package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	expectedJson := `
		[
			{
				"name": "Jonas",
				"favouriteSnack": "Geisha",
				"totalSnacks": 1800
			},
			{
				"name": "Annika",
				"favouriteSnack": "Geisha",
				"totalSnacks": 200
			},
			{
				"name": "Jane",
				"favouriteSnack": "Nötchoklad",
				"totalSnacks": 22
			},
			{
				"name": "Aadya",
				"favouriteSnack": "Center",
				"totalSnacks": 9
			}
		]
	`

	snacksJson, err := json.MarshalIndent(CustomerSnacksSorted(), "", "  ")

	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}

	jsonToString := string(snacksJson)

	require.JSONEq(t, expectedJson, jsonToString)
}
