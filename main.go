package main

import (
	"os"
	"encoding/json"

	"gopkg.in/yaml.v3"
	"github.com/santhosh-tekuri/jsonschema/v5"
)


func loadSchema(schemaPath string) (*jsonschema.Schema, error) {
	compiler := jsonschema.NewCompiler()

	// Register the schema file
	f, err := os.Open(schemaPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := compiler.AddResource(schemaPath, f); err != nil {
		return nil, err
	}

	// Compile it
	return compiler.Compile(schemaPath)
}

func sheetToJSONInstance(sheet CharacterSheet) (any, error) {
	data, err := json.Marshal(sheet)
	if err != nil {
		return nil, err
	}

	var instance any
	if err := json.Unmarshal(data, &instance); err != nil {
		return nil, err
	}

	return instance, nil
}

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
	characterSheetJsonInstance, err := sheetToJSONInstance(characterSheet)
	if err != nil {
		panic(err)
	}

	
	schema, err := loadSchema("schema/character_sheet.schema.json")
	if err != nil {
		panic(err)
	}
	schema.Validate(&characterSheetJsonInstance)
	
	// Validate that is has the starting information that is necessary


	// Make a copy of the YAML and fill it in with the available info

	// Return the YAML

}