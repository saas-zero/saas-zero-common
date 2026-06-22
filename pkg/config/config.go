// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package config

import (
	"github.com/zeromicro/go-zero/core/conf"
)

func LoadConfig(path string, v interface{}) error {
	return conf.LoadConfig(path, v)
}

func MustLoadConfig(path string, v interface{}) {
	conf.MustLoad(path, v)
}
