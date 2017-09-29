package uint

import "github.com/istyle-inc/gomni/types"

// NVL sub when s is nil
// Params s *uint
// Params sub uint
// Return uint
func NVL(s *uint, sub uint) uint {
	ni := types.NVL(s, sub)
	n, _ := ni.(uint)
	return n
}
