package dao

import (
	"fmt"
	"micro/global"
	"micro/model"
)

//select一条记录
func SelectOneUsers(Id uint64) (*model.Users, error) {
	fields := []string{"Id", "Name", "Password"}
	userOne := &model.Users{}
	err := global.DBLink.Select(fields).Where("Id=?", Id).First(&userOne).Error
	if err != nil {
		return nil, err
	} else {
		return userOne, nil
	}
}

//select一条记录
func InsertOneUsers(name, password, email, phone string) {
	userOne := &model.Users{
		Name:     name,
		Password: password,
		Email:    email,
		Phone:    phone,
	}

	fmt.Printf("%s,%s,%s,%s", userOne.Name, userOne.Password, userOne.Email, userOne.Phone)
	result := global.DBLink.Create(userOne)
	if result.Error != nil {
		fmt.Printf("%v", result.Error)
	}

}
