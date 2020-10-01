package resp

type User struct {
	Id        string `json:"id"`
	Key       string `json:"key"`
	UserId    string `json:"userId" gorm:"column:user_id"`
	NickName  string `json:"nickName" gorm:"column:nick_name"`
	Mobile    string `json:"mobile" gorm:"column:mobile" binding:"required"`
	Address   string `json:"address" gorm:"column:address"`
	IsDeleted bool   `json:"isDeleted" gorm:"column:is_deleted"`
	IsLocked  bool   `json:"isLocked" gorm:"column:is_locked"`
}
