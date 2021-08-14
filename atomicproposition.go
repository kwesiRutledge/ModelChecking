/*
atomicproposition.go
Description:
 	Basic implementation of an Atomic Proposition object.
*/
package main

import (
	"errors"

	combinations "github.com/mxschmitt/golang-combinations"
)

type AtomicProposition struct {
	Name string
}

/*
Equals
Description:
	Compares the atomic propositions by their names.
*/
func (ap1 AtomicProposition) Equals(ap2 AtomicProposition) bool {

	return ap1.Name == ap2.Name

}

/*
StringSliceToAPs
Description:
	Transforms a slice of int variables into a list of AtomicPropositions
*/
func StringSliceToAPs(stringSlice []string) []AtomicProposition {
	var APList []AtomicProposition
	for _, apName := range stringSlice {
		APList = append(APList, AtomicProposition{Name: apName})
	}

	return APList
}

/*
In
Description:
	Determines if the Atomic Proposition is in a slice of atomic propositions.
*/
func (ap AtomicProposition) In(apSliceIn []AtomicProposition) bool {
	for _, tempAP := range apSliceIn {
		if ap.Equals(tempAP) {
			return true
		}
	}
	return false
}

/*
ToSliceOfAtomicPropositions
Description:
	Attempts to convert an arbitrary slice into a slice of atomic propositions.
*/
func ToSliceOfAtomicPropositions(slice1 interface{}) ([]AtomicProposition, error) {
	// Attempt To Cast
	castedSlice1, ok1 := slice1.([]AtomicProposition)

	if ok1 {
		return castedSlice1, nil
	} else {
		return []AtomicProposition{}, errors.New("There was an issue casting the slice into a slice of Atomic Propositions.")
	}

}

/*
Powerset
Description:
	Creates all possible subsets of the input array of atomic propositions.
*/
func Powerset(setOfAPs []AtomicProposition) [][]AtomicProposition {
	var AllCombinations [][]AtomicProposition
	var AllCombinationsAsStrings [][]string

	var AllNames []string
	for _, ap := range setOfAPs {
		AllNames = append(AllNames, ap.Name)
	}

	AllCombinationsAsStrings = combinations.All(AllNames)

	for _, tempStringSlice := range AllCombinationsAsStrings {
		AllCombinations = append(AllCombinations, StringSliceToAPs(tempStringSlice))
	}

	AllCombinations = append(AllCombinations, []AtomicProposition{})

	return AllCombinations
}
