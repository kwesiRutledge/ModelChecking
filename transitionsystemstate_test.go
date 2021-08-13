package main

import (
	"testing"
)

/*
TestTransitionSystem_Satisfies1
Description:
	Tests if the Satisfies() member function correctly identifies when the system
	satisfies a given transition system.
*/
func TestTransitionSystem_Satisfies1(t *testing.T) {
	// Constants
	ts1 := GetSimpleTS1()
	ap2 := AtomicProposition{Name: "B"}

	// Test
	state2 := ts1.S[1]

	tf, err := state2.Satisfies(ap2)
	if err != nil {
		t.Errorf("There was an error while testing satisfies: %v", err.Error())
	}

	if !tf {
		t.Errorf("ap1 (%v) is supposed to be satisfied by ts1.", ap2.Name)
	}

}

/*
TestTransitionSystem_Satisfies2
Description:
	Tests if the Satisfies() member function correctly identifies when the system
	satisfies a given transition system.
*/
func TestTransitionSystem_Satisfies2(t *testing.T) {
	// Constants
	ts1 := GetSimpleTS1()
	ap2 := AtomicProposition{Name: "B"}

	// Test
	state2 := ts1.S[0]

	tf, err := state2.Satisfies(ap2)
	if err != nil {
		t.Errorf("There was an error while testing satisfies: %v", err.Error())
	}

	if tf {
		t.Errorf("ap1 (%v) is supposed to NOT be satisfied by ts1.", ap2.Name)
	}

}

/*
TestTransitionSystemState_Post1
Description:
	Tests if the Post() member function correctly identifies when the system
	satisfies a given transition system.
*/
func TestTransitionSystemState_Post1(t *testing.T) {
	// Constants
	ts1 := GetSimpleTS1()

	// Test
	state2 := ts1.S[1]

	nextStates, err := Post(state2, "1")
	if err != nil {
		t.Errorf("There was an error while testing Post: %v", err.Error())
	}

	if len(nextStates) != 2 {
		t.Errorf("Expected there to be 2 options for next state but received %v options. %v", len(nextStates), nextStates[0])
	}

}

/*
TestTransitionSystemState_Post2
Description:
	Tests if the Post() member function correctly creates the ancestors
	when there is no action given.
*/
func TestTransitionSystemState_Post2(t *testing.T) {
	// Constants
	ts1 := GetSimpleTS1()

	// Test
	state1 := ts1.S[0]
	state3 := ts1.S[2]

	nextStates, err := Post(state1)
	if err != nil {
		t.Errorf("There was an error while testing Post: %v", err.Error())
	}

	if len(nextStates) != 2 {
		t.Errorf("Expected there to be 2 options for next state but received %v options. %v", len(nextStates), nextStates[0])
	}

	if state3.In(nextStates) {
		t.Errorf("Expected for state3 to not be in the post result, but it was found there!")
	}

}

/*
TestTransitionSystemState_Post3
Description:
	Tests if the Post() member function correctly creates a post set which
	DOES NOT include all states in the system.
*/
func TestTransitionSystemState_Post3(t *testing.T) {
	// Constants
	ts1 := GetSimpleTS1()

	// Test
	state2 := ts1.S[1]

	nextStates, err := Post(state2)
	if err != nil {
		t.Errorf("There was an error while testing Post: %v", err.Error())
	}

	if len(nextStates) != 3 {
		t.Errorf("Expected there to be 3 options for next state but received %v options. %v", len(nextStates), nextStates[0])
	}

}

/*
TestTransitionSystemState_Pre1
Description:
	Tests if the Pre() member function correctly finds the proper state when given a
	state and action with only one possible predecessor.
*/
func TestTransitionSystemState_Pre1(t *testing.T) {
	// Constants
	ts1 := GetSimpleTS1()

	// Test
	state3 := ts1.S[2]

	predecessors, err := Pre(state3, "1")
	if err != nil {
		t.Errorf("There was an error while testing Post: %v", err.Error())
	}

	if len(predecessors) != 1 {
		t.Errorf("Expected there to be 1 options for next state but received %v options. %v", len(predecessors), predecessors[0])
	}

	if !predecessors[0].Equals(state3) {
		t.Errorf("Expected for the predecessor to be state3 but received the state%v.", predecessors[0].Name)
	}

}

/*
TestTransitionSystemState_Pre2
Description:
	Tests if the Pre() member function correctly finds the proper state when given a
	state only.
*/
func TestTransitionSystemState_Pre2(t *testing.T) {
	// Constants
	ts1 := GetSimpleTS1()

	// Test
	state3 := ts1.S[2]

	predecessors, err := Pre(state3)
	if err != nil {
		t.Errorf("There was an error while testing Post: %v", err.Error())
	}

	if len(predecessors) != 2 {
		t.Errorf("Expected there to be 2 options for next state but received %v options. %v", len(predecessors), predecessors[0])
	}

	if !ts1.S[1].In(predecessors) {
		t.Errorf("Expected for state2 to be in predecessors, but it was not!")
	}

	if ts1.S[0].In(predecessors) {
		t.Errorf("Expected for state1 to NOT be in predecessors, but it was!")
	}

}

/*
TestTransitionSystemState_Post4
Description:
	Tests if the Post() member function correctly finds an empty set when
	given the right system and inputs.
*/
func TestTransitionSystemState_Post4(t *testing.T) {
	// Constants
	ts1 := GetSimpleTS2()

	// Test
	state4 := ts1.S[3]

	ancestors, err := Post(state4)
	if err != nil {
		t.Errorf("There was an error while testing Post: %v", err.Error())
	}

	if len(ancestors) != 0 {
		t.Errorf("Expected there to be 0 options for next state but received %v options. %v", len(ancestors), ancestors[0])
	}

}

/*
TestTransitionSystemState_Pre3
Description:
	Tests if the Pre() member function correctly finds an empty set when
	given.
*/
func TestTransitionSystemState_Pre3(t *testing.T) {
	// Constants
	ts1 := GetSimpleTS2()

	// Test
	state5 := ts1.S[4]

	predecessors, err := Pre(state5)
	if err != nil {
		t.Errorf("There was an error while testing Post: %v", err.Error())
	}

	if len(predecessors) != 0 {
		t.Errorf("Expected there to be 0 options for next state but received %v options. %v", len(predecessors), predecessors[0])
	}

}

/*
TestTransitionSystemState_IsTerminal1
Description:
	Tests if the IsTerminal() function correctly identifies a terminal state.
*/
func TestTransitionSystemState_IsTerminal1(t *testing.T) {
	// Constants
	ts1 := GetSimpleTS2()

	// Test
	state4 := ts1.S[3]

	if !state4.IsTerminal() {
		t.Errorf("There was an error while verifying that state4 was terminal. Function claims that it is not!")
	}

}
