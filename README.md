# ModelChecking

An implementation of the algorithms in Baier and Katoen (and maybe more!). Feel free to message me about desired functionality or errors that you might find.

## Including this module in your code

```
import "github.com/kwesiRutledge/ModelChecking"
```

## Examples

An example file which uses the library:
```
package main

import (
	"fmt"
    "os"
	mc "github.com/kwesiRutledge/ModelChecking"
)

func main() {
	ts0, err := mc.GetTransitionSystem(
		[]string{"1", "2", "3"}, []string{"1", "2"},
		map[string]map[string][]string{
			"1": map[string][]string{
				"1": []string{"1"},
				"2": []string{"2"},
			},
			"2": map[string][]string{
				"1": []string{"1", "2"},
				"2": []string{"2", "3"},
			},
			"3": map[string][]string{
				"1": []string{"3"},
				"2": []string{"2"},
			},
		},
		[]string{"1"},
		[]string{"A", "B", "C", "D"},
		map[string][]string{
			"1": []string{"A"},
			"2": []string{"B", "D"},
			"3": []string{"C", "D"},
		},
	)

    if err != nil {
        fmt.Println(fmt.Sprintf("There was an issue creating the transition system: %v", err ))
        os.Exit(1)
    }

    // Create an atomic proposition
    ap2 := ts0.AP[1]
    state2 := ts0.S[1]

	tf, _ := state2.Satisfies(ap2)
    fmt.Println(tf)
}

```
