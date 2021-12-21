package controller

import (
	"fmt"
	"micro/pkg/page"
	"micro/pkg/result"
	"micro/pkg/validCheck"
	"micro/request"
	"micro/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func NewUsersController() UsersController {
	return UsersController{}
}

//得到一个用户的详情
func (a *UsersController) GetUsersOne(c *gin.Context) {
	result := result.NewResult(c)
	param := request.UsersRequest{ID: validCheck.StrTo(c.Param("userId")).MustUInt64()}
	fmt.Printf("userId:%v", &param)
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
}

//得到一篇文章的详情
func (a *ArticleController) GetUsersOne(c *gin.Context) {
	result := result.NewResult(c)
	param := request.UsersRequest{ID: validCheck.StrTo(c.Param("userId")).MustUInt64()}
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
}

//得到多个用户，按分页返回
func (a *UsersController) GetUserList(c *gin.Context) {
	result := result.NewResult(c)
	pageInt := 0
	//is exist?
	curPage := c.Query("page")
	//if curPage not exist
	if len(curPage) == 0 {
		pageInt = 1
	} else {
		param := request.UsersListRequest{Page: validCheck.StrTo(c.Param("page")).MustInt()}
		valid, errs := validCheck.BindAndValid(c, &param)
		if !valid {
			result.Error(400, errs.Error())
			return
		}
		pageInt = param.Page
	}

	pageSize := 4
	pageOffset := (pageInt - 1) * pageSize

	users, err := service.GetUsersList(pageOffset, pageSize)
	if err != nil {
		result.Error(404, "数据查询错误")
		fmt.Println(err.Error())
	} else {
		//sum,_ := dao.SelectcountAll()
		sum, _ := service.GetUsersSum()
		pageInfo, _ := page.GetPageInfo(pageInt, pageSize, sum)
		result.Success(gin.H{"list": &users, "pageinfo": pageInfo})
	}
}

//插入篇文章
func (a *UsersController) InsertUsersOne(c *gin.Context) {
	//result := result.NewResult(c)
	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	gender := c.PostForm("gender")
	introduce := c.PostForm("introduce")
	age := c.PostForm("age")
	ages, _ := strconv.Atoi(age)
	var fakeForm MyForm
	var hobbys string
	c.ShouldBind(&fakeForm)
	if fakeForm.Hobbys != nil {
		for _, i := range fakeForm.Hobbys {
			if len(i) > 0 {
				hobbys = hobbys + " " + i
				fmt.Println(hobbys)
			}

		}

	}
	fmt.Printf("%v,%v,%v,%v,%v,%v,%v", name, email, password, gender, phone, introduce, ages)
	_, err := service.InsertUsersOne(name, password, email, phone, gender, introduce, ages, hobbys)
	c.HTML(http.StatusOK, "/users/login.html", gin.H{
		"name":     name,
		"password": password,
		"email":    email,
		"phone":    phone,
	})
	if err != nil {
		fmt.Printf("404,数据插入错误")
	} else {
		fmt.Printf("success")
	}
	//return
}

type MyForm struct {
	Hobbys []string `form:"hobby[]"`
}

func (a *UsersController) RegistryUsersOne(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	gender := c.PostForm("gender")
	phone := c.PostForm("phone")
	age := string(c.PostForm("age"))
	introduce := c.PostForm("introduce")
	ages, _ := strconv.Atoi(age)

	var fakeForm MyForm
	var hobbys string
	c.ShouldBind(&fakeForm)
	if fakeForm.Hobbys != nil {
		for _, i := range fakeForm.Hobbys {
			if len(i) > 0 {
				hobbys = hobbys + " " + i
				fmt.Println(hobbys)
			}

		}

	}
	fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v", name, email, password, gender, phone, introduce, ages, hobbys)
	service.InsertUsersOne(name, password, email, phone, gender, introduce, ages, hobbys)
	c.JSON(http.StatusOK, gin.H{
		"name":      name,
		"password":  password,
		"email":     email,
		"phone":     phone,
		"gender":    gender,
		"age":       age,
		"introduce": introduce,
		"Hobby":     hobbys,
	})
}

func (a *UsersController) UserLogin(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	fmt.Printf("%v,%v", name, password)
	if name != "" && password != "" {
		user, err := service.GetOneUserbyName(name, password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"errors": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name":     user.Name,
				"password": user.Password,
			})
		}

	} else {
		c.String(http.StatusOK, "There are empty for username and password,please check... ")
	}

}
