package plugins

import (
	"strings"
)

// Args returns a slice of arguments from one or more interfaces.
func NewArgs(is ...interface{}) Args {
	var args []*Arg
	for n := range is {
		args = append(args, NewArg(n, is[n]))
	}
	return Args(args)
}

// Args is a slice of argumets that provides useful helpers.
type Args []*Arg

// Equals returns true IFF:
// 1. a and b are of the same length
// 2. a and b reference identical types in identical positions
func (a Args) Equals(b Args) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}
	return true
}

// String prints out the arguments and types.
// NOTE: this is good for usage information.
func (a Args) String() string {
	var s []string
	for _, arg := range a {
		s = append(s, arg.String())
	}
	return strings.Join(s, ",")
}
