// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package middleware

import (
	"net/http"

	"github.com/saas-zero/saas-zero-common/pkg/casbin"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// CasbinMiddleware Casbin 鉴权中间件
type CasbinMiddleware struct {
	Casbin *casbin.CasbinService
}

// NewCasbinMiddleware 创建 Casbin 中间件
func NewCasbinMiddleware(casbinService *casbin.CasbinService) *CasbinMiddleware {
	return &CasbinMiddleware{Casbin: casbinService}
}

// Handle 处理请求
func (m *CasbinMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取用户信息
		userId := GetUserIdFromCtx(r.Context())
		tenantId := GetTenantIdFromCtx(r.Context())

		if userId == "" || tenantId == 0 {
			httpx.Error(w, NewUnauthorizedError("未认证"))
			return
		}

		// 获取请求路径和方法
		obj := r.URL.Path
		act := r.Method

		// 检查权限
		hasPermission, err := m.Casbin.CheckPermission(userId, obj, act, tenantId)
		if err != nil {
			httpx.Error(w, NewForbiddenError("权限校验失败"))
			return
		}

		if !hasPermission {
			httpx.Error(w, NewForbiddenError("没有访问权限"))
			return
		}

		next(w, r)
	}
}
