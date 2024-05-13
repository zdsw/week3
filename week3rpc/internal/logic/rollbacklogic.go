package logic

import (
	"context"
	"zg4-1/src/rk/week3/models"

	"zg4-1/src/rk/week3/week3rpc/internal/svc"
	"zg4-1/src/rk/week3/week3rpc/week3rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackLogic {
	return &RollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 回滚
func (l *RollbackLogic) Rollback(in *week3rpc.RollbackRequest) (*week3rpc.RollbackResponse, error) {
	// todo: 回滚
	err := models.Rollback(in.Id, in.FileName, in.FileVersion, in.FileStatus)
	if err != nil {
		return &week3rpc.RollbackResponse{
			Code: 500,
			Msg:  "回滚失败",
			Data: nil,
		}, nil
	}
	
	return &week3rpc.RollbackResponse{
		Code: 200,
		Msg:  "回滚成功",
		Data: nil,
	}, nil
}
