// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package casbin

import (
	"fmt"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

// CasbinService Casbin 服务
type CasbinService struct {
	enforcer *casbin.Enforcer
	adapter  persist.Adapter
	mu       sync.RWMutex
}

// NewCasbinService 创建 Casbin 服务
func NewCasbinService(adapter persist.Adapter) (*CasbinService, error) {
	// 定义 RBAC with domains 模型
	m := `
[request_definition]
r = sub, obj, act, dom

[policy_definition]
p = sub, obj, act, dom

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.obj == p.obj && r.act == p.act || r.sub == "admin"
`

	model, err := model.NewModelFromString(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create model: %v", err)
	}

	enforcer, err := casbin.NewEnforcer(model, adapter)
	if err != nil {
		return nil, fmt.Errorf("failed to create enforcer: %v", err)
	}

	return &CasbinService{
		enforcer: enforcer,
		adapter:  adapter,
	}, nil
}

// LoadPolicy 加载策略
func (s *CasbinService) LoadPolicy() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.enforcer.LoadPolicy()
}

// SavePolicy 保存策略
func (s *CasbinService) SavePolicy() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.enforcer.SavePolicy()
}

// CheckPermission 检查权限
func (s *CasbinService) CheckPermission(userId, obj, act string, tenantId int64) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	domain := fmt.Sprintf("%d", tenantId)
	return s.enforcer.Enforce(userId, obj, act, domain)
}

// AddPolicy 添加策略
func (s *CasbinService) AddPolicy(sub, obj, act string, tenantId int64) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	domain := fmt.Sprintf("%d", tenantId)
	return s.enforcer.AddPolicy(sub, obj, act, domain)
}

// RemovePolicy 删除策略
func (s *CasbinService) RemovePolicy(sub, obj, act string, tenantId int64) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	domain := fmt.Sprintf("%d", tenantId)
	return s.enforcer.RemovePolicy(sub, obj, act, domain)
}

// AddRoleForUser 为用户添加角色
func (s *CasbinService) AddRoleForUser(userId, role string, tenantId int64) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	domain := fmt.Sprintf("%d", tenantId)
	return s.enforcer.AddRoleForUser(userId, role, domain)
}

// DeleteRoleForUser 删除用户角色
func (s *CasbinService) DeleteRoleForUser(userId, role string, tenantId int64) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	domain := fmt.Sprintf("%d", tenantId)
	return s.enforcer.DeleteRoleForUser(userId, role, domain)
}

// GetRolesForUser 获取用户角色
func (s *CasbinService) GetRolesForUser(userId string, tenantId int64) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	domain := fmt.Sprintf("%d", tenantId)
	return s.enforcer.GetRolesForUser(userId, domain)
}

// GetUsersForRole 获取角色用户
func (s *CasbinService) GetUsersForRole(role string, tenantId int64) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	domain := fmt.Sprintf("%d", tenantId)
	return s.enforcer.GetUsersForRole(role, domain)
}

// HasRoleForUser 检查用户是否有角色
func (s *CasbinService) HasRoleForUser(userId, role string, tenantId int64) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	domain := fmt.Sprintf("%d", tenantId)
	return s.enforcer.HasRoleForUser(userId, role, domain)
}

// AddPolicies 批量添加策略
func (s *CasbinService) AddPolicies(rules [][]string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.enforcer.AddPolicies(rules)
	if err != nil {
		return false, err
	}
	return true, nil
}

// RemovePolicies 批量删除策略
func (s *CasbinService) RemovePolicies(rules [][]string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.enforcer.RemovePolicies(rules)
	if err != nil {
		return false, err
	}
	return true, nil
}

// RemoveFilteredPolicy 按过滤条件删除策略
func (s *CasbinService) RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.enforcer.RemoveFilteredPolicy(fieldIndex, fieldValues...)
}
