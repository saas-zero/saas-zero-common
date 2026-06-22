// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package util

import (
	"fmt"
	"strconv"
)

// ToInt64 安全转换为 int64
func ToInt64(v interface{}) int64 {
	switch val := v.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case string:
		i, _ := strconv.ParseInt(val, 10, 64)
		return i
	default:
		return 0
	}
}

// ToInt 安全转换为 int
func ToInt(v interface{}) int {
	return int(ToInt64(v))
}

// ToString 安全转换为 string
func ToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	case float32, float64:
		return fmt.Sprintf("%f", val)
	case bool:
		return strconv.FormatBool(val)
	default:
		return ""
	}
}

// ToFloat64 安全转换为 float64
func ToFloat64(v interface{}) float64 {
	switch val := v.(type) {
	case float32:
		return float64(val)
	case float64:
		return val
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case string:
		f, _ := strconv.ParseFloat(val, 64)
		return f
	default:
		return 0
	}
}

// ToBool 安全转换为 bool
func ToBool(v interface{}) bool {
	switch val := v.(type) {
	case bool:
		return val
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return ToInt64(val) != 0
	case float32, float64:
		return ToFloat64(val) != 0
	case string:
		b, _ := strconv.ParseBool(val)
		return b
	default:
		return false
	}
}

// ToInt64Slice 转换为 int64 切片
func ToInt64Slice(v interface{}) []int64 {
	switch val := v.(type) {
	case []int64:
		return val
	case []int:
		result := make([]int64, len(val))
		for i, item := range val {
			result[i] = int64(item)
		}
		return result
	case []interface{}:
		result := make([]int64, 0, len(val))
		for _, item := range val {
			result = append(result, ToInt64(item))
		}
		return result
	default:
		return nil
	}
}

// ToStringSlice 转换为 string 切片
func ToStringSlice(v interface{}) []string {
	switch val := v.(type) {
	case []string:
		return val
	case []interface{}:
		result := make([]string, 0, len(val))
		for _, item := range val {
			result = append(result, ToString(item))
		}
		return result
	default:
		return nil
	}
}
