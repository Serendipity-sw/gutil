/*
*
redis帮助类
create by gloomysw 2017-3-24 14:11:284
*/
package gutil

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	rpool *redis.Pool
)

// redis通道开启
// create by gloomysw 2017-3-24 14:18:03
// addr IP地址+端口     idx 仓库数
func OpenRedis(addr string, idx int) error {
	rpool = &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: 600 * time.Second,
		Dial: func() (redis.Conn, error) {
			do := redis.DialDatabase(idx)
			c, err := redis.Dial("tcp", addr, do)
			if err != nil {
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	conn := rpool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	return err
}

// redis通道关闭
// create by gloomysw 2017-3-24 14:20:10
func CloseRedis() {
	rpool.Close()
}

// 设置redis缓存
// create by gloomysw 2017-3-24 14:24:31
// key 存储键名    value 存储值    cacheBssSeconds 存储时间(单位秒)
func SetRedisCache(key, value string, cacheBssSeconds int) error {
	c := rpool.Get()
	defer c.Close()
	_, err := c.Do("SETEX", key, cacheBssSeconds, value)
	return err
}

// 获取redis缓存
// create by gloomysw 2017-3-24 14:29:16
// key 存储键名
func GetRedisCache(key string) (string, error) {
	c := rpool.Get()
	defer c.Close()
	reply, err := c.Do("GET", key)
	if err != nil {
		return "", err
	}
	if reply == nil {
		return "", fmt.Errorf("Not found key: %s ", key)
	}
	return string(reply.([]byte)), nil
}
