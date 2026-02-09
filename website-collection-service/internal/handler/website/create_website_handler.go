package website

import (
	"net/http"

	"website-collection-service/internal/logic/website"
	"website-collection-service/internal/svc"
	"website-collection-service/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateWebsiteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateWebsiteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := website.NewCreateWebsiteLogic(r.Context(), svcCtx)
		resp, err := l.CreateWebsite(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
