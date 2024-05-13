package logic

import (
	"context"
	"encoding/json"
	"zg4-1/src/rk/week3/models"

	"zg4-1/src/rk/week3/week3rpc/internal/svc"
	"zg4-1/src/rk/week3/week3rpc/week3rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManySearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewManySearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManySearchLogic {
	return &ManySearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 多条件搜索接口
func (l *ManySearchLogic) ManySearch(in *week3rpc.ManySearchRequest) (*week3rpc.ManySearchResponse, error) {
	// todo: 多条件搜索接口
	data, _ := models.ManySearch(in.Index, in.Keyword, in.Keywords)
	a, _ := json.Marshal(data)
	return &week3rpc.ManySearchResponse{
		Code: 200,
		Msg:  "多条件搜索成功",
		Data: a,
	}, nil
}
