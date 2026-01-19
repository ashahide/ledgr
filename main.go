package main

import (
	"os"
	"encoding/json"
	"flag"
	"log"

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
	
	// Define the parser
	filePath := flag.String("f", "sheets/assets/schema/schema_v1_template.yaml", "Enter the file path") 

	// Parse inputs
	flag.Parse()

	// Take in a file and verify that it is a valid YAML
	log.Println("Reading file:", *filePath)
	file, err := os.ReadFile(*filePath)
	if err != nil {
		log.Println("ERROR reading file:", err)
		return
	}

	// Unmarshel YAML 
	var characterSheet sheets.CharacterSheet
	err = yaml.Unmarshal(file, &characterSheet)
	log.Println("Parsing input .yaml")
	if err != nil {
		log.Println("ERROR parsing file:", err)
		return
	}

	// Convert to a json
	log.Println("Creating .json from .yaml")
	characterSheetJsonInstance, err := sheetToJSONInstance(characterSheet)
	if err != nil {
		log.Println("ERROR creating json:", err)
	}

	// Load json schema
	log.Println("Loading .schema.json file")
	schema, err := loadSchema("sheets/assets/schema/character_sheet.schema.json")
	if err != nil {
		log.Println("ERROR loading .schema.json:", err)
	}

	// Validate schema
	log.Println("Validating schema")
	schema.Validate(&characterSheetJsonInstance)

	// Calculate ability modifiers
	log.Println("Calcuating ability score modifiers")
	sheets.CalcAllAbilityModifiers(&characterSheet.Attributes)

	// Set saving throw modifiers
	log.Println("Calculating saving throw modifiers")
	sheets.SetAllSavingThrowModifiers(&characterSheet.SavingThrows, &characterSheet.Attributes )

	// Set skill modifiers
	log.Println("Calculating skill modifiers")
	sheets.SetAllSkillModifiers(&characterSheet.Skills, &characterSheet.Attributes)

	// Convert back to YAML
	log.Println("Creating output .yaml")
	outputYaml, err := yaml.Marshal(&characterSheet)
	if err != nil {
		log.Println("ERROR generating .yaml:", err)
	}

	// Return the YAML
	os.WriteFile("test_character_out.yaml", outputYaml, 0644)

}
