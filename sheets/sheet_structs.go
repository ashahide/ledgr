package sheets

/*
Package ledgr defines the canonical data model for a YAML-based
D&D 5e character sheet.

Design philosophy:

- YAML is treated as a *human-authored configuration format*
- Go structs are the *authoritative schema*
- Some fields are user-authored "truths"
- Some fields are derived and may be omitted or validated
*/

/*
========================
ROOT CHARACTER DOCUMENT
========================
*/

type CharacterSheet struct {
	SchemaVersion int `yaml:"schema_version" json:"schema_version"`

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
	Age       int `yaml:"age" json:"age"`
	Class     string `yaml:"class" json:"class"`
	Level     int  `yaml:"level" json:"level"`
	Race      string `yaml:"race" json:"race"`
	Alignment string `yaml:"alignment" json:"alignment"`
}

/*
================
ABILITY SCORES
================
*/

type Ability string

const (
	Strength     Ability = "strength"
	Dexterity    Ability = "dexterity"
	Constitution Ability = "constitution"
	Intelligence Ability = "intelligence"
	Wisdom       Ability = "wisdom"
	Charisma     Ability = "charisma"
)

var AllAbilities = []Ability{
	Strength,
	Dexterity,
	Constitution,
	Intelligence,
	Wisdom,
	Charisma,
}

