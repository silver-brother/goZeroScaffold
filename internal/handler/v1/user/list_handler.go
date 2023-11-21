package user

import (
	"github.com/zeromicro/x/errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"goZeroScaffold/internal/logic/v1/user"
	"goZeroScaffold/internal/svc"
	"goZeroScaffold/internal/types"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserListReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.CodeMsg{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
			})
			return
		}

		l := user.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
