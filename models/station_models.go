package models

type Station struct {
	ID   int    `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name" gorm:"varchar(255)"`
}

type StationResponseModels struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func (StationResponseModels) TableName() string {
	return "stations"
}
