package handlers

import (
	resultdto "landtick/dto/result"
	transactiondto "landtick/dto/transaction"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:8000/uploads/"

type handlerTransaction struct {
	TransactionRepositories repositories.TransactionRepositories
}

type dataTrsansaction struct {
	Transaction interface{} `json:"transaction"`
}

func HandlerTransaction(TransactionRepositories repositories.TransactionRepositories) *handlerTransaction {
	return &handlerTransaction{TransactionRepositories}
}

func (h *handlerTransaction) FindTransactions(c echo.Context) error {
	transactions, err := h.TransactionRepositories.FindTransactions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	response := make([]transactiondto.TransactionResponseDTO, len(transactions))
	for i, t := range transactions {
		response[i] = convertResponseTransactionFindAll(t)
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: dataTrsansaction{
			Transaction: response,
		},
	})
}
func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepositories.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: dataTrsansaction{
			Transaction: convertResponseTransactionGet(transaction),
		},
	})
}

func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	request := new(transactiondto.CreateTransactionRequestDTO)

	UserID, _ := strconv.Atoi(c.FormValue("user_id"))
	TicketID, _ := strconv.Atoi(c.FormValue("ticket_id"))
	Image := c.Get("dataFile").(string)

	request.UserID = UserID
	request.TicketID = TicketID
	request.Image = Image

	validation := validator.New()

	validationErr := validation.Struct(request)
	if validationErr != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: "Failed", Message: validationErr.Error()})
	}

	newTransaction := models.Transaction{
		UserID:   request.UserID,
		TicketID: request.TicketID,
		Image:    request.Image,
	}

	data, err := h.TransactionRepositories.CreateTransaction(newTransaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: dataTrsansaction{
			Transaction: convertResponseTransactionGet(data),
		},
	})
}

func convertResponseTransactionFindAll(t models.Transaction) transactiondto.TransactionResponseDTO {
	return transactiondto.TransactionResponseDTO{
		ID:     t.ID,
		User:   convertUser(t.User),
		Ticket: convertTicket(t.Ticket),
		Image:  path_file + t.Image,
	}
}
func convertResponseTransactionGet(t models.Transaction) transactiondto.TransactionIdResponseDTO {
	return transactiondto.TransactionIdResponseDTO{
		ID:     t.ID,
		User:   convertUser(t.User),
		Ticket: convertTicket(t.Ticket),
		Image:  path_file + t.Image,
		Status: t.Status,
	}
}

func convertTicket(t models.TicketResponseModels) models.TicketResponseModelsTransaction {
	return models.TicketResponseModelsTransaction{
		ID:                   t.ID,
		NameTrain:            t.NameTrain,
		TypeTrain:            t.TypeTrain,
		StartDate:            t.StartDate.Format("02-01-2006"),
		StartStationID:       t.StartStationID,
		StartStation:         convertStation(t.StartStation),
		StartTime:            t.StartTime.Format("15:04"),
		DestinationStationID: t.DestinationStationID,
		DestinationStation:   convertStation(t.DestinationStation),
		ArrivalTime:          t.ArrivalTime.Format("15:04"),
		Price:                t.Price,
	}
}

func convertStation(s models.StationResponseModels) models.StationResponseModels {
	return models.StationResponseModels{
		ID:   s.ID,
		Name: s.Name,
	}
}

func convertUser(u models.UserResponseModelsTransaction) models.UserResponseModelsTransaction {
	return models.UserResponseModelsTransaction{
		Fullname: u.Fullname,
		Telp:     u.Telp,
		Email:    u.Email,
	}
}
