package utils

import (
	"context"
	"efeasy-gin/global"
	"github.com/go-redis/redis/v8"
	"time"
)

type Interface interface {
	Get() bool
	Block(seconds int64) bool
	Release() bool
	ForceRelease()
}

type lock struct {
	context context.Context
	name    string // 锁名称
	owner   string // 锁标识
	seconds int64  // 有效期
}

// 释放锁 Lua 脚本，防止任何客户端都能解锁
const releaseLockLuaScript = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
`

// Lock 生成锁
func Lock(name string, seconds int64) Interface {
	return &lock{
		context.Background(),
		name,
		RandString(16),
		seconds,
	}
}

// Get 获取锁
func (l *lock) Get() bool {
	return global.App.Redis.SetNX(l.context, l.name, l.owner, time.Duration(l.seconds)*time.Second).Val()
}

// Block 阻塞一段时间，尝试获取锁
func (l *lock) Block(seconds int64) bool {
	starting := time.Now().Unix()
	for {
		if !l.Get() {
			time.Sleep(time.Duration(1) * time.Second)
			if time.Now().Unix()-seconds >= starting {
				return false
			}
		} else {
			return true
		}
	}
}

// Release 释放锁
func (l *lock) Release() bool {
	luaScript := redis.NewScript(releaseLockLuaScript)
	result := luaScript.Run(l.context, global.App.Redis, []string{l.name}, l.owner).Val().(int64)
	return result != 0
}

// ForceRelease 强制释放锁
func (l *lock) ForceRelease() {
	global.App.Redis.Del(l.context, l.name).Val()
}
