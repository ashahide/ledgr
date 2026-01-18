package sheets

import "testing"

func TestCalcAbilityModifier(t *testing.T) {
	ans, err := CalcAbilityModifier(10)

	if err != nil {
		t.Errorf("CalcAbilityModifier failed during testing")
	}
	
	if ans != 0 {
		t.Errorf("The modifier for a score of 10 should be 0")
	}
}
