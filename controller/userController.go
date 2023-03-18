package controller

import (
	"fmt"
	"micro/gin_session"
	"micro/pkg/page"
	"micro/pkg/result"
	"micro/pkg/validCheck"
	"micro/request"
	"micro/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func NewUsersController() UsersController {
	return UsersController{}
}

// 得到一个用户的详情
func (a *UsersController) GetUsersOne(c *gin.Context) {
	result := result.NewResult(c)
	param := request.UsersRequest{ID: validCheck.StrTo(c.Param("userId")).MustUInt64()}
	fmt.Printf("userId:%v", &param)
	valid, errs := validCheck.BindAndValid(c, &param)
	if !valid {
		result.Error(400, errs.Error())
		return
	}

	userOne, err, source := service.GetOneUser(param.ID)
	if err != nil {
		result.Error(404, "数据查询错误")
	} else {
		result.Success(&userOne, source)
	}
}

// 得到一篇文章的详情
func (a *ArticleController) GetUsersOne(c *gin.Context) {
	result := result.NewResult(c)
	param := request.UsersRequest{ID: validCheck.StrTo(c.Param("userId")).MustUInt64()}
	valid, errs := validCheck.BindAndValid(c, &param)
	if !valid {
		result.Error(400, errs.Error())
		return
	}

	userOne, err, source := service.GetOneUser(param.ID)
	if err != nil {
		result.Error(404, "数据查询错误")
	} else {

		result.Success(&userOne, source)
	}
}

// 得到多个用户，按分页返回
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

	pageSize := 5
	pageOffset := (pageInt - 1) * pageSize

	users, err := service.GetUsersList(pageOffset, pageSize)
	if err != nil {
		result.Error(404, "数据查询错误")
		fmt.Println(err.Error())
	} else {
		//sum,_ := dao.SelectcountAll()
		sum, _ := service.GetUsersSum()
		pageInfo, _ := page.GetPageInfo(pageInt, pageSize, sum)
		source := "MySQL"
		result.Success(gin.H{"list": &users, "pageinfo": pageInfo}, source)
	}
}

// 插入篇文章
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
	path := c.Request.URL.RequestURI()
	if c.Request.Method == "POST" { //判断请求的方法，先判是否为post
		//toPath := c.DefaultQuery("next", "/index") //一个路径，用于后面的重定向
		fmt.Printf("Path---->:%s\n", path)
		var u UserInfo
		//绑定，并解析参数
		err := c.ShouldBind(&u)
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或密码不能为空",
			})
			return
		} else {
			name := strings.Trim(c.PostForm("username"), " ")
			password := strings.Trim(c.PostForm("password"), " ")
			fmt.Printf("user:%v,password:%v\n", name, password)
			if name != "" && password != "" {
				user, err := service.GetOneUserbyName(name, password)
				if err != nil {
					c.HTML(http.StatusOK, "login.html", gin.H{
						"err": err.Error(),
					})
					return
				}
				if user.Name == name && user.Password == password {
					tmpSD, ok := c.Get(gin_session.SessionContextName)
					fmt.Printf("tmpSD:%v\n", tmpSD)
					if !ok {
						panic("session middleware")
					}
					sd := tmpSD.(gin_session.SessionData)
					fmt.Printf("sd:%v\n", sd)
					// 2. 给session data设置isLogin = true
					sd.Set("isLogin", true)
					sd.Set("Username", user.Name)
					//调用Save，存储到Redis
					sd.Save()
					//跳转到index界面
					//fmt.Printf("toPath------>:%v", toPath)

					value, err := sd.Get("toPath")
					if err != nil {
						fmt.Printf("toPath error:%v\n", err)
					}
					fmt.Printf("toPath value:%v\n", value)
					//toPath := "/index"
					//c.Redirect(http.StatusMovedPermanently, toPath)
					//c.Redirect(http.StatusTemporaryRedirect, toPath)
					//c.Redirect(http.StatusPermanentRedirect,toPath)
					//return
					redirehtml := RedirecFunc(value)
					c.HTML(http.StatusOK, redirehtml, gin.H{
						"username": user.Name,
						"password": user.Password,
						"isLogin":  true,
					})

				}
			} else {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"err": "用户名或密码不能为空",
				})
				return
			}
		}

	} else { //get
		c.HTML(http.StatusOK, "login.html", nil)
	}

}
func RedirecFunc(toPath interface{}) string {
	var redirehtml string
	switch toPath {
	case "/vip":
		redirehtml = "vip.html"
	case "/home":
		redirehtml = "home.html"
	case "/users/list":
		redirehtml = "users-list.html"
	default:
		redirehtml = "index.html"
	}
	return redirehtml
}
