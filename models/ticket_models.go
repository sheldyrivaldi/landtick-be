package models

type Ticket struct {
	ID        int                `json:"id" form:"id" gorm:"primaryKey"`
	NameTrain string             `json:"name_train" form:"name_train" gorm:"varchar(255)"`
	TypeTrain string             `json:"type_train" form:"type_train" gorm:"varchar(255)"`
	UserId    int                `json:"user_id"`
	User      UserResponseModels `json:"user"`
}

type UserResponseTicket struct {
	ID        int                `json:"id"`
	NameTrain string             `json:"name_train"`
	TypeTrain string             `json:"type_train"`
	UserId    int                `json:"user_id"`
	User      UserResponseModels `json:"user"`
}

func (UserResponseTicket) TableName() string {
	return "tickets"
}
