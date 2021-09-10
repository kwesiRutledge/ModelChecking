/*
path.go
Description:
	Objects which are finite path fragments.
*/
package modelchecking

import "fmt"

/*
Type Declarations
*/

type FinitePathFragment struct {
	s []TransitionSystemState
}

type InfinitePathFragment struct {
	UniquePrefix    FinitePathFragment
	RepeatingSuffix FinitePathFragment
}

// Functions

func (pathIn FinitePathFragment) Check() error {
	// Verify that the transitions in the path fragment are okay
	for sIndex := 0; sIndex < len(pathIn.s)-1; sIndex++ {
		si := pathIn.s[sIndex]
		sip1 := pathIn.s[sIndex+1]

		siAncestors, err := Post(si)
		if err != nil {
			return fmt.Errorf("There was an issue computing the %vth post (Post(%v)): %v", sIndex, si, err)
		}

		if !sip1.In(siAncestors) {
			return fmt.Errorf(
				"The %vth state (%v) is not in the post of the %vth state (%v).",
				sIndex+1, sip1, sIndex, si,
			)
		}
	}

	// Return nothing if all transitions are correct
	return nil
}
