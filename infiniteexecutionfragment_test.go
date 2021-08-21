/*
infiniteexecutionfragment.go
Description:
	Implementation of functions for the Finite Execution Fragment object which is described in Principles of Model Checking.
*/

package modelchecking

import (
	"fmt"
	"testing"
)

func TestInfiniteExecutionFragment_Check1(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	a1 := ts0.Act[0]

	// Execution
	fef0 := FiniteExecutionFragment{
		s: []TransitionSystemState{s1, s1},
		a: []string{a1},
	}

	ief0 := InfiniteExecutionFragment{
		UniquePrefix:    fef0,
		RepeatingSuffix: fef0,
	}

	err := ief0.Check()

	if err == nil {
		t.Errorf("The check did not properly catch an error!")
	} else {
		if err.Error() != fmt.Sprintf("The number of states in ief.UniquePrefix is supposed to be equal to the number of actions, but there are %v states and %v actions.", len(fef0.s), len(fef0.a)) {
			t.Errorf("The value of err was unexpected: %v", err)
		}
	}

}

/*
TestInfiniteExecutionFragment_Check2
Description:
	Checks to verify that the Check() function correctly identifies errors in the size of the RepeatingSuffix.
*/
func TestInfiniteExecutionFragment_Check2(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	a1 := ts0.Act[0]

	// Execution

	ief0 := InfiniteExecutionFragment{
		UniquePrefix: FiniteExecutionFragment{
			s: []TransitionSystemState{s1, s1},
			a: []string{a1, a1},
		},
		RepeatingSuffix: FiniteExecutionFragment{
			s: []TransitionSystemState{s1, s1},
			a: []string{a1},
		},
	}

	err := ief0.Check()

	if err == nil {
		t.Errorf("The check did not properly catch an error!")
	} else {
		if err.Error() != fmt.Sprintf("The number of states in ief.RepeatingSuffix is supposed to be equal to the number of actions, but there are %v states and %v actions.", len(ief0.RepeatingSuffix.s), len(ief0.RepeatingSuffix.s)) {
			t.Errorf("The value of err was unexpected: %v", err)
		}
	}

}

/*
TestInfiniteExecutionFragment_Check3
Description:
	Checks to verify that the Check() function correctly identifies errors in the transitions of uniquePrefix.
*/
func TestInfiniteExecutionFragment_Check3(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	a1 := ts0.Act[0]

	// Execution

	ief0 := InfiniteExecutionFragment{
		UniquePrefix: FiniteExecutionFragment{
			s: []TransitionSystemState{s1, s2},
			a: []string{a1, a1},
		},
		RepeatingSuffix: FiniteExecutionFragment{
			s: []TransitionSystemState{s1, s1},
			a: []string{a1, a1},
		},
	}

	err := ief0.Check()

	if err == nil {
		t.Errorf("The check did not properly catch an error!")
	} else {
		if err.Error() != fmt.Sprintf(
			"There is an invalid transition in UniquePrefix between state %v and state %v with action %v. (i.e. %v not in Transition[%v][%v]).",
			s1.Name, s2.Name, a1, s2.Name, s1.Name, a1,
		) {
			t.Errorf("The value of err was unexpected: %v", err)
		}
	}

}

/*
TestInfiniteExecutionFragment_Check4
Description:
	Checks to verify that the Check() function correctly identifies errors in the transition between
	uniquePrefix and the suffix.
*/
func TestInfiniteExecutionFragment_Check4(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	// s2 := ts0.S[1]
	a1 := ts0.Act[0]
	a2 := ts0.Act[1]

	// Execution

	ief0 := InfiniteExecutionFragment{
		UniquePrefix: FiniteExecutionFragment{
			s: []TransitionSystemState{s1, s1},
			a: []string{a1, a2},
		},
		RepeatingSuffix: FiniteExecutionFragment{
			s: []TransitionSystemState{s1, s1},
			a: []string{a1, a1},
		},
	}

	err := ief0.Check()

	if err == nil {
		t.Errorf("The check did not properly catch an error!")
	} else {
		if err.Error() != fmt.Sprintf(
			"The transition from the prefix to the suffix was invalid! (i.e. %v not in Transition[%v][%v]).",
			s1.Name, s1.Name, a2,
		) {
			t.Errorf("The value of err was unexpected: %v", err)
		}
	}

}

