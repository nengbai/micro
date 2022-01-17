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
	return router
}

//404
func HandleNotFound(c *gin.Context) {
	result.NewResult(c).Error(404, "资源未找到")

}

//500
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

//用户信息
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
