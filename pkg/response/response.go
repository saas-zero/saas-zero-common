// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

func (r *Response) Error() string {
	return r.Message
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
		Success: true,
	}
}

func Error(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Success: false,
	}
}

type PageResponse struct {
	List      interface{} `json:"list"`
	Page      int         `json:"page"`
	PageSize  int         `json:"pageSize"`
	Total     int64       `json:"total"`
	TotalPage int         `json:"totalPage"`
}

func Page(list interface{}, page, pageSize int, total int64) *Response {
	totalPage := int(total+int64(pageSize)-1) / pageSize

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

func JSON(w http.ResponseWriter, code int, data interface{}) {
	httpx.OkJson(w, data)
}

func SuccessJson(w http.ResponseWriter, data interface{}) {
	httpx.OkJson(w, Success(data))
}

func ErrorJson(w http.ResponseWriter, code int, message string) {
	httpx.Error(w, Error(code, message))
}

func (r *Response) ToRPCResponse() *RPCResponse {
	return &RPCResponse{
		Code:    int32(r.Code),
		Message: r.Message,
		Success: r.Success,
	}
}

func NewRPCResponse(code int32, message string, success bool) *RPCResponse {
	return &RPCResponse{
		Code:    code,
		Message: message,
		Success: success,
	}
}

type RPCResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Success bool   `protobuf:"varint,3,opt,name=success,proto3" json:"success"`
}