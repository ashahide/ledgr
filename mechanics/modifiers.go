package mechanics

import (
	"math"
	"errors"
)

type ModifierNumeric interface {
	~int | ~uint
}

func CalcAbilityModifier[T ModifierNumeric](stat T) (modifier int32, err error) {

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
	modifier = int32(math.Floor((statAsFloat - 10.0) / 2.0))

	return modifier, nil
}
