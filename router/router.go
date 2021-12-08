package router

import (
	"log"
	"micro/controller"
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
	// router.Static("/static", "./static")
	// router.LoadHTMLGlob("templates/*")
	router.LoadHTMLFiles("static/login.html", "static/index.html", "static/registry.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// post

	router.POST("/getname", Posthandlefunc)
	// 路径映射
	/** router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
		fmt.Printf("c.Errors: %v\n", c.Errors)
	}) **/
	// router.GET("/index.html", IndexHandler)
	articlec := controller.NewArticleController()
	router.GET("/article/getone/:id", articlec.GetOne)
	router.GET("/article/list", articlec.GetList)
	router.POST("/", articlec.InsertArticleOne)

	userc := controller.NewUsersController()
	router.GET("/registry", HandleRegistry)
	router.POST("/registry", userc.RegistryUsersOne)
	router.GET("/login", HandleLogin)
	router.POST("/login", userc.InsertUsersOne)
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

func IndexHandler(c *gin.Context) {
	c.HTML(
		http.StatusOK, "index.tmpl", gin.H{
			"title": "Microcservices Demo",
		})
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
