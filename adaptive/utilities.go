/*
utilities.go
Description:
	Extra functions needed to reproduce the results of the paper
	'Formal Methods for Adaptive Control of Dynamical Systems' by Sadra Sadraddini and Calin Belta.
*/

package adaptive

import "fmt"

/*
SliceSubset
Description:
	Determines if slice1 is a subset of slice2
*/
func SliceSubset(slice1, slice2 interface{}) (bool, error) {

	switch x := slice1.(type) {
	case []DRAState:
		stateSlice1, ok1 := slice1.([]DRAState)
		stateSlice2, ok2 := slice2.([]DRAState)

		if (!ok1) || (!ok2) {
			return false, fmt.Errorf("Error converting slice1 (%v) or slice2 (%v).", ok1, ok2)
		}

		//Iterate through all TransitionSystemState in stateSlice1 and make sure that they are in 2.
		for _, stateFrom1 := range stateSlice1 {
			if !(stateFrom1.In(stateSlice2)) {
				return false, nil
			}
		}
		// If all elements of slice1 are in slice2 then return true!
		return true, nil

	default:
		return false, fmt.Errorf("Unexpected type given to SliceSubset(): %v", x)
	}

}
