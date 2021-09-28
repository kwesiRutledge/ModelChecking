/*
examples.go
Description:
	Defines a useful set of example transition systems for testing or other uses.
*/

package adaptive

import "fmt"

/*
GetBeverageVendingMachineTS
Description:
	Creates the beloved Vending Machine example which is used in a lot of Principles of Model Checking.
*/
func GetBeverageVendingMachineTS() TransitionSystem {

	ts0, err := GetTransitionSystem(
		[]string{"pay", "select", "beer", "soda"}, []string{"", "insert_coin", "get_beer", "get_soda"},
		map[string]map[string][]string{
			"pay": map[string][]string{
				"insert_coin": []string{"select"},
			},
			"select": map[string][]string{
				"": []string{"beer", "soda"},
			},
			"beer": map[string][]string{
				"get_beer": []string{"pay"},
			},
			"soda": map[string][]string{
				"get_soda": []string{"pay"},
			},
		},
		[]string{"paid", "drink"},
		map[string][]string{
			"pay":    []string{},
			"soda":   []string{"paid", "drink"},
			"beer":   []string{"paid", "drink"},
			"select": []string{"paid"},
		},
	)

	if err != nil {
		fmt.Println(fmt.Sprintf("There was an issue constructing the beverage vending machine! %v", err.Error()))
	}

	return ts0

}
