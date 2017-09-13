package int16

import (
	"gomni/types"
	"testing"
)

func TestNVL(t *testing.T) {
	type args struct {
		s   *int16
		sub int16
	}
	testees := []struct {
		name string
		args args
		want int16
	}{
		// Success: returning sub when s is nil
		{
			name: "success to returning-sub",
			args: args{s: nil, sub: 10},
			want: 10,
		},
		// Success：returning s when s is not nil
		{
			name: "success to returning-s",
			args: args{s: types.Int16ToPtr(2), sub: 10},
			want: 2,
		},
	}
	for _, tt := range testees {
		if got := NVL(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVL() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