/*
TestInfiniteExecutionFragment_Check5
Description:
	Checks to verify that the Check() function correctly identifies errors in the transitions of
	the suffix.
*/
func TestInfiniteExecutionFragment_Check5(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	a1 := ts0.Act[0]
	a2 := ts0.Act[1]

	// Execution

	ief0 := InfiniteExecutionFragment{
		UniquePrefix: FiniteExecutionFragment{
			s: []TransitionSystemState{s1, s1},
			a: []string{a1, a2},
		},
		RepeatingSuffix: FiniteExecutionFragment{
			s: []TransitionSystemState{s2, s1},
			a: []string{a2, a1},
		},
	}

	err := ief0.Check()

	if err == nil {
		t.Errorf("The check did not properly catch an error!")
	} else {
		if err.Error() != fmt.Sprintf(
			"There is an invalid transition in RepeatingSuffix between state %v and state %v with action %v. (i.e. %v not in Transition[%v][%v]).",
			s2.Name, s1.Name, a2, s1.Name, s2.Name, a2,
		) {
			t.Errorf("The value of err was unexpected: %v", err)
		}
	}

}

/*
TestInfiniteExecutionFragment_Check6
Description:
	Checks to verify that the Check() function correctly identifies errors in the transition between
	the end of the suffix and the beginning suffix.
*/
func TestInfiniteExecutionFragment_Check6(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	a1 := ts0.Act[0]
	a2 := ts0.Act[1]

	// Execution

	ief0 := InfiniteExecutionFragment{
		UniquePrefix: FiniteExecutionFragment{
			s: []TransitionSystemState{s1, s1},
			a: []string{a1, a2},
		},
		RepeatingSuffix: FiniteExecutionFragment{
			s: []TransitionSystemState{s2, s1},
			a: []string{a1, a1},
		},
	}

	err := ief0.Check()

	if err == nil {
		t.Errorf("The check did not properly catch an error!")
	} else {
		if err.Error() != fmt.Sprintf(
			"The transition from the suffix end to the suffix beginning was invalid! (i.e. %v not in Transition[%v][%v]).",
			s2.Name, s1.Name, a1,
		) {
			t.Errorf("The value of err was unexpected: %v", err)
		}
	}

}

/*
TestInfiniteExecutionFragment_Check7
Description:
	Checks to verify that the Check() function correctly identifies a good infinite execution.
*/
func TestInfiniteExecutionFragment_Check7(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	a1 := ts0.Act[0]
	a2 := ts0.Act[1]

	// Execution

	ief0 := InfiniteExecutionFragment{
		UniquePrefix: FiniteExecutionFragment{
			s: []TransitionSystemState{s1, s1},
			a: []string{a1, a2},
		},
		RepeatingSuffix: FiniteExecutionFragment{
			s: []TransitionSystemState{s2, s1},
			a: []string{a1, a2},
		},
	}

	err := ief0.Check()

	if err != nil {
		t.Errorf("The value of err was not nil! %v", err)
	}

}

/*
TestInfiniteExecutionFragment_IsMaximal1
Description:
	Most valid Infinite Execution Fragments are maximal.
	Given a valid infinite execution fragments, the function should correctly identify that the ief is maximal.
*/
func TestInfiniteExecutionFragment_IsMaximal1(t *testing.T) {
	// Create Simple Transition System
	ts0 := GetSimpleTS1()

	s1 := ts0.S[0]
	s2 := ts0.S[1]
	a1 := ts0.Act[0]
	a2 := ts0.Act[1]

	// Execution

	ief0 := InfiniteExecutionFragment{
		UniquePrefix: FiniteExecutionFragment{
			s: []TransitionSystemState{s1, s1},
			a: []string{a1, a2},
		},
		RepeatingSuffix: FiniteExecutionFragment{
			s: []TransitionSystemState{s2, s1},
			a: []string{a1, a2},
		},
	}

	maximalFlag, err := ief0.IsMaximal()
	if err != nil {
		t.Errorf("There was an issue checking if ief0 was maximal: %v", err.Error())
	}

	if !maximalFlag {
		t.Errorf("The function did not identify that ief0 is maximal!")
	}

}
