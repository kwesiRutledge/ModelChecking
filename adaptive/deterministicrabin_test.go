/*
deterministicrabin_test.go
Description:
	Tests the functions and objects created in deterministicrabin.go
*/

package adaptive

import (
	"fmt"
	"testing"

	mc "github.com/kwesiRutledge/ModelChecking"
)

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

/*
TestDeterministicRabin_Equals1
Description:
	Verifies that we can properly compare a state object with another.
	In this example, the two states are the same.
*/
func TestDeterministicRabin_Equals1(t *testing.T) {
	// Create a Simple DRAState
	q0 := DRAState{
		Name: "Quaren",
	}

	if !q0.Equals(q0) {
		t.Errorf("The two DRAState objects are the same, but the function claims they are different!")
	}
}

/*
TestDeterministicRabin_Equals2
Description:
	Verifies that we can properly compare a state object with another.
	In this example, the two states are different.
*/
func TestDeterministicRabin_Equals2(t *testing.T) {
	// Create two Simple DRAState
	q0 := DRAState{
		Name: "Quaren",
	}

	q1 := DRAState{
		Name: "Dr. Umar",
	}

	if q0.Equals(q1) {
		t.Errorf("The two DRAState objects are different, but the Equals() functions claims they are the same!")
	}
}

/*
TestDeterministicRabin_Find1
Description:
	Verifies that we can properly find a state in a slice of DRAState objects.
	In this example, the desired state has index 2.
*/
func TestDeterministicRabin_Find1(t *testing.T) {
	// Create two Simple DRAState
	q0 := DRAState{
		Name: "Quaren",
	}

	q1 := DRAState{
		Name: "Dr. Umar",
	}

	q2 := DRAState{
		Name: "Pegasus",
	}

	qSlice0 := []DRAState{q0, q1, q2}

	// Test
	tf, q2Index := q2.Find(qSlice0)
	if !tf {
		t.Errorf("The function Find() cannot find the desired state in the slice qSlice!")
	}

	if q2Index != 2 {
		t.Errorf("The function Find() found q2 at index %v, but expected index 2!", q2Index)
	}
}

/*
TestDeterministicRabin_Find2
Description:
	Verifies that we can properly find a state in a slice of DRAState objects.
	In this example, the desired state is not in the slice.
*/
func TestDeterministicRabin_Find2(t *testing.T) {
	// Create two Simple DRAState
	q0 := DRAState{
		Name: "Quaren",
	}

	q1 := DRAState{
		Name: "Dr. Umar",
	}

	q2 := DRAState{
		Name: "Pegasus",
	}

	q3 := DRAState{
		Name: "Hercules",
	}

	qSlice0 := []DRAState{q0, q1, q2}

	// Test
	tf, q2Index := q3.Find(qSlice0)
	if tf {
		t.Errorf("The function Find() cannot find the desired state in the slice qSlice!")
	}

	if q2Index != -1 {
		t.Errorf("The function Find() found q2 at index %v, but expected index 2!", q2Index)
	}
}

/*
TestDeterministicRabin_In1
Description:
	Verifies that we can properly find a state in a slice of DRAState objects.
	In this example, the desired state IS in the target slice.
*/
func TestDeterministicRabin_In1(t *testing.T) {
	// Create two Simple DRAState
	q0 := DRAState{
		Name: "Quaren",
	}

	q1 := DRAState{
		Name: "Dr. Umar",
	}

	q2 := DRAState{
		Name: "Pegasus",
	}

	qSlice0 := []DRAState{q0, q1, q2}

	// Test
	tf := q2.In(qSlice0)
	if !tf {
		t.Errorf("The function In() cannot find the desired state in the slice qSlice!")
	}
}

/*
TestDeterministicRabin_In2
Description:
	Verifies that we can properly find a state in a slice of DRAState objects.
	In this example, the desired state is not in the slice.
*/
func TestDeterministicRabin_In2(t *testing.T) {
	// Create two Simple DRAState
	q0 := DRAState{
		Name: "Quaren",
	}

	q1 := DRAState{
		Name: "Dr. Umar",
	}

	q2 := DRAState{
		Name: "Pegasus",
	}

	q3 := DRAState{
		Name: "Hercules",
	}

	qSlice0 := []DRAState{q0, q1, q2}

	// Test
	tf := q3.In(qSlice0)
	if tf {
		t.Errorf("The function Find() cannot find the desired state in the slice qSlice!")
	}

}

