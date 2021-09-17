/*
deterministicrabin.go
Description:
	Defines a Deterministic Rabin Automaton according to the definition provided in the paper
	'Formal Methods for Adaptive Control of Dynamical Systems' by Sadra Sadraddini and Calin Belta.
*/

package adaptive

import (
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
	alpha    map[DRAState]map[mc.AtomicProposition][]DeterministicRabinAutomaton
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
