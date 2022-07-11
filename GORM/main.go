package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gogorm/model"
	"log"
	"time"
)

func main() {
	db, err := gorm.Open("mysql", "root:123123@(localhost)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("GORM open mysql connection faild. err: %v", err)
	}
	// 全局禁用表名复数
	db.SingularTable(true)
	defer db.Close()
	// 1.根据结构体创建表
	//db.AutoMigrate(&model.User{}, &model.Department{})

	department := &model.Department{
		DepartmentName:    "RD",
		DepartmentAddress: "四川省 成都市",
	}
	user := &model.User{
		UserName:    "admin",
		Password:    "123456",
		Name:        "秦笑笑",
		DateOfBirth: time.Now(),
		Department:  *department,
		Tel:         "15735111403",
	}
	//if db.HasTable(department) {
	//	db.Create(department)
	//}
	//if db.HasTable(user) {
	//	db.Create(user)
	//}
	db.First(&user, 10)
	fmt.Printf("主键查询，结果： %v\n", user)
	db.Last(&department)
	fmt.Printf("主键顺序查询，结果： %v\n", department)
	var users []model.User
	db.Find(&users)
	fmt.Printf("查询所有记录，记录条数：%v\n", len(users))

	// 获取一个匹配记录
	db.Where("username = ?", "admin").Find(&user)
	fmt.Printf("条件查询，结果： %v\n", user)
	// 获取多条匹配记录
	db.Where("department_id = ?", "3").Find(&users)
	fmt.Printf("条件查询，结果： %v\n", len(users))
	db.Update()
}
