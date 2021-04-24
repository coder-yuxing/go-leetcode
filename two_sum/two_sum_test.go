package two_sum

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2,7,11,15},
				target: 9,
			},
			want: []int{0,1},
		},
		{
			name: "case2",
			args: args{
				nums: []int{3,2,3},
				target: 6,
			},
			want: []int{0,2},
		},
		{
			name: "case3",
			args: args{
				nums: []int{3,3},
				target: 6,
			},
			want: []int{0,1},
		},
		{
			name: "case4",
			args: args{
				nums: []int{-3,4,3,90},
				target: 0,
			},
			want: []int{0,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.args.nums, tt.args.target)
			sort.Ints(got)
			sort.Ints(tt.want)
			if  !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}

}
