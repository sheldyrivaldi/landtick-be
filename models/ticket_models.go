package models

import (
	"time"
)

type Ticket struct {
	ID                   int                   `json:"id"`
	NameTrain            string                `json:"name_train"`
	TypeTrain            string                `json:"type_train"`
	StartDate            time.Time             `json:"start_date" gorm:"type:date"`
	StartStationID       int                   `json:"start_station_id"`
	StartStation         StationResponseModels `json:"start_station" gorm:"foreignKey:StartStationID"`
	StartTime            time.Time             `json:"start_time" gorm:"type:time"`
	ArrivalTime          time.Time             `json:"arrival_time" gorm:"type:time"`
	DestinationStationID int                   `json:"destination_station_id" `
	DestinationStation   StationResponseModels `json:"destination_station" gorm:"foreignKey:DestinationStationID"`
	Price                int                   `json:"price"`
	Qty                  int                   `json:"-"`
}

type TicketResponseModels struct {
	ID                   int                   `json:"id"`
	NameTrain            string                `json:"name_train"`
	TypeTrain            string                `json:"type_train"`
	StartDate            time.Time             `json:"start_date" gorm:"type:date"`
	StartStationID       int                   `json:"start_station_id"`
	StartStation         StationResponseModels `json:"start_station"`
	StartTime            time.Time             `json:"start_time" gorm:"type:time"`
	DestinationStationID int                   `json:"destination_station_id"`
	DestinationStation   StationResponseModels `json:"destination_station"`
	ArrivalTime          time.Time             `json:"arrival_time" gorm:"type:time"`
	Price                int                   `json:"price"`
	Qty                  int                   `json:"-"`
}

type TicketResponseModelsTransaction struct {
	ID                   int                   `json:"-"`
	NameTrain            string                `json:"name_train" validate:"required"`
	TypeTrain            string                `json:"type_train" validate:"required"`
	StartDate            string                `json:"start_date" validate:"required"`
	StartStationID       int                   `json:"-"`
	StartStation         StationResponseModels `json:"start_station"`
	StartTime            string                `json:"start_time" validate:"required"`
	DestinationStationID int                   `json:"-" `
	DestinationStation   StationResponseModels `json:"destination_station"`
	ArrivalTime          string                `json:"arival_time" validate:"required"`
	Price                int                   `json:"price" validate:"required"`
}

func (TicketResponseModels) TableName() string {
	return "tickets"
}
func (TicketResponseModelsTransaction) TableName() string {
	return "tickets"
}
