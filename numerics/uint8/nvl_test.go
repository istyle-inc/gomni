package uint8

import (
	"testing"

	"github.com/istyle-inc/gomni/types"
)

func TestNVL(t *testing.T) {
	type args struct {
		s   *uint8
		sub uint8
	}
	testees := []struct {
		name string
		args args
		want uint8
	}{
		// Success: returning sub when s is nil
		{
			name: "success to returning-sub",
			args: args{s: nil, sub: uint8(10)},
			want: uint8(10),
		},
		// Success：returning s when s is not nil
		{
			name: "success to returning-s",
			args: args{s: types.Uint8ToPtr(2), sub: uint8(10)},
			want: uint8(2),
		},
	}
	for _, tt := range testees {
		if got := NVL(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVL() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
