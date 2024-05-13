package logic

import (
	"context"
	"encoding/json"
	"zg4-1/src/rk/week3/week3rpc/week3rpcclient"

	"zg4-1/src/rk/week3/week3api/internal/svc"
	"zg4-1/src/rk/week3/week3api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// todo: 关键词搜索
	res, err := l.svcCtx.Week3Server.Search(l.ctx, &week3rpcclient.SearchRequest{
		Index:   req.Index,
		Keyword: req.Keyword,
	})
	if res.Code != 200 {
		return &types.SearchResponse{
			Code: "400",
			Msg:  "搜索成功",
			Data: nil,
		}, nil
	}
	var a map[string]interface{}
	json.Unmarshal(res.Data, &a)
	return &types.SearchResponse{
		Code: "200",
		Msg:  "搜索成功",
		Data: a,
	}, nil
}
