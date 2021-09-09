/*
finiteexecutionfragment.go
Description:
	Implementation of functions for the Finite Execution Fragment object which is described in Principles of Model Checking.
*/

package modelchecking

import "fmt"

type FiniteExecutionFragment struct {
	s []TransitionSystemState
	a []string
}

/*
Check
Description:
	Check to make sure that the execution fragment is valid.
*/
func (fe FiniteExecutionFragment) Check() error {
	// Check that the length of each slice is appropriate.
	if len(fe.s) != (len(fe.a) + 1) {
		return fmt.Errorf("The number of states in the sequence is %v and so the sequence of actions should have length %v, but len(fe.a) = %v.", len(fe.s), len(fe.s)-1, len(fe.a))
	}

	system := fe.s[0].System

	// Verify that all of the transitions are okay
	for sIndex := 0; sIndex < len(fe.s)-1; sIndex++ {
		si := fe.s[sIndex]
		ai := fe.a[sIndex]
		sip1 := fe.s[sIndex+1]

		if !sip1.In(system.Transition[si][ai]) {
			return fmt.Errorf(
				"There is an invalid transition between state %v and state %v with action %v. (i.e. %v not in Transition[%v][%v]).",
				si, sip1, ai, sip1, si, ai,
			)

		}
	}

	// Everything is okay!
	return nil
}

/*
IsMaximal
Description:
	Identifies if the finite execution fragment is maximal. This is the case if
	the final state is terminal.
*/
func (fe FiniteExecutionFragment) IsMaximal() (bool, error) {
	// Check Execution Fragment
	err := fe.Check()
	if err != nil {
		return false, err
	}

	// Verify that the last state is terminal.
	n := len(fe.s)
	finalState := fe.s[n-1]

	return finalState.IsTerminal(), nil
}

/*
IsInitial
Description:
	Identifies if the finite execution fragment is initial. This is the case if
	the first state is from the initial state set of the transition system.
*/
func (fe FiniteExecutionFragment) IsInitial() (bool, error) {
	// Check Execution Fragment
	err := fe.Check()
	if err != nil {
		return false, err
	}

	// Verify that the last state is terminal.
	firstState := fe.s[0]
	System := firstState.System

	return firstState.In(System.I), nil
}
