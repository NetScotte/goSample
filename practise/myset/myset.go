package myset

import (
	"strings"
)

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return make(StringSet)
}

func (set StringSet) String() string {
	var keys []string
	for key, _ := range set {
		keys = append(keys, key)
	}
	return strings.Join(keys, ", ")
}

func (set StringSet) Add(s string) {
	set[s] = struct{}{}
}

func (set StringSet) Remove(s string) {
	delete(set, s)
}
