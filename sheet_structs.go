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
	SchemaVersion uint16 `yaml:"schema_version,omitempty"`

	Basics       CharacterBasicStats `yaml:"basics"`
	Attributes   AttributeStats      `yaml:"attributes"`
	SavingThrows SavingThrowStats    `yaml:"saving_throws"`
	Skills       SkillStats          `yaml:"skills"`
	Health       HealthStats         `yaml:"health"`

	// Combat contains AC, initiative, and speed values.
	Combat CombatStats `yaml:"combat,omitempty"`

	// Spellcasting is omitted for non-spellcasters.
	Spellcasting SpellcastingStats `yaml:"spellcasting,omitempty"`
}


/*
========================
BASIC CHARACTER METADATA
========================

CharacterBasicStats represents high-level identity and build metadata.

These fields are considered "authored input":
- they should normally be provided by the user
- they are not derived from other values
- validation should ensure they are internally consistent
*/
type CharacterBasicStats struct {
	// Name is the character's display name.
	// This is free-form and not validated beyond non-emptiness.
	Name string `yaml:"name"`

	// Age represents the character's age in years.
	// Optional: omitted if unknown or irrelevant.
	// Stored as uint16 to keep it human-scaled while allowing omission.
	Age uint16 `yaml:"age,omitempty"`

	// Class is the character's primary class name (e.g., "fighter").
	// NOTE: This is intentionally a string for now.
	// Multiclassing should eventually move this into a slice structure.
	Class string `yaml:"class"`

	// Level is the character's total level.
	// Must be >= 1 and <= system-defined maximum (usually 20).
	Level uint8 `yaml:"level"`

	// Race is the character's ancestry/species.
	// Free-form string, validated against known races if strict mode is enabled.
	Race string `yaml:"race"`

	// Alignment is optional narrative metadata.
	// It is not mechanically enforced by Ledgr.
	Alignment string `yaml:"alignment,omitempty"`
}


/*
================
ABILITY SCORES
================

Ability scores are core mechanical inputs.

Only Score is truly required input.
Modifier is typically derived from Score, but may be supplied
explicitly for verification, override, or legacy reasons.

Key design rule:
- If Modifier is present, Ledgr may validate or overwrite it.
- If Modifier is nil, Ledgr must compute it.
*/
type SingleAttribute struct {
	// Score is the raw ability score (e.g., 16).
	// Must be >= 1 and typically <= 30.
	Score uint8 `yaml:"score"`

	// Modifier is the derived ability modifier.
	// Pointer is used to distinguish:
	// - not provided (nil)
	// - explicitly provided (even if zero)
	Modifier *int8 `yaml:"modifier,omitempty"`
}

/*
AttributeStats represents the six standard D&D ability scores.

These are explicit fields (not a map) to guarantee:
- exactly six abilities
- predictable structure
- strong typing without stringly-typed keys
*/
type AttributeStats struct {
	Strength     SingleAttribute `yaml:"strength"`
	Dexterity    SingleAttribute `yaml:"dexterity"`
	Constitution SingleAttribute `yaml:"constitution"`
	Intelligence SingleAttribute `yaml:"intelligence"`
	Wisdom       SingleAttribute `yaml:"wisdom"`
	Charisma     SingleAttribute `yaml:"charisma"`
}


/*
================
SAVING THROWS
================

Saving throws depend on:
- the related ability modifier
- proficiency bonus (if proficient)
- optional expertise or miscellaneous bonuses

Only proficiency flags and misc bonuses should be authored.
Final Modifier is derived.
*/
type SingleSavingThrow struct {
	// Proficient indicates whether the character is proficient
	// in this saving throw.
	Proficient bool `yaml:"proficient,omitempty"`

	// Expertise is rare for saves but included for future-proofing
	// and homebrew compatibility.
	Expertise bool `yaml:"expertise,omitempty"`

	// MiscBonus represents bonuses from magic items, features,
	// or situational effects.
	MiscBonus int8 `yaml:"misc_bonus,omitempty"`

	// Modifier is the final computed saving throw bonus.
	// Pointer allows Ledgr to detect whether the user supplied it.
	Modifier *int8 `yaml:"modifier,omitempty"`
}

/*
SavingThrowStats explicitly defines all six saving throws.

As with attributes, explicit fields ensure completeness
and eliminate the need for map iteration or reflection.
*/
type SavingThrowStats struct {
	Strength     SingleSavingThrow `yaml:"strength"`
	Dexterity    SingleSavingThrow `yaml:"dexterity"`
	Constitution SingleSavingThrow `yaml:"constitution"`
	Intelligence SingleSavingThrow `yaml:"intelligence"`
	Wisdom       SingleSavingThrow `yaml:"wisdom"`
	Charisma     SingleSavingThrow `yaml:"charisma"`
}


/*
=========
SKILLS
=========

Skills are derived from:
- a related ability score
- proficiency / expertise
- misc bonuses

Only proficiency flags and misc bonuses should normally be authored.
*/
type SingleSkill struct {
	// Proficient indicates skill proficiency.
	Proficient bool `yaml:"proficient,omitempty"`

	// Expertise indicates doubled proficiency bonus.
	// Common for Rogues and Bards.
	Expertise bool `yaml:"expertise,omitempty"`

	// MiscBonus represents bonuses from features, items, or conditions.
	MiscBonus int8 `yaml:"misc_bonus,omitempty"`

	// RelatedAttribute allows overriding the default ability
	// used for this skill (useful for homebrew or special cases).
	RelatedAttribute string `yaml:"related_attribute,omitempty"`

	// Modifier is the final computed skill bonus.
	// Pointer distinguishes omitted vs user-provided values.
	Modifier *int8 `yaml:"modifier,omitempty"`
}

