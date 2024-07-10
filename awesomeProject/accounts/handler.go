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

func (h *Handler) CreateAccount_acc_to_params(name string, amount int64) error {
	if name == "" {
		return fmt.Errorf("empty name\n")
	}

	h.guard.Lock()

	if _, ok := h.accounts[name]; ok {
		h.guard.Unlock()
		return fmt.Errorf("account already exists\n")
	}

	h.accounts[name] = &models.Account{
		Name:   name,
		Amount: amount,
	}

	h.guard.Unlock()

	return nil
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

func (h *Handler) GetAccount_acc_to_params(name string) (*models.Account, error) {

	if len(name) == 0 {
		return nil, fmt.Errorf("empty name\n")
	}
	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return nil, fmt.Errorf("account not found\n")
	}

	return account, nil
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

func (h *Handler) DeleteAccount_acc_to_params(name string) error {

	if len(name) == 0 {
		return fmt.Errorf("empty name\n")
	}

	h.guard.Lock()
	if _, ok := h.accounts[name]; !ok {
		h.guard.Unlock()

		return fmt.Errorf("account doesn't exists\n")
	}

	delete(h.accounts, name)
	h.guard.Unlock()

	return nil
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

func (h *Handler) ChangeAccountBalance_acc_to_params(name string, new_amount int64) error {

	if len(name) == 0 {
		return fmt.Errorf("empty name\n")
	}

	h.guard.Lock()

	if _, ok := h.accounts[name]; !ok {
		h.guard.Unlock()

		return fmt.Errorf("account doesn't exists\n")
	}

	h.accounts[name].Amount = new_amount

	h.guard.Unlock()

	return nil
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

func (h *Handler) ChangeAccountName_acc_to_params(name string, newName string) error {

	h.guard.Lock()

	if len(name) == 0 {
		return fmt.Errorf("empty old name\n")
	}
	if len(newName) == 0 {
		return fmt.Errorf("empty new name\n")
	}

	if _, ok := h.accounts[name]; !ok {
		h.guard.Unlock()

		return fmt.Errorf("account doesn't exists\n")
	}

	h.accounts[newName] = &models.Account{
		Name:   newName,
		Amount: h.accounts[name].Amount,
	}

	delete(h.accounts, name)
	h.guard.Unlock()

	return nil
}

// Написать клиент консольный, который делает запросы
