package main

import (
	"awesomeProject/accounts/dto"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Port       int
	Host       string
	Cmd        string
	Name       string
	Amount     int64
	New_Amount int64
	New_Name   string
}

func main() {
	portVal := flag.Int("port", 1323, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newnameVal := flag.String("new_name", "", "new name of account")
	newamountVal := flag.Int("new_amount", 0, "new amount of account")

	flag.Parse()

	cmd := Command{
		Port:       *portVal,
		Host:       *hostVal,
		Cmd:        *cmdVal,
		Name:       *nameVal,
		Amount:     int64(*amountVal),
		New_Name:   *newnameVal,
		New_Amount: int64(*newamountVal),
	}

	for {
		fmt.Println("Введите комманду (get/create/change/delete) или break для выхода:")
		_, err := fmt.Scanln(&cmd.Cmd)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
		}

		switch cmd.Cmd {
		case "create":
			fmt.Print("Введите параметры. Name: ")
			fmt.Scanln(&cmd.Name)
			fmt.Print("Amount: ")
			fmt.Scanln(&cmd.Amount)
		case "get":
			fmt.Print("Введите параметры. Name: ")
			fmt.Scanln(&cmd.Name)
		case "change":
			var s string
			fmt.Println("Введите, что хотите изменить(name, amount): ")
			fmt.Scanln(&s)
			if s == "name" {
				cmd.Cmd = "change/name"
				fmt.Print("Введите параметры. Name: ")
				fmt.Scanln(&cmd.Name)
				fmt.Println("New name: ")
				fmt.Scanln(&cmd.New_Name)
			} else if s == "amount" {
				cmd.Cmd = "change/amount"
				fmt.Print("Введите параметры. Name: ")
				fmt.Scanln(&cmd.Name)
				fmt.Println("New amount: ")
				fmt.Scanln(&cmd.New_Amount)
			} else {
				cmd.Cmd += "/" + s
			}
		case "delete":
			fmt.Print("Введите параметры. Name: ")
			fmt.Scanln(&cmd.Name)
		case "break":
			return
		default:
			fmt.Printf("unknown command %s\n", cmd.Cmd)
			continue
		}

		err = cmd.Do()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (cmd *Command) Do() error {
	switch cmd.Cmd {
	case "create":
		if err := cmd.create(); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := cmd.get(); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "change/name":
		if err := cmd.change_name(); err != nil {
			return fmt.Errorf("change name failed: %w", err)
		}

		return nil
	case "change/amount":
		if err := cmd.change_balance(); err != nil {
			return fmt.Errorf("change amount failed: %w", err)
		}

		return nil
	case "delete":
		if err := cmd.delete(); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func (cmd *Command) get() error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		fmt.Println("body", string(body))
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d\n", response.Name, response.Amount)

	return nil
}

func (cmd *Command) create() error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (cmd *Command) delete() error {
	request := dto.DeleteAccountResponse{
		Name: cmd.Name,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	client := http.Client{}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s:%d/account/delete", cmd.Host, cmd.Port), bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("creating req failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (cmd *Command) change_balance() error {
	request := dto.ChangeAccountBalanceRequest{
		Name:       cmd.Name,
		New_Amount: cmd.New_Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("http://%s:%d/account/change/amount", cmd.Host, cmd.Port), bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("creating req failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (cmd *Command) change_name() error {
	request := dto.ChangeAccountNameRequest{
		Name:     cmd.Name,
		New_Name: cmd.New_Name,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("http://%s:%d/account/change/name", cmd.Host, cmd.Port), bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("creating req failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}
