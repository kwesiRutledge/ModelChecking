/*
transitionsystem.go
Description:
 	Basic implementation of a Transition System.
*/
package modelchecking

import (
	"fmt"
)

type TransitionSystem struct {
	S          []TransitionSystemState
	Act        []string
	Transition map[TransitionSystemState]map[string][]TransitionSystemState
	I          []TransitionSystemState // Set of Initial States
	AP         []AtomicProposition
	L          map[TransitionSystemState][]AtomicProposition
}

/*
GetTransitionSystem
Description:

*/
func GetTransitionSystem(stateNames []string, actionNames []string, transitionMap map[string]map[string][]string, initialStateList []string, atomicPropositionsList []string, labelMap map[string][]string) (TransitionSystem, error) {
	// Constants

	// Algorithm
	ts := TransitionSystem{
		Act: actionNames,
		AP:  StringSliceToAPs(atomicPropositionsList),
	}

	// Create List of States
	var S []TransitionSystemState //List Of States
	for _, stateName := range stateNames {
		S = append(S, TransitionSystemState{Value: stateName, System: &ts})
	}
	ts.S = S

	// Create List of Initial States
	var I []TransitionSystemState
	for _, stateName := range initialStateList {
		I = append(I, TransitionSystemState{Value: stateName, System: &ts})
	}
	ts.I = I

	err := ts.CheckI()
	if err != nil {
		return ts, err
	}

	// // Create List of Atomic Propositions
	// ts.AP = StringSliceToAPs(atomicPropositionsList)

	// Create Transition Map
	Transition := make(map[TransitionSystemState]map[string][]TransitionSystemState)
	for siName, perStateMap := range transitionMap {
		si := TransitionSystemState{Value: siName, System: &ts}
		tempActionMap := make(map[string][]TransitionSystemState)
		for actionName, stateArray := range perStateMap {
			var tempStates []TransitionSystemState
			for _, siPlus1Name := range stateArray {
				tempStates = append(tempStates, TransitionSystemState{Value: siPlus1Name, System: &ts})
			}
			tempActionMap[actionName] = tempStates
		}
		Transition[si] = tempActionMap
	}
	ts.Transition = Transition

	err = ts.CheckTransition()
	if err != nil {
		return ts, err
	}

	// Create Label Values
	fullLabelMap := make(map[TransitionSystemState][]AtomicProposition)
	for stateValue, associatedAPs := range labelMap {
		tempState := TransitionSystemState{Value: stateValue, System: &ts}
		fullLabelMap[tempState] = StringSliceToAPs(associatedAPs)
	}
	ts.L = fullLabelMap

	return ts, nil
}

/*
CheckI
Description:
	Checks that all of the states in the initial state set are from the state set S.
*/
func (ts TransitionSystem) CheckI() error {
	for _, Istate := range ts.I {
		if !Istate.In(ts.S) {
			return fmt.Errorf("The state %v is not in the state set of the transition system!", Istate.Value)
		}
	}
	// If we finish checking I,
	// then all states were satisfied
	return nil
}

/*
CheckTransition
Description:
	Checks that all of the transition states are correct.
*/
func (ts TransitionSystem) CheckTransition() error {
	// Checks that all source states are from the state set.
	for state1 := range ts.Transition {
		if !state1.In(ts.S) {
			return fmt.Errorf("One of the source states in the Transition was not in the state set: %v", state1.Value)
		}
	}

	// Checks that all input states are from the action set
	for _, actionMap := range ts.Transition {
		for tempAction, targetStates := range actionMap {
			if _, foundInAct := FindInSlice(tempAction, ts.Act); !foundInAct {
				return fmt.Errorf("The action \"%v\" was found in the transition map but is not in the action set!", tempAction)
			}
			// Search through all target states to see if they are in the state set
			for _, targetState := range targetStates {
				if !targetState.In(ts.S) {
					return fmt.Errorf("There is an ancestor state \"%v\" which is not part of the state.", targetState.Value)
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
	// Check Initial state Set
	err := ts.CheckI()
	if err != nil {
		return err
	}

	// Check Transition Map
	err = ts.CheckTransition()
	if err != nil {
		return err
	}

	return nil
}

/*
IsActionDeterministic
Description:
	Determines if the transition system is action deterministic:
	- Initial state set contains at most one element.
	- Post(s,a) contains one or fewer elements for all states s and actions a
*/
func (ts TransitionSystem) IsActionDeterministic() bool {

	// Check the Initial State Set
	if !(len(ts.I) <= 1) {
		return false
	}

	// Check the value of Post for each state, action pair.
	for _, state := range ts.S {
		for _, action := range ts.Act {
			tempPostValue, _ := Post(state, action)
			if !(len(tempPostValue) <= 1) {
				return false
			}
		}
	}

	return true
}

/*
IsAPDeterministic
Description:
	Determines if the transition system is action deterministic:
	- Initial state set contains at most one element.
	- Post(s,a) contains elements with unique labels (or none?), for all states s and actions a
*/
func (ts TransitionSystem) IsAPDeterministic() bool {

	// Check the Initial State Set
	if !(len(ts.I) <= 1) {
		return false
	}

	// Check the value of Post for each state, action pair.
	tempPowerset := Powerset(ts.AP)

	for _, state := range ts.S {
		for _, powersetElement := range tempPowerset {
			// Collect all elements of Post(state) that have label of powersetElement
			tempPostValue, _ := Post(state)
			// find all elements of tempPostValue with Label
			var postWithLabelPE []TransitionSystemState
			for _, postState := range tempPostValue {
				LOfPostState := ts.L[postState]
				equalsFlag, _ := SliceEquals(LOfPostState, powersetElement)
				if equalsFlag {
					postWithLabelPE = append(postWithLabelPE, postState)
				}
			}

			if !(len(postWithLabelPE) <= 1) {
				return false
			}
		}
	}

	return true
}

func (ts TransitionSystem) Interleave(ts2 TransitionSystem) (TransitionSystem, error) {
	// Check the Component Transition Systems
	err1 := ts.Check()
	err2 := ts2.Check()
	if (err1 != nil) || (err2 != nil) {
		return TransitionSystem{}, fmt.Errorf("There was an issue checking ts and ts2.\n From ts: %v\n From ts2: %v", err1, err2)
	}

	// Create initial interleaved TS
	var interleavedTS TransitionSystem

	// Create A The State Space from The Cartesian Product
	tempCartesianProduct, err := SliceCartesianProduct(ts.S, ts2.S)
	if err != nil {
		return interleavedTS, err
	}
	tempCPAsStateSlice := tempCartesianProduct.([][]TransitionSystemState)

	var S []TransitionSystemState
	for _, tempTuple := range tempCPAsStateSlice {
		S = append(S,
			TransitionSystemState{
				Value:  tempTuple,
				System: &interleavedTS,
			},
		)
	}
	interleavedTS.S = S

	// Work on Action Set

	return interleavedTS, nil
}
