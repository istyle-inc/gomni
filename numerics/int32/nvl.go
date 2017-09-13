package int32

import n "gomni/numerics"

// NVL sub when s is nil
// Params s *int32
// Params sub int32
// Return int32
func NVL(s *int32, sub int32) int32 {
	ni := n.CommonNVL(s, sub)
	n, _ := ni.(int32)
	return n
}
