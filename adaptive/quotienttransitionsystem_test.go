/*
quotienttransitionsystem_test.go
Description:
	Tests the functions and objects created in transitionsystem.go
*/

package adaptive

import (
	"testing"
)

/*
TestQTS_Equals1
Description:
	Compares an empty QTSState with a nonempty one.
*/
func TestQTS_Equals1(t *testing.T) {
	// Constants
	X := []TransitionSystemState{
		TransitionSystemState{Name: "1"},
		TransitionSystemState{Name: "2"},
		TransitionSystemState{Name: "3"},
		TransitionSystemState{Name: "4"},
	}

	q1 := &QTSState{Subset: X[0:2]}
	q2 := &QTSState{Subset: []TransitionSystemState{}}

	if q1.Equals(q2) {
		t.Errorf("The algorithm claims q2 and q1 are equal!")
	}
}

/*
TestQTS_Equals2
Description:
	Compares an empty QTSState with a nonempty one.
*/
func TestQTS_Equals2(t *testing.T) {
	// Constants
	X := []TransitionSystemState{
		TransitionSystemState{Name: "1"},
		TransitionSystemState{Name: "2"},
		TransitionSystemState{Name: "3"},
		TransitionSystemState{Name: "4"},
	}

	q1 := &QTSState{Subset: X[0:2]}
	q2 := &QTSState{Subset: []TransitionSystemState{}}

	if q2.Equals(q1) {
		t.Errorf("The algorithm claims q2 and q1 are equal!")
	}
}

/*
TestQTS_Equals3
Description:
	Compares different QTSState objects with the same subset in each.
*/
func TestQTS_Equals3(t *testing.T) {
	// Constants
	X := []TransitionSystemState{
		TransitionSystemState{Name: "1"},
		TransitionSystemState{Name: "2"},
		TransitionSystemState{Name: "3"},
		TransitionSystemState{Name: "4"},
	}

	q1 := &QTSState{Subset: X[0:2]}
	q2 := &QTSState{Subset: X[0:2]}

	if !q2.Equals(q1) {
		t.Errorf("The algorithm claims q2 and q1 are NOT equal!")
	}
}

/*
TestQTS_In1
Description:
	Determines if a state is in an empty list..
*/
func TestQTS_In1(t *testing.T) {
	// Constants
	X := []TransitionSystemState{
		TransitionSystemState{Name: "1"},
		TransitionSystemState{Name: "2"},
		TransitionSystemState{Name: "3"},
		TransitionSystemState{Name: "4"},
	}

	q1 := &QTSState{Subset: X[0:2]}
	// q2 := QTSState{Subset: X[0:2]}

	if q1.In([]*QTSState{}) {
		t.Errorf("The algorithm claims q1 is in an empty set!")
	}
}

/*
TestQTS_In2
Description:
	Determines if a state is in an single entry list..
*/
func TestQTS_In2(t *testing.T) {
	// Constants
	X := []TransitionSystemState{
		TransitionSystemState{Name: "1"},
		TransitionSystemState{Name: "2"},
		TransitionSystemState{Name: "3"},
		TransitionSystemState{Name: "4"},
	}

	q1 := &QTSState{Subset: X[0:2]}
	q2 := &QTSState{Subset: X[0:1]}

	if q1.In([]*QTSState{q2}) {
		t.Errorf("The algorithm claims q1 is equal to q2!")
	}
}
