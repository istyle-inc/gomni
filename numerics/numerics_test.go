package numerics

import (
	"gomni/types"
	"testing"
)

func TestNVLInt(t *testing.T) {
	type args struct {
		s   *int
		sub int
	}
	testees := []struct {
		name string
		args args
		want int
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
			args: args{s: types.IntToPtr(2), sub: 10},
			want: 2,
		},
	}
	for _, tt := range testees {
		if got := NVLInt(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVLInt() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNVLInt8(t *testing.T) {
	type args struct {
		s   *int8
		sub int8
	}
	testees := []struct {
		name string
		args args
		want int8
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
			args: args{s: types.Int8ToPtr(2), sub: 10},
			want: 2,
		},
	}
	for _, tt := range testees {
		if got := NVLInt8(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVLInt() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNVLInt16(t *testing.T) {
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
		if got := NVLInt16(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVLInt16() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNVLInt32(t *testing.T) {
	type args struct {
		s   *int32
		sub int32
	}
	testees := []struct {
		name string
		args args
		want int32
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
			args: args{s: types.Int32ToPtr(2), sub: 10},
			want: 2,
		},
	}
	for _, tt := range testees {
		if got := NVLInt32(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVLInt32() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNVLInt64(t *testing.T) {
	type args struct {
		s   *int64
		sub int64
	}
	testees := []struct {
		name string
		args args
		want int64
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
			args: args{s: types.Int64ToPtr(2), sub: 10},
			want: 2,
		},
	}
	for _, tt := range testees {
		if got := NVLInt64(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVLInt64() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNVLFloat32(t *testing.T) {
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
		if got := NVLFloat32(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVLFloat32() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNVLFloat64(t *testing.T) {
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
		if got := NVLFloat64(tt.args.s, tt.args.sub); got != tt.want {
			t.Errorf("%q. NVLFloat64() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
