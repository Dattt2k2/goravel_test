package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	ID 	 uint64 `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	Phone string
	Type string 
	Avatar string 
	orm.SoftDeletes

	
}