type SingleAttribute struct {
	Score    int `yaml:"score" json:"score"`
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

func (a *AttributeStats) GetModifier(ability Ability) int {
	switch ability {
	case Strength:
		return a.Strength.Modifier
	case Dexterity:
		return a.Dexterity.Modifier
	case Constitution:
		return a.Constitution.Modifier
	case Intelligence:
		return a.Intelligence.Modifier
	case Wisdom:
		return a.Wisdom.Modifier
	case Charisma:
		return a.Charisma.Modifier
	default:
		return -999 
	}
}

func (a *AttributeStats) SetModifier(ability Ability, mod int) {
	switch ability {
	case Strength:
		a.Strength.Modifier = mod
	case Dexterity:
		a.Dexterity.Modifier = mod
	case Constitution:
		a.Constitution.Modifier = mod
	case Intelligence:
		a.Intelligence.Modifier = mod
	case Wisdom:
		a.Wisdom.Modifier = mod
	case Charisma:
		a.Charisma.Modifier = mod
	}
}

func (a *AttributeStats) GetScore(ability Ability) int {
	switch ability {
	case Strength:
		return a.Strength.Score
	case Dexterity:
		return a.Dexterity.Score
	case Constitution:
		return a.Constitution.Score
	case Intelligence:
		return a.Intelligence.Score
	case Wisdom:
		return a.Wisdom.Score
	case Charisma:
		return a.Charisma.Score
	default:
		return -999
	}
}

func (a *AttributeStats) GetAttribute(ability Ability) *SingleAttribute {
	switch ability {
	case Strength:
		return &a.Strength
	case Dexterity:
		return &a.Dexterity
	case Constitution:
		return &a.Constitution
	case Intelligence:
		return &a.Intelligence
	case Wisdom:
		return &a.Wisdom
	case Charisma:
		return &a.Charisma
	default:
		return nil
	}
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


func (s *SavingThrowStats) GetModifier(ability Ability) int {
	switch ability {
	case Strength:
		return s.Strength.Modifier
	case Dexterity:
		return s.Dexterity.Modifier
	case Constitution:
		return s.Constitution.Modifier
	case Intelligence:
		return s.Intelligence.Modifier
	case Wisdom:
		return s.Wisdom.Modifier
	case Charisma:
		return s.Charisma.Modifier
	default:
		return -999
	}
}

func (s *SavingThrowStats) SetModifier(ability Ability, mod int) {
	switch ability {
	case Strength:
		s.Strength.Modifier = mod
	case Dexterity:
		s.Dexterity.Modifier = mod
	case Constitution:
		s.Constitution.Modifier = mod
	case Intelligence:
		s.Intelligence.Modifier = mod
	case Wisdom:
		s.Wisdom.Modifier = mod
	case Charisma:
		s.Charisma.Modifier = mod
	}
}

func (s *SavingThrowStats) GetSavingThrow(ability Ability) *SingleSavingThrow {
	switch ability {
	case Strength:
		return &s.Strength
	case Dexterity:
		return &s.Dexterity
	case Constitution:
		return &s.Constitution
	case Intelligence:
		return &s.Intelligence
	case Wisdom:
		return &s.Wisdom
	case Charisma:
		return &s.Charisma
	default:
		return nil
	}
}

/*
=========
SKILLS
=========
*/

type Skill string

const (
	Acrobatics     Skill = "acrobatics"
	AnimalHandling Skill = "animal_handling"
	Arcana         Skill = "arcana"
	Athletics      Skill = "athletics"
	Deception      Skill = "deception"
	History        Skill = "history"
	Insight        Skill = "insight"
	Intimidation   Skill = "intimidation"
	Investigation  Skill = "investigation"
	Medicine       Skill = "medicine"
	Nature         Skill = "nature"
	Perception     Skill = "perception"
	Performance    Skill = "performance"
	Persuasion     Skill = "persuasion"
	Religion       Skill = "religion"
	SleightOfHand  Skill = "sleight_of_hand"
	Stealth        Skill = "stealth"
	Survival       Skill = "survival"
)

var SkillToAbility = map[Skill]Ability{
	Acrobatics:     Dexterity,
	AnimalHandling: Wisdom,
	Arcana:         Intelligence,
	Athletics:      Strength,
	Deception:      Charisma,
	History:        Intelligence,
	Insight:        Wisdom,
	Intimidation:   Charisma,
	Investigation:  Intelligence,
	Medicine:       Wisdom,
	Nature:         Intelligence,
	Perception:     Wisdom,
	Performance:    Charisma,
	Persuasion:     Charisma,
	Religion:       Intelligence,
	SleightOfHand:  Dexterity,
	Stealth:        Dexterity,
	Survival:       Wisdom,
}



type SingleSkill struct {
	Proficient       bool   `yaml:"proficient" json:"proficient"`
	Expertise        bool   `yaml:"expertise" json:"expertise"`
	MiscBonus        int   `yaml:"misc_bonus" json:"misc_bonus"`
	Modifier         int  `yaml:"modifier" json:"modifier"`
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


func (s *SkillStats) GetModifier(skill Skill) int {
	switch skill {
	case Acrobatics:
		return s.Acrobatics.Modifier
	case AnimalHandling:
		return s.AnimalHandling.Modifier
	case Arcana:
		return s.Arcana.Modifier
	case Athletics:
		return s.Athletics.Modifier
	case Deception:
		return s.Deception.Modifier
	case History:
		return s.History.Modifier
	case Insight:
		return s.Insight.Modifier
	case Intimidation:
		return s.Intimidation.Modifier
	case Investigation:
		return s.Investigation.Modifier
	case Medicine:
		return s.Medicine.Modifier
	case Nature:
		return s.Nature.Modifier
	case Perception:
		return s.Perception.Modifier
	case Performance:
		return s.Performance.Modifier
	case Persuasion:
		return s.Persuasion.Modifier
	case Religion:
		return s.Religion.Modifier
	case SleightOfHand:
		return s.SleightOfHand.Modifier
	case Stealth:
		return s.Stealth.Modifier
	case Survival:
		return s.Survival.Modifier
	default:
		return -999
	}
}

func (s *SkillStats) SetModifier(skill Skill, mod int) {
	switch skill {
	case Acrobatics:
		s.Acrobatics.Modifier = mod
	case AnimalHandling:
		s.AnimalHandling.Modifier = mod
	case Arcana:
		s.Arcana.Modifier = mod
	case Athletics:
		s.Athletics.Modifier = mod
	case Deception:
		s.Deception.Modifier = mod
	case History:
		s.History.Modifier = mod
	case Insight:
		s.Insight.Modifier = mod
	case Intimidation:
		s.Intimidation.Modifier = mod
	case Investigation:
		s.Investigation.Modifier = mod
	case Medicine:
		s.Medicine.Modifier = mod
	case Nature:
		s.Nature.Modifier = mod
	case Perception:
		s.Perception.Modifier = mod
	case Performance:
		s.Performance.Modifier = mod
	case Persuasion:
		s.Persuasion.Modifier = mod
	case Religion:
		s.Religion.Modifier = mod
	case SleightOfHand:
		s.SleightOfHand.Modifier = mod
	case Stealth:
		s.Stealth.Modifier = mod
	case Survival:
		s.Survival.Modifier = mod
	}
}


func (s *SkillStats) GetSkill(skill Skill) *SingleSkill {
	switch skill {
	case Acrobatics:
		return &s.Acrobatics
	case AnimalHandling:
		return &s.AnimalHandling
	case Arcana:
		return &s.Arcana
	case Athletics:
		return &s.Athletics
	case Deception:
		return &s.Deception
	case History:
		return &s.History
	case Insight:
		return &s.Insight
	case Intimidation:
		return &s.Intimidation
	case Investigation:
		return &s.Investigation
	case Medicine:
		return &s.Medicine
	case Nature:
		return &s.Nature
	case Perception:
		return &s.Perception
	case Performance:
		return &s.Performance
	case Persuasion:
		return &s.Persuasion
	case Religion:
		return &s.Religion
	case SleightOfHand:
		return &s.SleightOfHand
	case Stealth:
		return &s.Stealth
	case Survival:
		return &s.Survival
	default:
		return nil
	}
}


/*
========
HEALTH
========
*/

type HealthStats struct {
	Current int `yaml:"current" json:"current"`
	Max     int `yaml:"max" json:"max"`
	Temp    int `yaml:"temp" json:"temp"`
}

/*
=========================
COMBAT / DEFENSE / MOVEMENT
=========================
*/

type SpeedStats struct {
	Walk   int `yaml:"walk" json:"walk"`
	Fly    int `yaml:"fly" json:"fly"`
	Swim   int `yaml:"swim" json:"swim"`
	Climb  int `yaml:"climb" json:"climb"`
	Burrow int `yaml:"burrow" json:"burrow"`
}

type InitiativeStats struct {
	MiscBonus int  `yaml:"misc_bonus" json:"misc_bonus"`
	Total     int `yaml:"total" json:"total"`
}

type ArmorClassStats struct {
	MiscBonus int   `yaml:"misc_bonus" json:"misc_bonus"`
	Total     *int `yaml:"total" json:"total"`
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
	SpellSaveDC         int  `yaml:"spell_save_dc" json:"spell_save_dc"`
	SpellAttackBonus    int  `yaml:"spell_attack_bonus" json:"spell_attack_bonus"`
	MiscBonus           int   `yaml:"misc_bonus" json:"misc_bonus"`
}
