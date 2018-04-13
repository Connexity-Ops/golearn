package mapsort

import (
	"sort"
)

// Pair creates a struct to hold Key/Value pairs
type Pair struct {
	Key   string
	Value int
}

// PairList initializes a slice of Pair structs
type PairList []Pair

// MapSort sorts a map's Keys by Value, and returns it as an ordered slice of []Pair
func MapSort(m map[string]int) PairList {
	var kvPairs PairList
	for k, v := range m {
		kvPairs = append(kvPairs, Pair{Key: k, Value: v})
	}
	sort.Slice(kvPairs, func(i, j int) bool {
		return kvPairs[i].Value < kvPairs[j].Value
	})
	return kvPairs
}
