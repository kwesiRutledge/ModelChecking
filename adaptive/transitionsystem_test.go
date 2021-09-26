/*
transitionsystem_test.go
Description:
	Tests the functions and objects created in transitionsystem.go
*/

package adaptive

import (
	"fmt"
	"testing"
)

func TestTransitionSystem_GetState1(t *testing.T) {
	// Create Simple Transition System
	ts0 := TransitionSystem{}
	ts0.X = []TransitionSystemState{
		TransitionSystemState{"1", &ts0},
		TransitionSystemState{"2", &ts0},
		TransitionSystemState{"3", &ts0},
		TransitionSystemState{"4", &ts0},
	}

	// Try to get a state which is outside of the allowable range.
	if len(ts0.X) != 4 {
		t.Errorf("There are not four states left.")
	}
}

func TestTransitionSystem_GetState2(t *testing.T) {
	// Create Simple Transition System
	ts0 := TransitionSystem{}
	ts0.X = []TransitionSystemState{
		TransitionSystemState{"1", &ts0},
		TransitionSystemState{"2", &ts0},
		TransitionSystemState{"3", &ts0},
		TransitionSystemState{"4", &ts0},
	}

	// Try to get a state which is outside of the allowable range.
	tempState := ts0.X[1]
	if tempState.Name != "2" {
		t.Errorf("The name of the correct state (2) was not saved in the state.")
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
	if len(ts0.Pi) != 4 {
		t.Errorf("The number of atomic propositions was expected to be 4, but it was %v.", len(ts0.Pi))
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

/*
TestTransitionSystem_Check1
Description:
	Testing that transition system check function works outside of the constructor.
*/
func TestTransitionSystem_Check1(t *testing.T) {
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
			"4": map[string][]string{
				"1": []string{"3"},
				"2": []string{"4"},
			},
		},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"B", "D"},
		},
	)

	err = ts0.Check()
	if err == nil {
		t.Errorf("The algorithm did not catch the Transition issue!")
	}

	if err.Error() != fmt.Sprintf("One of the source states in the Transition was not in the state set: 4") {
		t.Errorf("The error was not what we expected: %v", err.Error())
	}

}

/*
TestTransitionSystem_IsNonBlocking1
Description:
	When given a blocking transition system, tests to see if the function detects that it is blocking.
*/
func TestTransitionSystem_IsNonBlocking1(t *testing.T) {
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
				"2": []string{},
			},
		},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"B", "D"},
		},
	)

	err = ts0.Check()
	if err != nil {
		t.Errorf("There was an error in the construction of ts0! %v", err)
	}

	if ts0.IsNonBlocking() {
		t.Errorf("The function IsNonBlocking() did not recognize that ts0 is blocking!")
	}

}

/*
TestTransitionSystem_IsNonBlocking2
Description:
	When given a non-blocking transition system, tests to see if the function detects that it is blocking.
*/
func TestTransitionSystem_IsNonBlocking2(t *testing.T) {
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
				"2": []string{"1"},
			},
		},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"B", "D"},
		},
	)

	err = ts0.Check()
	if err != nil {
		t.Errorf("There was an error in the construction of ts0! %v", err)
	}

	if !ts0.IsNonBlocking() {
		t.Errorf("The function IsNonBlocking() incorrectly claims that ts0 is blocking!")
	}

}

/*
TestTransitionSystem_UnionOfStates1
Description:
	Verifies that UnionOfStates works when only one set is given to it.
*/
func TestTransitionSystem_UnionOfStates1(t *testing.T) {

	// Constants
	x0 := TransitionSystemState{Name: "Steven"}
	x1 := TransitionSystemState{Name: "Gabe"}
	x2 := TransitionSystemState{Name: "Nosa"}

	slice1 := []TransitionSystemState{x0, x1, x2}

	// Algorithm
	slice2 := UnionOfStates(slice1)

	if len(slice2) == 0 {
		t.Errorf("slice2 is empty!")
	}

	if !x0.In(slice2) {
		t.Errorf("x0 not found in slice2!")
	}

	if !x1.In(slice2) {
		t.Errorf("x1 not found in slice2!")
	}

	if !x2.In(slice2) {
		t.Errorf("x2 not found in slice2!")
	}

}

/*
TestTransitionSystem_UnionOfStates2
Description:
	Verifies that UnionOfStates works when two overlapping sets are given to it.
*/
func TestTransitionSystem_UnionOfStates2(t *testing.T) {

	// Constants
	x0 := TransitionSystemState{Name: "Steven"}
	x1 := TransitionSystemState{Name: "Gabe"}
	x2 := TransitionSystemState{Name: "Nosa"}
	x3 := TransitionSystemState{Name: "Desnor"}
	x4 := TransitionSystemState{Name: "Nailah"}

	slice1 := []TransitionSystemState{x0, x1, x2}
	slice2 := []TransitionSystemState{x2, x3, x4}

	// Algorithm
	slice3 := UnionOfStates(slice1, slice2)

	if len(slice3) != 5 {
		t.Errorf("slice3 is empty!")
	}

	if !x0.In(slice3) {
		t.Errorf("x0 not found in slice3!")
	}

	if !x1.In(slice3) {
		t.Errorf("x1 not found in slice3!")
	}

	if !x2.In(slice3) {
		t.Errorf("x2 not found in slice3!")
	}

}

/*
TestTransitionSystem_UnionOfStates3
Description:
	Verifies that UnionOfStates works when three overlapping sets are given to it.
*/
func TestTransitionSystem_UnionOfStates3(t *testing.T) {

	// Constants
	x0 := TransitionSystemState{Name: "Steven"}
	x1 := TransitionSystemState{Name: "Gabe"}
	x2 := TransitionSystemState{Name: "Nosa"}
	x3 := TransitionSystemState{Name: "Desnor"}
	x4 := TransitionSystemState{Name: "Nailah"}
	x5 := TransitionSystemState{Name: "Elon Musk"}
	x6 := TransitionSystemState{Name: "Jeff Bezos"}

	slice1 := []TransitionSystemState{x0, x1, x2}
	slice2 := []TransitionSystemState{x2, x3, x4}
	slice3 := []TransitionSystemState{x0, x3, x4, x5, x6}

	// Algorithm
	slice4 := UnionOfStates(slice1, slice2, slice3)

	if len(slice4) != 7 {
		t.Errorf("slice4 does not contain the proper number of elements!")
	}

	if !x0.In(slice4) {
		t.Errorf("x0 not found in slice4!")
	}

	if !x1.In(slice4) {
		t.Errorf("x1 not found in slice4!")
	}

	if !x2.In(slice4) {
		t.Errorf("x2 not found in slice4!")
	}

}
