package uint8

import "github.com/istyle-inc/gomni/types"

// NVL sub when s is nil
// Params s *uint8
// Params sub uint8
// Return uint8
func NVL(s *uint8, sub uint8) uint8 {
	ni := types.NVL(s, sub)
	n, _ := ni.(uint8)
	return n
}
