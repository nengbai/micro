package controller

import (
	"micro/pkg/result"
	"micro/pkg/validCheck"
	"micro/request"
	"micro/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func NewUsersController() UsersController {
	return UsersController{}
}

//得到一个用户的详情
func (a *UsersController) GetUsersOne(c *gin.Context) {
	result := result.NewResult(c)
	param := request.UsersRequest{ID: validCheck.StrTo(c.Param("id")).MustUInt64()}
	valid, errs := validCheck.BindAndValid(c, &param)
	if !valid {
		result.Error(400, errs.Error())
		return
	}

	userOne, err := service.GetOneUser(param.ID)
	if err != nil {
		result.Error(404, "数据查询错误")
	} else {
		result.Success(&userOne)
	}
	return
}

//插入篇文章
func (a *UsersController) InsertUsersOne(c *gin.Context) {
	//result := result.NewResult(c)
	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	service.InsertUsersOne(name, password, email, phone)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"name":     name,
		"password": password,
		"email":    email,
		"phone":    phone,
	})
	// if err != nil {
	// 	result.Error(404, "数据插入错误")
	// } else {
	// 	//result.Success(&articleOne)
	// 	fmt.Printf("success")
	// }
	// return
}

func (a *UsersController) RegistryUsersOne(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	service.InsertUsersOne(name, password, email, phone)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"name":     name,
		"password": password,
		"email":    email,
		"phone":    phone,
	})
}
