// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// OperationLogMiddleware 操作日志中间件
type OperationLogMiddleware struct {
}

// NewOperationLogMiddleware 创建操作日志中间件
func NewOperationLogMiddleware() *OperationLogMiddleware {
	return &OperationLogMiddleware{}
}

// OperationLog 操作日志结构
type OperationLog struct {
	UserId     int64     `json:"userId"`
	Username   string    `json:"username"`
	TenantId   int64     `json:"tenantId"`
	Module     string    `json:"module"`
	Operation  string    `json:"operation"`
	Method     string    `json:"method"`
	Url        string    `json:"url"`
	Param      string    `json:"param"`
	Ip         string    `json:"ip"`
	UserAgent  string    `json:"userAgent"`
	Status     int       `json:"status"`
	ErrorMsg   string    `json:"errorMsg"`
	Duration   int64     `json:"duration"`
	OperTime   time.Time `json:"operTime"`
}

// Handle 处理请求
func (m *OperationLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 读取请求参数
		var param string
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil {
				param = string(bodyBytes)
				// 重新设置 Body，以便后续处理
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// 获取用户信息
		userId := GetUserIdFromCtx(r.Context())
		tenantId := GetTenantIdFromCtx(r.Context())

		// 包装 ResponseWriter 以捕获状态码
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// 调用下一个处理器
		next(wrapped, r)

		// 计算耗时
		duration := time.Since(start).Milliseconds()

		// 构建操作日志
		operLog := &OperationLog{
			UserId:    parseUserId(userId),
			TenantId:  tenantId,
			Method:    r.Method,
			Url:       r.URL.Path,
			Param:     truncate(param, 2000),
			Ip:        getClientIp(r),
			UserAgent: r.UserAgent(),
			Status:    wrapped.statusCode,
			Duration:  duration,
			OperTime:  start,
		}

		// 异步记录日志（可以改为写入数据库或消息队列）
		go m.saveLog(operLog)
	}
}

// saveLog 保存日志
func (m *OperationLogMiddleware) saveLog(log *OperationLog) {
	// TODO: 实现日志保存逻辑，可以写入数据库或发送到消息队列
	logx.Infof("OperationLog: %+v", log)
}

// responseWriter 包装 ResponseWriter 以捕获状态码
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// getClientIp 获取客户端 IP
func getClientIp(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

// truncate 截断字符串
func truncate(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen] + "..."
	}
	return s
}

// parseUserId 解析用户 ID
func parseUserId(userId string) int64 {
	var id int64
	json.Unmarshal([]byte(userId), &id)
	return id
}