/*
TestDeterministicRabin_CheckS0
Description:
	Verifies that the initial state set does not include the state set.
	Catches the appropriate value of the error.
*/
func TestDeterministicRabin_CheckS0(t *testing.T) {
	// Create Constants
	// Create two Simple DRAState
	q0 := DRAState{Name: "Quaren"}
	q1 := DRAState{Name: "Dr. Umar"}
	q2 := DRAState{Name: "Pegasus"}
	q3 := DRAState{Name: "Hercules"}

	var dra0 DeterministicRabinAutomaton
	dra0.S = []DRAState{q0, q1, q2}
	dra0.s0 = q3

	// Algorithm
	err := dra0.CheckS0()
	if err.Error() != fmt.Sprintf("The initial state \"%v\" was not in the state space S.", dra0.s0) {
		t.Errorf("The error created by err was not what we expected: %v", err)
	}
}

/*
TestDeterministicRabin_CheckAlpha0
Description:
	Verifies that the transition map has source states which are all a part of the state space.
*/
func TestDeterministicRabin_CheckAlpha1(t *testing.T) {
	// Create Constants
	// Create two Simple DRAState
	q0 := DRAState{Name: "Quaren"}
	q1 := DRAState{Name: "Dr. Umar"}
	q2 := DRAState{Name: "Pegasus"}
	q3 := DRAState{Name: "Hercules"}
	q4 := DRAState{Name: "Temptation"}

	ap0 := mc.AtomicProposition{Name: "red"}
	ap1 := mc.AtomicProposition{Name: "blue"}

	var dra0 DeterministicRabinAutomaton
	dra0.S = []DRAState{q0, q1, q2, q3}
	dra0.s0 = q3
	dra0.Alphabet = []mc.AtomicProposition{ap0, ap1}
	dra0.alpha = map[DRAState]map[mc.AtomicProposition]DRAState{
		q0: map[mc.AtomicProposition]DRAState{
			ap0: q0,
			ap1: q1,
		},
		q1: map[mc.AtomicProposition]DRAState{
			ap0: q1,
			ap1: q2,
		},
		q2: map[mc.AtomicProposition]DRAState{
			ap0: q2,
			ap1: q3,
		},
		q3: map[mc.AtomicProposition]DRAState{
			ap0: q3,
			ap1: q0,
		},
		q4: map[mc.AtomicProposition]DRAState{
			ap0: q3,
			ap1: q0,
		},
	}

	// Algorithm
	err := dra0.CheckAlpha()
	if err.Error() != fmt.Sprintf("The state \"%v\" is not in the state space.", q4) {
		t.Errorf("The error created by err was not what we expected: %v", err)
	}
}

/*
TestDeterministicRabin_CheckAlpha2
Description:
	Verifies that the transition map does not have a value alpha (or atomic proposition) in the function.
*/
func TestDeterministicRabin_CheckAlpha2(t *testing.T) {
	// Create Constants
	// Create two Simple DRAState
	q0 := DRAState{Name: "Quaren"}
	q1 := DRAState{Name: "Dr. Umar"}
	q2 := DRAState{Name: "Pegasus"}
	q3 := DRAState{Name: "Hercules"}

	ap0 := mc.AtomicProposition{Name: "red"}
	ap1 := mc.AtomicProposition{Name: "blue"}
	ap2 := mc.AtomicProposition{Name: "yellow"}

	var dra0 DeterministicRabinAutomaton
	dra0.S = []DRAState{q0, q1, q2, q3}
	dra0.s0 = q3
	dra0.Alphabet = []mc.AtomicProposition{ap0, ap1}
	dra0.alpha = map[DRAState]map[mc.AtomicProposition]DRAState{
		q0: map[mc.AtomicProposition]DRAState{
			ap0: q0,
			ap1: q1,
		},
		q1: map[mc.AtomicProposition]DRAState{
			ap0: q1,
			ap1: q2,
		},
		q2: map[mc.AtomicProposition]DRAState{
			ap0: q2,
			ap1: q3,
			ap2: q0,
		},
		q3: map[mc.AtomicProposition]DRAState{
			ap0: q3,
			ap1: q0,
		},
	}

	// Algorithm
	err := dra0.CheckAlpha()
	if err.Error() != fmt.Sprintf("The atomic proposition \"%v\" is not in the Alphabet.", ap2) {
		t.Errorf("The error created by err was not what we expected: %v", err)
	}
}

