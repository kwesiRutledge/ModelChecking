/*
transitionsystem.go
Description:
 	Basic implementation of a Transition System.
*/
package main

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
		S = append(S, TransitionSystemState{Name: stateName, System: &ts})
	}
	ts.S = S

	// Create List of Initial States
	var I []TransitionSystemState
	for _, stateName := range initialStateList {
		I = append(I, TransitionSystemState{Name: stateName, System: &ts})
	}
	ts.I = I

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

	// Create Label Values
	fullLabelMap := make(map[TransitionSystemState][]AtomicProposition)
	for stateValue, associatedAPs := range labelMap {
		tempState := TransitionSystemState{Name: stateValue, System: &ts}
		fullLabelMap[tempState] = StringSliceToAPs(associatedAPs)
	}
	ts.L = fullLabelMap

	return ts, nil
}

func (ap AtomicProposition) SatisfactionHelper(stateIn TransitionSystemState) (bool, error) {
	// Constants

	SystemPointer := stateIn.System
	LOfState := SystemPointer.L[stateIn]

	// Find If ap is in LOfState
	tf := false
	for _, tempAP := range LOfState {
		tf = tf || tempAP.Equals(ap)
	}

	return tf, nil

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
