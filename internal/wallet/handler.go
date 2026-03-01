package wallet

import (
	"net/http"
	"strconv"

	"github.com/archit-batra/fintech-wallet-backend/internal/events"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

type addMoneyRequest struct {
	Amount int64 `json:"amount"`
}

func (h *Handler) CreateWallet(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Param("userId"))

	err := h.service.CreateWallet(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "wallet created"})
}

func (h *Handler) AddMoney(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Param("userId"))

	var req addMoneyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.AddMoney(userID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "money added"})
}

func (h *Handler) GetWallet(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Param("userId"))

	wallet, err := h.service.GetWallet(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
		return
	}

	c.JSON(http.StatusOK, wallet)
}

type transferRequest struct {
	FromUser int   `json:"from_user"`
	ToUser   int   `json:"to_user"`
	Amount   int64 `json:"amount"`
}

func (h *Handler) Transfer(c *gin.Context) {

	var req transferRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Transfer(req.FromUser, req.ToUser, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	events.EventQueue <- events.Event{
		Type: "transfer_completed",
		Data: "transfer executed",
	}

	c.JSON(http.StatusOK, gin.H{"status": "transfer successful"})
}
