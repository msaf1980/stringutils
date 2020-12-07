package stringutils

import (
	"reflect"
	"strings"
	"testing"
)

func Test_Split2(t *testing.T) {
	tests := []struct {
		s           string
		sep         string
		want0       string
		want1       string
		wantStrings int
	}{
		{"", "", "", "", 1},
		{"", "&", "", "", 1},
		{"test", "&", "test", "", 1},
		{"test&", "&", "test", "", 2},
		{"test&after", "&", "test", "after", 2},
	}
	for _, tt := range tests {
		t.Run(tt.s+" -> "+tt.sep, func(t *testing.T) {
			s0, s1, n := Split2(tt.s, tt.sep)
			if s0 != tt.want0 {
				t.Errorf("Split2() s[0] = %v, want %v", s0, tt.want0)
			}
			if s1 != tt.want1 {
				t.Errorf("Split2() s[1] = %v, want %v", s1, tt.want1)
			}
			if n != tt.wantStrings {
				t.Errorf("Split2() count = %v, want %v", n, tt.wantStrings)
			}
		})
	}
}

// for compare with Benchmark_Split2
func Benchmark_Split(b *testing.B) {
	s := "teststring&where"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = strings.SplitN(s, "&", 1)
	}
}

func Benchmark_Split2(b *testing.B) {
	s := "teststring&where"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = Split2(s, "&")
	}
}

func TestSplitN(t *testing.T) {
	buf := make([]string, 4)
	tests := []struct {
		s       string
		sep     string
		want    []string
		wantPos int
	}{
		{"", ",", []string{""}, 0},
		{"test", ",", []string{"test"}, 4},
		{"test1,2", ",", []string{"test1", "2"}, 7},
		{"test1.2.test3.", ".", []string{"test1", "2", "test3", ""}, 14},
		{"test1.2.test3.4", ".", []string{"test1", "2", "test3", "4"}, 15},
		{"test1.2.test3.4.", ".", []string{"test1", "2", "test3", "4"}, 16},
		{"test1.2.test3.4.5", ".", []string{"test1", "2", "test3", "4"}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got, gotPos := SplitN(tt.s, tt.sep, buf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitN() = %v, want %v", got, tt.want)
			}
			if gotPos != tt.wantPos {
				t.Errorf("SplitN() pos = %d, want %d", gotPos, tt.wantPos)
			}
		})
	}
}

func Benchmark_SplitN(b *testing.B) {
	buf := make([]string, 4)
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = SplitN(s, ".", buf)
	}
}