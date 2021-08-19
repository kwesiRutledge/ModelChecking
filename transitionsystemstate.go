/*
transitionsystemstate.go
Description:
 	Basic implementation of a Transition System's state.
*/
package modelchecking

import (
	"errors"
	"fmt"
)

type TransitionSystemState struct {
	Name   string
	System *TransitionSystem
}

func (s1 TransitionSystemState) Equals(s2 TransitionSystemState) bool {
	return s1.Name == s2.Name
}

/*
In
Description:
	Determines if the state is in a given slice of TransitionSystemState objects.
Usage:
	tf := s1.In( sliceIn )
*/
func (stateIn TransitionSystemState) In(stateSliceIn []TransitionSystemState) bool {

	for _, tempState := range stateSliceIn {
		if tempState.Equals(stateIn) {
			return true //If there is a match, then return true.
		}
	}

	//If there is no match in the slice,
	//then return false
	return false

}

/*
Satisfies
Description:
	The state of the transition system satisfies the given formula.
*/

func (stateIn TransitionSystemState) Satisfies(formula interface{}) (bool, error) {

	// Input Processing
	if stateIn.System == nil {
		return false, errors.New("The system pointer is not defined for the input state.")
	}

	// Algorithm
	var tf = false
	var err error = nil

	if singleAP, ok := formula.(AtomicProposition); ok {
		tf, err = singleAP.SatisfactionHelper(stateIn)
	}

	return tf, err
}

/*
AppendIfUnique
Description:
	Appends to the input slice sliceIn if and only if the new state
	is actually a unique state.
*/
func AppendIfUnique(sliceIn []TransitionSystemState, stateIn TransitionSystemState) []TransitionSystemState {
	// Check to see if the State is equal to any of the ones in the list.
	for _, stateFromSlice := range sliceIn {
		if stateFromSlice.Equals(stateIn) {
			return sliceIn
		}
	}

	// If none of the states in sliceIn are equal to stateIn, then return the appended version of sliceIn.
	return append(sliceIn, stateIn)

}

/*
Post
Description:
	Finds the set of states that can follow a given state (or set of states).
Usage:

*/

func Post(SorSA ...interface{}) ([]TransitionSystemState, error) {
	switch len(SorSA) {
	case 1:
		// Only State Is Given
		stateIn, ok := SorSA[0].(TransitionSystemState)
		if !ok {
			return []TransitionSystemState{}, errors.New("The first input to post is not of type TransitionSystemState.")
		}

		System := stateIn.System
		allActions := System.Act

		var nextStates []TransitionSystemState
		var tempPost []TransitionSystemState
		var err error
		for _, action := range allActions {
			tempPost, err = Post(stateIn, action)
			if err != nil {
				return []TransitionSystemState{}, err
			}
			for _, postElt := range tempPost {
				nextStates = AppendIfUnique(nextStates, postElt)
			}
		}

		return nextStates, nil

	case 2:
		// State and Action is Given
		stateIn, ok := SorSA[0].(TransitionSystemState)
		if !ok {
			return []TransitionSystemState{}, errors.New("The first input to post is not of type TransitionSystemState.")
		}

		actionIn, ok := SorSA[1].(string)
		if !ok {
			return []TransitionSystemState{}, errors.New("The second input to post is not of type string!")
		}

		// Get Transition value
		System := stateIn.System
		tValues := System.Transition[stateIn][actionIn]
		var nextStates []TransitionSystemState
		for _, nextState := range tValues {
			nextStates = AppendIfUnique(nextStates, nextState)
		}

		return nextStates, nil
	}

	// Return error
	return []TransitionSystemState{}, errors.New(fmt.Sprintf("Unexpected number of inputs to post (%v).", len(SorSA)))
}

/*
Pre
Description:
	Finds the set of states that can precede a given state (or set of states).
Usage:

*/

func Pre(SorSA ...interface{}) ([]TransitionSystemState, error) {
	switch len(SorSA) {
	case 1:
		// Only State Is Given
		stateIn, ok := SorSA[0].(TransitionSystemState)
		if !ok {
			return []TransitionSystemState{}, errors.New("The first input to post is not of type TransitionSystemState.")
		}

		System := stateIn.System
		allActions := System.Act

		var predecessors []TransitionSystemState
		var tempPre []TransitionSystemState
		var err error
		for _, action := range allActions {
			tempPre, err = Pre(stateIn, action)
			if err != nil {
				return []TransitionSystemState{}, err
			}
			for _, preElt := range tempPre {
				predecessors = AppendIfUnique(predecessors, preElt)
			}
		}

		return predecessors, nil

	case 2:
		// State and Action is Given
		stateIn, ok := SorSA[0].(TransitionSystemState)
		if !ok {
			return []TransitionSystemState{}, errors.New("The first input to post is not of type TransitionSystemState.")
		}

		actionIn, ok := SorSA[1].(string)
		if !ok {
			return []TransitionSystemState{}, errors.New("The second input to post is not of type string!")
		}

		// Get Transition value
		System := stateIn.System
		var matchingPredecessors []TransitionSystemState
		for predecessor, actionMap := range System.Transition {
			if stateIn.In(actionMap[actionIn]) {
				// If the target state is in the result of (predecessor,actionIn) -> stateIn,
				// then save the value of stateIn
				matchingPredecessors = append(matchingPredecessors, predecessor)
			}
		}

		return matchingPredecessors, nil
	}

	// Return error
	return []TransitionSystemState{}, errors.New(fmt.Sprintf("Unexpected number of inputs to post (%v).", len(SorSA)))
}

/*
IsTerminal
Description:
	Identifies if the given state is terminal. (i.e. if Post(stateIn) is empty)
*/
func (stateIn TransitionSystemState) IsTerminal() bool {
	// Construct the Ancestor States
	ancestors, _ := Post(stateIn)

	return len(ancestors) == 0
}
