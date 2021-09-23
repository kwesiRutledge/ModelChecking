/*
utilities_test.go
Description:
	Tests the functions and objects created in utilities.go
*/

package adaptive

import "testing"

/*
TestUtilties_SliceSubset1
Description:
	Verifies that one slice is NOT a subset of the other.
*/
func TestUtilties_SliceSubset1(t *testing.T) {
	// Constants
	q0 := DRAState{Name: "Quaren"}
	q1 := DRAState{Name: "Conspiracy Brother"}
	q2 := DRAState{Name: "Dave Chappelle"}

	qSlice0 := []DRAState{q0, q1, q2}
	qSlice1 := []DRAState{q0, q2}

	// Algorithm
	subsetFlag, err := SliceSubset(qSlice0, qSlice1)
	if err != nil {
		t.Errorf("There was an error computing the SliceSubset(): %v", err)
	}

	if subsetFlag {
		t.Errorf("The function incorrectly claims that qSlice0 is a subset of qSlice1")
	}

}

/*
TestUtilties_SliceSubset2
Description:
	Verifies that one slice is a subset of the other.
*/
func TestUtilties_SliceSubset2(t *testing.T) {
	// Constants
	q0 := DRAState{Name: "Quaren"}
	q1 := DRAState{Name: "Conspiracy Brother"}
	q2 := DRAState{Name: "Dave Chappelle"}

	qSlice0 := []DRAState{q0, q1, q2}
	qSlice1 := []DRAState{q0, q2}

	// Algorithm
	subsetFlag, err := SliceSubset(qSlice1, qSlice0)
	if err != nil {
		t.Errorf("There was an error computing the SliceSubset(): %v", err)
	}

	if !subsetFlag {
		t.Errorf("The function incorrectly claims that qSlice1 is NOT a subset of qSlice0")
	}

}
