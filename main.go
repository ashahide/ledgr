package main

import (
	"os"
	"encoding/json"
	"fmt"

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

func calcAllAbilityModifiers(sheet *sheets.AttributeStats) error {

	for _, ability := range sheets.AllAbilities {

		// Get the score
		score := sheet.GetScore(ability)

		// Calculate Modifier
		mod, err := mechanics.CalcAbilityModifier(score)

		if err != nil {
			return err
		}

		// Set the Modifier
		sheet.SetModifier(ability, mod)

	}

	return nil
}


func setAllSavingThrowModifiers(savingThrow *sheets.SavingThrowStats, abilities *sheets.AttributeStats) {

	for _, ability := range sheets.AllAbilities {

		// Get the modifier
		mod := abilities.GetModifier(ability)

		// Set the Modifier
		savingThrow.SetModifier(ability, mod)

	}
}


func setAllSkillModifiers(skills *sheets.SkillStats, abilities *sheets.AttributeStats) {

	for skill, ability := range sheets.SkillToAbility {
		// Get the ability modifier 
		abilityMod := abilities.GetModifier(ability)

		// Update the skill
		skills.SetModifier(skill, abilityMod)
	}


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

	// Calculate ability modifiers
	calcAllAbilityModifiers(&characterSheet.Attributes)

	// Set saving throw modifiers
	setAllSavingThrowModifiers(&characterSheet.SavingThrows, &characterSheet.Attributes )

	// Set skill modifiers
	setAllSkillModifiers(&characterSheet.Skills, &characterSheet.Attributes)

	// Convert back to YAML
	outputYaml, err := yaml.Marshal(&characterSheet)
	if err != nil {
		panic(err)
	}

	// Return the YAML
	os.WriteFile("character_out.yaml", outputYaml, 0644)

}
