package service

import (
	"micro/cache"
	"micro/dao"
	"micro/model"

	"github.com/go-redis/redis"
)

//得到一篇文章的详情
func GetOneUser(ID uint64) (*model.Users, error, string) {
	//get from cache
	user, err := cache.GetOneUsersCache(ID)
	if err == redis.Nil || err != nil {
		//get from mysql
		user, errSel := dao.SelectOneUsers(ID)
		if errSel != nil {
			return nil, errSel, ""
		} else {
			//set cache
			errSet := cache.SetOneUsersCache(ID, user)
			if errSet != nil {
				return nil, errSet, ""
			} else {
				source := "mysql"
				return user, errSel, source
			}
		}
	} else {
		source := "Redis"
		return user, err, source
	}

}

//得到多篇文章，按分页返回
func GetUsersList(page int, pageSize int) ([]*model.Users, error) {
	users, err := dao.SelectListUsers(page, pageSize)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func GetUsersSum() (int, error) {
	return dao.SelectUserscountAll()
}

//插入一篇文章
func InsertUsersOne(name, password, email, phone, gender, introduce string, ages int, hobbys string) (status bool, err error) {
	dao.InsertOneUsers(name, password, email, phone, gender, introduce, ages, hobbys)
	if err != nil {
		return false, err
	}
	return true, nil

}

func GetOneUserbyName(name, password string) (*model.Users, error) {
	user, err := dao.GetOneUser(name, password)
	if err != nil {
		//fmt.Printf("errors:%s", err)
		return nil, err
	}
	return user, err
}
