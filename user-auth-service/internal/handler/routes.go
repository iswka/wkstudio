package handler

import (
	"net/http"

	authHandler "user-auth-service/internal/handler/auth"
	userHandler "user-auth-service/internal/handler/user"
	"user-auth-service/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/auth/register",
				Handler: authHandler.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/auth/login",
				Handler: authHandler.LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/user/info",
				Handler: userHandler.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/change-password",
				Handler: userHandler.ChangePasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/"),
	)
}
