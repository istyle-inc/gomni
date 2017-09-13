package numerics

// NVLInt sub when s is nil
// Params s *int
// Params sub int
// Return int
func NVLInt(s *int, sub int) int {
	if s == nil {
		return sub
	}
	return *s
}

// NVLInt8 sub when s is nil
// Params s *int8
// Params sub int8
// Return int8
func NVLInt8(s *int8, sub int8) int8 {
	if s == nil {
		return sub
	}
	return *s
}

// NVLInt16 sub when s is nil
// Params s *int16
// Params sub int16
// Return int16
func NVLInt16(s *int16, sub int16) int16 {
	if s == nil {
		return sub
	}
	return *s
}

// NVLInt32 sub when s is nil
// Params s *int32
// Params sub int32
// Return int32
func NVLInt32(s *int32, sub int32) int32 {
	if s == nil {
		return sub
	}
	return *s
}

// NVLInt64 sub when s is nil
// Params s *int64
// Params sub int64
// Return int64
func NVLInt64(s *int64, sub int64) int64 {
	if s == nil {
		return sub
	}
	return *s
}

// NVLFloat32 sub when s is nil
// Params s *float32
// Params sub float32
// Return float32
func NVLFloat32(s *float32, sub float32) float32 {
	if s == nil {
		return sub
	}
	return *s
}

// NVLFloat64 sub when s is nil
// Params s *float64
// Params sub float64
// Return float64
func NVLFloat64(s *float64, sub float64) float64 {
	if s == nil {
		return sub
	}
	return *s
}
