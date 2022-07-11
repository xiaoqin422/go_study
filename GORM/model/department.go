package model

type Department struct {
	ID                int    `gorm:"primary_key;AUTO_INCREMENT"`
	DepartmentName    string `gorm:"type:varchar(12);column:department_name"`
	DepartmentAddress string `gorm:"type:varchar(24);column:department_address"`
	Users             []User
}
