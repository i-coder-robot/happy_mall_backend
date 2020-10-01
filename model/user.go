package model

import "time"

type User struct {
	UserId    string    `json:"userId" gorm:"column:user_id"`
	NickName  string    `json:"nickName" gorm:"column:nick_name"`
	Mobile    string    `json:"mobile" gorm:"column:mobile" binding:"required"`
	Password  string    `json:"password" gorm:"column:password"`
	Address   string    `json:"address" gorm:"column:address"`
	IsDeleted bool      `json:"isDeleted" gorm:"column:is_deleted"`
	IsLocked  bool      `json:"isLocked" gorm:"column:is_locked"`
	CreateAt  time.Time `json:"createAt" gorm:"column:create_at;default:null"`
	UpdateAt  time.Time `json:"updateAt" gorm:"column:update_at;default:null"`
}
