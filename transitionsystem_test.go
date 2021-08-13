package main

import (
	"testing"
)

func TestTransitionSystem_GetState1(t *testing.T) {
	// Create Simple Transition System
	ts0 := TransitionSystem{}
	ts0.S = []TransitionSystemState{
		TransitionSystemState{"1", &ts0},
		TransitionSystemState{"2", &ts0},
		TransitionSystemState{"3", &ts0},
		TransitionSystemState{"4", &ts0},
	}

	// Try to get a state which is outside of the allowable range.
	if len(ts0.S) != 4 {
		t.Errorf("There are not four states left.")
	}
}

func TestTransitionSystem_GetState2(t *testing.T) {
	// Create Simple Transition System
	ts0 := TransitionSystem{}
	ts0.S = []TransitionSystemState{
		TransitionSystemState{"1", &ts0},
		TransitionSystemState{"2", &ts0},
		TransitionSystemState{"3", &ts0},
		TransitionSystemState{"4", &ts0},
	}

	// Try to get a state which is outside of the allowable range.
	tempState := ts0.S[1]
	if tempState.Name != "2" {
		t.Errorf("The value of the correct state (2) was not saved in the state.")
	}
}

func TestTransitionSystem_GetTransitionSystem1(t *testing.T) {
	ts0, err := GetTransitionSystem(
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

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Tests
	if len(ts0.AP) != 4 {
		t.Errorf("The number of atomic propositions was expected to be 4, but it was %v.", len(ts0.AP))
	}
}

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
TestTransitionSystem_IsActionDeterministic1
Description:
	Testing that the function correctly recognizes that a transition system IS NOT action deterministic.
*/
func TestTransitionSystem_IsActionDeterministic1(t *testing.T) {
	ts1 := GetSimpleTS2()

	if ts1.IsActionDeterministic() {
		t.Errorf("Test is given a transition system that is not action deterministic, but function claims that it is!")
	}
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
TestTransitionSystem_IsActionDeterministic2
Description:
	Testing that the function correctly recognizes that a transition system IS action deterministic.
*/
func TestTransitionSystem_IsActionDeterministic2(t *testing.T) {
	ts1 := TransitionSystem_GetSimpleTS3()

	if !ts1.IsActionDeterministic() {
		t.Errorf("Test is given a transition system that is action deterministic, but function claims that it is not!")
	}
}
