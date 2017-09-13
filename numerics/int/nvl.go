package int

import "github.com/istyle-inc/gomni/types"

// NVL sub when s is nil
// Params s *int
// Params sub int
// Return int
func NVL(s *int, sub int) int {
	ni := types.NVL(s, sub)
	n, _ := ni.(int)
	return n
}
