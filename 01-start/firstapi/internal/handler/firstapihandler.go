package handler

import (
	"net/http"

	"firstapi/internal/logic"
	"firstapi/internal/svc"
	"firstapi/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FirstapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFirstapiLogic(r.Context(), svcCtx)
		resp, err := l.Firstapi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
