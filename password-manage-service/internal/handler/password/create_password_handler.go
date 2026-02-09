package password

import (
	"net/http"

	"password-manage-service/internal/logic/password"
	"password-manage-service/internal/svc"
	"password-manage-service/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreatePasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreatePasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := password.NewCreatePasswordLogic(r.Context(), svcCtx)
		resp, err := l.CreatePassword(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
