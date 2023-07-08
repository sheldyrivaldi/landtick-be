package ticketdto

import (
	"landtick/models"
)

type TicketResponseDTOCreate struct {
	ID                   int    `json:"id"`
	NameTrain            string `json:"name_train" validate:"required"`
	TypeTrain            string `json:"type_train" validate:"required"`
	StartDate            string `json:"start_date" validate:"required"`
	StartStationID       int    `json:"start_station_id"`
	StartTime            string `json:"start_time" validate:"required"`
	DestinationStationID int    `json:"destination_station_id" `
	ArrivalTime          string `json:"arival_time" validate:"required"`
	Price                int    `json:"price" validate:"required"`
	Qty                  int    `json:"qty"`
}

type TicketResponseDTOGet struct {
	ID                   int                          `json:"id"`
	NameTrain            string                       `json:"name_train" validate:"required"`
	TypeTrain            string                       `json:"type_train" validate:"required"`
	StartDate            string                       `json:"start_date" validate:"required"`
	StartStationID       int                          `json:"-"`
	StartStation         models.StationResponseModels `json:"start_station"`
	StartTime            string                       `json:"start_time" validate:"required"`
	DestinationStationID int                          `json:"-" `
	DestinationStation   models.StationResponseModels `json:"destination_station"`
	ArrivalTime          string                       `json:"arival_time" validate:"required"`
	Price                int                          `json:"price" validate:"required"`
}

type TicketMyResponseDTO struct {
	ID                   int                                  `json:"id"`
	NameTrain            string                               `json:"name_train" validate:"required"`
	TypeTrain            string                               `json:"type_train" validate:"required"`
	StartDate            string                               `json:"start_date" validate:"required"`
	StartStationID       int                                  `json:"-"`
	StartStation         models.StationResponseModels         `json:"start_station"`
	StartTime            string                               `json:"start_time" validate:"required"`
	DestinationStationID int                                  `json:"-" `
	DestinationStation   models.StationResponseModels         `json:"destination_station"`
	ArrivalTime          string                               `json:"arival_time" validate:"required"`
	Price                int                                  `json:"price" validate:"required"`
	UserID               int                                  `json:"-"`
	User                 models.UserResponseModelsTransaction `json:"user"`
}
