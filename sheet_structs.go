package main

/*
Package ledgr defines the canonical data model for a YAML-based
D&D 5e character sheet.

Design philosophy:

- YAML is treated as a *human-authored configuration format*
- Go structs are the *authoritative schema*
- Some fields are user-authored "truths"
- Some fields are derived and may be omitted or validated
- Pointer fields (*T) indicate optional or overrideable values
*/

/*
========================
ROOT CHARACTER DOCUMENT
========================
*/

type CharacterSheet struct {
	SchemaVersion uint16 `yaml:"schema_version,omitempty" json:"schema_version,omitempty"`

	Basics       CharacterBasicStats `yaml:"basics" json:"basics"`
	Attributes   AttributeStats      `yaml:"attributes" json:"attributes"`
	SavingThrows SavingThrowStats    `yaml:"saving_throws" json:"saving_throws"`
	Skills       SkillStats          `yaml:"skills" json:"skills"`
	Health       HealthStats         `yaml:"health" json:"health"`

	// Combat contains AC, initiative, and speed values.
	Combat CombatStats `yaml:"combat,omitempty" json:"combat,omitempty"`

	// Spellcasting is omitted for non-spellcasters.
	Spellcasting SpellcastingStats `yaml:"spellcasting,omitempty" json:"spellcasting,omitempty"`
}

/*
========================
BASIC CHARACTER METADATA
========================
*/

type CharacterBasicStats struct {
	Name      string `yaml:"name" json:"name"`
	Age       uint16 `yaml:"age,omitempty" json:"age,omitempty"`
	Class     string `yaml:"class" json:"class"`
	Level     uint8  `yaml:"level" json:"level"`
	Race      string `yaml:"race" json:"race"`
	Alignment string `yaml:"alignment,omitempty" json:"alignment,omitempty"`
}

/*
================
ABILITY SCORES
================
*/

type SingleAttribute struct {
	Score    uint8 `yaml:"score" json:"score"`
	Modifier *int8 `yaml:"modifier,omitempty" json:"modifier,omitempty"`
}

type AttributeStats struct {
	Strength     SingleAttribute `yaml:"strength" json:"strength"`
	Dexterity    SingleAttribute `yaml:"dexterity" json:"dexterity"`
	Constitution SingleAttribute `yaml:"constitution" json:"constitution"`
	Intelligence SingleAttribute `yaml:"intelligence" json:"intelligence"`
	Wisdom       SingleAttribute `yaml:"wisdom" json:"wisdom"`
	Charisma     SingleAttribute `yaml:"charisma" json:"charisma"`
}

/*
================
SAVING THROWS
================
*/

type SingleSavingThrow struct {
	Proficient bool  `yaml:"proficient,omitempty" json:"proficient,omitempty"`
	Expertise  bool  `yaml:"expertise,omitempty" json:"expertise,omitempty"`
	MiscBonus  int8  `yaml:"misc_bonus,omitempty" json:"misc_bonus,omitempty"`
	Modifier   *int8 `yaml:"modifier,omitempty" json:"modifier,omitempty"`
}

type SavingThrowStats struct {
	Strength     SingleSavingThrow `yaml:"strength" json:"strength"`
	Dexterity    SingleSavingThrow `yaml:"dexterity" json:"dexterity"`
	Constitution SingleSavingThrow `yaml:"constitution" json:"constitution"`
	Intelligence SingleSavingThrow `yaml:"intelligence" json:"intelligence"`
	Wisdom       SingleSavingThrow `yaml:"wisdom" json:"wisdom"`
	Charisma     SingleSavingThrow `yaml:"charisma" json:"charisma"`
}

/*
=========
SKILLS
=========
*/

type SingleSkill struct {
	Proficient       bool   `yaml:"proficient,omitempty" json:"proficient,omitempty"`
	Expertise        bool   `yaml:"expertise,omitempty" json:"expertise,omitempty"`
	MiscBonus        int8   `yaml:"misc_bonus,omitempty" json:"misc_bonus,omitempty"`
	RelatedAttribute string `yaml:"related_attribute,omitempty" json:"related_attribute,omitempty"`
	Modifier         *int8  `yaml:"modifier,omitempty" json:"modifier,omitempty"`
}

