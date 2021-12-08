package handler

import (
	"net/http"

	"micro/internal/logic"
	"micro/internal/svc"
	"micro/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MicroHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMicroLogic(r.Context(), ctx)
		resp, err := l.Micro(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
