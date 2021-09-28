/*
transitionsystem.go
Description:
	Defines a Transition System according to the definition provided in the paper
	'Formal Methods for Adaptive Control of Dynamical Systems' by Sadra Sadraddini and Calin Belta.
*/

package adaptive

import (
	"errors"
	"fmt"

	mc "github.com/kwesiRutledge/ModelChecking"
)

type TransitionSystem struct {
	X          []TransitionSystemState
	U          []string
	Transition map[TransitionSystemState]map[string][]TransitionSystemState
	Pi         []mc.AtomicProposition
	O          map[TransitionSystemState][]mc.AtomicProposition
}

/*
GetTransitionSystem
Description:

*/
func GetTransitionSystem(stateNames []string, actionNames []string, transitionMap map[string]map[string][]string, atomicPropositionsList []string, labelMap map[string][]string) (TransitionSystem, error) {
	// Constants

	// Algorithm
	ts := TransitionSystem{
		U:  actionNames,
		Pi: mc.StringSliceToAPs(atomicPropositionsList),
	}

	// Create List of States
	var X []TransitionSystemState //List Of States
	for _, stateName := range stateNames {
		X = append(X, TransitionSystemState{Name: stateName, System: &ts})
	}
	ts.X = X

	// // Create List of Atomic Propositions
	// ts.AP = StringSliceToAPs(atomicPropositionsList)

	// Create Transition Map
	Transition := make(map[TransitionSystemState]map[string][]TransitionSystemState)
	for siName, perStateMap := range transitionMap {
		si := TransitionSystemState{Name: siName, System: &ts}
		tempActionMap := make(map[string][]TransitionSystemState)
		for actionName, stateArray := range perStateMap {
			var tempStates []TransitionSystemState
			for _, siPlus1Name := range stateArray {
				tempStates = append(tempStates, TransitionSystemState{Name: siPlus1Name, System: &ts})
			}
			tempActionMap[actionName] = tempStates
		}
		Transition[si] = tempActionMap
	}
	ts.Transition = Transition

	err := ts.CheckTransition()
	if err != nil {
		return ts, err
	}

	// Create Label Values
	fullLabelMap := make(map[TransitionSystemState][]mc.AtomicProposition)
	for stateName, associatedAPs := range labelMap {
		tempState := TransitionSystemState{Name: stateName, System: &ts}
		fullLabelMap[tempState] = mc.StringSliceToAPs(associatedAPs)
	}
	ts.O = fullLabelMap

	return ts, nil
}

/*
CheckTransition
Description:
	Checks that all of the transition states are correct.
*/
func (ts TransitionSystem) CheckTransition() error {
	// Checks that all source states are from the state set.
	for state1 := range ts.Transition {
		if !state1.In(ts.X) {
			return fmt.Errorf("One of the source states in the Transition was not in the state set: %v", state1)
		}
	}

	// Checks that all input states are from the action set
	for _, actionMap := range ts.Transition {
		for tempAction, targetStates := range actionMap {
			if _, foundInAct := mc.FindInSlice(tempAction, ts.U); !foundInAct {
				return fmt.Errorf("The action \"%v\" was found in the transition map but is not in the action set!", tempAction)
			}
			// Search through all target states to see if they are in the state set
			for _, targetState := range targetStates {
				if !targetState.In(ts.X) {
					return fmt.Errorf("There is an ancestor state \"%v\" which is not part of the state.", targetState)
				}
			}
		}
	}

	// Completed with no errors
	return nil
}

/*
Check
Description:
	Checks the following components of the transition system object:
	- Initial State Set is a Subset of the State Space
	- Transitions map properly between the state space and the action space to the next state space
*/
func (ts TransitionSystem) Check() error {

	// Check Transition Map
	err := ts.CheckTransition()
	if err != nil {
		return err
	}

	return nil
}

/*
IsNonBlocking
Description:
	Determines if the transition system is non-blocking, that is for any state and action pair the value of Post is not zero.
*/
func (ts TransitionSystem) IsNonBlocking() bool {

	// Check the value of Post for each state, action pair.
	for _, state := range ts.X {
		for _, action := range ts.U {
			tempPostValue, _ := Post(state, action)
			if len(tempPostValue) == 0 {
				return false
			}
		}
	}

	return true
}

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

	if singleAP, ok := formula.(mc.AtomicProposition); ok {

		System := stateIn.System
		return singleAP.In(System.O[stateIn]), nil
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
		allActions := System.U

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
		allActions := System.U

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
String
Description:
	Returns the value of the state if the value is of type string.
	Otherwise, this returns the value of each part of the Value
*/
func (stateIn TransitionSystemState) String() string {
	return stateIn.Name
}

/*
HasStateSpacePartition
Description:
	Determines if a collection of subsets of states (slice of slices of states)
	forms a partition of the state space for a transition system.
	According to the paper Q is a partition if:
	- Empty set is not in Q
	- Union of the subsets in Q makes X
	- There is no overlap of sets in Q
*/
func (ts TransitionSystem) HasStateSpacePartition(Q [][]TransitionSystemState) bool {
	// Verify that empty set is not in Q
	for _, tempSubset := range Q {
		if len(tempSubset) == 0 {
			return false // Q contains the empty set. Return false.
		}
	}

	// Verify that the union of all of these sets in Q
	// make up the state space when considered in total.
	unionOfQ := UnionOfStates(Q...) // Unrolls Q
	tf, _ := SliceEquals(unionOfQ, ts.X)
	if !tf {
		return false // union of Q does not make X
	}

	// Verify that for each pair, the two sets are disjoint
	for subsetIndex1, subset1 := range Q {
		for subsetIndex2 := subsetIndex1 + 1; subsetIndex2 < len(Q); subsetIndex2++ {
			subset2 := Q[subsetIndex2]
			subsetsAreDisjoint := len(IntersectionOfStates(subset1, subset2)) == 0
			if !subsetsAreDisjoint {
				return false
			}
		}
	}

	return true
}

/*
UnionOfStates
Description:
	Computes the union of a collection of TransitionSystemState slices.
*/
func UnionOfStates(stateSlicesIn ...[]TransitionSystemState) []TransitionSystemState {
	// Constants
	//numSubsets := len(stateSlicesIn)

	// Algorithm
	var unionOut []TransitionSystemState
	for _, tempSubset := range stateSlicesIn {
		for _, tempState := range tempSubset {
			unionOut = tempState.AppendIfUniqueTo(unionOut)
		}
	}

	return unionOut

}

/*
IntersectionOfStates
Description:
	Computes the intersection of a collection of TransitionSystemState slices.
*/
func IntersectionOfStates(stateSlicesIn ...[]TransitionSystemState) []TransitionSystemState {
	// Constants
	//numSubsets := len(stateSlicesIn)

	// Algorithm
	var intersectionOut []TransitionSystemState
	if len(stateSlicesIn) == 0 {
		return []TransitionSystemState{}
	} else {
		slice0 := stateSlicesIn[0]

		for _, xi := range slice0 {
			// Search through all other slices
			var xiExistsInAll = true

			for _, sliceI := range stateSlicesIn {
				if !xi.In(sliceI) {
					xiExistsInAll = false
				}
			}

			if xiExistsInAll {
				intersectionOut = append(intersectionOut, xi)
			}
		}

		return intersectionOut
	}

}
