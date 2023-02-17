package q206

import (
	"testing"

	"github.com/konomichael/leetcode/datastructure"
)

func Test_reverseList(t *testing.T) {
	l := datastructure.MakeListNode(1, 3, 5, 7)
	r := reverseList(l)
	w := datastructure.MakeListNode(7, 5, 3, 1)
	if !datastructure.IsListEqual(w, r) {
		t.Errorf("want: %s, got: %s", w, r)
	}
}