type SkillStats struct {
	Acrobatics     SingleSkill `yaml:"acrobatics" json:"acrobatics"`
	AnimalHandling SingleSkill `yaml:"animal_handling" json:"animal_handling"`
	Arcana         SingleSkill `yaml:"arcana" json:"arcana"`
	Athletics      SingleSkill `yaml:"athletics" json:"athletics"`
	Deception      SingleSkill `yaml:"deception" json:"deception"`
	History        SingleSkill `yaml:"history" json:"history"`
	Insight        SingleSkill `yaml:"insight" json:"insight"`
	Intimidation   SingleSkill `yaml:"intimidation" json:"intimidation"`
	Investigation  SingleSkill `yaml:"investigation" json:"investigation"`
	Medicine       SingleSkill `yaml:"medicine" json:"medicine"`
	Nature         SingleSkill `yaml:"nature" json:"nature"`
	Perception     SingleSkill `yaml:"perception" json:"perception"`
	Performance    SingleSkill `yaml:"performance" json:"performance"`
	Persuasion     SingleSkill `yaml:"persuasion" json:"persuasion"`
	Religion       SingleSkill `yaml:"religion" json:"religion"`
	SleightOfHand  SingleSkill `yaml:"sleight_of_hand" json:"sleight_of_hand"`
	Stealth        SingleSkill `yaml:"stealth" json:"stealth"`
	Survival       SingleSkill `yaml:"survival" json:"survival"`
}

/*
========
HEALTH
========
*/

type HealthStats struct {
	Current uint16 `yaml:"current,omitempty" json:"current,omitempty"`
	Max     uint16 `yaml:"max,omitempty" json:"max,omitempty"`
	Temp    uint16 `yaml:"temp,omitempty" json:"temp,omitempty"`
}

/*
=========================
COMBAT / DEFENSE / MOVEMENT
=========================
*/

type SpeedStats struct {
	Walk   uint16 `yaml:"walk,omitempty" json:"walk,omitempty"`
	Fly    uint16 `yaml:"fly,omitempty" json:"fly,omitempty"`
	Swim   uint16 `yaml:"swim,omitempty" json:"swim,omitempty"`
	Climb  uint16 `yaml:"climb,omitempty" json:"climb,omitempty"`
	Burrow uint16 `yaml:"burrow,omitempty" json:"burrow,omitempty"`
}

type InitiativeStats struct {
	MiscBonus int8  `yaml:"misc_bonus,omitempty" json:"misc_bonus,omitempty"`
	Total     *int8 `yaml:"total,omitempty" json:"total,omitempty"`
}

type ArmorClassStats struct {
	MiscBonus int8   `yaml:"misc_bonus,omitempty" json:"misc_bonus,omitempty"`
	Total     *uint8 `yaml:"total,omitempty" json:"total,omitempty"`
}

type CombatStats struct {
	ArmorClass ArmorClassStats `yaml:"armor_class,omitempty" json:"armor_class,omitempty"`
	Initiative InitiativeStats `yaml:"initiative,omitempty" json:"initiative,omitempty"`
	Speed      SpeedStats      `yaml:"speed,omitempty" json:"speed,omitempty"`
}

/*
===================
SPELLCASTING STATS
===================
*/

type SpellcastingStats struct {
	SpellcastingAbility string `yaml:"spellcasting_ability,omitempty" json:"spellcasting_ability,omitempty"`
	SpellSaveDC         *int8  `yaml:"spell_save_dc,omitempty" json:"spell_save_dc,omitempty"`
	SpellAttackBonus    *int8  `yaml:"spell_attack_bonus,omitempty" json:"spell_attack_bonus,omitempty"`
	MiscBonus           int8   `yaml:"misc_bonus,omitempty" json:"misc_bonus,omitempty"`
}
