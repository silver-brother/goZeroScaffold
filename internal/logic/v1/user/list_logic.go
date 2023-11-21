package user

import (
	"context"
	"github.com/zeromicro/x/errors"
	"goZeroScaffold/internal/common"
	"net/http"

	"goZeroScaffold/internal/svc"
	"goZeroScaffold/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.UserListReq) (resp *types.UserListRes, err error) {
	userDo := l.svcCtx.DB.User.WithContext(l.ctx)
	if req.Sex != "" {
		userDo = userDo.Where(l.svcCtx.DB.User.Sex.Eq(req.Sex))
	}

	count, err := userDo.Count()
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, errors.New(http.StatusInternalServerError, "系统内部错误")
	}
	if count <= 0 {
		return &types.UserListRes{}, nil
	}

	users, err := userDo.Scopes(common.Paginate(req.Page, req.PageSize)).Find()
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, errors.New(http.StatusInternalServerError, "系统内部错误")
	}
	list := make([]*types.UserListItem, 0)
	for _, user := range users {
		list = append(list, &types.UserListItem{
			ID: user.ID,
		})
	}
	return &types.UserListRes{
		Total: count,
		Data:  list,
	}, nil
}
