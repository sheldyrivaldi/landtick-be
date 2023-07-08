package transactiondto

import (
	ticketdto "landtick/dto/ticket"
	"landtick/models"
)

type TransactionResponseDTO struct {
	ID     int                                    `json:"id"`
	User   models.UserResponseModelsTransaction   `json:"user"`
	Ticket models.TicketResponseModelsTransaction `json:"ticket"`
	Image  string                                 `json:"image" form:"image"`
}

type TransactionIdResponseDTO struct {
	ID     int                                    `json:"id"`
	User   models.UserResponseModelsTransaction   `json:"user"`
	Ticket models.TicketResponseModelsTransaction `json:"ticket"`
	Image  string                                 `json:"image" form:"image"`
	Status string                                 `json:"status" form:"status"`
}
type TransactionTicketResponse struct {
	ID     int                                  `json:"-"`
	Ticket ticketdto.TicketResponseDTOGet       `json:"ticket"`
	User   models.UserResponseModelsTransaction `json:"user"`
}
