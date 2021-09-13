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
		S = append(S, TransitionSystemState{Name: stateName, System: &ts})
	}
	ts.S = S

	// Create List of Initial States
	var I []TransitionSystemState
	for _, stateName := range initialStateList {
		I = append(I, TransitionSystemState{Name: stateName, System: &ts})
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

	err = ts.CheckTransition()
	if err != nil {
		return ts, err
	}

	// Create Label Values
	fullLabelMap := make(map[TransitionSystemState][]AtomicProposition)
	for stateName, associatedAPs := range labelMap {
		tempState := TransitionSystemState{Name: stateName, System: &ts}
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
			return fmt.Errorf("The state %v is not in the state set of the transition system!", Istate)
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
			return fmt.Errorf("One of the source states in the Transition was not in the state set: %v", state1)
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

	// Create The State Space from The Cartesian Product of these two state spaces
	var S []TransitionSystemState
	for _, s1 := range ts.S {
		for _, s2 := range ts2.S {
			S = append(S,
				TransitionSystemState{
					Name:   fmt.Sprintf("(%v,%v)", s1, s2),
					System: &interleavedTS,
				})
		}
	}
	interleavedTS.S = S

	// Work on Action Set
	var Act []string
	Act = AppendIfUnique(Act, ts.Act...)
	Act = AppendIfUnique(Act, ts2.Act...)
	interleavedTS.Act = Act

	// Work on Transition based on the rules
	//	1. First state s1 changes to s1_prime if s1_prime in Post(s1,a)

	Transition := make(map[TransitionSystemState]map[string][]TransitionSystemState)
	// Create transitions for the first system actions.
	for _, s1 := range ts.S {
		for _, s2 := range ts2.S {
			tempProductState := TransitionSystemState{
				Name:   fmt.Sprintf("(%v,%v)", s1, s2),
				System: &interleavedTS,
			}

			tempActionMap := make(map[string][]TransitionSystemState)
			for _, actionName := range interleavedTS.Act {
				// Check to see if actionName is from ts.
				if _, tf := FindInSlice(actionName, ts.Act); tf {
					// Find all ancestors under this action
					s1Ancestors, err := Post(s1, actionName)
					if err != nil {
						return interleavedTS, err
					}

					// Use ancestors to create tuple
					for _, s1Ancestor := range s1Ancestors {
						tempActionMap[actionName] = append(tempActionMap[actionName],
							TransitionSystemState{
								Name:   fmt.Sprintf("(%v,%v)", s1Ancestor, s2),
								System: &interleavedTS,
							})
					}
				}
				// Check to see if action Name is from ts2
				if _, tf := FindInSlice(actionName, ts2.Act); tf {
					// Find all ancestors under this action
					s2Ancestors, err := Post(s2, actionName)
					if err != nil {
						return interleavedTS, err
					}

					// Use ancestors to create tuple
					for _, s2Ancestor := range s2Ancestors {
						tempActionMap[actionName] = append(tempActionMap[actionName],
							TransitionSystemState{
								Name:   fmt.Sprintf("(%v,%v)", s1, s2Ancestor),
								System: &interleavedTS,
							},
						)
					}
				}
			}
			// Append map to the new transition
			Transition[tempProductState] = tempActionMap

		}
	}
	interleavedTS.Transition = Transition

	// Initial States
	var I []TransitionSystemState
	for _, iState1 := range ts.I {
		for _, iState2 := range ts2.I {
			I = append(I,
				TransitionSystemState{
					Name:   fmt.Sprintf("(%v,%v)", iState1, iState2),
					System: &interleavedTS,
				})
		}
	}
	interleavedTS.I = I

	// AP
	var AP []AtomicProposition
	AP = append(AP, ts.AP...)
	AP = append(AP, ts2.AP...)
	interleavedTS.AP = AP

	// Label Map
	L := make(map[TransitionSystemState][]AtomicProposition)
	for _, s1 := range ts.S {
		for _, s2 := range ts2.S {
			// Create product state.
			tempProductState := TransitionSystemState{
				Name:   fmt.Sprintf("(%v,%v)", s1, s2),
				System: &interleavedTS,
			}

			// Create label
			L[tempProductState] = append(ts.L[s1], ts2.L[s2]...)
		}
	}
	interleavedTS.L = L

	return interleavedTS, nil
}

/*
Examples
*/

/*
GetSimpleTS
Description:
	Get a transition system to test satisfies.
*/
func GetSimpleTS1() TransitionSystem {
	ts0, _ := GetTransitionSystem(
		[]string{"1", "2", "3"}, []string{"1", "2"},
		map[string]map[string][]string{
			"1": map[string][]string{
				"1": []string{"1"},
				"2": []string{"2"},
			},
			"2": map[string][]string{
				"1": []string{"1", "2"},
				"2": []string{"2", "3"},
			},
			"3": map[string][]string{
				"1": []string{"3"},
				"2": []string{"2"},
			},
		},
		[]string{"1"},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"C", "D"},
		},
	)

	return ts0
}

