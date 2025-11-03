// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package config

import (
	"github.com/saas-zero/saas-zero-common/config"
	"github.com/spf13/viper"
)

func LoadConfig() (config.AppConfig, error) {
	var config config.AppConfig

	viper.SetConfigName("config")
	//viper.SetConfigName("demo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	//viper.AutomaticEnv() //将环境变量与配置绑定
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
