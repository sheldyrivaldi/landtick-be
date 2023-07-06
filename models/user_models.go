package models

type User struct {
	ID       int    `json:"id" form:"id" gorm:"primaryKey"`
	Fullname string `json:"fullname" form:"fullname" gorm:"varchar(255)"`
	Username string `json:"username" form:"username" gorm:"varchar(255)"`
	Email    string `json:"email" form:"email" gorm:"varchar(255)"`
	Password string `json:"password" form:"password" gorm:"varchar(255)"`
}

type UserResponseModels struct {
	ID       int    `json:"id" form:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
}

func (UserResponseModels) TableName() string {
	return "users"
}
