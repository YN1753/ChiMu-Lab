package handlers

import (
	"testing"
	"time"
)

func TestParseFlexTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		hasError bool
	}{
		{
			name:     "HTML5 datetime-local without seconds",
			input:    "2026-10-09T14:30",
			expected: time.Date(2026, 10, 9, 14, 30, 0, 0, time.Local),
			hasError: false,
		},
		{
			name:     "HTML5 datetime-local with seconds",
			input:    "2026-10-09T14:30:45",
			expected: time.Date(2026, 10, 9, 14, 30, 45, 0, time.Local),
			hasError: false,
		},
		{
			name:     "Standard datetime with space",
			input:    "2026-10-09 14:30",
			expected: time.Date(2026, 10, 9, 14, 30, 0, 0, time.Local),
			hasError: false,
		},
		{
			name:     "Standard date only",
			input:    "2026-10-09",
			expected: time.Date(2026, 10, 9, 0, 0, 0, 0, time.Local),
			hasError: false,
		},
		{
			name:     "Dot separated date",
			input:    "2026.10.09",
			expected: time.Date(2026, 10, 9, 0, 0, 0, 0, time.Local),
			hasError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseFlexTime(tc.input)
			if tc.hasError && err == nil {
				t.Fatalf("expected error for %s, got nil", tc.input)
			}
			if !tc.hasError && err != nil {
				t.Fatalf("unexpected error for %s: %v", tc.input, err)
			}
			if got.Year() != tc.expected.Year() ||
				got.Month() != tc.expected.Month() ||
				got.Day() != tc.expected.Day() ||
				got.Hour() != tc.expected.Hour() ||
				got.Minute() != tc.expected.Minute() {
				t.Errorf("input %s: got %v, expected %v", tc.input, got, tc.expected)
			}
		})
	}
}
