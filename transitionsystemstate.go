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

/*
TransitionSystemState
Description:
	This type is an object which contains the Transition System's State.
*/
type TransitionSystemState struct {
	Name   string
	System *TransitionSystem
}

/*
Equals
Description:
	Checks to see if two states in the transition system have the same value.
*/
func (stateIn TransitionSystemState) Equals(s2 TransitionSystemState) bool {
	return stateIn.Name == s2.Name
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
AppendIfUniqueTo
Description:
	Appends to the input slice sliceIn if and only if the new state
	is actually a unique state.
Usage:
	newSlice := stateIn.AppendIfUniqueTo(sliceIn)
*/
func (stateIn TransitionSystemState) AppendIfUniqueTo(sliceIn []TransitionSystemState) []TransitionSystemState {

	for _, tempState := range sliceIn {
		if tempState.Equals(stateIn) {
			return sliceIn
		}

	}

	//If all checks were passed
	return append(sliceIn, stateIn)

}

/*
Post
Description:
	Finds the set of states that can follow a given state (or set of states).
Usage:
	tempPost, err := Post( s0 )
	tempPost, err := Post( s0 , a0 )
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
				nextStates = postElt.AppendIfUniqueTo(nextStates)
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
			nextStates = nextState.AppendIfUniqueTo(nextStates)
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
	tempPreSet , err := Pre( s0 )
	tempPreSet , err := Pre( s0 , a0 )
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
				predecessors = preElt.AppendIfUniqueTo(predecessors)
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

/*
IsReachable
Description:
	Identifies if the given state is reachable or not.
Usage:
	reachableFlag := state1.IsReachable()
*/
func (stateIn TransitionSystemState) IsReachable() bool {
	// Get Transition System
	System := stateIn.System

	// Check to see if state is initial.
	// If it is, then we are done.
	if stateIn.In(System.I) {
		return true
	}

	// Create loop variables
	SI := []TransitionSystemState{stateIn}
	SIm1 := []TransitionSystemState{}
	Reachable := SI

	// Create Loop Which Checks:
	//	- If a predecessor of stateIn is in System.I
	//	- Or if the predecessor set is equal to the
	for {

		// Compute the predecessors.
		for _, si := range SI {
			tempPre, _ := Pre(si)
			for _, sPre := range tempPre {
				Reachable = sPre.AppendIfUniqueTo(Reachable)
				SIm1 = sPre.AppendIfUniqueTo(SIm1)
			}
		}

		// Check to see if any predecessors are in the initial state set.
		for _, sIm1 := range SIm1 {
			if sIm1.In(System.I) {
				return true
			}
		}

		// Check if the predecessor set is a subset of the current set at i
		// If it is, then exit. And the state is not reachable
		if tf, _ := SliceSubset(SIm1, SI); tf {
			break
		}

		// Prepare for next iteration of loop
		SI = SIm1
		SIm1 = []TransitionSystemState{}
	}

	return false
}

/*
String
Description:
	Returns the value of the state if the value is of type string.
	Otherwise, this returns the value of each part of the Value
*/
func (stateIn TransitionSystemState) String() string {
	return stateIn.Name
}
