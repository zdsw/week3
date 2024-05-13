package logic

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"zg4-1/src/rk/week3/common"
	"zg4-1/src/rk/week3/week3rpc/week3rpcclient"

	"zg4-1/src/rk/week3/week3api/internal/svc"
	"zg4-1/src/rk/week3/week3api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadRequest, w http.ResponseWriter, r *http.Request) (resp *types.UploadResponse, err error) {
	// todo: 上传
	file, head, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("文件获取失败"))
		return
	}
	defer file.Close()

	fileExt := filepath.Ext(head.Filename)
	filetype := []string{".jpg", ".png", ".jpeg", ".gif"}
	isFile := false
	for _, val := range filetype {
		if val == fileExt {
			isFile = true
		}
	}

	if !isFile {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("文件格式错误"))
	}

	key := uuid.New().String() + head.Filename
	qn, _ := common.Qn(file, fileExt)

	res, _ := l.svcCtx.Week3Server.Upload(l.ctx, &week3rpcclient.UploadRequest{
		FileName:    qn,
		FileUrl:     key,
		FileType:    fileExt,
		FileVersion: req.FileVersion,
		FileStatus:  req.FileStatus,
	})
	if err != nil {
		return &types.UploadResponse{
			Code: "400",
			Msg:  "文件入库失败",
			Data: nil,
		}, nil
	}
	return &types.UploadResponse{
		Code: "200",
		Msg:  "文件入库成功",
		Data: map[string]interface{}{
			"fileName":    key,
			"url":         head.Filename,
			"filename":    res.FileId,
			"fileType":    fileExt,
			"fileVersion": req.FileVersion,
			"fileStatus":  req.FileStatus,
		},
	}, nil
}
