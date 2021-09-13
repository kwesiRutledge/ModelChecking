/*
path_test.go
Description:
	Tests for the objects and functions defined in path.go
*/
package sequences

import (
	"strings"
	"testing"

	mc "github.com/kwesiRutledge/ModelChecking"
)

/*
TestPath_Check1
Description:
	Tests whether or not a finite path fragment with a bad transition is identified as invalid.
*/
func TestPath_Check1(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[0]},
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
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	// Test Fragment
	if fpf0.Check() != nil {
		t.Errorf("The Check did not identify that there is an invalid transition in path. %v", fpf0.Check())
	}
}

/*
TestPath_Check3
Description:
	Tests whether or not an infinite path fragment with all good transitions is identified as valid.
*/
func TestPath_Check3(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	fpf1 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[3]},
	}

	ipf0 := InfinitePathFragment{
		UniquePrefix:    fpf0,
		RepeatingSuffix: fpf1,
	}

	// Test Fragment
	if ipf0.Check() != nil {
		t.Errorf("The Check did not identify that this is a valid infinite path fragment. %v", fpf0.Check())
	}
}

/*
TestPath_Check4
Description:
	Checks an infinite path fragment where the prefix is bad.
*/
func TestPath_Check4(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[1]},
	}

	fpf1 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[3]},
	}

	ipf0 := InfinitePathFragment{
		UniquePrefix:    fpf0,
		RepeatingSuffix: fpf1,
	}

	// Test Fragment
	err := ipf0.Check()
	if !strings.Contains(err.Error(), "There was an issue while checking the prefix of the path fragment:") {
		t.Errorf("The Check did not identify that this is an invalid infinite path fragment. %v", err)
	}
}

/*
TestPath_Check5
Description:
	Checks an infinite path fragment where the suffix is bad (contains bad transition in its sequence).
*/
func TestPath_Check5(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	fpf1 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[0]},
	}

	ipf0 := InfinitePathFragment{
		UniquePrefix:    fpf0,
		RepeatingSuffix: fpf1,
	}

	// Test Fragment
	err := ipf0.Check()
	if !strings.Contains(err.Error(), "There was an issue while checking the suffix of the path fragment:") {
		t.Errorf("The Check did not identify that this is an invalid infinite path fragment. %v", err)
	}
}

/*
TestPath_Check6
Description:
	Checks an infinite path fragment where the transition from
	prefix to suffix is bad.
*/
func TestPath_Check6(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	fpf1 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[1], ts0.S[3], ts0.S[0]},
	}

	ipf0 := InfinitePathFragment{
		UniquePrefix:    fpf0,
		RepeatingSuffix: fpf1,
	}

	// Test Fragment
	err := ipf0.Check()
	if !strings.Contains(err.Error(), "The first state in the suffix \"select\" was not an ancestor of the last state in the prefix \"beer\".") {
		t.Errorf("The Check did not identify that this is an invalid infinite path fragment. %v", err)
	}
}

/*
TestPath_Check7
Description:
	Checks an infinite path fragment where the transition from
	prefix to suffix is bad.
*/
func TestPath_Check7(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	fpf1 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[3], ts0.S[0]},
	}

	ipf0 := InfinitePathFragment{
		UniquePrefix:    fpf0,
		RepeatingSuffix: fpf1,
	}

	// Test Fragment
	err := ipf0.Check()
	if !strings.Contains(err.Error(), "The first state in the suffix \"pay\" was not an ancestor of the last state in the suffix \"pay\".") {
		t.Errorf("The Check did not identify that this is an invalid infinite path fragment. %v", err)
	}
}

/*
TestPath_IsMaximal1
Description:
	Checks to see if a finitepathfragment which is known to be not maximal is maximal.
*/
func TestPath_IsMaximal1(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	// Test
	if fpf0.IsMaximal() {
		t.Errorf("The FinitePathFragment is expected to NOT BE maximal, but the function claims it is.")
	}
}

