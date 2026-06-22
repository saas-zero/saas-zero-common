// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package util

import (
	"fmt"
	"sync"
	"time"
)

// Snowflake 雪花 ID 生成器
type Snowflake struct {
	mu            sync.Mutex
	epoch         int64
	machineId     int64
	sequence      int64
	lastTimestamp int64
}

var defaultSnowflake *Snowflake
var once sync.Once

// InitSnowflake 初始化默认雪花生成器
func InitSnowflake(machineId int64) {
	once.Do(func() {
		defaultSnowflake = NewSnowflake(machineId)
	})
}

// NewSnowflake 创建雪花生成器
func NewSnowflake(machineId int64) *Snowflake {
	if machineId < 0 || machineId > 1023 {
		panic("machineId must be between 0 and 1023")
	}
	return &Snowflake{
		epoch:     1609459200000, // 2021-01-01 00:00:00
		machineId: machineId,
	}
}

// NextId 生成下一个 ID
func (s *Snowflake) NextId() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixMilli()

	if timestamp < s.lastTimestamp {
		return 0, fmt.Errorf("clock moved backwards, refusing to generate id")
	}

	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & 4095
		if s.sequence == 0 {
			timestamp = s.waitNextMillis()
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	id := ((timestamp - s.epoch) << 22) | (s.machineId << 12) | s.sequence
	return id, nil
}

func (s *Snowflake) waitNextMillis() int64 {
	timestamp := time.Now().UnixMilli()
	for timestamp <= s.lastTimestamp {
		timestamp = time.Now().UnixMilli()
	}
	return timestamp
}

// NextIdStr 生成下一个 ID 字符串
func NextIdStr() string {
	if defaultSnowflake == nil {
		InitSnowflake(1)
	}
	id, _ := defaultSnowflake.NextId()
	return fmt.Sprintf("%d", id)
}

// NextId 生成下一个 ID
func NextId() int64 {
	if defaultSnowflake == nil {
		InitSnowflake(1)
	}
	id, _ := defaultSnowflake.NextId()
	return id
}
