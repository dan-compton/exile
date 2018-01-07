package plugins

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArgSet(t *testing.T) {
	tests := []struct {
		arg      *Arg
		Got      interface{} // type that is set during test
		Expected interface{} // pre-set type
	}{
		// test 1...n
		{
			NewArg(0, "hi"),
			"1",
			"hi",
		},
		{
			NewArg(0, 1.0),
			0.0,
			1.0,
		},
	}

	for n, test := range tests {
		err := test.arg.Set(&test.Got)
		if err != nil {
			t.Errorf("test %d: set error %s, expected nil", n+1, err)
		}
		if !reflect.DeepEqual(test.Got, test.Expected) {
			Got := fmt.Sprintf("%v", test.Got)
			exp := fmt.Sprintf("%v", test.Expected)
			t.Errorf("test %d: set resulted in %v, Expected %v", n+1, Got, exp)
		}
	}
}
