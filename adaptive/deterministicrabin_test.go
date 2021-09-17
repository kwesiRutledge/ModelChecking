/*
deterministicrabin_test.go
Description:
	Tests the functions and objects created in deterministicrabin.go
*/

package adaptive

import "testing"

/*
TestDeterministicRabin_String1
Description:
	Creates a basic DRAState and then verifies that the name is what we initialized it as.
*/
func TestDeterministicRabin_String1(t *testing.T) {
	// Create a Simple DRAState
	q0 := DRAState{
		Name: "Quaren",
	}

	// Test String
	if q0.String() != "Quaren" {
		t.Errorf("The name of q0 was \"%v\", not \"Quaren\"!", q0)
	}
}
