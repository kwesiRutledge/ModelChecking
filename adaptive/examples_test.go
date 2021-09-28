/*

 */
package adaptive

import "testing"

func TestExamples_GetBeverageVendingMachineTS1(t *testing.T) {

	// Constants

	// Algorithm
	ts0 := GetBeverageVendingMachineTS()

	if len(ts0.X) != 4 {
		t.Errorf("Expected the Beverage Vending Machine Transition System to have 4 states, but it has %v.", len(ts0.X))
	}

}
