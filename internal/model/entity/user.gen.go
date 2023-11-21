// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID           int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
	Name         string         `gorm:"column:name;type:varchar(255);not null;comment:医生姓名" json:"name"`                      // 医生姓名
	Sex          string         `gorm:"column:sex;type:enum('F','M','O');not null;default:O;comment:性别 1男 2女 3未知" json:"sex"` // 性别 1男 2女 3未知
	Mobile       string         `gorm:"column:mobile;type:varchar(32);not null;comment:手机号" json:"mobile"`                    // 手机号
	Photo        string         `gorm:"column:photo;type:varchar(255);not null;default:照片" json:"photo"`
	Description  string         `gorm:"column:description;type:varchar(255);not null;comment:医生简介" json:"description"` // 医生简介
	Title        int32          `gorm:"column:title;type:tinyint unsigned;not null;comment:职称" json:"title"`           // 职称
	State        int32          `gorm:"column:state;type:tinyint unsigned;not null;default:1" json:"state"`
	Password     string         `gorm:"column:password;type:varchar(255);not null" json:"password"`
	LoginName    string         `gorm:"column:login_name;type:varchar(255);not null" json:"login_name"`
	Role         int32          `gorm:"column:role;type:int unsigned;not null;comment:角色" json:"role"`                      // 角色
	DepartmentID int32          `gorm:"column:department_id;type:int unsigned;not null;comment:部门id" json:"department_id"`  // 部门id
	StaffGroupID int32          `gorm:"column:staff_group_id;type:int unsigned;not null;comment:组id" json:"staff_group_id"` // 组id
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
