// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	v1user "goZeroScaffold/internal/handler/v1/user"
	"goZeroScaffold/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthInterceptor},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: v1user.ListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/user"),
		rest.WithTimeout(3000*time.Millisecond),
		rest.WithMaxBytes(1048576),
	)
}
