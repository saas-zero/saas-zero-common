package middleware

import "context"

type contextKey string

const (
	UserIdKey   contextKey = "userId"
	TenantIdKey contextKey = "tenantId"
	UserNameKey contextKey = "username"
	RoleIdKey   contextKey = "roleId"
	DeptIdKey   contextKey = "deptId"
)

// GetUserIdFromCtx 从上下文中获取用户 ID
func GetUserIdFromCtx(ctx context.Context) string {
	if val := ctx.Value(UserIdKey); val != nil {
		if userId, ok := val.(string); ok {
			return userId
		}
	}
	return ""
}

// GetTenantIdFromCtx 从上下文中获取租户 ID
func GetTenantIdFromCtx(ctx context.Context) int64 {
	if val := ctx.Value(TenantIdKey); val != nil {
		if tenantId, ok := val.(int64); ok {
			return tenantId
		}
	}
	return 0
}
