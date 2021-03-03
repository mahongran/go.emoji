package emoji

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestUnicode(t *testing.T) {
	printf := t.Logf
	buff := bytes.Buffer{}

	// ==== Emoji_presentation_sequence ====

	// thunder default mode
	buff.WriteRune(0x26A1)
	printf("%s", buff.String())
	buff.Reset()

	// thunder text mode
	buff.WriteRune(0x26A1)
	buff.WriteRune(0xFE0E)
	printf("%s", buff.String())
	buff.Reset()

	// thunder emoji mode
	buff.WriteRune(0x26A1)
	buff.WriteRune(0xFE0F)
	printf("%s", buff.String())
	buff.Reset()

	// ==== Modfier_Base_Sequence ====
	buff.WriteRune(0x270D)
	buff.WriteRune(0xFE0F)
	printf("%s", buff.String())
	buff.Reset()

	buff.WriteRune(0x270D)
	buff.WriteRune(0x1F3FF)
	printf("%s", buff.String())
	buff.Reset()

	// ==== Flag Sequence ====
	// reference: https://en.wikipedia.org/wiki/Regional_Indicator_Symbol
	// from 0x1F1E6 to 0x1F1FF
	buff.WriteRune(0x1F1E8)
	buff.WriteRune(0x1F1F3)
	// buff.WriteRune(0x1F1FA)
	// buff.WriteRune(0x1F1F8)
	printf("%s", buff.String())
	buff.Reset()

	// ==== KeyCap Sequence ====
	// 012345678*#
	buff.WriteRune('#')
	buff.WriteRune(0xFE0F)
	buff.WriteRune(0x20E3)
	printf("%s", buff.String())
	buff.Reset()

	// ==== ZWJ sequence ====
	buff.WriteRune(0x1F468)
	buff.WriteRune(0x200D)
	buff.WriteRune(0x1F469)
	buff.WriteRune(0x200D)
	buff.WriteRune(0x1F467)
	buff.WriteRune(0x200D)
	buff.WriteRune(0x1F466)
	buff.WriteRune(0x1F46a)
	printf("%s", buff.String())
	buff.Reset()

	// ==== Tag Sequence ====
	buff.WriteRune(0x1F3F4)
	buff.WriteRune(0xE0067)
	buff.WriteRune(0xE0062)
	buff.WriteRune(0xE0065)
	buff.WriteRune(0xE006E)
	buff.WriteRune(0xE0067)
	buff.WriteRune(0xE007F)
	printf("%s", buff.String())
	buff.Reset()
}

func TestReplaceAllEmojiFunc(t *testing.T) {
	printf := t.Logf

	s := "👩‍👩‍👦🇨🇳"
	i := 0

	final := replaceAllEmojiFunc(s, func(emoji string) string {
		i++
		printf("%02d - %s - len %d", i, emoji, len(emoji))
		return fmt.Sprintf("%d-", i)
	})

	printf("final: <%s>", final)
}

/*
goos: darwin
goarch: amd64
pkg: github.com/go-xman/go.emoji
BenchmarkReplaceAllEmojiFunc
BenchmarkReplaceAllEmojiFunc-8   	   14910	     79685 ns/op
PASS
*/
func BenchmarkReplaceAllEmojiFunc(b *testing.B) {
	s := "👩‍👩‍👦🇨🇳" + strings.Repeat("abc", 1000)
	for i := 0; i < b.N; i++ {
		_ = replaceAllEmojiFunc(s, func(emoji string) string {
			return ""
		})
	}
}

func TestHasEmoji(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			s:    "一",
			want: false,
		},
		{
			s:    "1⃣️",
			want: true,
		},
		{
			s:    "👩‍👩‍👦",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasEmoji(tt.s); got != tt.want {
				t.Errorf("HasEmoji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterEmoji(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			s:    "1⃣️",
			want: "",
		},
		{
			s:    "1⃣️23",
			want: "23",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterEmoji(tt.s); got != tt.want {
				t.Errorf("FilterEmoji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceEmoji(t *testing.T) {
	type args struct {
		s string
		f func(emoji string) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "1⃣️234⃣️",
				f: func(emoji string) string {
					return "#"
				},
			},
			want: "#23#",
		},
		{
			args: args{
				s: "1⃣️234⃣️",
				f: func(emoji string) string {
					if emoji == "1⃣️" {
						return "1"
					} else {
						return ""
					}
				},
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceEmoji(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("ReplaceEmoji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHumanReadLen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			s:    "123",
			want: 3,
		},
		{
			s:    "1⃣️23",
			want: 3,
		},
		{
			s:    "👩‍👩‍👦🇨🇳3",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HumanReadLen(tt.s); got != tt.want {
				t.Errorf("HumanReadLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDump(t *testing.T) {
	Dump("👨‍👩‍👧‍👦")
}
