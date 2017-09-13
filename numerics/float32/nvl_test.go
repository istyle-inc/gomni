package float32

import (
	"gomni/types"
	"testing"
)

func TestNVL(t *testing.T) {
	type args struct {
		s   *float32
		sub float32
	}
	testees := []struct {
		name string
		args args
		want float32
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
			args: args{s: types.Float32ToPtr(2), sub: 2.0},
			want: 2.0,
		},
	}
	for _, tt := range testees {
		if got := NVL(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVL() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
