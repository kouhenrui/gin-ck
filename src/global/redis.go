package global

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

/**
 * @ClassName redis
 * @Description TODO
 * @Author khr
 * @Date 2023/7/31 11:02
 * @Version 1.0
 */

func Redisinit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: RedisConfig.Host + ":" + RedisConfig.Port,
		//Username:   redisCon.UserName,
		//Password:   redisCon.PassWord,
		DB:         RedisConfig.Db,
		PoolSize:   RedisConfig.PoolSize,
		MaxRetries: RedisConfig.MaxRetries,
	})
	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		//fmt.Printf("redis connect get failed.%v", err.Error())
		log.Fatalf("redis connect get failed.%v", err.Error())
		return
	}
	fmt.Printf("redis 初始化连接成功")
	log.Printf("redis 初始化连接成功")
}

// 添加数据
func SetRedis(key string, value []byte, t time.Duration) error {
	return RedisClient.Set(ctx, key, value, t).Err()
	//expire := time.Duration(t)
	//if err = RedisClient.Set(ctx, key, value, t).Err(); err != nil {
	//	fmt.Println(err, "redis存放错误")
	//	return false
	//}
	//fmt.Println("redis存放时间", t)
	//return true
}

// set 中是否存在某个成员
func ExistRedis(key string) error {
	return RedisClient.Exists(ctx, key).Err()

}

// 获取数据
func GetRedis(key string) string {
	result, _ := RedisClient.Get(ctx, key).Result()
	return result
}

// 获取数据
func GetLimitRedis(key string) int {
	result, _ := RedisClient.Get(ctx, key).Int()
	return result
}

// 删除数据
func DelRedis(key string) error {
	err = RedisClient.Del(ctx, key).Err()
	return err
}

// 延长过期时间
func ExpireRedis(key string, t time.Duration) error {
	if err := RedisClient.Expire(ctx, key, t).Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

/*
 * @MethodName
 * @Description redis自增
 * @Author khr
 * @Date 2023/7/31 15:25
 */
func AutoInc(key string) error {
	return RedisClient.Incr(ctx, key).Err()
}

/*
 * @MethodName
 * @Description
 * @Author khr
 * @Date 2023/7/31 16:21
 */
func RpushRedis(name string, key string) error {
	return RedisClient.RPush(ctx, name, key).Err()
}
