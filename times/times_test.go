package times

import (
	"strconv"
	"testing"
	"time"
)

// TestNowf Test for Nowf
func TestNowf(t *testing.T) {
	year := strconv.Itoa(time.Now().Year())
	testee := &Time{}
	if testee.Nowf("2006") != year {
		t.Error("Failed: something wrong with getting time ")
	}
}
