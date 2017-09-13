package times

import (
	"time"
)

// Time implements of TimeAdaptor interface
type Time struct{}

// Nowf return now time in accordance with the format
func (t *Time) Nowf(format string) string {
	return t.Now().Format(format)
}

// Now return now time
func (t *Time) Now() time.Time {
	return time.Now()
}

var (
	// Timer public variable
	Timer TimeAdaptor
)

// initialize
func init() {
	Timer = &Time{}
}
