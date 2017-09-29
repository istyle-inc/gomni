package types

import (
	"reflect"
	"time"
)

// StringToPtr return string pointer
func StringToPtr(s string) *string {
	return &s
}

// BoolToPtr return bool pointer
func BoolToPtr(b bool) *bool {
	return &b
}

// TimeToPtr return time pointer
func TimeToPtr(t time.Time) *time.Time {
	return &t
}

// IntToPtr return int pointer
func IntToPtr(i int) *int {
	return &i
}

// Int8ToPtr return int8 pointer
func Int8ToPtr(i int8) *int8 {
	return &i
}

// Int16ToPtr return int16 pointer
func Int16ToPtr(i int16) *int16 {
	return &i
}

// Int32ToPtr return int32 pointer
func Int32ToPtr(i int32) *int32 {
	return &i
}

// Int64ToPtr return int64 pointer
func Int64ToPtr(i int64) *int64 {
	return &i
}

// Float32ToPtr return float32 pointer
func Float32ToPtr(f float32) *float32 {
	return &f
}

// Float64ToPtr return float64 pointer
func Float64ToPtr(f float64) *float64 {
	return &f
}

// UintToPtr return uint pointer
func UintToPtr(i uint) *uint {
	return &i
}

// Uint8ToPtr return uint8 pointer
func Uint8ToPtr(i uint8) *uint8 {
	return &i
}

// Uint16ToPtr return uint16 pointer
func Uint16ToPtr(i uint16) *uint16 {
	return &i
}

// Uint32ToPtr return uint32 pointer
func Uint32ToPtr(i uint32) *uint32 {
	return &i
}

// Uint64ToPtr return uint64 pointer
func Uint64ToPtr(i uint64) *uint64 {
	return &i
}

// NVL return sub when s is nil
func NVL(s interface{}, sub interface{}) interface{} {
	if reflect.ValueOf(s).IsNil() {
		return sub
	}
	return DereferenceIfPtr(s)
}

// DereferenceIfPtr return dereference when v is interface
func DereferenceIfPtr(v interface{}) interface{} {
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		return reflect.ValueOf(v).Elem().Interface()
	}
	return v
}
