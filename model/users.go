package model

type Users struct {
	UserId   uint64 `gorm:"column:userId" json:"Id"` // 自增
	Name     string `gorm:"column:name" json:"name"` //
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
	Phone    string `gorm:"column:phone" json:"phone"`
}

func (Users) TableName() string {
	return "Users"
}
