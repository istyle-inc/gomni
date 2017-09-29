package uint32

import "github.com/istyle-inc/gomni/types"

// NVL sub when s is nil
// Params s *uint32
// Params sub uint32
// Return uint32
func NVL(s *uint32, sub uint32) uint32 {
	ni := types.NVL(s, sub)
	n, _ := ni.(uint32)
	return n
}
