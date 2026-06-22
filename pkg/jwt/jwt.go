// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package ryjwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenConf struct {
	AccessSecret  string
	AccessExpire  int64
	RefreshExpire int64
}

func Sign(conf *TokenConf, k, v string, exp int64) (string, error) {
	if exp == 0 {
		exp = conf.AccessExpire
	}

	claims := jwt.MapClaims{
		k:     v,
		"exp": time.Now().Add(time.Duration(exp) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.AccessSecret))
}

func Valid(conf *TokenConf, k, tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(conf.AccessSecret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		val, exists := claims[k]
		if !exists {
			return "", errors.New("key " + k + " not found in token")
		}

		strVal, ok := val.(string)
		if !ok {
			return "", errors.New("key " + k + " is not a string")
		}

		return strVal, nil
	}

	return "", errors.New("token invalid")
}

func RefreshToken(conf *TokenConf, userId string) (string, error) {
	return Sign(conf, "userId", userId, conf.AccessExpire)
}

func GenerateTokens(conf *TokenConf, userId string) (accessToken, refreshToken string, err error) {
	accessToken, err = Sign(conf, "userId", userId, conf.AccessExpire)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = Sign(conf, "userId", userId, conf.RefreshExpire)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
