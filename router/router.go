package router

import (
	"fmt"
	"log"
	"micro/controller"
	"micro/gin_session"
	"micro/pkg/result"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(response())
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(Recover)

	// 加载静态页面
	router.Static("/static", "./static")
	router.LoadHTMLGlob("static/templates/*")
	// router.LoadHTMLFiles("static/html/index.html", "static/html/users/login.html", "static/html/index.html", "static/html/users/registry.html", "static/html/users/test.html")
	// router.LoadHTMLFiles("static/html/index.html", "static/html/users/registry.html", "static/html/users/test.html")
	gin_session.InitMgr("redis", "")
	router.Use(gin_session.SessionMiddleware(gin_session.MgrObj))
	router.Use(cors())
	//router.Any("/login", loginHandler)
	router.GET("/", indexHandler)
	router.Any("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	router.GET("/vip", AuthMiddleware, vipHandler)
	router.GET("/home", AuthMiddleware, homeHandler)
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	//router.POST("/getname", Posthandlefunc)
	// 路径映射
	articlec := controller.NewArticleController()
	router.GET("/article/getone/:id", articlec.GetOneArticle)
	router.GET("/article/list", articlec.GetList)
	router.POST("/", articlec.InsertArticleOne)
	userc := controller.NewUsersController()
	v1 := router.Group("/users")
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "users v1's homepage")
		})
		v1.POST("/registry", userc.RegistryUsersOne)
		v1.GET("/list", AuthMiddleware, userc.GetUserList)
		v1.GET("/getone/:userId", userc.GetUsersOne)
		v1.Any("/login", userc.UserLogin)

	}
	router.GET("/registry", HandleRegistry)
	//router.GET("/login", HandleLogin)

	type Student struct {
		Username string `json:"name"`
		Password string `json:"passwd"`
	}

	router.Handle("POST", "/hello", func(ctx *gin.Context) {
		fmt.Println(ctx.FullPath())
		var name Student
		err := ctx.ShouldBindJSON(&name)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(name.Username)
		fmt.Println(name.Password)
		n := ctx.PostForm(name.Username)
		p := ctx.PostForm(name.Password)
		ctx.JSON(200, gin.H{
			"passwd": n,
			"name":   p,
		})

	})
	// curl "http://localhost:8080/test?name=Tom&role=student"
	router.GET("/test", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	// 动态的路由 匹配 /test/geektutu
	router.GET("/test/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	// POST curl http://localhost:8080/form  -X POST -d 'username=geektutu&password=1234'
	router.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	//使用ShouldBindQuery
	router.GET("/student1", func(c *gin.Context) {
		var student Student
		if err := c.ShouldBindQuery(&student); err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "fail"})
		} else {
			fmt.Println(student)
			c.JSON(http.StatusOK, gin.H{
				"msg":      "success",
				"username": student.Username,
				"password": student.Password,
			})
		}
	})

	// POST json  ShouldBind curl http://localhost:8080/json  -X POST -d '{"name":"baineng","passwd":"1234"}'
	router.POST("/json", func(c *gin.Context) {
		var u Student
		// u.Username = "baineng"
		// u.Password = "madd11"
		err := c.ShouldBindJSON(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			u.Username = c.Param("name")
			u.Password = c.Param("passwd")
			fmt.Printf("%#v\n", u.Username)
			c.JSON(http.StatusOK, gin.H{
				"message":  "ok",
				"username": u.Username,
				"passwd":   u.Password,
			})
		}

	})

	//使用BindQuery
	router.GET("/student2", func(c *gin.Context) {
		var stu Student
		if err := c.BindQuery(&stu); err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "fail"})
		}
	})

	// GET 和 POST 混合 curl "http://localhost:8080/posts?id=9876&page=7"  -X POST -d 'name=geektutu&passwd=1234'
	router.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("name")
		password := c.DefaultPostForm("passwd", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})
	// Map参数(字典参数) curl -g "http://localhost:8080/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
	// 重定向(Redirect) curl -i http://localhost:8080/redirect
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
	// curl "http://localhost:8080/goindex"
	router.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		router.HandleContext(c)
	})

	// group routes 分组路由
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	// group: v1 curl http://localhost:8080/v1/posts
	v1 = router.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	// group: v2 curl http://localhost:8080/v2/posts
	v2 := router.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}
	// post 单个文件
	router.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})
	// post 多个文件
	router.POST("/upload2", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})

	// HTML模板(Template) curl http://localhost:8080/arr
	type students struct {
		Name string
		Age  int8
	}
	router.LoadHTMLGlob("static/templates/*")
	stu1 := &students{Name: "Geektutu", Age: 20}
	stu2 := &students{Name: "Jack", Age: 22}
	router.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*students{stu1, stu2},
		})
	})

	return router
}

