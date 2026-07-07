package redis

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"
	gzredis "github.com/zeromicro/go-zero/core/stores/redis"
)

type Conf struct {
	Host string
	Pass string
	Type string `json:",default=node,options=node|cluster"`
	DB   int    `json:",default=0"`
}

type Client struct {
	gz *gzredis.Redis
	gr *goredis.Client
}

func NewClient(c Conf) (*Client, error) {
	if c.DB == 0 {
		rds, err := gzredis.NewRedis(gzredis.RedisConf{
			Host: c.Host,
			Pass: c.Pass,
			Type: c.Type,
		})
		if err != nil {
			return nil, err
		}
		return &Client{gz: rds}, nil
	}
	gr := goredis.NewClient(&goredis.Options{
		Addr:     c.Host,
		Password: c.Pass,
		DB:       c.DB,
	})
	return &Client{gr: gr}, nil
}

func (c *Client) Get(key string) (string, error) {
	if c.gz != nil {
		return c.gz.Get(key)
	}
	return c.gr.Get(context.Background(), key).Result()
}

func (c *Client) Setex(key, value string, seconds int) error {
	if c.gz != nil {
		return c.gz.Setex(key, value, seconds)
	}
	return c.gr.Set(context.Background(), key, value, time.Duration(seconds)*time.Second).Err()
}

func (c *Client) Exists(key string) (bool, error) {
	if c.gz != nil {
		return c.gz.Exists(key)
	}
	n, err := c.gr.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func (c *Client) Incr(key string) (int64, error) {
	if c.gz != nil {
		return c.gz.Incr(key)
	}
	return c.gr.Incr(context.Background(), key).Result()
}

func (c *Client) Del(key string) (int, error) {
	if c.gz != nil {
		return c.gz.Del(key)
	}
	n, err := c.gr.Del(context.Background(), key).Result()
	return int(n), err
}

func (c *Client) Expire(key string, seconds int) error {
	if c.gz != nil {
		return c.gz.Expire(key, seconds)
	}
	return c.gr.Expire(context.Background(), key, time.Duration(seconds)*time.Second).Err()
}
