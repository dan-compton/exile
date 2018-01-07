package plugins

import "testing"

// TestEquals test the Equals method for each arg in the argument list.
func TestArgsEquals(t *testing.T) {
	var r, x = &[]int{1}[0], &[]int{2}[0]
	var z = &[]float64{1.0}[0]
	tests := []struct {
		a        []*Arg
		b        []*Arg
		expected []bool
	}{
		// test 1...n
		{
			// all true, primitive equality
			NewArgs("zero", 1, true, 3.0, "four"),
			NewArgs("zero", 1, true, 3.0, "four"),
			[]bool{true, true, true, true, true},
		},
		{
			// mixed primitive equality
			NewArgs("zero", 1, 2, 3),
			NewArgs(0, 1, "two", new(int)),
			[]bool{false, true, false, false},
		},
		{
			// nil should always return false
			NewArgs("a"),
			NewArgs(nil),
			[]bool{false},
		},
		{
			// pointer equality
			NewArgs(r),
			NewArgs(x),
			[]bool{true},
		},
		{
			// pointer inequality
			NewArgs(r),
			NewArgs(z),
			[]bool{true},
		},
	}

	for n, test := range tests {
		for i := range test.a {
			got := test.a[i].Equals(test.b[i])
			if got != test.expected[i] {
				ai := test.a[i].Value().Interface()
				bi := test.b[i].Value().Interface()
				t.Errorf("test #%d: got %t expected %t when comparing %v and %v", n+1, got, test.expected[i], ai, bi)
			}
		}
	}
}
