package uint16

import "github.com/istyle-inc/gomni/types"

// NVL sub when s is nil
// Params s *uint16
// Params sub uint16
// Return uint16
func NVL(s *uint16, sub uint16) uint16 {
	ni := types.NVL(s, sub)
	n, _ := ni.(uint16)
	return n
}
