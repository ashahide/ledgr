package main

import (
	"fmt"
	"os"
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func main() {
	
	// Take in a file and verify that it is a valid YAML
	file, err := os.ReadFile("assets/schema_v1_template.yaml")
	if err != nil {
		panic(err)
	}

	// Unmarshel YAML 
	var characterSheet CharacterSheet
	err = yaml.Unmarshal(file, &characterSheet)
	if err != nil {
		panic(err)
	}

	// Convert to a json
	characterSheetJson, err := json.MarshalIndent(characterSheet, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(characterSheetJson))

	// Validate that is has the starting information that is necessary

	// Make a copy of the YAML and fill it in with the available info

	// Return the YAML

}