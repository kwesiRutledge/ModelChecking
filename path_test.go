/*
path_test.go
Description:
	Tests for the objects and functions defined in path.go
*/
package modelchecking

import "testing"

/*
TestPath_Check1
Description:
	Tests whether or not a finite path fragment with a bad transition is identified as invalid.
*/
func TestPath_Check1(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[0]},
	}

	// Test Fragment
	if fpf0.Check() == nil {
		t.Errorf("The Check did not identify that there is an invalid transition in path. %v", fpf0.Check())
	}
}

/*
TestPath_Check2
Description:
	Tests whether or not a finite path fragment with all good transitions is identified as valid.
*/
func TestPath_Check2(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	// Test Fragment
	if fpf0.Check() != nil {
		t.Errorf("The Check did not identify that there is an invalid transition in path. %v", fpf0.Check())
	}
}
