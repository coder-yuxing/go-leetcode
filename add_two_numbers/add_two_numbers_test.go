package add_two_numbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				l1: convert2ListNode([]int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}),
				l2: convert2ListNode([]int{5,6,4}),
			},
			want: convert2ListNode([]int{6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}),
		},
		{
			name: "case2",
			args: args{
				l1: convert2ListNode([]int{5}),
				l2: convert2ListNode([]int{5}),
			},
			want: convert2ListNode([]int{0,1}),
		},
		{
			name: "case4",
			args: args{
				l1: convert2ListNode([]int{0}),
				l2: convert2ListNode([]int{1}),
			},
			want: convert2ListNode([]int{1}),
		},
		{
			name: "case5",
			args: args{
				l1: convert2ListNode([]int{0}),
				l2: convert2ListNode([]int{0}),
			},
			want: convert2ListNode([]int{0}),
		},
		{
			name: "case6",
			args: args{
				l1: convert2ListNode([]int{2,4,3}),
				l2: convert2ListNode([]int{5,6,4}),
			},
			want: convert2ListNode([]int{7,0,8}),
		},
		{
			name: "case7",
			args: args{
				l1: convert2ListNode([]int{9,9,9,9,9,9,9}),
				l2: convert2ListNode([]int{9,9,9,9}),
			},
			want: convert2ListNode([]int{8,9,9,9,0,0,0,1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}


