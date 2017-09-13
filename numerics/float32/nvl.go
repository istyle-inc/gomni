package float32

import "github.com/istyle-inc/gomni/types"

// NVL sub when s is nil
// Params s *float32
// Params sub float32
// Return float32
func NVL(s *float32, sub float32) float32 {
	ni := types.NVL(s, sub)
	n, _ := ni.(float32)
	return n
}
