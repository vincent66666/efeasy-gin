package model

import (
	"efeasy-gin/global/model"
	"strconv"
	"time"
)

type User struct {
	model.ID
	Name     string     `json:"name" gorm:"not null;comment:登录用户名"`
	Mobile   string     `json:"mobile" gorm:"not null;index;comment:用户手机号"`
	Password string     `json:"password" gorm:"not null;default:'';comment:用户密码"`
	NickName string     `json:"nick_name" gorm:"default:'';comment:用户名称"`
	Birthday *time.Time `json:"birthday" gorm:"type:date;comment:生日"`
	Address  string     `json:"address"`
	Desc     string     `json:"desc"`
	Gender   string     `json:"gender"`
	Role     int        `json:"role"`
	model.Timestamps
}

// TableName 设置User的表名为`profiles`
func (User) TableName() string {
	return "users"
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}