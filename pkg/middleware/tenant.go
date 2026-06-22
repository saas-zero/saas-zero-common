// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// TenantMiddleware 租户中间件
type TenantMiddleware struct {
}

// NewTenantMiddleware 创建租户中间件
func NewTenantMiddleware() *TenantMiddleware {
	return &TenantMiddleware{}
}

// Handle 处理请求
func (m *TenantMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头获取租户 ID
		tenantIdStr := r.Header.Get("X-Tenant-Id")
		if tenantIdStr == "" {
			// 从查询参数获取
			tenantIdStr = r.URL.Query().Get("tenantId")
		}

		if tenantIdStr == "" {
			httpx.Error(w, NewBadRequestError("缺少租户标识"))
			return
		}

		tenantId, err := strconv.ParseInt(tenantIdStr, 10, 64)
		if err != nil || tenantId <= 0 {
			httpx.Error(w, NewBadRequestError("无效的租户标识"))
			return
		}

		// 将租户 ID 存入上下文
		ctx := context.WithValue(r.Context(), TenantIdKey, tenantId)
		next(w, r.WithContext(ctx))
	}
}
