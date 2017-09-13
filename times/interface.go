package times

import "time"

// TimeAdaptor interface to mock Date
type TimeAdaptor interface {
	Nowf(format string) string
	Now() time.Time
}
