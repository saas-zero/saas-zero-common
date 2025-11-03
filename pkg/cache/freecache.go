// Copyright (c) [2025] Kong All rights reserved.
// Use of this source code is governed by a Apache 2.0 license that can be found in the LICENSE file.
// Author: Kong See：https://github.com/saas-zero/saas-zero or https://gitee.com/saas-zero/saas-zero
// Email: hot_kun@hotmail.com

package cache

import (
	"github.com/coocood/freecache"
)

type FreeCacheClient struct {
	cache *freecache.Cache
}

func NewFreeCacheClient(size int) *FreeCacheClient {
	return &FreeCacheClient{cache: freecache.NewCache(size)}
}

func (c *FreeCacheClient) Set(key, value []byte, expireSeconds int) (err error) {
	return c.cache.Set(key, value, expireSeconds)
}

func (c *FreeCacheClient) Get(key []byte) (value []byte, err error) {
	return c.cache.Get(key)
}

func (c *FreeCacheClient) Del(key string) bool {
	return c.cache.Del([]byte(key))
}

func (c *FreeCacheClient) Clear() {
	c.cache.Clear()
}
