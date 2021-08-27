package modelchecking

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
	if tempState.Value != "2" {
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

/*
TestTransitionSystem_IsAPDeterministic1
Description:
	Testing that the function correctly recognizes that a transition system IS AP-deterministic.
*/
func TestTransitionSystem_IsAPDeterministic1(t *testing.T) {
	ts1 := GetSimpleTS2()

	if !ts1.IsAPDeterministic() {
		t.Errorf("Test is given a transition system that is AP-deterministic, but function claims that it is not!")
	}
}

/*
TestTransitionSystem_IsAPDeterministic2
Description:
	Testing that the function correctly recognizes that a transition system IS AP-deterministic.
*/
func TestTransitionSystem_IsAPDeterministic2(t *testing.T) {
	ts1 := GetSimpleTS1()

	if !ts1.IsAPDeterministic() {
		t.Errorf("Test is given a transition system that is AP-deterministic, but function claims that it is not!")
	}
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

/*
TestTransitionSystem_IsAPDeterministic3
Description:
	Testing that the function correctly recognizes that a transition system IS NOT AP-deterministic.
*/
func TestTransitionSystem_IsAPDeterministic3(t *testing.T) {
	ts1 := TransitionSystem_GetSimpleTS4()

	if ts1.IsAPDeterministic() {
		t.Errorf("Test is given a transition system that is NOT AP-deterministic, but function claims that it is!")
	}
}

/*
TestTransitionSystem_CheckI1
Description:
	Testing that transition system constructor catches a bad initial state set is given.
*/
func TestTransitionSystem_CheckI1(t *testing.T) {
	_, err := GetTransitionSystem(
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
				"2": []string{"4"},
			},
		},
		[]string{"4"},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"B", "D"},
		},
	)

	if err == nil {
		t.Errorf("There was not an error getting a transition system! There should have been")
	}

	if err.Error() != "The state 4 is not in the state set of the transition system!" {
		t.Errorf("The error was not what we expected: %v", err.Error())
	}

}

/*
TestTransitionSystem_CheckI2
Description:
	Testing that transition system constructor catches a good initial state set is given.
*/
func TestTransitionSystem_CheckI2(t *testing.T) {
	_, err := GetTransitionSystem(
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
				"2": []string{"1"},
			},
		},
		[]string{"1", "2", "3"},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"B", "D"},
		},
	)

	if err != nil {
		t.Errorf("There was an error getting the transition system: %v", err)
	}

}

/*
TestTransitionSystem_CheckTransition1
Description:
	Testing that transition system constructor catches a bad transition with bad
	initial state.
*/
func TestTransitionSystem_CheckTransition1(t *testing.T) {
	_, err := GetTransitionSystem(
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
			"4": map[string][]string{
				"1": []string{"3"},
				"2": []string{"4"},
			},
		},
		[]string{"1", "2", "3"},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"B", "D"},
		},
	)

	if err == nil {
		t.Errorf("The algorithm did not catch the Transition issue!")
	}

	if err.Error() != "One of the source states in the Transition was not in the state set: 4" {
		t.Errorf("The error was not what we expected: %v", err.Error())
	}

}
