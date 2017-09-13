package int

import n "gomni/numerics"

// NVL sub when s is nil
// Params s *int
// Params sub int
// Return int
func NVL(s *int, sub int) int {
	ni := n.CommonNVL(s, sub)
	n, _ := ni.(int)
	return n
}
