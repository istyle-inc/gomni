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

// TestFloat32ToPtr Float32ToPtrのテスト
func TestFloat32ToPtr(t *testing.T) {
	var value float32 = 6.413241
	ptr := Float32ToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}

// TestFloat64ToPtr Float64ToPtrのテスト
func TestFloat64ToPtr(t *testing.T) {
	value := 6.413241
	ptr := Float64ToPtr(value)
	if *ptr != value {
		t.Errorf("Pointer has pointed wrong value expected %v, actually %v", value, *ptr)
	}
}
