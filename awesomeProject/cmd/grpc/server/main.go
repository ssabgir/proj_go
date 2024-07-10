package main

import (
	"awesomeProject/accounts"
	"awesomeProject/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedPersonServer
	acc *accounts.Handler
}

func (s *server) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetPerson, error) {
	name := req.GetName()
	if name == "" {
		return nil, fmt.Errorf("invalid request")
	}
	person, err := s.acc.GetAccount_acc_to_params(name)
	if err != nil {
		return nil, fmt.Errorf("account doesn't exsist")
	}
	return &proto.GetPerson{
		Name:   person.Name,
		Amount: person.Amount,
	}, nil
}

func (s *server) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreatePerson, error) {
	name := req.GetName()
	amount := req.GetAmount()
	if name == "" {
		return nil, fmt.Errorf("invalid request")
	}
	s.acc.CreateAccount_acc_to_params(name, amount)
	return &proto.CreatePerson{Ok: true}, nil

}

func (s *server) ChangeName(ctx context.Context, req *proto.ChangeNameRequest) (*proto.ChangePersonName, error) {
	name := req.GetName()
	if name == "" {
		return nil, fmt.Errorf("invalid request")
	}
	new_name := req.GetNewName()
	err := s.acc.ChangeAccountName_acc_to_params(name, new_name)
	if err != nil {
		return nil, fmt.Errorf("invalid params")
	}
	return &proto.ChangePersonName{Ok: true}, nil
}

func (s *server) ChangeAmount(ctx context.Context, req *proto.ChangeAmountRequest) (*proto.ChangePersonAmount, error) {
	name := req.GetName()
	if name == "" {
		return nil, fmt.Errorf("invalid request")
	}
	new_amount := req.GetNewAmount()
	err := s.acc.ChangeAccountBalance_acc_to_params(name, new_amount)
	if err != nil {
		return nil, fmt.Errorf("invalid params")
	}
	return &proto.ChangePersonAmount{Ok: true}, nil
}

func (s *server) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeletePerson, error) {
	name := req.GetName()
	if name == "" {
		return nil, fmt.Errorf("invalid request")
	}
	err := s.acc.DeleteAccount_acc_to_params(name)
	if err != nil {
		return nil, fmt.Errorf("invalid params")
	}
	return &proto.DeletePerson{Ok: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterPersonServer(s, &server{
		acc: accounts.New(),
	})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
