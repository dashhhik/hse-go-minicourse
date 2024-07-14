package main

import (
	"context"
	"fmt"
	"fourth-hwk/errs"
	pb "fourth-hwk/proto"
	"fourth-hwk/src/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	querier *db.Queries
	pb.UnimplementedBankAccountServiceServer
}

func (s *server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	name := req.GetName()
	if name == "" {
		return nil, errs.ErrInvalidRequest
	}

	balance := pgtype.Numeric{}
	err := balance.Scan("0.0")
	if err != nil {
		log.Println("Error scanning balance:", err)
		return nil, fmt.Errorf("failed to scan balance: %w", err)
	}

	err = s.querier.CreateBankAccount(ctx, db.CreateBankAccountParams{Username: name, Balance: balance})
	if err != nil {
		log.Println("Error creating account:", err)
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	return &pb.CreateResponse{Name: name, Balance: 0}, nil
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	name := req.GetName()
	account, err := s.querier.GetBankAccount(ctx, name)
	if err != nil {
		log.Println("Error getting account:", err)
		return nil, errs.ErrAccountNotFound
	}

	balance, err := account.Balance.Float64Value()
	if err != nil {
		log.Println("Error converting balance to float64:", err)
		return nil, fmt.Errorf("failed to convert balance: %w", err)
	}

	return &pb.GetResponse{Name: account.Username, Balance: balance.Float64}, nil
}

func (s *server) UpdateBalance(ctx context.Context, req *pb.UpdateBalanceRequest) (*pb.UpdateBalanceResponse, error) {
	name := req.GetName()
	balance := req.GetBalance()

	numericBalance := pgtype.Numeric{}
	err := numericBalance.Scan(fmt.Sprintf("%f", balance))
	if err != nil {
		log.Println("Error scanning balance:", err)
		return nil, fmt.Errorf("failed to scan balance: %w", err)
	}

	err = s.querier.UpdateBankAccountBalance(ctx, db.UpdateBankAccountBalanceParams{Username: name, Balance: numericBalance})
	if err != nil {
		log.Println("Error updating balance:", err)
		return nil, errs.ErrAccountNotFound
	}

	return &pb.UpdateBalanceResponse{Name: name, Balance: balance}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	name := req.GetName()
	err := s.querier.DeleteBankAccount(ctx, name)
	if err != nil {
		log.Println("Error deleting account:", err)
		return nil, errs.ErrAccountNotFound
	}

	return &pb.DeleteResponse{Name: name}, nil
}

func (s *server) UpdateName(ctx context.Context, req *pb.UpdateNameRequest) (*pb.UpdateNameResponse, error) {
	name := req.GetName()
	newName := req.GetNewName()
	err := s.querier.UpdateBankAccountName(ctx, db.UpdateBankAccountNameParams{Username: name, Username_2: newName})
	if err != nil {
		log.Println("Error updating account name:", err)
		return nil, fmt.Errorf("failed to update account name: %w", err)
	}
	return &pb.UpdateNameResponse{Name: name, NewName: newName}, nil
}

func main() {
	// db
	databaseUrl := "postgresql://postgres:postgres@localhost:5432/bank"
	dbpool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	// grpc
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBankAccountServiceServer(s, &server{
		querier: db.New(dbpool),
	})
	log.Printf("Server listening on port %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
