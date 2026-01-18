package sheets

import (
	"math"
	"errors"
)

type ModifierNumeric interface {
	~int | ~uint
}

func CalcAbilityModifier[T ModifierNumeric](stat T) (modifier int, err error) {

	/*
	This takes an ability scores and returns a modifier
	*/

	if stat < 0 {
		err = errors.New("Cannot use a negative ability score")
		return 0, err
	}

	// Convert stat to float64 as required by floor
	statAsFloat := float64(stat)

	// Calculate the modifier
	modifier = int(math.Floor((statAsFloat - 10.0) / 2.0))

	return modifier, nil
}


func CalcAllAbilityModifiers(sheet *AttributeStats) error {

	for _, ability := range AllAbilities {

		// Get the score
		score := sheet.GetScore(ability)

		// Calculate Modifier
		mod, err := CalcAbilityModifier(score)

		if err != nil {
			return err
		}

		// Set the Modifier
		sheet.SetModifier(ability, mod)

	}

	return nil
}


func SetAllSavingThrowModifiers(savingThrow *SavingThrowStats, abilities *AttributeStats) {

	for _, ability := range AllAbilities {

		// Get the modifier
		mod := abilities.GetModifier(ability)

		// Set the Modifier
		savingThrow.SetModifier(ability, mod)

	}
}


func SetAllSkillModifiers(skills *SkillStats, abilities *AttributeStats) {

	for skill, ability := range SkillToAbility {
		// Get the ability modifier 
		abilityMod := abilities.GetModifier(ability)

		// Update the skill
		skills.SetModifier(skill, abilityMod)
	}


}