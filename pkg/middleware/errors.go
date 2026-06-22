// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type UnauthorizedError struct {
	Message string
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{Message: message}
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

type ForbiddenError struct {
	Message string
}

func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{Message: message}
}

func (e *ForbiddenError) Error() string {
	return e.Message
}

type BadRequestError struct {
	Message string
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{Message: message}
}

func (e *BadRequestError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{Message: message}
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type InternalError struct {
	Message string
}

func NewInternalError(message string) *InternalError {
	return &InternalError{Message: message}
}

func (e *InternalError) Error() string {
	return e.Message
}

func HandleError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *UnauthorizedError:
		httpx.Error(w, err)
	case *ForbiddenError:
		httpx.Error(w, err)
	case *BadRequestError:
		httpx.Error(w, err)
	case *NotFoundError:
		httpx.Error(w, err)
	default:
		httpx.Error(w, err)
	}
}