package handler

import (
	"net/http"

	websiteHandler "website-collection-service/internal/handler/website"
	"website-collection-service/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/website/create",
				Handler: websiteHandler.CreateWebsiteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/website/list",
				Handler: websiteHandler.GetWebsiteListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/website/detail",
				Handler: websiteHandler.GetWebsiteDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/website/update",
				Handler: websiteHandler.UpdateWebsiteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/website/delete",
				Handler: websiteHandler.DeleteWebsiteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/"),
	)
}
