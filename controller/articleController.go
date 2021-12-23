package controller

import (
	"fmt"
	"micro/pkg/page"
	"micro/pkg/result"
	"micro/pkg/validCheck"
	"micro/request"
	"micro/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func NewArticleController() ArticleController {
	return ArticleController{}
}

//得到一篇文章的详情
func (a *ArticleController) GetOneArticle(c *gin.Context) {
	result := result.NewResult(c)
	param := request.ArticleRequest{ID: validCheck.StrTo(c.Param("id")).MustUInt64()}
	valid, errs := validCheck.BindAndValid(c, &param)
	if !valid {
		result.Error(400, errs.Error())
		return
	}

	articleOne, err, source := service.GetOneArticle(param.ID)
	if err != nil {
		result.Error(404, "数据查询错误")
	} else {
		result.Success(&articleOne, source)
	}

}

//得到多篇文章，按分页返回
func (a *ArticleController) GetList(c *gin.Context) {
	result := result.NewResult(c)
	pageInt := 0
	//is exist?
	curPage := c.Query("page")
	//if curPage not exist
	if len(curPage) == 0 {
		pageInt = 1
	} else {
		param := request.ArticleListRequest{Page: validCheck.StrTo(c.Param("page")).MustInt()}
		valid, errs := validCheck.BindAndValid(c, &param)
		if !valid {
			result.Error(400, errs.Error())
			return
		}
		pageInt = param.Page
	}

	pageSize := 5
	pageOffset := (pageInt - 1) * pageSize

	articles, err := service.GetArticleList(pageOffset, pageSize)
	if err != nil {
		result.Error(404, "数据查询错误")
		fmt.Println(err.Error())
	} else {
		//sum,_ := dao.SelectcountAll()
		sum, _ := service.GetArticleSum()
		pageInfo, _ := page.GetPageInfo(pageInt, pageSize, sum)
		source := "MySQL"
		result.Success(gin.H{"list": &articles, "pageinfo": pageInfo}, source)
	}

}

//插入篇文章
func (a *ArticleController) InsertArticleOne(c *gin.Context) {
	//result := result.NewResult(c)
	name := c.PostForm("Subject")
	link := c.PostForm("Url")
	article, err := service.InsertArticleOne(name, link)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errors": err.Error(),
		})
	} else {
		// c.HTML(http.StatusOK, "index.html", gin.H{
		// 	"name": article.Subject,
		// 	"url":  article.Url,
		// })
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"name":   article.Subject,
			"url":    article.Url,
		})
	}
}
