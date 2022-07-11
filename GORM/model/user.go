package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName     string    `gorm:"type:varchar(20);unique_index;not null;column:username"`
	Password     string    `gorm:"type:varchar(20);not null;column:password"`
	Name         string    `gorm:"type:varchar(6);column:name"`
	DateOfBirth  time.Time `gorm:"column:birthday"`
	Department   Department
	DepartmentId int    `gorm:"type:int;column:department_id"`
	Tel          string `gorm:"type:varchar(11);column:tel"`
}

//TableName 设置User的表名为`user`
func (User) TableName() string {
	return "user"
}