/*
TestPath_IsMaximal2
Description:
	Checks to see if a finitepathfragment which is known to be maximal is maximal.
*/
func TestPath_IsMaximal2(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.TransitionSystem_GetSimpleTS3()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2], ts0.S[3]},
	}

	// Test
	if fpf0.Check() != nil {
		t.Errorf("fpf0 failed the check! %v", fpf0.Check())
	}

	if !fpf0.IsMaximal() {
		t.Errorf("The FinitePathFragment is expected to BE maximal, but the function claims it is not.")
	}
}

/*
TestPath_IsMaximal3
Description:
	Checks to see if an infinitepathfragment is maximal.
*/
func TestPath_IsMaximal3(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	fpf1 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[3]},
	}

	ipf0 := InfinitePathFragment{
		UniquePrefix:    fpf0,
		RepeatingSuffix: fpf1,
	}

	// Test Fragment
	if ipf0.Check() != nil {
		t.Errorf("ipf0 failed the check! %v", ipf0.Check())
	}

	if !ipf0.IsMaximal() {
		t.Errorf("ipf0 is maximal, but IsMaximal() claims it is not!")
	}

}

/*
TestPath_IsInitial1
Description:
	Checks to see if a finitepathfragment which is known to be initial is initial.
*/
func TestPath_IsInitial1(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	// Test
	if !fpf0.IsInitial() {
		t.Errorf("The FinitePathFragment is expected to BE initial, but the function claims it is not.")
	}
}

/*
TestPath_IsInitial2
Description:
	Checks to see if a finitepathfragment which is known to be not initial is initial.
*/
func TestPath_IsInitial2(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[1], ts0.S[2], ts0.S[0]},
	}

	// Test
	if fpf0.IsInitial() {
		t.Errorf("The FinitePathFragment is expected to NOT BE initial, but the function claims it is.")
	}
}

/*
TestPath_IsInitial3
Description:
	Checks to see if an infinitepathfragment is initial.
*/
func TestPath_IsInitial3(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	fpf1 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[3]},
	}

	ipf0 := InfinitePathFragment{
		UniquePrefix:    fpf0,
		RepeatingSuffix: fpf1,
	}

	// Test Fragment
	if ipf0.Check() != nil {
		t.Errorf("ipf0 failed the check! %v", ipf0.Check())
	}

	if !ipf0.IsInitial() {
		t.Errorf("ipf0 is initial, but IsInitial() claims it is not!")
	}

}

/*
TestPath_IsInitial4
Description:
	Checks to see if an infinitepathfragment is initial.
*/
func TestPath_IsInitial4(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[1], ts0.S[2]},
	}

	fpf1 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[3]},
	}

	ipf0 := InfinitePathFragment{
		UniquePrefix:    fpf0,
		RepeatingSuffix: fpf1,
	}

	// Test Fragment
	if ipf0.Check() != nil {
		t.Errorf("ipf0 failed the check! %v", ipf0.Check())
	}

	if ipf0.IsInitial() {
		t.Errorf("ipf0 is initial, but IsInitial() claims it is not!")
	}

}

/*
TestPath_IsPath1
Description:
	Checks to see if a finitepathfragment which is known to be initial but not maximal, is a path.
	Answer should be no.
*/
func TestPath_IsPath1(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	// Test
	if IsPath(fpf0) {
		t.Errorf("The FinitePathFragment is expected to BE initial, but the function claims it is not.")
	}
}

/*
TestPath_IsPath2
Description:
	Checks to see if a finitepathfragment which is known to be initial but not maximal, is a path.
	Answer should be no.
*/
func TestPath_IsPath2(t *testing.T) {
	// Create an example FinitePathFragment object
	ts0 := mc.GetBeverageVendingMachineTS()

	fpf0 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[2]},
	}

	fpf1 := FinitePathFragment{
		s: []mc.TransitionSystemState{ts0.S[0], ts0.S[1], ts0.S[3]},
	}

	ipf0 := InfinitePathFragment{
		UniquePrefix:    fpf0,
		RepeatingSuffix: fpf1,
	}

	// Test
	if !IsPath(ipf0) {
		t.Errorf("The InfinitePathFragment is expected to be a path, as it is both initial and maximal, but the function claims it is not.")
	}
}