/*
TestDeterministicRabin_CheckAlpha3
Description:
	Verifies that the transition map does not have a post state value s1 which
	is in the state space..
*/
func TestDeterministicRabin_CheckAlpha3(t *testing.T) {
	// Create Constants
	// Create two Simple DRAState
	q0 := DRAState{Name: "Quaren"}
	q1 := DRAState{Name: "Dr. Umar"}
	q2 := DRAState{Name: "Pegasus"}
	q3 := DRAState{Name: "Hercules"}
	q4 := DRAState{Name: "Temptation"}

	ap0 := mc.AtomicProposition{Name: "red"}
	ap1 := mc.AtomicProposition{Name: "blue"}

	var dra0 DeterministicRabinAutomaton
	dra0.S = []DRAState{q0, q1, q2, q3}
	dra0.s0 = q3
	dra0.Alphabet = []mc.AtomicProposition{ap0, ap1}
	dra0.alpha = map[DRAState]map[mc.AtomicProposition]DRAState{
		q0: map[mc.AtomicProposition]DRAState{
			ap0: q0,
			ap1: q1,
		},
		q1: map[mc.AtomicProposition]DRAState{
			ap0: q1,
			ap1: q2,
		},
		q2: map[mc.AtomicProposition]DRAState{
			ap0: q2,
			ap1: q4,
		},
	}

	// Algorithm
	err := dra0.CheckAlpha()
	if err.Error() != fmt.Sprintf("The state \"%v\" is not in the state space.", q4) {
		t.Errorf("The error created by err was not what we expected: %v", err)
	}
}

/*
TestDeterministicRabin_CheckAlpha4
Description:
	Verifies that the alpha is correct.
*/
func TestDeterministicRabin_CheckAlpha4(t *testing.T) {
	// Create Constants
	// Create two Simple DRAState
	q0 := DRAState{Name: "Quaren"}
	q1 := DRAState{Name: "Dr. Umar"}
	q2 := DRAState{Name: "Pegasus"}
	q3 := DRAState{Name: "Hercules"}

	ap0 := mc.AtomicProposition{Name: "red"}
	ap1 := mc.AtomicProposition{Name: "blue"}

	var dra0 DeterministicRabinAutomaton
	dra0.S = []DRAState{q0, q1, q2, q3}
	dra0.s0 = q3
	dra0.Alphabet = []mc.AtomicProposition{ap0, ap1}
	dra0.alpha = map[DRAState]map[mc.AtomicProposition]DRAState{
		q0: map[mc.AtomicProposition]DRAState{
			ap0: q0,
			ap1: q1,
		},
		q1: map[mc.AtomicProposition]DRAState{
			ap0: q1,
			ap1: q2,
		},
		q2: map[mc.AtomicProposition]DRAState{
			ap0: q2,
			ap1: q3,
		},
		q3: map[mc.AtomicProposition]DRAState{
			ap0: q3,
			ap1: q0,
		},
	}

	// Algorithm
	err := dra0.CheckAlpha()
	if err != nil {
		t.Errorf("Expected error to be nil but it was %v", err)
	}

}

/*
TestDeterministicRabin_CheckOmega1
Description:
	Verifies that the Omega pairs are correct
*/
func TestDeterministicRabin_CheckOmega1(t *testing.T) {
	// Create Constants
	// Create two Simple DRAState
	q0 := DRAState{Name: "Quaren"}
	q1 := DRAState{Name: "Dr. Umar"}
	q2 := DRAState{Name: "Pegasus"}
	q3 := DRAState{Name: "Hercules"}

	ap0 := mc.AtomicProposition{Name: "red"}
	ap1 := mc.AtomicProposition{Name: "blue"}

	var dra0 DeterministicRabinAutomaton
	dra0.S = []DRAState{q0, q1, q2, q3}
	dra0.s0 = q3
	dra0.Alphabet = []mc.AtomicProposition{ap0, ap1}
	dra0.alpha = map[DRAState]map[mc.AtomicProposition]DRAState{
		q0: map[mc.AtomicProposition]DRAState{
			ap0: q0,
			ap1: q1,
		},
		q1: map[mc.AtomicProposition]DRAState{
			ap0: q1,
			ap1: q2,
		},
		q2: map[mc.AtomicProposition]DRAState{
			ap0: q2,
			ap1: q3,
		},
		q3: map[mc.AtomicProposition]DRAState{
			ap0: q3,
			ap1: q0,
		},
	}
	dra0.Omega = [][2][]DRAState{
		[2][]DRAState{[]DRAState{q0, q1}, []DRAState{q2, q3}},
		[2][]DRAState{[]DRAState{q0, q1, q2}, []DRAState{q3}},
		[2][]DRAState{[]DRAState{q0, q1, q2}, []DRAState{q1, q2, q3}},
	}

	// Algorithm
	err := dra0.CheckOmega()
	if err != nil {
		t.Errorf("Expected error to be nil but it was %v", err)
	}

}

