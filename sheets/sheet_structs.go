package sheets

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
	SchemaVersion uint16 `yaml:"schema_version" json:"schema_version"`

	Basics       CharacterBasicStats `yaml:"basics" json:"basics"`
	Attributes   AttributeStats      `yaml:"attributes" json:"attributes"`
	SavingThrows SavingThrowStats    `yaml:"saving_throws" json:"saving_throws"`
	Skills       SkillStats          `yaml:"skills" json:"skills"`
	Health       HealthStats         `yaml:"health" json:"health"`

	// Combat contains AC, initiative, and speed values.
	Combat CombatStats `yaml:"combat" json:"combat"`

	// Spellcasting is omitted for non-spellcasters.
	Spellcasting SpellcastingStats `yaml:"spellcasting" json:"spellcasting"`
}

/*
========================
BASIC CHARACTER METADATA
========================
*/

type CharacterBasicStats struct {
	Name      string `yaml:"name" json:"name"`
	Age       uint16 `yaml:"age" json:"age"`
	Class     string `yaml:"class" json:"class"`
	Level     uint  `yaml:"level" json:"level"`
	Race      string `yaml:"race" json:"race"`
	Alignment string `yaml:"alignment" json:"alignment"`
}

/*
================
ABILITY SCORES
================
*/

type SingleAttribute struct {
	Score    uint `yaml:"score" json:"score"`
	Modifier int `yaml:"modifier" json:"modifier"`
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
	Proficient bool  `yaml:"proficient" json:"proficient"`
	Expertise  bool  `yaml:"expertise" json:"expertise"`
	MiscBonus  int  `yaml:"misc_bonus" json:"misc_bonus"`
	Modifier   int `yaml:"modifier" json:"modifier"`
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
	Proficient       bool   `yaml:"proficient" json:"proficient"`
	Expertise        bool   `yaml:"expertise" json:"expertise"`
	MiscBonus        int   `yaml:"misc_bonus" json:"misc_bonus"`
	Modifier         *int  `yaml:"modifier" json:"modifier"`
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
	Current uint `yaml:"current" json:"current"`
	Max     uint `yaml:"max" json:"max"`
	Temp    uint `yaml:"temp" json:"temp"`
}

/*
=========================
COMBAT / DEFENSE / MOVEMENT
=========================
*/

type SpeedStats struct {
	Walk   uint `yaml:"walk" json:"walk"`
	Fly    uint `yaml:"fly" json:"fly"`
	Swim   uint `yaml:"swim" json:"swim"`
	Climb  uint `yaml:"climb" json:"climb"`
	Burrow uint `yaml:"burrow" json:"burrow"`
}

type InitiativeStats struct {
	MiscBonus int  `yaml:"misc_bonus" json:"misc_bonus"`
	Total     *int `yaml:"total" json:"total"`
}

type ArmorClassStats struct {
	MiscBonus int   `yaml:"misc_bonus" json:"misc_bonus"`
	Total     *uint `yaml:"total" json:"total"`
}

type CombatStats struct {
	ArmorClass ArmorClassStats `yaml:"armor_class" json:"armor_class"`
	Initiative InitiativeStats `yaml:"initiative" json:"initiative"`
	Speed      SpeedStats      `yaml:"speed" json:"speed"`
}

/*
===================
SPELLCASTING STATS
===================
*/

type SpellcastingStats struct {
	SpellcastingAbility string `yaml:"spellcasting_ability" json:"spellcasting_ability"`
	SpellSaveDC         *int  `yaml:"spell_save_dc" json:"spell_save_dc"`
	SpellAttackBonus    *int  `yaml:"spell_attack_bonus" json:"spell_attack_bonus"`
	MiscBonus           int   `yaml:"misc_bonus" json:"misc_bonus"`
}
