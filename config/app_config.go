// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package config

import "go.uber.org/zap/zapcore"

type AppConfig struct {
	App AppInfo   `mapstructure:"appInfo"`   // 应用语言
	Log LogConfig `mapstructure:"logConfig"` // 应用语言
}

type AppInfo struct {
	Language string `mapstructure:"language"` // 应用语言
	AppName  string `mapstructure:"appName"`
}

type LogConfig struct {
	LogLevel zapcore.Level `mapstructure:"logLevel"`
	LogPath  string        `mapstructure:"logPath"`
}
