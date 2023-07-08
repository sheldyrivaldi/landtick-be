package transactiondto

type CreateTransactionRequestDTO struct {
	UserID   int    `json:"user_id" form:"user_id"`
	TicketID int    `json:"ticket_id" form:"ticket_id"`
	Image    string `json:"image" form:"image"`
}

type UpdateTransactionRequestDTO struct {
	UserID   int    `json:"user_id" form:"user_id"`
	TicketID int    `json:"ticket_id" form:"ticket_id"`
	Image    string `json:"image" form:"image"`
}
