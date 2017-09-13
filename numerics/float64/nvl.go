package float64

import n "gomni/numerics"

// NVL sub when s is nil
// Params s *float64
// Params sub float64
// Return float64
func NVL(s *float64, sub float64) float64 {
	ni := n.CommonNVL(s, sub)
	n, _ := ni.(float64)
	return n
}
