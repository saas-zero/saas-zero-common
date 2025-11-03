// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package captcha

import (
	"github.com/mojocn/base64Captcha"
)

// 设置自带的 store 存在服务器内存中
var store = base64Captcha.DefaultMemStore

func GenerateCaptcha() (string, string, string, error) {
	//设置验证码的配置
	driverString := base64Captcha.NewDriverString(80, 240, 2, base64Captcha.OptionShowSlimeLine,
		4, base64Captcha.TxtNumbers+base64Captcha.TxtAlphabet, nil, nil, nil)

	//生成base64图片
	c := base64Captcha.NewCaptcha(driverString, store)
	//验证码id base64图片字符串 验证码字符串 error
	return c.Generate()
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	//验证验证码，参数1是验证码的id，参数2是用户输入的验证码
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
