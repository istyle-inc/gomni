package int64

import n "gomni/numerics"

// NVL sub when s is nil
// Params s *int64
// Params sub int64
// Return int64
func NVL(s *int64, sub int64) int64 {
	ni := n.CommonNVL(s, sub)
	n, _ := ni.(int64)
	return n
}
