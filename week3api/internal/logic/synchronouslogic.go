package logic

import (
	"context"
	"zg4-1/src/rk/week3/week3rpc/week3rpcclient"

	"zg4-1/src/rk/week3/week3api/internal/svc"
	"zg4-1/src/rk/week3/week3api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SynchronousLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSynchronousLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SynchronousLogic {
	return &SynchronousLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SynchronousLogic) Synchronous(req *types.SynchronousRequest) (resp *types.SynchronousResponse, err error) {
	// todo: 数据同步
	res, err := l.svcCtx.Week3Server.Synchronous(l.ctx, &week3rpcclient.SynchronousRequest{
		CourseName:          req.CourseName,
		Score:               float32(req.Score),
		MediumOfInstruction: req.MediumOfInstruction,
		CourseDescription:   req.CourseDescription,
	})
	if res.Code != 200 {
		return &types.SynchronousResponse{
			Code: "400",
			Msg:  "数据同步成功",
			Data: nil,
		}, nil
	}
	return &types.SynchronousResponse{
		Code: "200",
		Msg:  "数据同步成功",
	}, nil
}
