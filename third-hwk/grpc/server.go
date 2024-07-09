package main

import (
	"HSEGoCourse/third-hwk/grpc/accounts/db"
	"HSEGoCourse/third-hwk/grpc/errs"
	pb "HSEGoCourse/third-hwk/grpc/protobufs"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	storage *db.AccountStorage
	pb.UnimplementedBankAccountServiceServer
}

func (s *server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	name := req.GetName()
	if name == "" {
		return nil, errs.ErrInvalidRequest
	}

	s.storage.CreateAccount(&db.Account{Name: name, Balance: 0})

	return &pb.CreateResponse{Name: name, Balance: 0}, nil

}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	name := req.GetName()
	account, err := s.storage.GetAccount(name)
	if err != nil {
		return nil, errs.ErrAccountNotFound
	}

	return &pb.GetResponse{Name: account.Name, Balance: account.Balance}, nil
}

func (s *server) UpdateBalance(ctx context.Context, req *pb.UpdateBalanceRequest) (*pb.UpdateBalanceResponse, error) {
	name := req.GetName()
	balance := req.GetBalance()
	account, err := s.storage.GetAccount(name)
	if err != nil {
		return nil, errs.ErrAccountNotFound
	}

	account.Balance = balance

	return &pb.UpdateBalanceResponse{Name: account.Name, Balance: account.Balance}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	name := req.GetName()
	err := s.storage.DeleteAccount(name)
	if err != nil {
		return nil, errs.ErrAccountNotFound
	}

	return &pb.DeleteResponse{Name: name}, nil
}

func (s *server) UpdateName(ctx context.Context, req *pb.UpdateNameRequest) (*pb.UpdateNameResponse, error) {
	name := req.GetName()
	newName := req.GetNewName()
	err := s.storage.UpdateAccountName(name, newName)
	if err != nil {
		return nil, errs.ErrAccountNotFound
	}

	return &pb.UpdateNameResponse{Name: newName}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBankAccountServiceServer(s, &server{
		storage: db.NewAccountStorage(),
	})
	log.Printf("Server listening on port %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
