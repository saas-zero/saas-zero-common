package util

import (
	"testing"
	"time"
)

func TestPtrInt64(t *testing.T) {
	v := int64(123)
	p := PtrInt64(v)
	if *p != v {
		t.Errorf("PtrInt64(%d) = %d, want %d", v, *p, v)
	}
}

func TestPtrString(t *testing.T) {
	v := "hello"
	p := PtrString(v)
	if *p != v {
		t.Errorf("PtrString(%q) = %q, want %q", v, *p, v)
	}
}

func TestPtrBool(t *testing.T) {
	v := true
	p := PtrBool(v)
	if *p != v {
		t.Errorf("PtrBool(%v) = %v, want %v", v, *p, v)
	}
}

func TestDerefInt64(t *testing.T) {
	v := int64(123)
	p := &v

	tests := []struct {
		input    *int64
		def      int64
		expected int64
	}{
		{p, 0, 123},
		{nil, 0, 0},
		{nil, 456, 456},
	}

	for _, tt := range tests {
		result := DerefInt64(tt.input, tt.def)
		if result != tt.expected {
			t.Errorf("DerefInt64(%v, %d) = %d, want %d", tt.input, tt.def, result, tt.expected)
		}
	}
}

func TestDerefString(t *testing.T) {
	v := "hello"
	p := &v

	tests := []struct {
		input    *string
		def      string
		expected string
	}{
		{p, "", "hello"},
		{nil, "", ""},
		{nil, "default", "default"},
	}

	for _, tt := range tests {
		result := DerefString(tt.input, tt.def)
		if result != tt.expected {
			t.Errorf("DerefString(%v, %q) = %q, want %q", tt.input, tt.def, result, tt.expected)
		}
	}
}

func TestDerefTime(t *testing.T) {
	now := time.Now()
	p := &now

	result := DerefTime(p, time.Time{})
	if result != now {
		t.Errorf("DerefTime should return the pointed value")
	}

	def := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	result = DerefTime(nil, def)
	if result != def {
		t.Errorf("DerefTime(nil, def) should return default")
	}
}

func TestNowTimestamp(t *testing.T) {
	ts := NowTimestamp()
	if ts <= 0 {
		t.Error("NowTimestamp() should return positive value")
	}

	time.Sleep(10 * time.Millisecond)
	ts2 := NowTimestamp()
	if ts2 <= ts {
		t.Error("NowTimestamp() should increase over time")
	}
}

func TestFormatTimestamp(t *testing.T) {
	ts := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC).UnixMilli()

	tests := []struct {
		layout   string
		expected string
	}{
		{"2006-01-02", "2024-01-15"},
	}

	for _, tt := range tests {
		result := FormatTimestamp(ts, tt.layout)
		if result != tt.expected {
			t.Errorf("FormatTimestamp(%d, %q) = %q, want %q", ts, tt.layout, result, tt.expected)
		}
	}
}

func TestFormatTime(t *testing.T) {
	tm := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	tests := []struct {
		layout   string
		expected string
	}{
		{"", "2024-01-15 10:30:00"},
		{"2006-01-02", "2024-01-15"},
	}

	for _, tt := range tests {
		result := FormatTime(tm, tt.layout)
		if result != tt.expected {
			t.Errorf("FormatTime(%v, %q) = %q, want %q", tm, tt.layout, result, tt.expected)
		}
	}
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		s        string
		layout   string
		wantErr  bool
	}{
		{"2024-01-15 10:30:00", "", false},
		{"2024-01-15", "2006-01-02", false},
		{"invalid", "", true},
	}

	for _, tt := range tests {
		_, err := ParseTime(tt.s, tt.layout)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseTime(%q, %q) error = %v, wantErr %v", tt.s, tt.layout, err, tt.wantErr)
		}
	}
}
