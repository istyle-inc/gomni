package int16

import "github.com/istyle-inc/gomni/types"

// NVL sub when s is nil
// Params s *int16
// Params sub int16
// Return int16
func NVL(s *int16, sub int16) int16 {
	ni := types.NVL(s, sub)
	n, _ := ni.(int16)
	return n
}
