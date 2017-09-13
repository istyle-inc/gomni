package numerics

import (
	"gomni/types"
	"testing"
)

func TestCommonNVL(t *testing.T) {
	type args struct {
		s   interface{}
		sub interface{}
	}
	var nili *interface{}
	testees := []struct {
		name string
		args args
		want interface{}
	}{
		// Success: returning sub when s is nil interface
		{
			name: "success to returning-sub",
			args: args{s: nili, sub: 10},
			want: 10,
		},
		// Success：returning s when s is not nil interface
		{
			name: "success to returning-s",
			args: args{s: types.IntToPtr(2), sub: 10},
			want: 2,
		},
	}
	for _, tt := range testees {
		if got := CommonNVL(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. CommonNVL() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
