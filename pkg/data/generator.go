package data

import (
	"mock-data-generator/internal/csv"
	"mock-data-generator/pkg/utils"
	"os"
)

// GenerateData generates mock Latin American people's data and saves it to a CSV file.
func GenerateData(numRecords int) error {
	// Separate male and female names
	maleNames := []string{"Carlos", "Pedro", "Luis", "Diego", "Mateo"}
	femaleNames := []string{"Maria", "Ana", "Sofia", "Camila", "Valentina"}
	lastNames := []string{"Gomez", "Rodriguez", "Lopez", "Perez", "Martinez", "Gonzalez", "Hernandez", "Silva", "Torres", "Ramirez"}

	// Shuffle names to add some randomness
	utils.ShuffleStringSlice(maleNames)
	utils.ShuffleStringSlice(femaleNames)
	utils.ShuffleStringSlice(lastNames)

	// Create a CSV file
	file, err := os.Create("generated/people.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewCSVWriter(file)
	defer writer.Close()

	// Write header row
	header := []string{"First Name 1", "First Name 2", "Last Name 1", "Last Name 2", "Email"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Generate and write data rows
	for i := 0; i < numRecords; i++ {
		var firstName1, firstName2, lastName1, lastName2 string

		if i%2 == 0 {
			// Male names for even records
			firstName1 = maleNames[i%len(maleNames)]
			firstName2 = maleNames[(i+1)%len(maleNames)]
		} else {
			// Female names for odd records
			firstName1 = femaleNames[i%len(femaleNames)]
			firstName2 = femaleNames[(i+1)%len(femaleNames)]
		}

		lastName1 = lastNames[i%len(lastNames)]
		lastName2 = lastNames[(i+1)%len(lastNames)]

		email := utils.GenerateRandomEmail()

		dataRow := []string{firstName1, firstName2, lastName1, lastName2, email}

		if err := writer.Write(dataRow); err != nil {
			return err
		}
	}

	return nil
}
