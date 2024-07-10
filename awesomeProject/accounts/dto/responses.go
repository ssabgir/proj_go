package dto

type GetAccountResponse struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}

type CreateAccountResponse struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}

type ChangeAccountBalanceResponse struct {
	Name       string `json:"name"`
	New_Amount int64  `json:"new_amount"`
}

type ChangeAccountNameResponse struct {
	Name     string `json:"name"`
	New_Name string `json:"new_name"`
}

type DeleteAccountResponse struct {
	Name string `json:"name"`
}
