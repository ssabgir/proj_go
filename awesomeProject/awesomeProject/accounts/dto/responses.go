package dto

type GetAccountResponse struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type CreateAccountResponse struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountBalanceResponse struct {
	Name       string `json:"name"`
	New_Amount int    `json:"new_amount"`
}

type ChangeAccountNameResponse struct {
	Name     string `json:"name"`
	New_Name string `json:"new_name"`
}

type DeleteAccountResponse struct {
	Name string `json:"name"`
}
