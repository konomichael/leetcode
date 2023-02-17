package datastructure

import "testing"

func TestIsListEqual(t *testing.T) {
	type args struct {
		a *ListNode
		b *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "both nil",
			args: args{
				a: nil,
				b: nil,
			},
			want: true,
		},
		{
			name: "should equal",
			args: args{
				a: MakeListNode(1, 3, 5),
				b: MakeListNode(1, 3, 5),
			},
			want: true,
		},
		{
			name: "should not equal",
			args: args{
				a: MakeListNode(1, 3, 5),
				b: MakeListNode(2, 4, 6),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsListEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsListEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
