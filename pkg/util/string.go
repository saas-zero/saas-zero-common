// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

// MD5 计算 MD5 值
func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateSalt 生成随机盐
func GenerateSalt(length int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		result[i] = chars[num.Int64()]
	}
	return string(result), nil
}

// EncryptPassword 加密密码
func EncryptPassword(password, salt string) string {
	return MD5(password + salt)
}

// VerifyPassword 验证密码
func VerifyPassword(password, salt, encryptedPassword string) bool {
	return EncryptPassword(password, salt) == encryptedPassword
}

// IsBlank 判断字符串是否为空
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotBlank 判断字符串是否不为空
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// DefaultIfEmpty 如果为空返回默认值
func DefaultIfEmpty(s, defaultValue string) string {
	if IsBlank(s) {
		return defaultValue
	}
	return s
}

// Join 拼接字符串
func Join(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}

// Contains 判断字符串是否在切片中
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Remove 从切片中移除元素
func Remove(slice []string, item string) []string {
	result := make([]string, 0, len(slice))
	for _, s := range slice {
		if s != item {
			result = append(result, s)
		}
	}
	return result
}

// Unique 去重
func Unique(slice []string) []string {
	result := make([]string, 0, len(slice))
	seen := make(map[string]bool)
	for _, s := range slice {
		if !seen[s] {
			seen[s] = true
			result = append(result, s)
		}
	}
	return result
}

// Substring 截取子字符串
func Substring(s string, start, length int) string {
	runes := []rune(s)
	if start < 0 {
		start = 0
	}
	if start >= len(runes) {
		return ""
	}
	end := start + length
	if end > len(runes) {
		end = len(runes)
	}
	return string(runes[start:end])
}

// FormatString 格式化字符串
func FormatString(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
