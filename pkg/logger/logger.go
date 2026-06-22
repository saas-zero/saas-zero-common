// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package logger

import (
	"github.com/zeromicro/go-zero/core/logx"
)

func InitLogger(name string) {
	logx.MustSetup(logx.LogConf{
		ServiceName: name,
		Mode:        "console",
		Path:        "logs",
		Level:       "info",
		MaxBackups:  3,
		MaxSize:     100,
	})
}

func Debug(v ...interface{}) {
	logx.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	logx.Debugf(format, v...)
}

func Info(v ...interface{}) {
	logx.Info(v...)
}

func Infof(format string, v ...interface{}) {
	logx.Infof(format, v...)
}

func Error(v ...interface{}) {
	logx.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	logx.Errorf(format, v...)
}

