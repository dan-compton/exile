package plugins

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

// NewArg creates a new argument from the given position and interface, i.
func NewArg(position int, i interface{}) *Arg {
	return &Arg{
		position: position,
		t:        reflect.TypeOf(i),
		value:    reflect.ValueOf(i),
	}
}

// Arg represents a function argument.
// NOTE: Value should be a primitive, not a pointer.
type Arg struct {
	position int
	t        reflect.Type
	value    reflect.Value
}

// Type returns the argument's reflect.Type.
func (a *Arg) Type() reflect.Type {
	return a.t
}

// Value returns the arguments' reflect.Value.
func (a *Arg) Value() reflect.Value {
	return a.value
}

// Kind returns the argument's reflect.Kind.
func (a *Arg) Kind() reflect.Kind {
	return a.value.Kind()
}

// Position returns the argument's position.
func (a *Arg) Position() int {
	return a.position
}

// String satisfies the fmt.Stringer interface.
func (a *Arg) String() string {
	return fmt.Sprintf("argv%d<%v>", a.Position(), a.Kind())
}

// Equals returns true IFF:
// 1. a and b must have identical positions in the arguments list.
// 2. a and b must have the same reflect.Type.
// 3. nil values are never considered equal.
func (a *Arg) Equals(b *Arg) bool {
	if a.Position() != b.Position() {
		return false
	}
	if a.Kind() != b.Kind() {
		return false
	}
	return true
}

func isElemableKind(k reflect.Kind) bool {
	switch k {
	case reflect.Ptr:
		fallthrough
	case reflect.Interface:
		return true
	}
	return false
}

// Set sets the Value of i from the reflect.Value of a.
// If the value of i is invalid or cannot be set, false is returned.
// If the value of i is valid and settable, true is returned.
func (a *Arg) Set(i interface{}) error {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return errors.New("value to set is invalid")
	}
	if !isElemableKind(iv.Kind()) {
		return errors.Errorf("expected value to be of type reflect.Ptr or reflect.Kind, got %v", iv.Kind())
	}

	v := iv.Elem()
	if !v.CanSet() {
		return errors.New("value.Elem() to set is not settable")
	}
	v.Set(a.Value())
	return nil
}
