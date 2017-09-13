package types

import (
	"testing"
	"time"
)

// TestStringToPtr Test for StringToPtr
func TestStringToPtr(t *testing.T) {
	value := "test"
	ptr := StringToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestBoolToPtr Test for BoolToPtr
func TestBoolToPtr(t *testing.T) {
	value := true
	ptr := BoolToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestTimeToPtr Test for TimeToPtr
func TestTimeToPtr(t *testing.T) {
	value := time.Now()
	ptr := TimeToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestIntToPtr Test for IntToPtr
func TestIntToPtr(t *testing.T) {
	value := 1
	ptr := IntToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestInt8ToPtr Test for Int8ToPtr
func TestInt8ToPtr(t *testing.T) {
	var value int8 = 1
	ptr := Int8ToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestInt16ToPtr Test for Int16ToPtr
func TestInt16ToPtr(t *testing.T) {
	var value int16 = 1
	ptr := Int16ToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestInt32ToPtr Test for Int32ToPtr
func TestInt32ToPtr(t *testing.T) {
	var value int32 = 1
	ptr := Int32ToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestInt64ToPtr Test for Int64ToPtr
func TestInt64ToPtr(t *testing.T) {
	var value int64 = 1
	ptr := Int64ToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestFloat32ToPtr Test for Float32ToPtr
func TestFloat32ToPtr(t *testing.T) {
	var value float32 = 6.413241
	ptr := Float32ToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestFloat64ToPtr Test for Float64ToPtr
func TestFloat64ToPtr(t *testing.T) {
	value := 6.413241
	ptr := Float64ToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestNVL Test for NVL
func TestNVL(t *testing.T) {
	type args struct {
		s   interface{}
		sub interface{}
	}
	var nili *interface{}
	testees := []struct {
		name string
		args args
		want interface{}
	}{
		// Success: returning sub when s is nil interface
		{
			name: "success to returning-sub",
			args: args{s: nili, sub: 10},
			want: 10,
		},
		// Success：returning s when s is not nil interface
		{
			name: "success to returning-s",
			args: args{s: IntToPtr(2), sub: 10},
			want: 2,
		},
	}
	for _, tt := range testees {
		if got := NVL(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVL() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestDereferenceIfPtr(t *testing.T) {
	testees := []struct {
		name string
		args interface{}
		want interface{}
	}{
		// Success: returning dereference when v is pointer
		{
			name: "success to returning dereference value",
			args: IntToPtr(10),
			want: 10,
		},
		// Success：returning v when v is not pointer
		{
			name: "success to returning v",
			args: 2,
			want: 2,
		},
	}
	for _, tt := range testees {
		if got := DereferenceIfPtr(tt.args); got != tt.want {
			t.Errorf("%q. DereferenceIfPtr() = %v, want %v", tt.name, got, tt.want)
		}
	}

}
