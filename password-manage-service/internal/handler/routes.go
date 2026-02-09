package handler

import (
	"net/http"

	passwordHandler "password-manage-service/internal/handler/password"
	"password-manage-service/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/password/create",
				Handler: passwordHandler.CreatePasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/password/list",
				Handler: passwordHandler.GetPasswordListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/password/detail",
				Handler: passwordHandler.GetPasswordDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/password/update",
				Handler: passwordHandler.UpdatePasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/password/delete",
				Handler: passwordHandler.DeletePasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/"),
	)
}
