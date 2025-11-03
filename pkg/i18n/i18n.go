// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package ryi18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// 获取Localizer实例
func LoadLocalizer(lang string) *i18n.Localizer {
	var bundle = i18n.NewBundle(language.SimplifiedChinese)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// 使用LoadMessageFile加载YAML文件，并检查错误
	bundle.LoadMessageFile("./locales/en-US.yaml")
	bundle.LoadMessageFile("./locales/zh-CN.yaml")

	if bundle != nil {
		return i18n.NewLocalizer(bundle, lang)
	} else {
		return nil
	}
}
