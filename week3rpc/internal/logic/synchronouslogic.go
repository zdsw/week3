package logic

import (
	"context"
	"zg4-1/src/rk/week3/models"

	"zg4-1/src/rk/week3/week3rpc/internal/svc"
	"zg4-1/src/rk/week3/week3rpc/week3rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SynchronousLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSynchronousLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SynchronousLogic {
	return &SynchronousLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// es同步
func (l *SynchronousLogic) Synchronous(in *week3rpc.SynchronousRequest) (*week3rpc.SynchronousResponse, error) {
	// todo: es同步
	data := models.TeachingInfoTable{
		CourseName:          in.CourseName,
		Score:               float64(in.Score),
		MediumOfInstruction: in.MediumOfInstruction,
		CourseDescription:   in.CourseDescription,
	}
	err := models.Synchronous(&data)
	if err != nil {
		return &week3rpc.SynchronousResponse{
			Code: 400,
			Msg:  "数据同步失败",
			Data: nil,
		}, nil
	}
	return &week3rpc.SynchronousResponse{
		Code: 200,
		Msg:  "数据同步失败",
	}, nil
}
