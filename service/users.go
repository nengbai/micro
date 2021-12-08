package service

import (
	"micro/cache"
	"micro/dao"
	"micro/model"

	"github.com/go-redis/redis"
)

//得到一篇文章的详情
func GetOneUser(userId uint64) (*model.Users, error) {
	//get from cache
	user, err := cache.GetOneUsersCache(userId)
	if err == redis.Nil || err != nil {
		//get from mysql
		user, errSel := dao.SelectOneUsers(userId)
		if errSel != nil {
			return nil, errSel
		} else {
			//set cache
			errSet := cache.SetOneUsersCache(userId, user)
			if errSet != nil {
				return nil, errSet
			} else {
				return user, errSel
			}
		}
	} else {
		return user, err
	}
}

//插入一篇文章
func InsertUsersOne(name, password, email, phone string) {
	dao.InsertOneUsers(name, password, email, phone)
	// if err != nil {
	// 	return err
	// } else {
	// 	return nil
	// }

}
