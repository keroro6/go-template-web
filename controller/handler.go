package controller

import (
	types2 "go-template-web/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-template-web/service"
)

func GoTemplateWebHandler(svcCtx *service.SrvContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types2.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := service.NewGoTemplateWebLogic(r.Context(), svcCtx)
		resp, err := l.GoTemplateWeb(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
