// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package util

import (
	"fmt"
	"time"
)

// PtrInt64 返回 int64 指针
func PtrInt64(v int64) *int64 {
	return &v
}

// PtrString 返回 string 指针
func PtrString(v string) *string {
	return &v
}

// PtrTime 返回 time.Time 指针
func PtrTime(v time.Time) *time.Time {
	return &v
}

// PtrBool 返回 bool 指针
func PtrBool(v bool) *bool {
	return &v
}

// PtrInt32 返回 int32 指针
func PtrInt32(v int32) *int32 {
	return &v
}

// PtrUint64 返回 uint64 指针
func PtrUint64(v uint64) *uint64 {
	return &v
}

// PtrFloat64 返回 float64 指针
func PtrFloat64(v float64) *float64 {
	return &v
}

// DerefInt64 安全解引用 int64 指针
func DerefInt64(p *int64, defaultVal int64) int64 {
	if p == nil {
		return defaultVal
	}
	return *p
}

// DerefString 安全解引用 string 指针
func DerefString(p *string, defaultVal string) string {
	if p == nil {
		return defaultVal
	}
	return *p
}

// DerefTime 安全解引用 time.Time 指针
func DerefTime(p *time.Time, defaultVal time.Time) time.Time {
	if p == nil {
		return defaultVal
	}
	return *p
}

// DerefBool 安全解引用 bool 指针
func DerefBool(p *bool, defaultVal bool) bool {
	if p == nil {
		return defaultVal
	}
	return *p
}

// DerefInt32 安全解引用 int32 指针
func DerefInt32(p *int32, defaultVal int32) int32 {
	if p == nil {
		return defaultVal
	}
	return *p
}

// DerefUint64 安全解引用 uint64 指针
func DerefUint64(p *uint64, defaultVal uint64) uint64 {
	if p == nil {
		return defaultVal
	}
	return *p
}

// DerefFloat64 安全解引用 float64 指针
func DerefFloat64(p *float64, defaultVal float64) float64 {
	if p == nil {
		return defaultVal
	}
	return *p
}

// PageOffset 计算分页偏移量
func PageOffset(page, pageSize int) int {
	if page <= 0 {
		page = 1
	}
	return (page - 1) * pageSize
}

// PageTotal 计算总页数
func PageTotal(total int64, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	return int((total + int64(pageSize) - 1) / int64(pageSize))
}

// FormatTime 格式化时间
func FormatTime(t time.Time, layout string) string {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return t.Format(layout)
}

// ParseTime 解析时间
func ParseTime(s string, layout string) (time.Time, error) {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return time.Parse(layout, s)
}

// NowTimestamp 获取当前时间戳（毫秒）
func NowTimestamp() int64 {
	return time.Now().UnixMilli()
}

// FormatTimestamp 格式化时间戳
func FormatTimestamp(timestamp int64, layout string) string {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	t := time.UnixMilli(timestamp)
	return t.Format(layout)
}

// Println 打印输出
func Println(args ...interface{}) {
	fmt.Println(args...)
}
