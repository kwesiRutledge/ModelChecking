/*
atomicproposition_test.go
Description:
	Tests of the functions pertaining to the AtomicProposition object.
*/
package main

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
func TestAtomicPropositionIn1(t *testing.T) {
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
func TestAtomicPropositionIn2(t *testing.T) {
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
TestAtomicProposition_Subset1
Description:
	Checks that the function correctly identifies when one set is a slice is a subset of another.
*/
func TestAtomicProposition_Subset1(t *testing.T) {
	// Create Slice of AtomicPropositions
	ap1 := AtomicProposition{Name: "A"}
	ap2 := AtomicProposition{Name: "B"}
	ap3 := AtomicProposition{Name: "C"}

	apSlice1 := []AtomicProposition{ap1, ap2, ap3}
	apSlice2 := []AtomicProposition{ap1, ap3}

	// Test
	if !Subset(apSlice2, apSlice1) {
		t.Errorf("The function does not correctly identify that slice 2 is a subset of slice 1.")
	}
}

/*
TestAtomicProposition_Subset2
Description:
	Checks that the function correctly identifies when one set is NOT a slice is a subset of another.
*/
func TestAtomicProposition_Subset2(t *testing.T) {
	// Create Slice of AtomicPropositions
	ap1 := AtomicProposition{Name: "A"}
	ap2 := AtomicProposition{Name: "B"}
	ap3 := AtomicProposition{Name: "C"}

	apSlice1 := []AtomicProposition{ap1, ap2, ap3}
	apSlice2 := []AtomicProposition{ap1, ap3}

	// Test
	if Subset(apSlice1, apSlice2) {
		t.Errorf("The function does not correctly identify that slice 1 is NOT a subset of slice 2.")
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
	if len(powerset) != 0 {
		t.Errorf("idk")
	}

}
