package modelcontroller

import (
	"github.com/zeromicro/go-zero/rest"
	"go-template-web/service"
	"net/http"
)

func RegisterHandlers(server *rest.Server, serverCtx *service.SrvContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: GoTemplateWebHandler(serverCtx),
			},
		},
	)
}
