package numerics

import "gomni/types"

import "reflect"

// CommonNVL sub when s is nil
// Params s interface
// Params sub interface
// Return interface
func CommonNVL(s interface{}, sub interface{}) interface{} {
	if reflect.ValueOf(s).IsNil() {
		return sub
	}
	return types.DereferenceIfPtr(s)
}
