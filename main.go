package main

import (
	"os"
	"encoding/json"
	"fmt"
	"errors"

	"ledgr/sheets"
	"ledgr/mechanics"

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

func addAttributeModifier(attribute *sheets.SingleAttribute) error {
	// Check that attribute exists
	if attribute == nil {
		return errors.New("Attribute is nil")
	}
	// Check that the score is >= 0
	if attribute.Score < 0 {
		return errors.New("Input attribute score was <0 and is not allowed.")
	}

	// Calculate modifier
	mod, err := mechanics.CalcAbilityModifier(attribute.Score)

	if err != nil {
		return err
	}

	// Add modifier
	attribute.Modifier = mod

	return nil
	
}

func updateAllAttributeModifiers(sheet *sheets.AttributeStats) error {

	err := addAttributeModifier(&sheet.Strength)
	err = addAttributeModifier(&sheet.Dexterity)
	err = addAttributeModifier(&sheet.Constitution)
	err = addAttributeModifier(&sheet.Intelligence)
	err = addAttributeModifier(&sheet.Wisdom)
	err = addAttributeModifier(&sheet.Charisma)

	return err
}

func updateSavingThrowModifiers(savingThrows *sheets.SavingThrowStats, attributes *sheets.AttributeStats) error {
	savingThrows.Strength.Modifier = attributes.Strength.Modifier
	savingThrows.Dexterity.Modifier = attributes.Dexterity.Modifier
	savingThrows.Constitution.Modifier = attributes.Constitution.Modifier
	savingThrows.Intelligence.Modifier = attributes.Intelligence.Modifier
	savingThrows.Wisdom.Modifier = attributes.Wisdom.Modifier
	savingThrows.Charisma.Modifier = attributes.Charisma.Modifier
}

func main() {
	
	// Take in a file and verify that it is a valid YAML
	filePath := "sheets/assets/schema/schema_v1_template.yaml"
	fmt.Println(">>> Reading sheet:", filePath)
	file, err := os.ReadFile(filePath)
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

	// Update all ability modifiers
	err = updateAllAttributeModifiers(&characterSheet.Attributes)
	if err != nil {
		panic(err)
	}	
	
	// Convert back to YAML
	outputYaml, err := yaml.Marshal(&characterSheet)
	if err != nil {
		panic(err)
	}

	// Return the YAML
	os.WriteFile("character_out.yaml", outputYaml, 0644)

}
