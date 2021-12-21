package service

import (
	"micro/cache"
	"micro/dao"
	"micro/model"

	"github.com/go-redis/redis"
)

//得到一篇文章的详情
func GetOneArticle(articleId uint64) (*model.Article, error) {
	//get from cache
	article, err := cache.GetOneArticleCache(articleId)
	if err == redis.Nil || err != nil {
		//get from mysql
		article, errSel := dao.SelectOneArticle(articleId)
		if errSel != nil {
			return nil, errSel
		} else {
			//set cache
			errSet := cache.SetOneArticleCache(articleId, article)
			if errSet != nil {
				return nil, errSet
			} else {
				return article, errSel
			}
		}
	} else {
		return article, err
	}
}

func GetArticleSum() (int, error) {
	return dao.SelectcountAll()
}

//得到多篇文章，按分页返回
func GetArticleList(page int, pageSize int) ([]*model.Article, error) {
	articles, err := dao.SelectAllArticle(page, pageSize)
	if err != nil {
		return nil, err
	} else {
		return articles, nil
	}
}

//插入一篇文章
func InsertArticleOne(subject string, url string) (*model.ArticleBase, error) {
	article, err := dao.InsertOneArticle(subject, url)
	if err != nil {
		//fmt.Printf("errors:%s", err)
		return nil, err
	}
	return article, err

}
