package float64

import "github.com/istyle-inc/gomni/types"

// NVL sub when s is nil
// Params s *float64
// Params sub float64
// Return float64
func NVL(s *float64, sub float64) float64 {
	ni := types.NVL(s, sub)
	n, _ := ni.(float64)
	return n
}
