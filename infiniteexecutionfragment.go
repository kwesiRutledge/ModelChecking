/*
infiniteexecutionfragment.go
Description:
	Implementation of functions for the Finite Execution Fragment object which is described in Principles of Model Checking.
*/

package modelchecking

import "fmt"

type InfiniteExecutionFragment struct {
	UniquePrefix    FiniteExecutionFragment
	RepeatingSuffix FiniteExecutionFragment
}

/*
Check
Description:
	Check to make sure that the execution fragment is valid.
*/
func (ief InfiniteExecutionFragment) Check() error {
	// Check that the length of each slice is appropriate for the prefix and then the suffix
	pref := ief.UniquePrefix
	if len(pref.s) != len(pref.a) {
		return fmt.Errorf("The number of states in ief.UniquePrefix is supposed to be equal to the number of actions, but there are %v states and %v actions.", len(pref.s), len(pref.a))
	}
	suffix := ief.RepeatingSuffix
	if len(suffix.s) != len(suffix.a) {
		return fmt.Errorf("The number of states in ief.RepeatingSuffix is supposed to be equal to the number of actions, but there are %v states and %v actions.", len(pref.s), len(pref.a))
	}

	system := pref.s[0].System

	// Verify that the UniquePrefix of the transitions are okay
	for sIndex := 0; sIndex < len(pref.s)-1; sIndex++ {
		si := pref.s[sIndex]
		ai := pref.a[sIndex]
		sip1 := pref.s[sIndex+1]

		if !sip1.In(system.Transition[si][ai]) {
			return fmt.Errorf(
				"There is an invalid transition in UniquePrefix between state %v and state %v with action %v. (i.e. %v not in Transition[%v][%v]).",
				si.Value, sip1.Value, ai, sip1.Value, si.Value, ai,
			)

		}
	}

	prefixFinalState := pref.s[len(pref.s)-1]
	prefixFinalAction := pref.a[len(pref.a)-1]
	if !suffix.s[0].In(system.Transition[prefixFinalState][prefixFinalAction]) {
		return fmt.Errorf(
			"The transition from the prefix to the suffix was invalid! (i.e. %v not in Transition[%v][%v]).",
			suffix.s[0].Value, prefixFinalState.Value, prefixFinalAction,
		)
	}

	// Verify that the RepeatingSuffix of the transitions are okay
	for sIndex := 0; sIndex < len(suffix.s)-1; sIndex++ {
		si := suffix.s[sIndex]
		ai := suffix.a[sIndex]
		sip1 := suffix.s[sIndex+1]

		if !sip1.In(system.Transition[si][ai]) {
			return fmt.Errorf(
				"There is an invalid transition in RepeatingSuffix between state %v and state %v with action %v. (i.e. %v not in Transition[%v][%v]).",
				si.Value, sip1.Value, ai, sip1.Value, si.Value, ai,
			)

		}
	}

	suffixFinalState := suffix.s[len(pref.s)-1]
	suffixFinalAction := suffix.a[len(pref.a)-1]
	if !suffix.s[0].In(system.Transition[suffixFinalState][suffixFinalAction]) {
		return fmt.Errorf(
			"The transition from the suffix end to the suffix beginning was invalid! (i.e. %v not in Transition[%v][%v]).",
			suffix.s[0].Value, suffixFinalState.Value, suffixFinalAction,
		)
	}

	// Everything is okay!
	return nil
}

/*
IsMaximal
Description:
	Determines whether or not the infinite execution fragment is maximal.
*/
func (ief InfiniteExecutionFragment) IsMaximal() (bool, error) {
	// Check Execution Fragment
	err := ief.Check()
	if err != nil {
		return false, err
	}

	// Return true
	return true, nil
}
