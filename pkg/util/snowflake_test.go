package util

import (
	"testing"
)

func TestSnowflake_NextId(t *testing.T) {
	sf := NewSnowflake(1)

	id1, err := sf.NextId()
	if err != nil {
		t.Fatalf("NextId() error = %v", err)
	}

	id2, err := sf.NextId()
	if err != nil {
		t.Fatalf("NextId() error = %v", err)
	}

	if id1 == id2 {
		t.Error("NextId() should generate unique IDs")
	}

	if id1 <= 0 || id2 <= 0 {
		t.Error("NextId() should generate positive IDs")
	}
}

func TestSnowflake_MachineIdRange(t *testing.T) {
	tests := []struct {
		machineId int64
		wantPanic bool
	}{
		{0, false},
		{512, false},
		{1023, false},
		{-1, true},
		{1024, true},
	}

	for _, tt := range tests {
		func() {
			defer func() {
				if r := recover(); (r != nil) != tt.wantPanic {
					t.Errorf("NewSnowflake(%d) panic = %v, wantPanic %v", tt.machineId, r, tt.wantPanic)
				}
			}()
			NewSnowflake(tt.machineId)
		}()
	}
}

func TestNextId(t *testing.T) {
	InitSnowflake(1)

	id1 := NextId()
	id2 := NextId()

	if id1 == id2 {
		t.Error("NextId() should generate unique IDs")
	}

	if id1 <= 0 {
		t.Error("NextId() should generate positive IDs")
	}
}

func TestNextIdStr(t *testing.T) {
	InitSnowflake(1)

	id1 := NextIdStr()
	id2 := NextIdStr()

	if id1 == id2 {
		t.Error("NextIdStr() should generate unique IDs")
	}

	if id1 == "" {
		t.Error("NextIdStr() should not return empty string")
	}
}