/*
TestDeterministicRabin_CheckOmega2
Description:
	Verifies that the Omega pairs are not correct (first element in a pair is not a subset)
*/
func TestDeterministicRabin_CheckOmega2(t *testing.T) {
	// Create Constants
	// Create two Simple DRAState
	q0 := DRAState{Name: "Quaren"}
	q1 := DRAState{Name: "Dr. Umar"}
	q2 := DRAState{Name: "Pegasus"}
	q3 := DRAState{Name: "Hercules"}
	q4 := DRAState{Name: "Temptation"}

	ap0 := mc.AtomicProposition{Name: "red"}
	ap1 := mc.AtomicProposition{Name: "blue"}

	var dra0 DeterministicRabinAutomaton
	dra0.S = []DRAState{q0, q1, q2, q3}
	dra0.s0 = q3
	dra0.Alphabet = []mc.AtomicProposition{ap0, ap1}
	dra0.alpha = map[DRAState]map[mc.AtomicProposition]DRAState{
		q0: map[mc.AtomicProposition]DRAState{
			ap0: q0,
			ap1: q1,
		},
		q1: map[mc.AtomicProposition]DRAState{
			ap0: q1,
			ap1: q2,
		},
		q2: map[mc.AtomicProposition]DRAState{
			ap0: q2,
			ap1: q3,
		},
		q3: map[mc.AtomicProposition]DRAState{
			ap0: q3,
			ap1: q0,
		},
	}
	dra0.Omega = [][2][]DRAState{
		[2][]DRAState{[]DRAState{q0, q1}, []DRAState{q2, q3}},
		[2][]DRAState{[]DRAState{q0, q1, q2}, []DRAState{q3}},
		[2][]DRAState{[]DRAState{q0, q1, q2, q4}, []DRAState{q1, q2, q3}},
	}

	// Algorithm
	err := dra0.CheckOmega()
	if err.Error() != fmt.Sprintf("The %vth pair's first element is not a subset of the state space.", 2) {
		t.Errorf("Expected error in the 2th pair's first element, but received: %v", err)
	}

}

/*
TestDeterministicRabin_CheckOmega3
Description:
	Verifies that the Omega pairs are not correct (second element in a pair is not a subset)
*/
func TestDeterministicRabin_CheckOmega3(t *testing.T) {
	// Create Constants
	// Create two Simple DRAState
	q0 := DRAState{Name: "Quaren"}
	q1 := DRAState{Name: "Dr. Umar"}
	q2 := DRAState{Name: "Pegasus"}
	q3 := DRAState{Name: "Hercules"}
	q4 := DRAState{Name: "Temptation"}

	ap0 := mc.AtomicProposition{Name: "red"}
	ap1 := mc.AtomicProposition{Name: "blue"}

	var dra0 DeterministicRabinAutomaton
	dra0.S = []DRAState{q0, q1, q2, q3}
	dra0.s0 = q3
	dra0.Alphabet = []mc.AtomicProposition{ap0, ap1}
	dra0.alpha = map[DRAState]map[mc.AtomicProposition]DRAState{
		q0: map[mc.AtomicProposition]DRAState{
			ap0: q0,
			ap1: q1,
		},
		q1: map[mc.AtomicProposition]DRAState{
			ap0: q1,
			ap1: q2,
		},
		q2: map[mc.AtomicProposition]DRAState{
			ap0: q2,
			ap1: q3,
		},
		q3: map[mc.AtomicProposition]DRAState{
			ap0: q3,
			ap1: q0,
		},
	}
	dra0.Omega = [][2][]DRAState{
		[2][]DRAState{[]DRAState{q0, q1}, []DRAState{q2, q3}},
		[2][]DRAState{[]DRAState{q0, q1, q2}, []DRAState{q3, q4}},
		[2][]DRAState{[]DRAState{q0, q1, q2}, []DRAState{q1, q2, q3}},
	}

	// Algorithm
	err := dra0.CheckOmega()
	if err.Error() != fmt.Sprintf("The %vth pair's second element is not a subset of the state space.", 1) {
		t.Errorf("Expected error in the 1th pair's second element, but received: %v", err)
	}

}
