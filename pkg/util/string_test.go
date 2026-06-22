package util

import (
	"testing"
)

func TestMD5(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"world", "7d793037a0760186574b0282f2f435e7"},
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
	}

	for _, tt := range tests {
		result := MD5(tt.input)
		if result != tt.expected {
			t.Errorf("MD5(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt(16)
	if err != nil {
		t.Fatalf("GenerateSalt() error = %v", err)
	}
	if len(salt) != 16 {
		t.Errorf("GenerateSalt() length = %d, want 16", len(salt))
	}

	salt2, _ := GenerateSalt(16)
	if salt == salt2 {
		t.Error("GenerateSalt() should generate unique salts")
	}
}

func TestEncryptPassword(t *testing.T) {
	password := "test123"
	salt := "randomsalt"

	encrypted := EncryptPassword(password, salt)
	if encrypted == "" {
		t.Error("EncryptPassword() should not return empty")
	}

	encrypted2 := EncryptPassword(password, salt)
	if encrypted != encrypted2 {
		t.Error("EncryptPassword() should return same result for same input")
	}

	encrypted3 := EncryptPassword(password, "differentsalt")
	if encrypted == encrypted3 {
		t.Error("EncryptPassword() should return different result for different salt")
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "test123"
	salt := "randomsalt"
	encrypted := EncryptPassword(password, salt)

	if !VerifyPassword(password, salt, encrypted) {
		t.Error("VerifyPassword() should return true for correct password")
	}

	if VerifyPassword("wrong", salt, encrypted) {
		t.Error("VerifyPassword() should return false for wrong password")
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{" ", true},
		{"  ", true},
		{"\t", true},
		{"\n", true},
		{"hello", false},
		{" hello ", false},
	}

	for _, tt := range tests {
		result := IsBlank(tt.input)
		if result != tt.expected {
			t.Errorf("IsBlank(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsNotBlank(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", false},
		{" ", false},
		{"hello", true},
	}

	for _, tt := range tests {
		result := IsNotBlank(tt.input)
		if result != tt.expected {
			t.Errorf("IsNotBlank(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestDefaultIfEmpty(t *testing.T) {
	tests := []struct {
		input        string
		defaultValue string
		expected     string
	}{
		{"", "default", "default"},
		{" ", "default", "default"},
		{"hello", "default", "hello"},
	}

	for _, tt := range tests {
		result := DefaultIfEmpty(tt.input, tt.defaultValue)
		if result != tt.expected {
			t.Errorf("DefaultIfEmpty(%q, %q) = %q, want %q", tt.input, tt.defaultValue, result, tt.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		slice    []string
		item     string
		expected bool
	}{
		{[]string{"a", "b", "c"}, "a", true},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{}, "a", false},
		{nil, "a", false},
	}

	for _, tt := range tests {
		result := Contains(tt.slice, tt.item)
		if result != tt.expected {
			t.Errorf("Contains(%v, %q) = %v, want %v", tt.slice, tt.item, result, tt.expected)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		slice    []string
		item     string
		expected []string
	}{
		{[]string{"a", "b", "c"}, "b", []string{"a", "c"}},
		{[]string{"a", "b", "c"}, "d", []string{"a", "b", "c"}},
		{[]string{}, "a", []string{}},
	}

	for _, tt := range tests {
		result := Remove(tt.slice, tt.item)
		if len(result) != len(tt.expected) {
			t.Errorf("Remove(%v, %q) length = %d, want %d", tt.slice, tt.item, len(result), len(tt.expected))
		}
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"a", "b", "a", "c", "b"}, 3},
		{[]string{"a", "b", "c"}, 3},
		{[]string{}, 0},
	}

	for _, tt := range tests {
		result := Unique(tt.input)
		if len(result) != tt.expected {
			t.Errorf("Unique(%v) length = %d, want %d", tt.input, len(result), tt.expected)
		}
	}
}

func TestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		start    int
		length   int
		expected string
	}{
		{"hello", 0, 3, "hel"},
		{"hello", 1, 3, "ell"},
		{"hello", 0, 10, "hello"},
		{"hello", 10, 3, ""},
		{"hello", -1, 3, "hel"},
		{"你好世界", 0, 2, "你好"},
	}

	for _, tt := range tests {
		result := Substring(tt.input, tt.start, tt.length)
		if result != tt.expected {
			t.Errorf("Substring(%q, %d, %d) = %q, want %q", tt.input, tt.start, tt.length, result, tt.expected)
		}
	}
}
