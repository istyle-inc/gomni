package float64

import (
	"github.com/istyle-inc/gomni/types"
	"testing"
)

func TestNVL(t *testing.T) {
	type args struct {
		s   *float64
		sub float64
	}
	testees := []struct {
		name string
		args args
		want float64
	}{
		// Success: returning sub when s is nil
		{
			name: "success to returning-sub",
			args: args{s: nil, sub: 10.0},
			want: 10.0,
		},
		// Success：returning s when s is not nil
		{
			name: "success to returning-s",
			args: args{s: types.Float64ToPtr(2.0), sub: 10.0},
			want: 2.0,
		},
	}
	for _, tt := range testees {
		if got := NVL(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVL() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