// 404
func HandleNotFound(c *gin.Context) {
	result.NewResult(c).Error(404, "资源未找到")

}

// 500
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			result.NewResult(c).Error(500, "服务器内部错误")
		}
	}()
	//继续后续接口调用
	c.Next()
}

func IndexProcess(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("<h1>404</h1>"))
	} else {
		w.Write([]byte("index"))
	}

}

func HandleLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Microcservices Demo",
	})
}

func HandleRegistry(c *gin.Context) {
	c.HTML(http.StatusOK, "registry.html", gin.H{
		"title": "Microcservices Demo",
	})
}

func Posthandlefunc(c *gin.Context) {
	name := c.PostForm("Subject")
	link := c.PostForm("Url")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name": name,
		"link": link,
	})
}

func response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Written() {
			return
		}

		params := c.Keys
		if len(params) == 0 {
			return
		}
		c.JSON(http.StatusOK, params)
	}
}

// 用户信息
type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// 编写一个校验用户是否登录的中间件
// 其实就是从上下文中取到session data,从session data取到isLogin
func AuthMiddleware(c *gin.Context) {
	// 1. 从上下文中取到session data
	// 1. 先从上下文中获取session data
	toPath := c.Request.URL.Path
	fmt.Println("Checking Auth")
	tmpSD, _ := c.Get(gin_session.SessionContextName)
	sd := tmpSD.(gin_session.SessionData)
	// 2. 从session data取到isLogin
	fmt.Printf("sd---->:%v\n", sd)
	value, err := sd.Get("isLogin")
	if err != nil {
		fmt.Printf("error ----->:%v\n", err)
		// 取不到就是没有登录
		sd.Set("toPath", toPath)
		c.Redirect(http.StatusFound, "/users/login")
		return
	}
	fmt.Printf("value----->:%v \n", value)
	isLogin, ok := value.(bool) //类型断言
	if !ok {
		fmt.Println("!ok")
		sd.Set("toPath", toPath)
		c.Redirect(http.StatusFound, "/users/login")
		return
	}
	if !isLogin {
		c.Redirect(http.StatusFound, "/users/login")
		return
	}
	c.Next()
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func homeHandler(c *gin.Context) {
	tmpSD, ok := c.Get(gin_session.SessionContextName)
	if !ok {
		panic("session middleware")
	}
	sd := tmpSD.(gin_session.SessionData)
	username, err := sd.Get("Username")
	if err != nil {
		fmt.Printf("error ----->:%v\n", err)
		// 取不到就是没有登录
		c.Redirect(http.StatusFound, "/users/login")
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
		"isLogin":  true,
	})
}

func vipHandler(c *gin.Context) {
	tmpSD, ok := c.Get(gin_session.SessionContextName)
	if !ok {
		panic("session middleware")
	}
	sd := tmpSD.(gin_session.SessionData)
	username, err := sd.Get("Username")
	if err != nil {
		fmt.Printf("error ----->:%v\n", err)
		// 取不到就是没有登录
		c.Redirect(http.StatusFound, "/users/login")
		return
	}
	c.HTML(http.StatusOK, "vip.html", gin.H{
		"username": username,
	})
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
