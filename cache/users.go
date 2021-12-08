package cache

import (
	"encoding/json"
	"fmt"
	"micro/global"
	"micro/model"

	"strconv"
	"time"

	"github.com/go-redis/redis"
)

//token的过期时长
const UsersDuration = time.Minute * 5

//cache的名字
func getUsersCacheName(userId uint64) string {
	return "user_" + strconv.FormatUint(userId, 10)
}

//从cache得到一篇文章
func GetOneUsersCache(userId uint64) (*model.Users, error) {
	key := getArticleCacheName(userId)
	val, err := global.RedisDb.Get(key).Result()

	if err == redis.Nil || err != nil {
		return nil, err
	} else {
		user := model.Users{}
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			//t.Error(target)
			return nil, err
		}
		return &user, nil
	}
}

//向cache保存一篇文章
func SetOneUsersCache(userId uint64, user *model.Users) error {
	key := getArticleCacheName(userId)
	content, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	errSet := global.RedisDb.Set(key, content, UsersDuration).Err()
	if errSet != nil {
		return errSet
	}
	return nil
}
