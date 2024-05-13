package models

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
)

// 数据同步方法
func Synchronous(data *TeachingInfoTable) error {
	tx := Db.Begin()
	err = tx.Create(data).Error
	if err != nil {
		fmt.Println("数据库添加失败", err)
		tx.Rollback()
		return nil
	}
	_, err = EsCreate("week3", strconv.Itoa(int(data.ID)), data)
	if err != nil {
		fmt.Println("Es添加失败", err)
		tx.Rollback()
		return nil
	}
	tx.Commit()
	return nil
}

// es添加
func EsCreate(index, id string, data interface{}) (*elastic.IndexResponse, error) {
	return Client.Index().Index(index).Id(id).BodyJson(data).Do(context.Background())
}

// 关键词搜索方法
func SearchHighlight(index, keyword string) (*elastic.SearchResult, error) {
	return Client.Search().Index(index).Query(elastic.NewMatchQuery("CourseName", keyword)).
		// 设置高亮
		Highlight(
			elastic.NewHighlight().
				Field("CourseName").
				PreTags("<span color = `red`>").
				PostTags("</span>"),
		).Do(context.Background())
}

// 多条件搜索
func ManySearch(index, keyword, keywords string) (*elastic.SearchResult, error) {
	return Client.Search().Index(index).Query(elastic.NewBoolQuery().
		Must(elastic.NewMatchQuery("CourseName", keyword)).
		Must(elastic.NewMatchQuery("CourseDescription", keywords)),
	).Do(context.Background())
}

// 分页搜索方法
func PaginationSearchHighlight(index, keyword string, page, pageSize int) (*elastic.SearchResult, error) {
	return Client.Search().Index(index).Query(elastic.NewMatchQuery("CourseName", keyword)).Highlight(
		elastic.NewHighlight().
			Field("CourseName").
			PreTags("<span color = `red`>").
			PostTags("</span>"),
	).From((page - 1) * pageSize).
		Size(pageSize).Do(context.Background())
}

// 上传方法
func CreateUpload(data *UploadTable) error {
	return Db.Create(data).Error
}

// 回滚方法
func Rollback(id int64, url, version string, status int64) error {
	return Db.Model(&UploadTable{}).Where("id", id).Updates(UploadTable{FileUrl: url, FileVersion: version, FileStatus: status}).Error
}
