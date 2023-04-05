package models

type User struct {
	ID       int64  `gorm:"id" json:"id"`
	Name     string `gorm:"name" json:"name"`
	Surname  string `gorm:"surname" json:"surname"`
	Email    string `gorm:"email" json:"email"`
	Password string `gorm:"password" json:"password"`
}

func (u User) TableName() string {
	return "users"
}
