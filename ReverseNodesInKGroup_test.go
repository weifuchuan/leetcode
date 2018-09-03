package leetcode

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head:
				&ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 3, Next: &ListNode{
								Val: 4, Next: &ListNode{
									Val: 5, Next: &ListNode{
										Val: 6, Next: &ListNode{
											Val: 7, Next: &ListNode{
												Val: 8}}}}}}}},
				k: 2,
			},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 6, Next: &ListNode{Val: 5, Next: &ListNode{Val: 8, Next: &ListNode{Val: 7}}}}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseKGroup(tt.args.head, tt.args.k); !listEqual(got, tt.want) {
				t.Errorf("ReverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func Test_reverse(t *testing.T) {
//	list := &ListNode{
//		Val: 1, Next: &ListNode{
//			Val: 2, Next: &ListNode{
//				Val: 3, Next: &ListNode{
//					Val: 4, Next: &ListNode{
//						Val: 5, Next: &ListNode{
//							Val: 6, Next: &ListNode{
//								Val: 7, Next: &ListNode{
//									Val: 8}}}}}}}}
//	b,e:=reverse(list, list.Next.Next)
//	if !listEqual(got, tt.want) {
//		t.Errorf("ReverseKGroup() = %v, want %v", got, tt.want)
//	}
//}
