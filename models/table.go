package models

import "gorm.io/gorm"

type TeachingInfoTable struct {
	gorm.Model
	CourseName          string  //课程名称
	Score               float64 //评分
	MediumOfInstruction string  //授课语言
	CourseDescription   string  //课程描述
}

type UploadTable struct {
	gorm.Model
	FileName    string
	FileUrl     string
	FileType    string
	FileVersion string
	FileStatus  int64
}
