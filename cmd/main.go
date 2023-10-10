package cmd

import (
	"fmt"
	"mock-data-generator/pkg/data"
)

// Run starts the application
func Run(numRecords int) error {
	err := data.GenerateData(numRecords)
	if err != nil {
		return err
	}
	fmt.Printf("%d mock Latin American people's names with two names and two last names (separated by gender) generated and saved to 'generated/people.csv'.\n", numRecords)
	return nil
}
