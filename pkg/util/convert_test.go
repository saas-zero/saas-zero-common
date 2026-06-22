package util

import (
	"testing"
)

func TestToInt64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int64
	}{
		{int(1), 1},
		{int8(2), 2},
		{int16(3), 3},
		{int32(4), 4},
		{int64(5), 5},
		{uint(6), 6},
		{float32(7.5), 7},
		{float64(8.5), 8},
		{"9", 9},
		{"abc", 0},
		{nil, 0},
		{true, 0},
	}

	for _, tt := range tests {
		result := ToInt64(tt.input)
		if result != tt.expected {
			t.Errorf("ToInt64(%v) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int
	}{
		{int64(1), 1},
		{"2", 2},
		{float64(3.5), 3},
	}

	for _, tt := range tests {
		result := ToInt(tt.input)
		if result != tt.expected {
			t.Errorf("ToInt(%v) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{"hello", "hello"},
		{123, "123"},
		{int64(456), "456"},
		{true, "true"},
		{nil, ""},
	}

	for _, tt := range tests {
		result := ToString(tt.input)
		if result != tt.expected {
			t.Errorf("ToString(%v) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToFloat64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float64
	}{
		{float64(1.5), 1.5},
		{int(2), 2.0},
		{"3.14", 3.14},
		{"abc", 0},
	}

	for _, tt := range tests {
		result := ToFloat64(tt.input)
		if result != tt.expected {
			t.Errorf("ToFloat64(%v) = %f, want %f", tt.input, result, tt.expected)
		}
	}
}

func TestToBool(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{true, true},
		{false, false},
		{1, true},
		{0, false},
		{"true", true},
		{"false", false},
		{"1", true},
		{"0", false},
	}

	for _, tt := range tests {
		result := ToBool(tt.input)
		if result != tt.expected {
			t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestPageOffset(t *testing.T) {
	tests := []struct {
		page     int
		pageSize int
		expected int
	}{
		{1, 10, 0},
		{2, 10, 10},
		{3, 20, 40},
		{0, 10, 0},
	}

	for _, tt := range tests {
		result := PageOffset(tt.page, tt.pageSize)
		if result != tt.expected {
			t.Errorf("PageOffset(%d, %d) = %d, want %d", tt.page, tt.pageSize, result, tt.expected)
		}
	}
}

func TestPageTotal(t *testing.T) {
	tests := []struct {
		total    int64
		pageSize int
		expected int
	}{
		{100, 10, 10},
		{101, 10, 11},
		{0, 10, 0},
		{10, 0, 0},
	}

	for _, tt := range tests {
		result := PageTotal(tt.total, tt.pageSize)
		if result != tt.expected {
			t.Errorf("PageTotal(%d, %d) = %d, want %d", tt.total, tt.pageSize, result, tt.expected)
		}
	}
}
