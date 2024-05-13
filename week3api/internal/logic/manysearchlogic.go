package logic

import (
	"context"
	"encoding/json"
	"zg4-1/src/rk/week3/week3rpc/week3rpcclient"

	"zg4-1/src/rk/week3/week3api/internal/svc"
	"zg4-1/src/rk/week3/week3api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManySearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManySearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManySearchLogic {
	return &ManySearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManySearchLogic) ManySearch(req *types.ManySearchRequest) (resp *types.ManySearchResponse, err error) {
	// todo: 多条件搜索
	res, err := l.svcCtx.Week3Server.ManySearch(l.ctx, &week3rpcclient.ManySearchRequest{
		Index:    req.Index,
		Keyword:  req.Keyword,
		Keywords: req.Keywords,
	})
	if res.Code != 200 {
		return &types.ManySearchResponse{
			Code: "400",
			Msg:  "多条件搜索成功",
			Data: nil,
		}, nil
	}
	var a map[string]interface{}
	json.Unmarshal(res.Data, &a)
	return &types.ManySearchResponse{
		Code: "200",
		Msg:  "多条件搜索成功",
		Data: a,
	}, nil
}
