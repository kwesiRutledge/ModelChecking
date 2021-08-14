/*
utilities_test.go
Description:
	Tests of the extra functions located in utilities.go.
*/
package main

import (
	"testing"
)

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
	tf, err := SliceSubset(apSlice2, apSlice1)
	if err != nil {
		t.Errorf("SliceSubset encountered an error! %v", err)
	}
	if !tf {
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
	tf, err := SliceSubset(apSlice1, apSlice2)
	if err != nil {
		t.Errorf("SliceSubset encountered an error! %v", err)
	}
	if tf {
		t.Errorf("The function does not correctly identify that slice 1 is NOT a subset of slice 2.")
	}
}

/*
TestUtilities_SliceSubset3
Description:
	Verifies that SliceSubset knows how to fail when there are two different types given to it.
*/
func TestUtilities_SliceSubset3(t *testing.T) {
	// Create Slice of AtomicPropositions
	ap1 := AtomicProposition{Name: "A"}
	ap2 := AtomicProposition{Name: "B"}
	ap3 := AtomicProposition{Name: "C"}

	apSlice1 := []AtomicProposition{ap1, ap2, ap3}
	tempSlice := []string{"2", "36"}

	_, err := SliceSubset(apSlice1, tempSlice)
	if err == nil {
		t.Errorf("SliceSubset did not throw an error for invalid combination of slices.")
	}
}

/*
TestUtilities_SliceEquals1
Description:
	Verifies that SliceEquals can understand when two slices ARE NOT equal.
*/
func TestUtilities_SliceEquals1(t *testing.T) {
	// Create Slice of AtomicPropositions
	ap1 := AtomicProposition{Name: "A"}
	ap2 := AtomicProposition{Name: "B"}
	ap3 := AtomicProposition{Name: "C"}

	apSlice1 := []AtomicProposition{ap1, ap2, ap3}
	apSlice2 := []AtomicProposition{ap1, ap3}

	// Test
	tf, err := SliceEquals(apSlice1, apSlice2)
	if err != nil {
		t.Errorf("There was an issue computing SliceEquals(): %v", err)
	}

	if tf {
		t.Errorf("The function incorrectly claims that the slices are equal!")
	}
}

/*
TestUtilities_SliceEquals2
Description:
	Verifies that SliceEquals can understand when two slices ARE  equal.
*/
func TestUtilities_SliceEquals2(t *testing.T) {
	// Create Slice of AtomicPropositions
	ap1 := AtomicProposition{Name: "A"}
	ap2 := AtomicProposition{Name: "B"}
	ap3 := AtomicProposition{Name: "C"}

	apSlice1 := []AtomicProposition{ap1, ap2, ap3}
	// apSlice2 := []AtomicProposition{ap1, ap3}

	// Test
	tf, err := SliceEquals(apSlice1, apSlice1)
	if err != nil {
		t.Errorf("There was an issue computing SliceEquals(): %v", err)
	}

	if !tf {
		t.Errorf("The function incorrectly claims that the slices are not equal!")
	}
}