/*
GetSimpleTS2
Description:
	Get a transition system to test Pre and Post.
	It should have states that generates empty sets for Pre() and Post() respectively.
*/
func GetSimpleTS2() TransitionSystem {
	ts0, _ := GetTransitionSystem(
		[]string{"1", "2", "3", "4", "5"}, []string{"1", "2"},
		map[string]map[string][]string{
			"1": map[string][]string{
				"1": []string{"1"},
				"2": []string{"2"},
			},
			"2": map[string][]string{
				"1": []string{"1", "2", "4"},
				"2": []string{"2", "3"},
			},
			"3": map[string][]string{
				"1": []string{"3"},
				"2": []string{"2", "4"},
			},
			"4": map[string][]string{
				"1": []string{},
				"2": []string{},
			},
			"5": map[string][]string{
				"1": []string{"4"},
				"2": []string{"1", "2", "3"},
			},
		},
		[]string{"1"},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"C", "D"},
			"4": []string{"A", "C"},
			"5": []string{"A", "B", "C", "D"},
		},
	)

	return ts0
}

/*
GetSimpleTS3
Description:
	Get a transition system to test Pre and Post.
	It should have states that generates empty sets for Pre() and Post() respectively.
*/
func TransitionSystem_GetSimpleTS3() TransitionSystem {
	ts0, _ := GetTransitionSystem(
		[]string{"1", "2", "3", "4", "5"}, []string{"1", "2"},
		map[string]map[string][]string{
			"1": map[string][]string{
				"1": []string{"1"},
				"2": []string{"2"},
			},
			"2": map[string][]string{
				"1": []string{"1"},
				"2": []string{"3"},
			},
			"3": map[string][]string{
				"1": []string{"3"},
				"2": []string{"4"},
			},
			"4": map[string][]string{
				"1": []string{},
				"2": []string{},
			},
			"5": map[string][]string{
				"1": []string{"4"},
				"2": []string{},
			},
		},
		[]string{"1"},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"C", "D"},
			"4": []string{"A", "C"},
			"5": []string{"A", "B", "C", "D"},
		},
	)

	return ts0
}

/*
TransitionSystem_GetSimpleTS4
Description:
	Get a transition system to test Pre and Post.
	It should have states that generates empty sets for Pre() and Post() respectively.
*/
func TransitionSystem_GetSimpleTS4() TransitionSystem {
	ts0, _ := GetTransitionSystem(
		[]string{"1", "2", "3", "4", "5"}, []string{"1", "2"},
		map[string]map[string][]string{
			"1": map[string][]string{
				"1": []string{"1"},
				"2": []string{"2"},
			},
			"2": map[string][]string{
				"1": []string{"1", "2"},
				"2": []string{"2", "3"},
			},
			"3": map[string][]string{
				"1": []string{"3"},
				"2": []string{"4"},
			},
			"4": map[string][]string{
				"1": []string{},
				"2": []string{},
			},
			"5": map[string][]string{
				"1": []string{"4"},
				"2": []string{},
			},
		},
		[]string{"1"},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"B", "D"},
			"4": []string{"A", "C"},
			"5": []string{"A", "B", "C", "D"},
		},
	)

	return ts0
}
