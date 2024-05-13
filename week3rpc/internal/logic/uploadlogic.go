package logic

import (
	"context"
	"zg4-1/src/rk/week3/models"

	"zg4-1/src/rk/week3/week3rpc/internal/svc"
	"zg4-1/src/rk/week3/week3rpc/week3rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传
func (l *UploadLogic) Upload(in *week3rpc.UploadRequest) (*week3rpc.UploadResponse, error) {
	// todo: 上传
	data := models.UploadTable{
		FileName:    in.FileUrl,
		FileUrl:     in.FileName,
		FileType:    in.FileType,
		FileVersion: in.FileVersion,
		FileStatus:  in.FileStatus,
	}
	models.CreateUpload(&data)
	return &week3rpc.UploadResponse{
		FileId: int64(data.ID),
	}, nil
}
