package common

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"mime/multipart"
)

func Qn(file multipart.File, key string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: "wgdzg42",
	}

	accessKey := "IiKGf2R_Z-rbD1cALzeo64IlOyEDvMEU3ZirmeAj"
	secretKey := "PrN7B0ChPrR4FBl1jXhoBjTBieOjw7nnbMZOOzpo"
	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	data, _ := io.ReadAll(file)
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err

	}
	fmt.Println(ret.Key, ret.Hash)
	return "http://scc757swq.hd-bkt.clouddn.com/" + key, nil
}
