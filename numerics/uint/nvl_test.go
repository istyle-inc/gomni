package uint

import (
	"testing"

	"github.com/istyle-inc/gomni/types"
)

func TestNVL(t *testing.T) {
	type args struct {
		s   *uint
		sub uint
	}
	testees := []struct {
		name string
		args args
		want uint
	}{
		// Success: returning sub when s is nil
		{
			name: "success to returning-sub",
			args: args{s: nil, sub: uint(10)},
			want: uint(10),
		},
		// Success：returning s when s is not nil
		{
			name: "success to returning-s",
			args: args{s: types.UintToPtr(2), sub: uint(10)},
			want: uint(2),
		},
	}
	for _, tt := range testees {
		if got := NVL(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVL() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
