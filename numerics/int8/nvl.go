package int8

import n "gomni/numerics"

// NVL sub when s is nil
// Params s *int8
// Params sub int8
// Return int8
func NVL(s *int8, sub int8) int8 {
	ni := n.CommonNVL(s, sub)
	n, _ := ni.(int8)
	return n
}
