//package sequences
/*
path.go
Description:
	Objects which are finite path fragments.
*/
package sequences

import (
	"fmt"

	mc "github.com/kwesiRutledge/ModelChecking"
)

/*
Type Declarations
*/

type FinitePathFragment struct {
	s []mc.TransitionSystemState
}

type InfinitePathFragment struct {
	UniquePrefix    FinitePathFragment
	RepeatingSuffix FinitePathFragment
}

type PathFragment interface {
	Check() error
	IsMaximal() bool
	IsInitial() bool
}

// Functions

func (fragmentIn FinitePathFragment) Check() error {
	// Verify that the transitions in the path fragment are okay
	for sIndex := 0; sIndex < len(fragmentIn.s)-1; sIndex++ {
		si := fragmentIn.s[sIndex]
		sip1 := fragmentIn.s[sIndex+1]

		siAncestors, err := mc.Post(si)
		if err != nil {
			return fmt.Errorf("There was an issue computing the %vth post (Post(%v)): %v", sIndex, si, err)
		}

		if !sip1.In(siAncestors) {
			return fmt.Errorf(
				"The %vth state (%v) is not in the post of the %vth state (%v).",
				sIndex+1, sip1, sIndex, si,
			)
		}
	}

	// Return nothing if all transitions are correct
	return nil
}

/*
Check
Description:
	For the InfinitePathFragment object, this checks that:
	- The sequences within UniquePrefix and RepeatingSuffix are independently valid
	- The transition from UniquePrefix to RepeatingSuffix is valid
	- The transition from the end of RepeatingSuffix to the beginning of RepeatingSuffix is valid
*/
func (fragmentIn InfinitePathFragment) Check() error {
	// Make sure that the suffix contains at least one element
	if len(fragmentIn.RepeatingSuffix.s) == 0 {
		return fmt.Errorf("The RepeatingSuffix value has length 0. If this is a FinitePathFragment, then use that object instead!")
	}

	//Verify that Both the prefix and suffix are valid by themselves
	err := fragmentIn.UniquePrefix.Check()
	if err != nil {
		return fmt.Errorf("There was an issue while checking the prefix of the path fragment: %v", err)
	}

	err = fragmentIn.RepeatingSuffix.Check()
	if err != nil {
		return fmt.Errorf("There was an issue while checking the suffix of the path fragment: %v", err)
	}

	//Check that the first state in the suffix is in the Post of the last state from the prefix
	lastStateInPrefix := fragmentIn.UniquePrefix.s[len(fragmentIn.UniquePrefix.s)-1]
	firstStateInSuffix := fragmentIn.RepeatingSuffix.s[0]
	ancestorsOfLastState, err := mc.Post(lastStateInPrefix)
	if err != nil {
		return fmt.Errorf("There was an error computing the Post of the last state in the prefix: %v", err)
	}

	if !firstStateInSuffix.In(ancestorsOfLastState) {
		return fmt.Errorf("The first state in the suffix \"%v\" was not an ancestor of the last state in the prefix \"%v\".", firstStateInSuffix, lastStateInPrefix)
	}

	//Check that the first state in the suffix is in the Post of the last state in the suffix
	lastStateInSuffix := fragmentIn.RepeatingSuffix.s[len(fragmentIn.RepeatingSuffix.s)-1]
	ancestorsOfLastState, err = mc.Post(lastStateInSuffix)
	if err != nil {
		return fmt.Errorf("There was an error computing the Post of the last state in the suffix: %v", err)
	}

	if !firstStateInSuffix.In(ancestorsOfLastState) {
		return fmt.Errorf("The first state in the suffix \"%v\" was not an ancestor of the last state in the suffix \"%v\".", firstStateInSuffix, lastStateInSuffix)
	}

	// If you made it this far, then everything is fine.
	return nil
}

/*
IsMaximal
Description:
	For the FinitePathFragment, this is true if the last state is terminal.
Assumption:
	Assumes that you've already checked the FinitePathFragment.
*/
func (fragmentIn FinitePathFragment) IsMaximal() bool {
	//Grabs last state in path.
	finalState := fragmentIn.s[len(fragmentIn.s)-1]

	return finalState.IsTerminal()
}

/*
IsMaximal
Description:
	For the InfinitePathFragment, this is always true.
Assumption:
	This assumes that the InfinitePathFragment was already checked.
*/
func (fragmentIn InfinitePathFragment) IsMaximal() bool {
	return true
}

/*
IsInitial
Description
	For any type of execution, this is true if the first state of the execution is in the transition system's initial state set.
*/
func (fragmentIn FinitePathFragment) IsInitial() bool {
	initialState := fragmentIn.s[0]
	System := *&initialState.System

	return initialState.In(System.I)
}

func (fragmentIn InfinitePathFragment) IsInitial() bool {
	var initialState mc.TransitionSystemState
	if len(fragmentIn.UniquePrefix.s) > 0 {
		initialState = fragmentIn.UniquePrefix.s[0]
	} else {
		initialState = fragmentIn.RepeatingSuffix.s[0]
	}

	System := *&initialState.System

	return initialState.In(System.I)
}

/*
IsPath()
Description:
	This function should work for either FinitePathFragment or InitialPathFragment objects and will check to
	see if the given fragment is initial and maximal before returning true or false.
Assumption:
	This function assumes that fragmentIn has previously been checked.
*/
func IsPath(fragmentIn PathFragment) bool {
	return fragmentIn.IsInitial() && fragmentIn.IsMaximal()
}

/*
ToTrace
Description:
	Uses the labels of this transition system to compute the trace
	of a finite or infinite path.
Assumptions:
	This function assumes that you've run fragmentIn.Check() beforehand.
*/
func (fragmentIn FinitePathFragment) ToTrace() FiniteTrace {
	// Get System
	firstState := fragmentIn.s[0]
	ts := *&firstState.System

	var SequenceOfAPSubsets [][]mc.AtomicProposition
	for _, tempState := range fragmentIn.s {
		SequenceOfAPSubsets = append(SequenceOfAPSubsets, ts.L[tempState])
	}

	// Return final answer.
	return FiniteTrace{L: SequenceOfAPSubsets}

}

func (fragmentIn InfinitePathFragment) ToTrace() InfiniteTrace {
	return InfiniteTrace{
		UniquePrefix:    fragmentIn.UniquePrefix.ToTrace(),
		RepeatingSuffix: fragmentIn.RepeatingSuffix.ToTrace(),
	}
}
