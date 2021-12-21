package dao

import (
	"fmt"
	"micro/global"
	"micro/model"
)

//select一条记录
func SelectOneUsers(Id uint64) (*model.Users, error) {
	fields := []string{"userId", "name", "password", "introduce", "hobby", "email", "phone", "gender", "age"}
	userOne := &model.Users{}
	err := global.DBLink.Select(fields).Where("userId=?", Id).First(&userOne).Error
	if err != nil {
		return nil, err
	} else {
		return userOne, nil
	}
}

//select总数
func SelectUserscountAll() (int, error) {
	var count int
	err := global.DBLink.Table(model.Users{}.TableName()).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

//select所有记录
func SelectListUsers(pageOffset int, pageSize int) ([]*model.Users, error) {
	fields := []string{"userId", "name", "password", "introduce", "hobby", "email", "phone", "gender", "age"}
	userOne := &model.Users{}
	rows, err := global.DBLink.Select(fields).Table(userOne.TableName()).Offset(pageOffset).Limit(pageSize).Rows()
	//rows, err := global.DBLink.Select(fields).Table(userOne.TableName()).Rows()
	if err != nil {
		fmt.Println("sql is error:")
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	var users []*model.Users
	for rows.Next() {
		fmt.Println("rows.next:")
		r := &model.Users{}
		if err := rows.Scan(&r.UserId, &r.Name, &r.Password, &r.Introduce, &r.Hobby, &r.Email, &r.Phone, &r.Gender, &r.Age); err != nil {
			fmt.Println("rows.next:")
			fmt.Println(err)
			return nil, err
		}
		users = append(users, r)
	}
	return users, nil
}

//select一条记录
func InsertOneUsers(name, password, email, phone, gender, introduce string, ages int, hobbys string) {
	userOne := &model.Users{
		UserId:    0,
		Name:      name,
		Password:  password,
		Email:     email,
		Phone:     phone,
		Gender:    gender,
		Age:       ages,
		Introduce: introduce,
		Hobby:     hobbys,
	}

	fmt.Printf("%s,%s,%s,%s,%v,%v,%v,%v", userOne.Name, userOne.Password, userOne.Email, userOne.Phone, userOne.Age, userOne.Gender, userOne.Hobby, userOne.Introduce)
	result := global.DBLink.Create(userOne)
	if result.Error != nil {
		fmt.Printf("%v", result.Error)
	}

}

func GetOneUser(name, password string) (*model.Users, error) {
	fields := []string{"name", "password"}
	userOne := &model.Users{}
	err := global.DBLink.Select(fields).Where("name=? and password=?", name, password).First(&userOne).Error
	if err != nil {
		return nil, err
	} else {
		return userOne, nil
	}
}
