package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Ctx *gin.Context
}

// 返回的结果：
type ResultCont struct {
	Code   int         `json:"code"`   //提示代码
	Msg    string      `json:"msg"`    //提示信息
	Data   interface{} `json:"data"`   //出错
	Source string      `json:"source"` // 数据来源
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

// 成功
func (r *Result) Success(data interface{}, source string) {
	if data == nil {
		data = gin.H{}
	}
	res := ResultCont{}
	res.Code = 0
	res.Msg = ""
	res.Data = data
	res.Source = source
	r.Ctx.JSON(http.StatusOK, res)
}

// 出错
func (r *Result) Error(code int, msg string) {
	res := ResultCont{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusOK, res)
}
