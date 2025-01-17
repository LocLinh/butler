package models

import "time"

type User struct {
	Id         int64     `gorm:"primaryKey;column:id"`
	Name       string    `gorm:"column:name;default:(-)" jenk:"filter:many"`
	Email      string    `gorm:"column:email;default:(-)" jenk:"filter:many"`
	RoleId     int64     `gorm:"column:role_id"`
	Status     int64     `gorm:"column:status"`
	Config     int64     `gorm:"column:config"`
	ConfigZone int64     `gorm:"column:config_zone"`
	CreatedAt  time.Time `gorm:"autoCreateTime" jenk:"filter:-;range:from-to;type:datetime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	AuthUserId int64     `gorm:"column:auth_user_id"`
}

func (e *User) TableName() string {
	return "wms_user"
}

type GetRequest struct {
	UserId int64
	Email  string
}
