package model

type Users struct {
	UserId    uint64 `gorm:"column:userId" json:"userId" form:"userId"` // 自增
	Name      string `gorm:"column:name" json:"name"`                   //
	Password  string `gorm:"column:password" json:"password"`
	Email     string `gorm:"column:email" json:"email"`
	Phone     string `gorm:"column:phone" json:"phone"`
	Gender    string `gorm:"gender" json:"gender"`
	Age       int    `gorm:"age" json:"age"`
	Introduce string `gorm:"introduce" json:"introduce"`
	Hobby     string `gorm:"hobby" json:"hobby"`
}

func (Users) TableName() string {
	return "users"
}
