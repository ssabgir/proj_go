package main

import (
	"awesomeProject/proto"
	"context"
	"flag"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Command struct {
	Cmd        string
	Name       string
	Amount     int64
	New_Amount int64
	New_Name   string
}

func main() {

	name := flag.String("name", "", "name of account")
	amount := flag.Int64("amount", 0, "amount of account")
	newname := flag.String("new_name", "", "new name of account")
	newamount := flag.Int64("new_amount", 0, "new amount of account")
	cmdV := flag.String("cmd", "", "command to execute")

	flag.Parse()

	conn, err := grpc.NewClient("0.0.0.0:4567", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewPersonClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	cmd := Command{
		Cmd:        *cmdV,
		Name:       *name,
		Amount:     *amount,
		New_Name:   *newname,
		New_Amount: *newamount,
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

		err = cmd.Do(ctx, c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (cmd *Command) Do(ctx context.Context, c proto.PersonClient) error {
	switch cmd.Cmd {
	case "create":
		_, err := c.Create(ctx, &proto.CreateRequest{Name: cmd.Name, Amount: cmd.Amount})
		if err != nil {
			return fmt.Errorf("panic :%w", err)
		}
		fmt.Println("acc was created")
		return nil
	case "get":
		res, err := c.Get(ctx, &proto.GetRequest{Name: cmd.Name})
		if err != nil {
			return fmt.Errorf("panic :%w", err)
		}
		fmt.Printf("account name: %s and amount: %d\n", res.Name, res.Amount)
		return nil
	case "change/name":
		_, err := c.ChangeName(ctx, &proto.ChangeNameRequest{Name: cmd.Name, NewName: cmd.New_Name})
		if err != nil {
			return fmt.Errorf("panic :%w", err)
		}
		fmt.Printf("name was changed\n")
		return nil
	case "change/amount":
		_, err := c.ChangeAmount(ctx, &proto.ChangeAmountRequest{Name: cmd.Name, NewAmount: cmd.New_Amount})
		if err != nil {
			return fmt.Errorf("panic :%w", err)
		}
		fmt.Printf("amount was changed\n")
		return nil
	case "delete":
		_, err := c.Delete(ctx, &proto.DeleteRequest{Name: cmd.Name})
		if err != nil {
			return fmt.Errorf("panic :%w", err)
		}
		fmt.Printf("acc was deleted\n")
		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}
