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
	s := "pwwkewdawdaddvdsdsfnodawkdnajchv"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 10
	b.Logf("len(s) = %d", len(s))
	b.ResetTimer() //タイマーのリセット

	for i := 0; i < b.N; i++ {
		actual := LengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for ;"+"expected %d", actual, ans)
		}
	}
}
