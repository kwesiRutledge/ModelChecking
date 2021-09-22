/*
deterministicrabin.go
Description:
	Defines a Deterministic Rabin Automaton according to the definition provided in the paper
	'Formal Methods for Adaptive Control of Dynamical Systems' by Sadra Sadraddini and Calin Belta.
*/

package adaptive

import (
	"fmt"

	mc "github.com/kwesiRutledge/ModelChecking"
)

/*
Types
*/

/*
DeterministicRabinAutomaton
Description:
	An object to represent the Deterministic Rabin Automaton.
*/
type DeterministicRabinAutomaton struct {
	S        []DRAState
	s0       DRAState
	Alphabet []mc.AtomicProposition
	alpha    map[DRAState]map[mc.AtomicProposition]DRAState
	Omega    [][2][]DRAState
}

type DRAState struct {
	Name      string
	Automaton *DeterministicRabinAutomaton
}

/*
Functions for DRAState
*/

/*
String
Description:
	This is called whenever one tries to "print" a DRAState.
	It provides the name of the state.
*/
func (stateIn DRAState) String() string {
	return stateIn.Name
}

/*
Equals
Description:
	Returns true if the name of the two objects are equal.
*/
func (stateIn DRAState) Equals(state2 DRAState) bool {
	return stateIn.Name == state2.Name
}

/*
In
Description:
	Determines whether or not a DRAState is in a slice of DRAState objects
*/
func (stateIn DRAState) In(stateList []DRAState) bool {

	// Search through each element of stateList
	tf, _ := stateIn.Find(stateList)

	// Return whether or not stateIn was found.
	return tf
}

/*
Find
Description:
	Determines whether or not the state stateIn is in the list stateList and returns
	a boolean flag reflecting whether it is (true) or isn't (false) as well as the index
	in the list.
Assumptions:
	If the state is not found, the index -1 will be returned along with false.
	Returns the FIRST index where stateIn is found.
*/
func (stateIn DRAState) Find(stateList []DRAState) (bool, int) {

	// Search through each element of stateList
	for stateIndex, tempState := range stateList {
		if stateIn.Equals(tempState) {
			return true, stateIndex
		}
	}

	// If nothing was found during the loop, then
	// stateIn is not in stateList.
	return false, -1
}

/*
GetDRA
Description:
	This function will create a DeterministicRabinAutomaton object for you using simple strings and
	maps of strings.
*/
func GetDRA(SNames []string, initialStateName string, alphabetNames []string, transitionMap map[string]map[string]string, omegaSlice [][2][]string) (DeterministicRabinAutomaton, error) {

	draOut := DeterministicRabinAutomaton{}

	//Use SNames to Create State Space
	var S []DRAState
	for _, stateName := range SNames {
		S = append(S,
			DRAState{Name: stateName, Automaton: &draOut},
		)
	}
	draOut.S = S

	// Use initialStateName to create initial state
	s0 := DRAState{Name: initialStateName, Automaton: &draOut}
	draOut.s0 = s0

	err := draOut.CheckS0()
	if err != nil {
		return draOut, err
	}

	// Create Alphabet which is a list of AtomicPropositions
	var Alphabet []mc.AtomicProposition
	for _, alphabetName := range alphabetNames {
		Alphabet = append(Alphabet,
			mc.AtomicProposition{Name: alphabetName},
		)
	}
	draOut.Alphabet = Alphabet

	// Create Transition Map
	alpha := make(map[DRAState]map[mc.AtomicProposition]DRAState)
	for s0Name, apMap := range transitionMap {
		s0 := DRAState{Name: s0Name, Automaton: &draOut}
		tempAPMap := make(map[mc.AtomicProposition]DRAState)
		for apName, s1Name := range apMap {
			tempAP := mc.AtomicProposition{Name: apName}
			s1 := DRAState{Name: s1Name, Automaton: &draOut}
			tempAPMap[tempAP] = s1
		}
		alpha[s0] = tempAPMap
	}
	draOut.alpha = alpha

	err = draOut.CheckAlpha()
	if err != nil {
		return draOut, err
	}

	// Creates the Omega pairs of sets.

	return draOut, nil

}

/*
CheckS0
Description:
	Checks to make sure that s0 state satisfies:
	- s0 is in draIn.S
*/
func (draIn *DeterministicRabinAutomaton) CheckS0() error {

	// Define Constants
	S := draIn.S
	s0 := draIn.s0

	// If s0 is not in space, return an error
	if !s0.In(S) {
		return fmt.Errorf("The initial state \"%v\" was not in the state space S.", s0)
	}

	// Everything is Fine
	return nil
}

/*
CheckAlpha
Description:
	Checks the transition map for the Deterministic Rabin Automaton.
*/
func (draIn *DeterministicRabinAutomaton) CheckAlpha() error {
	// Constants
	alpha := draIn.alpha

	// Algorithms
	for s0, apMap := range alpha {
		if !s0.In(draIn.S) {
			return fmt.Errorf("The state \"%v\" is not in the state space.", s0)
		}

		for ap0, s1 := range apMap {
			if !ap0.In(draIn.Alphabet) {
				return fmt.Errorf("The atomic proposition \"%v\" is not in the Alphabet.", ap0)
			}

			// State Space
			if !s1.In(draIn.S) {
				return fmt.Errorf("The state \"%v\" is not in the state space.", s1)
			}

		}
	}

	// Finished checking all transitions
	return nil
}
