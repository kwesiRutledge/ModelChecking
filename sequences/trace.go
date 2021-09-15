//package sequences
/*
trace.go
Description:
	This is the labeled version of a given path.
*/

package sequences

import (
	mc "github.com/kwesiRutledge/ModelChecking"
)

/*
Type Definitions
*/

type FiniteTrace struct {
	L [][]mc.AtomicProposition
}

/*
InfiniteTrace
Description:
	This is made from a finite prefix (i.e. UniquePrefix) that contains a suffix which is
	RepeatingSuffix repeated infinitely.
*/
type InfiniteTrace struct {
	UniquePrefix    FiniteTrace
	RepeatingSuffix FiniteTrace
}

type Trace interface {
	SatisfiesAPInvariant(mc.AtomicProposition) bool
}

/*
Functions
*/

/*
SatisfiesAPInvariant()

*/
func (traceIn FiniteTrace) SatisfiesAPInvariant(apIn mc.AtomicProposition) bool {
	//If any subset in the trace does not contain apIn, then return false.
	for _, Ai := range traceIn.L {
		if !apIn.In(Ai) {
			return false
		}
	}
	// Otherwise return true.
	return true
}

func (traceIn InfiniteTrace) SatisfiesAPInvariant(apIn mc.AtomicProposition) bool {

	return traceIn.UniquePrefix.SatisfiesAPInvariant(apIn) && traceIn.RepeatingSuffix.SatisfiesAPInvariant()
}
