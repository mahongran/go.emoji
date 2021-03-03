package official

import (
	"strings"
	"testing"
)

func TestAllSequences(t *testing.T) {
	AllSequences.printAllSequences()
	return
}

func TestSequences_HasEmojiPrefix(t *testing.T) {
	tests := []struct {
		name       string
		s          string
		wantHas    bool
		wantLength int
	}{
		{
			s:          "1⃣️",
			wantHas:    true,
			wantLength: 7,
		},
		{
			s:          "123",
			wantHas:    false,
			wantLength: 0,
		},
		{
			s:          "👨‍👩‍👧‍👦123",
			wantHas:    true,
			wantLength: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHas, gotLength := AllSequences.HasEmojiPrefix(tt.s)
			if gotHas != tt.wantHas {
				t.Errorf("HasEmojiPrefix() gotHas = %v, want %v", gotHas, tt.wantHas)
			}
			if gotLength != tt.wantLength {
				t.Errorf("HasEmojiPrefix() gotLength = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}

func BenchmarkSequences_HasEmojiPrefix_BigString(b *testing.B) {
	s := "👨‍👩‍👧‍👦" + strings.Repeat("1234567890", 1000)
	for i := 0; i < b.N; i++ {
		AllSequences.HasEmojiPrefix(s)
	}
}

func BenchmarkSequences_HasEmojiPrefix_ShortString(b *testing.B) {
	s := "👨‍👩‍👧‍👦123"
	for i := 0; i < b.N; i++ {
		AllSequences.HasEmojiPrefix(s)
	}
}
