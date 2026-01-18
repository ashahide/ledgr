package sheets

import "ledgr/mechanics"

func CalcAllAbilityModifiers(sheet *AttributeStats) error {

	for _, ability := range AllAbilities {

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