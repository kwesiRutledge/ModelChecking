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
				si.Name, sip1.Name, ai, sip1.Name, si.Name, ai,
			)

		}
		fmt.Sprintf("%v not in Transition[%v][%v]", sip1.Name, si.Name, ai)
	}

	fmt.Sprintf("%v", len(fe.s)-1)

	// Everything is okay!
	return nil
}