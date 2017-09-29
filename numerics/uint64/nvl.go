package uint64

import "github.com/istyle-inc/gomni/types"

// NVL sub when s is nil
// Params s *uint64
// Params sub uint64
// Return uint64
func NVL(s *uint64, sub uint64) uint64 {
	ni := types.NVL(s, sub)
	n, _ := ni.(uint64)
	return n
}
