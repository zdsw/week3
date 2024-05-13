package main

import (
	"fmt"
	"math/rand"
	"zg4-1/src/rk/week3/models"
)

func Nums() float64 {
	Max := 9.9
	Min := 0.0
	return float64(rand.Intn(int(float64(Max-Min+1) + Min)))
}

func main() {
	// 创建一个切片来存储搜索记录
	var records []models.TeachingInfoTable
	// 生成1万条搜索记录
	for i := 0; i < 5000; i++ {
		// 随机生成一个查询字符串
		course_name := fmt.Sprintf("课程名称%d", rand.Intn(5000))
		medium := fmt.Sprintf("授课语言%d", rand.Intn(5000))
		course := fmt.Sprintf("课程描述%d", rand.Intn(5000))
		// 随机生成评分
		num := Nums()
		// 创建一个搜索记录
		record := models.TeachingInfoTable{
			CourseName:          course_name,
			Score:               num,
			MediumOfInstruction: medium,
			CourseDescription:   course,
		}

		// 将记录添加到切片中
		records = append(records, record)
	}
	models.Db.Create(records)
}
