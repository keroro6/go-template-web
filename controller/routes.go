package controller

import (
	"go-template-web/service"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *service.SrvContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: GoTemplateWebHandler(serverCtx),
			},
		},
	)
}