/*
SkillStats defines all standard D&D 5e skills.

Explicit fields guarantee:
- no missing skills
- consistent YAML structure
- schema stability
*/
type SkillStats struct {
	Acrobatics     SingleSkill `yaml:"acrobatics"`
	AnimalHandling SingleSkill `yaml:"animal_handling"`
	Arcana         SingleSkill `yaml:"arcana"`
	Athletics      SingleSkill `yaml:"athletics"`
	Deception      SingleSkill `yaml:"deception"`
	History        SingleSkill `yaml:"history"`
	Insight        SingleSkill `yaml:"insight"`
	Intimidation   SingleSkill `yaml:"intimidation"`
	Investigation  SingleSkill `yaml:"investigation"`
	Medicine       SingleSkill `yaml:"medicine"`
	Nature         SingleSkill `yaml:"nature"`
	Perception     SingleSkill `yaml:"perception"`
	Performance    SingleSkill `yaml:"performance"`
	Persuasion     SingleSkill `yaml:"persuasion"`
	Religion       SingleSkill `yaml:"religion"`
	SleightOfHand  SingleSkill `yaml:"sleight_of_hand"`
	Stealth        SingleSkill `yaml:"stealth"`
	Survival       SingleSkill `yaml:"survival"`
}


/*
========
HEALTH
========

HealthStats represents current and maximum hit points.

Current and Temp are optional to allow:
- build-only sheets
- out-of-combat snapshots
*/
type HealthStats struct {
	// Current is the character's current HP.
	Current uint16 `yaml:"current,omitempty"`

	// Max is the character's maximum HP.
	Max uint16 `yaml:"max,omitempty"`

	// Temp is temporary hit points.
	Temp uint16 `yaml:"temp,omitempty"`
}


/*
=========================
COMBAT / DEFENSE / MOVEMENT
=========================

CombatStats groups "sheet-facing" combat numbers that are commonly displayed together:
- Armor Class (AC)
- Initiative
- Speed(s)

These are intentionally separated from:
- Attributes (raw ability inputs)
- Skills/Saves (derived checks)
- Health (HP tracking)

Rationale:
- These values are often derived from multiple sources (gear, features, spells)
- They frequently need "misc bonus" hooks
- They are commonly overridden on sheets for convenience

Ledgr can choose to:
- compute these values when sufficient inputs exist, and/or
- accept user-provided totals and validate them when possible
*/


/*
SpeedStats models character movement.

In D&D 5e, "speed" is not always a single number:
- Most characters have a walking speed
- Some have additional movement modes: fly, swim, climb, burrow
- Effects can temporarily grant or modify these speeds

Design choices:
- Walk is the most common baseline and is included explicitly.
- Other movement modes are optional and omitted if not present.

Validation notes:
- Speeds should be non-negative.
- If a speed mode is not available, it should be omitted rather than set to 0
  (to keep YAML clean and semantically meaningful).
*/
type SpeedStats struct {
	Walk   uint16 `yaml:"walk,omitempty"`
	Fly    uint16 `yaml:"fly,omitempty"`
	Swim   uint16 `yaml:"swim,omitempty"`
	Climb  uint16 `yaml:"climb,omitempty"`
	Burrow uint16 `yaml:"burrow,omitempty"`
}


/*
InitiativeStats models initiative.

In default 5e rules:
- Initiative total = Dexterity modifier + misc bonuses (rare but possible)

Design choices:
- MiscBonus is authored input and supports items/features (e.g., Alert feat, magic).
- Total is derived in normal play but can be user-supplied for verification/override.
*/
type InitiativeStats struct {
	MiscBonus int8  `yaml:"misc_bonus,omitempty"`
	Total     *int8 `yaml:"total,omitempty"`
}


/*
ArmorClassStats models Armor Class.

In 5e, AC can be derived from:
- Armor type (light/medium/heavy), Dex modifier caps
- Shields
- Class features (Unarmored Defense)
- Spells (Mage Armor, Shield)
- Magic item bonuses
- Situational effects

For a starter schema:
- We allow a user-provided Total AC for simplicity.
- We include MiscBonus as a hook for bonuses that are not otherwise modeled.
- Later versions could add breakdown sources (armor name, shield, dex cap, etc.).
*/
type ArmorClassStats struct {
	MiscBonus int8   `yaml:"misc_bonus,omitempty"`
	Total     *uint8 `yaml:"total,omitempty"`
}


/*
CombatStats is the parent struct for combat-facing stats.
All fields are optional to support "build-only" character files
that may not include combat values yet.
*/
type CombatStats struct {
	ArmorClass ArmorClassStats `yaml:"armor_class,omitempty"`
	Initiative InitiativeStats `yaml:"initiative,omitempty"`
	Speed      SpeedStats      `yaml:"speed,omitempty"`
}


/*
========================
ROOT CHARACTER 

/*
===================
SPELLCASTING STATS
===================

SpellcastingStats applies only to spellcasting characters.

All numeric fields are derived in normal play, but optional
for verification or override purposes.
*/
type SpellcastingStats struct {
	// SpellcastingAbility is the ability used for spellcasting
	// (e.g., "intelligence", "wisdom", "charisma").
	SpellcastingAbility string `yaml:"spellcasting_ability,omitempty"`

	// SpellSaveDC is the DC for spell saving throws.
	SpellSaveDC *int8 `yaml:"spell_save_dc,omitempty"`

	// SpellAttackBonus is the bonus to spell attack rolls.
	SpellAttackBonus *int8 `yaml:"spell_attack_bonus,omitempty"`

	// MiscBonus represents bonuses from items or features
	// affecting spellcasting.
	MiscBonus int8 `yaml:"misc_bonus,omitempty"`
}

