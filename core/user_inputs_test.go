package core

import (
	"testing"
)

func TestUserInputToBool(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"lowercase y", "y", true},
		{"uppercase Y", "Y", true},
		{"mixed case Y", "y", true},
		{"lowercase yes", "yes", true},
		{"uppercase YES", "YES", true},
		{"mixed case Yes", "Yes", true},
		{"mixed case yEs", "yEs", true},
		{"with spaces", "  yes  ", true},
		{"with tabs", "\tyes\t", true},

		{"lowercase n", "n", false},
		{"uppercase N", "N", false},
		{"lowercase no", "no", false},
		{"uppercase NO", "NO", false},
		{"mixed case No", "No", false},
		{"empty string", "", false},
		{"whitespace only", "   ", false},
		{"other word", "maybe", false},
		{"partial match", "ye", false},
		{"number", "1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserInputToBool(tt.input); got != tt.expected {
				t.Errorf("UserInputToBool(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
