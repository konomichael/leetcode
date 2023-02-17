package q3

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	l := lengthOfLongestSubstring(s)
	if l != 3 {
		t.Errorf("got: %d, want: %d", l, 3)
	}
}
