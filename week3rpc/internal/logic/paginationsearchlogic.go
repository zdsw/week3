package logic

import (
	"context"
	"encoding/json"
	"zg4-1/src/rk/week3/models"

	"zg4-1/src/rk/week3/week3rpc/internal/svc"
	"zg4-1/src/rk/week3/week3rpc/week3rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaginationSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaginationSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaginationSearchLogic {
	return &PaginationSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaginationSearchLogic) PaginationSearch(in *week3rpc.PaginationSearchRequest) (*week3rpc.PaginationSearchResponse, error) {
	// todo: 分页搜索
	data, _ := models.PaginationSearchHighlight(in.Index, in.Keyword, int(in.Page), int(in.PageSize))
	a, _ := json.Marshal(data)
	return &week3rpc.PaginationSearchResponse{
		Code: 200,
		Msg:  "分页搜索成功",
		Data: a,
	}, nil
}
