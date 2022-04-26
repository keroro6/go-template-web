package controller

import (
	"github.com/zeromicro/go-zero/rest/httpx"

	"go-template-web/service"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *service.SrvContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: healthHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/metrics",
				Handler: metricsHandler(serverCtx),
			},
		},
	)
}

// healthHandler heart beat
func healthHandler(svcCtx *service.SrvContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.Ok(w)
	}
}

// metricsHandler prometheus report
func metricsHandler(svcCtx *service.SrvContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.Ok(w)
	}
}
