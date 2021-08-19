/*
atomicproposition_test.go
Description:
	Tests of the functions pertaining to the AtomicProposition object.
*/
package modelchecking

import (
	"testing"
)

/*
TestAtomicProposition_Equals1
Description:
	Tests if the Equals() member function for AtomicProposition works.
*/
func TestAtomicProposition_Equals1(t *testing.T) {
	// Constants
	ap1 := AtomicProposition{Name: "A"}
	ap2 := AtomicProposition{Name: "B"}
	ap3 := AtomicProposition{Name: "A"}

	if ap1.Equals(ap2) {
		t.Errorf("ap1 (%v) is supposed to be different from ap2 (%v).", ap1.Name, ap2.Name)
	}

	if !ap1.Equals(ap3) {
		t.Errorf("ap1 (%v) is supposed to be the same as ap3 (%v).", ap1.Name, ap3.Name)
	}

}

/*
TestAtomicProposition_In1
Description:
	A good test that correctly identifies if an ap is in a list.
*/
func TestAtomicProposition_In1(t *testing.T) {
	// Create Slice of AtomicPropositions
	ap1 := AtomicProposition{Name: "A"}
	ap2 := AtomicProposition{Name: "B"}

	apSlice1 := []AtomicProposition{ap1, ap2}

	// Test
	if !ap1.In(apSlice1) {
		t.Errorf("Atomic proposition 1 is in the slice, but the function doesn't recognize this!")
	}
}

/*
TestAtomicProposition_In2
Description:
	A good test that correctly identifies if an ap is not in a slice.
*/
func TestAtomicProposition_In2(t *testing.T) {
	// Create Slice of AtomicPropositions
	ap1 := AtomicProposition{Name: "A"}
	ap2 := AtomicProposition{Name: "B"}
	ap3 := AtomicProposition{Name: "C"}

	apSlice1 := []AtomicProposition{ap1, ap2}

	// Test
	if ap3.In(apSlice1) {
		t.Errorf("Atomic proposition 3 is not in the slice, but the function doesn't recognize this!")
	}
}

/*
TestAtomicProposition_Powerset1
Description:
	Checks that the powerset correctly creates all of the values that we know it should.
*/
func TestAtomicProposition_Powerset1(t *testing.T) {
	// Create Slice of AtomicPropositions
	ap1 := AtomicProposition{Name: "A"}
	ap2 := AtomicProposition{Name: "B"}

	apSlice := []AtomicProposition{ap1, ap2}

	// Algorithm
	powerset := Powerset(apSlice)
	if len(powerset) != 4 {
		t.Errorf("There are supposed to be 4 elements in the powerset, but there are only %v.", len(powerset))
	}

	for _, powersetElement := range powerset {
		equalsEmpty, _ := SliceEquals(powersetElement, []AtomicProposition{})
		equals2, _ := SliceEquals(powersetElement, []AtomicProposition{ap2})
		equals1, _ := SliceEquals(powersetElement, []AtomicProposition{ap1})
		equals12, _ := SliceEquals(powersetElement, []AtomicProposition{ap1, ap2})
		if !(equalsEmpty || equals2 || equals1 || equals12) {
			t.Errorf("The element %v of the powerset is unexpected.", powersetElement)
		}
	}

}
