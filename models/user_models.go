package models

type User struct {
	ID       int    `json:"id" form:"id" gorm:"primaryKey"`
	Fullname string `json:"fullname" form:"fullname" gorm:"varchar(255)"`
	Username string `json:"username" form:"username" gorm:"varchar(255)"`
	Email    string `json:"email" form:"email" gorm:"varchar(255)"`
	Password string `json:"password" form:"password" gorm:"varchar(255)"`
	Gender   string `json:"gender" form:"gender" gorm:"varchar(255)"`
	Telp     string `json:"no_hp" form:"no_hp" gorm:"varchar(255)"`
	Address  string `json:"address" form:"address"`
}

type UserResponseModels struct {
	ID       int    `json:"id" form:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
}

type UserResponseModelsTicket struct {
	ID       int    `json:"id" form:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Telp     string `json:"no_hp" form:"no_hp"`
	Email    string `json:"email" form:"email"`
}

type UserResponseModelsTransaction struct {
	ID       int    `json:"id" form:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Telp     string `json:"no_hp" form:"no_hp"`
	Email    string `json:"email" form:"email"`
}

func (UserResponseModels) TableName() string {
	return "users"
}

func (UserResponseModelsTicket) TableName() string {
	return "users"
}
func (UserResponseModelsTransaction) TableName() string {
	return "users"
}
