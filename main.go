package main

import (
	"os"
	"encoding/json"
	"fmt"

	"ledgr/sheets"

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

func sheetToJSONInstance(sheet sheets.CharacterSheet) (any, error) {
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
	file, err := os.ReadFile("sheets/assets/schema/schema_v1_template.yaml")
	if err != nil {
		panic(err)
	}

	// Unmarshel YAML 
	var characterSheet sheets.CharacterSheet
	err = yaml.Unmarshal(file, &characterSheet)
	if err != nil {
		panic(err)
	}

	// Convert to a json
	characterSheetJsonInstance, err := sheetToJSONInstance(characterSheet)
	if err != nil {
		panic(err)
	}

	// Load json schema
	schema, err := loadSchema("sheets/assets/schema/character_sheet.schema.json")
	if err != nil {
		panic(err)
	}

	// Validate schema
	schema.Validate(&characterSheetJsonInstance)

	// Validate that is has the starting information that is necessary
	data, err := json.Marshal(characterSheetJsonInstance)
	
	fmt.Println(string(data))

	// Make a copy of the YAML and fill it in with the available info
	

	// Return the YAML

}
