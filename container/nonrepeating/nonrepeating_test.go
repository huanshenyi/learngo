package main

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},
	}
	for _, tt := range tests {
		if res := LengthOfNonRepeatingSubStr(tt.s); res != tt.ans {
			t.Errorf("結果はあってません、文字列%s,重複しない長さ%d", tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "pwwkew"
	ans := 3

	for i := 0; i < b.N; i++ {
		actual := LengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s;"+"expected %d", actual, s, ans)
		}
	}
}
