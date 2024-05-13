package logic

import (
	"context"
	"encoding/json"
	"zg4-1/src/rk/week3/models"

	"zg4-1/src/rk/week3/week3rpc/internal/svc"
	"zg4-1/src/rk/week3/week3rpc/week3rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关键词搜索接口
func (l *SearchLogic) Search(in *week3rpc.SearchRequest) (*week3rpc.SearchResponse, error) {
	// todo: 搜索接口
	data, _ := models.SearchHighlight(in.Index, in.Keyword)
	a, _ := json.Marshal(data)
	return &week3rpc.SearchResponse{
		Code: 200,
		Msg:  "关键词搜索成功",
		Data: a,
	}, nil
}
