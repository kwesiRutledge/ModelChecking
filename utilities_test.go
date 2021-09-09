/*
utilities_test.go
Description:
	Tests of the extra functions located in utilities.go.
*/
package modelchecking

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

/*
TestUtilities_GetBeverageVendingMachineTS1
Description:
	Verifies that the beverage vending machine transition system has the appropriate number of states.
*/
func TestUtilities_GetBeverageVendingMachineTS1(t *testing.T) {
	// Create
	ts0 := GetBeverageVendingMachineTS()

	// Test
	if len(ts0.S) != 4 {
		t.Errorf("There were not four states in the transition system's state space.")
	}
}

/*
TestUtilities_SliceCartesianProduct1
Description:
	Verifies that a simple Cartesian product has 4 elements.
*/
func TestUtilities_SliceCartesianProduct1(t *testing.T) {
	// Create
	ts0 := GetBeverageVendingMachineTS()

	s0 := ts0.S[0]
	s1 := ts0.S[1]

	S1 := []TransitionSystemState{s0, s1}

	// Test
	cp1, err := SliceCartesianProduct(S1, S1)
	if err != nil {
		t.Errorf("There was an error computing the cartesian product: %v", err)
	}

	cp1Converted := cp1.([][]TransitionSystemState)

	if len(cp1Converted) != 4 {
		t.Errorf("There were not 4 tuples in the cartesian product. Instead there were %v.", len(cp1Converted))
	}
}

/*
TestUtilities_SliceCartesianProduct1
Description:
	Verifies that a simple Cartesian product has the correct tuple as its first element.
*/
func TestUtilities_SliceCartesianProduct2(t *testing.T) {
	// Create
	ts0 := GetBeverageVendingMachineTS()

	s0 := ts0.S[0]
	s1 := ts0.S[1]

	S1 := []TransitionSystemState{s0, s1}

	// Test
	cp1, err := SliceCartesianProduct(S1, S1)
	if err != nil {
		t.Errorf("There was an error computing the cartesian product: %v", err)
	}

	cp1Converted := cp1.([][]TransitionSystemState)

	if (!s0.Equals(cp1Converted[0][0])) || (!s0.Equals(cp1Converted[0][1])) {
		t.Errorf("There were not 4 tuples in the cartesian product. Instead there were %v.", len(cp1Converted))
	}
}

/*
TestUtilities_AppendIfUnique1
Description:
	Verifies that the Append function works properly for a single string to AppendIfUnique.
	The extra string is NOT part of the original list.
*/
func TestUtilities_AppendIfUnique1(t *testing.T) {
	// Create a simple slice of strings
	slice1 := []string{"A", "B", "C"}
	string1 := "Temp"

	// Append string.
	slice2 := AppendIfUnique(slice1, string1)

	if len(slice2) != 4 {
		t.Errorf("The new slice was expected to have length %v, but was %v instead.", 4, len(slice2))
	}

	if slice2[len(slice2)-1] != string1 {
		t.Errorf("The new slice should have final element \"%v\", but was \"%v\".", string1, slice2[len(slice2)-1])
	}
}

/*
TestUtilities_AppendIfUnique2
Description:
	Verifies that the Append function works properly for a single string to AppendIfUnique.
	The extra string IS part of the original list.
*/
func TestUtilities_AppendIfUnique2(t *testing.T) {
	// Create a simple slice of strings
	slice1 := []string{"A", "B", "C"}
	string1 := "A"

	// Append string.
	slice2 := AppendIfUnique(slice1, string1)

	if len(slice2) != 3 {
		t.Errorf("The new slice was expected to have length %v, but was %v instead.", 3, len(slice2))
	}

	if slice2[len(slice2)-1] != "C" {
		t.Errorf("The new slice should have final element \"%v\", but was \"%v\".", "C", slice2[len(slice2)-1])
	}
}

/*
TestUtilities_AppendIfUnique3
Description:
	Verifies that the Append function works properly for a string slice to AppendIfUnique.
	The string slice IS NOT COMPLETELY part of the original list.
*/
func TestUtilities_AppendIfUnique3(t *testing.T) {
	// Create a simple slice of strings
	slice1 := []string{"A", "B", "C"}
	slice2 := []string{"B", "D"}

	// Append string.
	slice3 := AppendIfUnique(slice1, slice2...)

	if len(slice3) != 4 {
		t.Errorf("The new slice was expected to have length %v, but was %v instead.", 4, len(slice3))
	}

	if slice3[len(slice3)-1] != slice2[len(slice2)-1] {
		t.Errorf("The new slice should have final element \"%v\", but was \"%v\".", slice2[len(slice2)-1], slice3[len(slice3)-1])
	}
}

/*
TestUtilities_AppendIfUnique4
Description:
	Verifies that the Append function works properly for a string slice to AppendIfUnique.
	The string slice IS COMPLETELY part of the original list.
*/
func TestUtilities_AppendIfUnique4(t *testing.T) {
	// Create a simple slice of strings
	slice1 := []string{"A", "B", "C"}
	slice2 := []string{"B", "A"}

	// Append string.
	slice3 := AppendIfUnique(slice1, slice2...)

	if len(slice3) != 3 {
		t.Errorf("The new slice was expected to have length %v, but was %v instead.", 3, len(slice3))
	}

	if slice3[len(slice3)-1] != "C" {
		t.Errorf("The new slice should have final element \"%v\", but was \"%v\".", "C", slice3[len(slice3)-1])
	}
}

/*
TestUtilities_AppendIfUnique5
Description:
	Verifies that the Append function works properly for a single string to AppendIfUnique.
	The string slice IS COMPLETELY part of the original list.
*/
func TestUtilities_AppendIfUnique5(t *testing.T) {
	// Create a simple slice of strings
	slice1 := []string{}
	slice2 := []string{"B", "A"}

	// Append string.
	slice3 := AppendIfUnique(slice1, slice2...)

	if len(slice3) != 2 {
		t.Errorf("The new slice was expected to have length %v, but was %v instead.", 2, len(slice3))
	}

	if slice3[len(slice3)-1] != slice2[len(slice2)-1] {
		t.Errorf("The new slice should have final element \"%v\", but was \"%v\".", slice2[len(slice2)-1], slice3[len(slice3)-1])
	}
}

/*
TestUtilities_FindInSlice1
Description:
	Verifies that FindInSlice() works for a given slice of strings.
*/
func TestUtilities_FindInSlice1(t *testing.T) {
	//Create a simple slice.
	slice1 := []string{"A", "B", "C", "E"}
	target1 := "E"

	foundIndex, tf := FindInSlice(target1, slice1)

	if !tf {
		t.Errorf("The function FindInSlice() could not find the target even though it exists in slice1!")
	}

	if foundIndex != 3 {
		t.Errorf("Expected that the foundIndex would be 3, but received %v.", foundIndex)
	}
}

/*
TestUtilities_FindInSlice2
Description:
	Verifies that FindInSlice() works for a given slice of strings. The target string is NOT in the slice.
*/
func TestUtilities_FindInSlice2(t *testing.T) {
	//Create a simple slice.
	slice1 := []string{"A", "B", "C", "E"}
	target1 := "D"

	foundIndex, tf := FindInSlice(target1, slice1)

	if tf {
		t.Errorf("The function FindInSlice() found the target even though it DOES NOT exist in slice1!")
	}

	if foundIndex != -1 {
		t.Errorf("Expected that the foundIndex would be -1, but received %v.", foundIndex)
	}
}
