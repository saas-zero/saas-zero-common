// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package response

import (
	"net/http"
)

// Response 定义统一的API响应结构
type Response struct {
	Code    int         `json:"code"`           // 状态码
	Message string      `json:"message"`        // 响应消息
	Data    interface{} `json:"data,omitempty"` // 响应数据
	Success bool        `json:"success"`        // 是否成功
}

// Success 返回成功的响应
func Success(data interface{}) *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
		Success: true,
	}
}

// Error 返回错误的响应
func Error(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Success: false,
	}
}

// PageResponse 分页响应结构
type PageResponse struct {
	List      interface{} `json:"list"`      // 数据列表
	Page      int         `json:"page"`      // 当前页码
	PageSize  int         `json:"pageSize"`  // 每页数量
	Total     int64       `json:"total"`     // 总数
	TotalPage int         `json:"totalPage"` // 总页数
}

// Page 返回分页响应
func Page(list interface{}, page, pageSize int, total int64) *Response {
	totalPage := int(total+int64(pageSize)-1) / pageSize // 向上取整计算总页数

	pageResp := &PageResponse{
		List:      list,
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: totalPage,
	}

	return &Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    pageResp,
		Success: true,
	}
}

// ToRPCResponse 转换为RPC响应格式
func (r *Response) ToRPCResponse() *RPCResponse {
	return &RPCResponse{
		Code:    int32(r.Code),
		Message: r.Message,
		Success: r.Success,
	}
}

// RPCResponse 定义RPC通用响应结构
type RPCResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Success bool   `protobuf:"varint,3,opt,name=success,proto3" json:"success"`
}

// NewRPCResponse 创建RPC响应
func NewRPCResponse(code int32, message string, success bool) *RPCResponse {
	return &RPCResponse{
		Code:    code,
		Message: message,
		Success: success,
	}
}
