/*
quotienttransitionsystem.go
Description:
	Defines a Quotient Transition System according to the definition provided in the paper
	'Formal Methods for Adaptive Control of Dynamical Systems' by Sadra Sadraddini and Calin Belta.
*/

package adaptive

import mc "github.com/kwesiRutledge/ModelChecking"

/*
Type Definitions
*/

type QTSState struct {
	Subset []TransitionSystemState
	System *QuotientTransitionSystem
}

type QuotientTransitionSystem struct {
	Q     []QTSState
	U     []string
	BetaQ map[*QTSState]map[string][]*QTSState
	Pi    []mc.AtomicProposition
	O_Q   map[*QTSState][]mc.AtomicProposition
}

// =========
// Functions
// =========

func (stateIn *QTSState) Equals(stateCompare *QTSState) bool {
	// Compare the Two Subsets
	tf, _ := SliceEquals(stateIn.Subset, stateCompare.Subset)

	return tf
}

/*
In
Description:
	Determines if the given state is in a list of states.
Usage:
	tf := q1.In( qList )
*/
func (stateIn *QTSState) In(stateList []*QTSState) bool {
	//Verify if state is in the state list
	for _, tempState := range stateList {
		if stateIn.Equals(tempState) {
			return true
		}
	}
	// If the state is not found in stateList, return true.
	return false
}

/*
AppendIfUniqueTo
Description:
	Appends to the input slice sliceIn if and only if the new state
	is actually a unique state.
Usage:
	newSlice := stateIn.AppendIfUniqueTo(sliceIn)
*/
func (stateIn *QTSState) AppendIfUniqueTo(sliceIn []*QTSState) []*QTSState {

	for _, tempState := range sliceIn {
		if tempState.Equals(stateIn) {
			return sliceIn
		}

	}

	//If all checks were passed
	return append(sliceIn, stateIn)

}
