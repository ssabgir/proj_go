package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountBalanceRequest struct {
	Name       string `json:"name"`
	New_Amount int    `json:"new_amount"`
}

type ChangeAccountNameRequest struct {
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	New_Name string `json:"new_name"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
