package logic

import (
	"context"
	"zg4-1/src/rk/week3/week3rpc/week3rpcclient"

	"zg4-1/src/rk/week3/week3api/internal/svc"
	"zg4-1/src/rk/week3/week3api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackLogic {
	return &RollbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RollbackLogic) Rollback(req *types.RollbackRequest) (resp *types.RollbackResponse, err error) {
	// todo: 回滚
	res, err := l.svcCtx.Week3Server.Rollback(l.ctx, &week3rpcclient.RollbackRequest{
		Id:          req.Id,
		FileName:    req.FileName,
		FileVersion: req.FileVersion,
		FileStatus:  req.FileStatus,
	})
	if res.Code != 200 {
		return &types.RollbackResponse{
			Code: "400",
			Msg:  "回滚成功",
			Data: nil,
		}, nil
	}
	return &types.RollbackResponse{
		Code: "200",
		Msg:  "回滚成功",
	}, nil
}
