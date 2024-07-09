package accounts

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/accounts/models"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

func New() *Handler {
	return &Handler{
		accounts: generateData(),
		guard:    &sync.RWMutex{},
	}
}

func generateData() map[string]*models.Account {
	result := make(map[string]*models.Account)

	result["sonya"] = &models.Account{
		Name:   "sonya",
		Amount: 1000,
	}

	result["max"] = &models.Account{
		Name:   "max",
		Amount: 123,
	}
	result["john"] = &models.Account{
		Name:   "john",
		Amount: 444,
	}

	return result
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request\n")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name\n")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists\n")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	if len(name) == 0 {
		return c.String(http.StatusBadRequest, "empty name\n")
	}
	fmt.Println(h.accounts)
	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found\n")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

// Удаляет аккаунт
func (h *Handler) DeleteAccount(c echo.Context) error {
	var request dto.DeleteAccountRequest
	fmt.Println(h.accounts)
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request\n")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name\n")
	}

	h.guard.Lock()
	fmt.Println(h.accounts, request.Name)
	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account doesn't exists\n")
	}

	delete(h.accounts, request.Name)
	h.guard.Unlock()

	return c.String(http.StatusOK, "account was deleted\n")
}

// Меняет баланс
func (h *Handler) ChangeAccountBalance(c echo.Context) error {
	var request dto.ChangeAccountBalanceRequest

	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request\n")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name\n")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account doesn't exists\n")
	}

	h.accounts[request.Name].Amount = request.New_Amount

	h.guard.Unlock()

	return c.String(http.StatusOK, "balance was changed\n")
}

// Меняет имя
func (h *Handler) ChangeAccountName(c echo.Context) error {
	var request dto.ChangeAccountNameRequest

	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request\n")
	}

	h.guard.Lock()

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name\n")
	}

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account doesn't exists\n")
	}

	h.accounts[request.New_Name] = &models.Account{
		Name:   request.New_Name,
		Amount: h.accounts[request.Name].Amount,
	}

	delete(h.accounts, request.Name)
	h.guard.Unlock()

	return c.String(http.StatusOK, "name was changed\n")
}

// Написать клиент консольный, который делает запросы
