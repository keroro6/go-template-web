package modelcontroller

import (
	"go-template-web/internalapi"
	"go-template-web/service/modelservice"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-template-web/service"
)

func GoTemplateWebHandler(svcCtx *service.SrvContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req internalapi.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := modelservice.NewModelSrvLogic(r.Context(), svcCtx)
		resp, err := l.GoTemplateWeb(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
