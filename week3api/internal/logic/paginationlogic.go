package logic

import (
	"context"
	"encoding/json"
	"zg4-1/src/rk/week3/week3rpc/week3rpcclient"

	"zg4-1/src/rk/week3/week3api/internal/svc"
	"zg4-1/src/rk/week3/week3api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaginationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaginationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaginationLogic {
	return &PaginationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaginationLogic) Pagination(req *types.PaginationSearchRequest) (resp *types.PaginationSearchResponse, err error) {
	// todo: 分页
	res, err := l.svcCtx.Week3Server.PaginationSearch(l.ctx, &week3rpcclient.PaginationSearchRequest{
		Index:    req.Index,
		Keyword:  req.Keyword,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if res.Code != 200 {
		return &types.PaginationSearchResponse{
			Code: "400",
			Msg:  "分页搜索成功",
			Data: nil,
		}, nil
	}
	var a map[string]interface{}
	json.Unmarshal(res.Data, &a)
	return &types.PaginationSearchResponse{
		Code: "200",
		Msg:  "分页搜索成功",
		Data: a,
	}, nil
}
