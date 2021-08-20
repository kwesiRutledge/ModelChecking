/*
finiteexecutionfragment.go
Description:
	Implementation of functions for the Finite Execution Fragment object which is described in Principles of Model Checking.
*/

package modelchecking

import (
	"fmt"
	"testing"
)

func TestFiniteExecutionFragment_Check1(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	a1 := ts0.Act[0]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []TransitionSystemState{s1, s1},
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
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	a1 := ts0.Act[0]
	//a2 := ts0.Act[1]

	// Execution
	e0 := FiniteExecutionFragment{
		s: []TransitionSystemState{s1, s2},
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
