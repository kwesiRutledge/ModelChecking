/*
finiteexecutionfragment.go
Description:
	Implementation of functions for the Finite Execution Fragment object which is described in Principles of Model Checking.
*/

package sequences

import (
	"fmt"
	"testing"

	mc "github.com/kwesiRutledge/ModelChecking"
)

func TestFiniteExecutionFragment_Check1(t *testing.T) {
	// Create Simple Transition System
	ts0 := mc.GetSimpleTS1()

	s1 := ts0.S[0]
	a1 := ts0.Act[0]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []mc.TransitionSystemState{s1, s1},
		a: []string{a1, a1},
	}

	err := e0.Check()

	if err.Error() != fmt.Sprintf("The number of states in the sequence is %v and so the sequence of actions should have length %v, but len(fe.a) = %v.", len(e0.s), len(e0.s)-1, len(e0.a)) {
		t.Errorf("The value of err was unexpected: %v", err)
	}

}

/*
TestFiniteExectutionFragment_Check2
Description:
	Checks to see whether or not there exists a transition for all states in the sequence given the provided actions.
*/
func TestFiniteExecutionFragment_Check2(t *testing.T) {
	// Create Simple Transition System
	ts0 := mc.GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	a1 := ts0.Act[0]
	//a2 := ts0.Act[1]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []mc.TransitionSystemState{s1, s2},
		a: []string{a1},
	}

	err := e0.Check()

	if err == nil {
		t.Errorf("The check function did not find anything wrong!")
	} else {
		if err.Error() != fmt.Sprintf("There is an invalid transition between state 1 and state 2 with action 1. (i.e. 2 not in Transition[1][1]).") {
			t.Errorf("The value of err was unexpected: %v", err)
		}
	}

}

/*
TestFiniteExectutionFragment_Check3
Description:
	Checks to see whether or not there exists a transition for all states in the sequence given the provided actions.
*/
func TestFiniteExecutionFragment_Check3(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	a1 := ts0.Act[0]
	//a2 := ts0.Act[1]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []TransitionSystemState{s1, s1, s2},
		a: []string{a1, a1},
	}

	err := e0.Check()

	if err == nil {
		t.Errorf("The check function did not find anything wrong!")
	} else {
		if err.Error() != fmt.Sprintf("There is an invalid transition between state 1 and state 2 with action 1. (i.e. 2 not in Transition[1][1]).") {
			t.Errorf("The value of err was unexpected: %v", err)
		}
	}

}

/*
TestFiniteExectutionFragment_Check4
Description:
	Checks to see whether or not there exists a transition for all states in the sequence given the provided actions.
	Correct execution is defined.
*/
func TestFiniteExecutionFragment_Check4(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	a1 := ts0.Act[0]
	//a2 := ts0.Act[1]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []TransitionSystemState{s1, s1},
		a: []string{a1},
	}

	err := e0.Check()

	if err != nil {
		t.Errorf("The check function found an error! %v", err)
	}

}

/*
TestFiniteExectutionFragment_Check5
Description:
	Checks to see whether or not there exists a transition for all states in the sequence given the provided actions.
	Good execution is given.
*/
func TestFiniteExecutionFragment_Check5(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	s3 := ts0.S[2]
	a1 := ts0.Act[0]
	a2 := ts0.Act[1]
	//a2 := ts0.Act[1]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []TransitionSystemState{s1, s2, s2, s3},
		a: []string{a2, a1, a2},
	}

	err := e0.Check()

	if err != nil {
		t.Errorf("The check function found an error! %v", err)
	}

}

/*
TestFiniteExectutionFragment_IsMaximal1
Description:
	Checks to see whether or not the final state in the execution is maximal.
	The execution is maximal.
*/
func TestFiniteExectutionFragment_IsMaximal1(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS2()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	s3 := ts0.S[2]
	s4 := ts0.S[3]
	a1 := ts0.Act[0]
	a2 := ts0.Act[1]
	//a2 := ts0.Act[1]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []TransitionSystemState{s1, s2, s2, s3, s4},
		a: []string{a2, a1, a2, a2},
	}

	maximalFlag, err := e0.IsMaximal()
	if err != nil {
		t.Errorf("There was an error running IsMaximal(): %v", err.Error())
	}

	if !maximalFlag {
		t.Errorf("The function IsMaximal() did not correctly identify that the execution is maximal.")
	}

}

/*
TestFiniteExectutionFragment_IsMaximal2
Description:
	Checks whether or not the execution is maximal or not.
	The input execution is NOT maximal.
*/
func TestFiniteExectutionFragment_IsMaximal2(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	s3 := ts0.S[2]
	a1 := ts0.Act[0]
	a2 := ts0.Act[1]
	//a2 := ts0.Act[1]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []TransitionSystemState{s1, s2, s2, s3},
		a: []string{a2, a1, a2},
	}

	maximalFlag, err := e0.IsMaximal()
	if err != nil {
		t.Errorf("There was an error running IsMaximal(): %v", err.Error())
	}

	if maximalFlag {
		t.Errorf("The function IsMaximal() did not correctly identify that the execution is NOT maximal.")
	}

}

/*
TestFiniteExectutionFragment_IsInitial1
Description:
	Checks whether or not the execution is initial or not.
	The input execution is NOT initial.
*/
func TestFiniteExectutionFragment_IsInitial1(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s2 := ts0.S[1]
	s3 := ts0.S[2]
	a1 := ts0.Act[0]
	a2 := ts0.Act[1]
	//a2 := ts0.Act[1]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []TransitionSystemState{s2, s2, s3},
		a: []string{a1, a2},
	}

	initialFlag, err := e0.IsInitial()
	if err != nil {
		t.Errorf("There was an error running IsInitial(): %v", err.Error())
	}

	if initialFlag {
		t.Errorf("The function IsInitial() did not correctly identify that the execution is NOT initial.")
	}

}

/*
TestFiniteExectutionFragment_IsInitial2
Description:
	Checks whether or not the execution is initial or not.
	The input execution is NOT initial.
*/
func TestFiniteExectutionFragment_IsInitial2(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	s3 := ts0.S[2]
	a1 := ts0.Act[0]
	a2 := ts0.Act[1]
	//a2 := ts0.Act[1]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []TransitionSystemState{s1, s2, s2, s3},
		a: []string{a2, a1, a2},
	}

	initialFlag, err := e0.IsInitial()
	if err != nil {
		t.Errorf("There was an error running IsInitial(): %v", err.Error())
	}

	if !initialFlag {
		t.Errorf("The function IsInitial() did not correctly identify that the execution is initial.")
	}

}
