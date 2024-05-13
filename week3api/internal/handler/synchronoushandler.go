package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zg4-1/src/rk/week3/week3api/internal/logic"
	"zg4-1/src/rk/week3/week3api/internal/svc"
	"zg4-1/src/rk/week3/week3api/internal/types"
)

func SynchronousHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SynchronousRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSynchronousLogic(r.Context(), svcCtx)
		resp, err := l.Synchronous(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
