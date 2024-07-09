package main

import (
	pb "HSEGoCourse/third-hwk/grpc/protobufs"
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "defaultName"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(CreateAccountCmd, GetAccountCmd, UpdateAmountCmd, GetAllAccountsCmd, DeleteAccountCmd, ChangeAccountNameCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func getClient() (pb.BankAccountServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewBankAccountServiceClient(conn)
	return client, conn
}

var CreateAccountCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new account",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			log.Println("Account name is required")
			return
		}
		client, conn := getClient()
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := &pb.CreateRequest{Name: name}
		res, err := client.Create(ctx, req)
		if err != nil {
			log.Fatalf("could not create account: %v", err)
		}
		log.Printf("Account created: %s, Balance: %f", res.Name, res.Balance)
	},
}

var GetAccountCmd = &cobra.Command{
	Use:   "get",
	Short: "Get account by name",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			log.Println("Account name is required")
			return
		}
		client, conn := getClient()
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := &pb.GetRequest{Name: name}
		res, err := client.Get(ctx, req)
		if err != nil {
			log.Fatalf("could not get account: %v", err)
		}
		log.Printf("Account: %s, Balance: %f", res.Name, res.Balance)
	},
}

var UpdateAmountCmd = &cobra.Command{
	Use:   "update",
	Short: "Update account amount",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		balance, _ := cmd.Flags().GetFloat64("balance")
		if name == "" {
			log.Println("Account name is required")
			return
		}
		if balance < 0 {
			log.Println("Balance must be non-negative")
			return
		}
		client, conn := getClient()
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := &pb.UpdateBalanceRequest{Name: name, Balance: balance}
		res, err := client.UpdateBalance(ctx, req)
		if err != nil {
			log.Fatalf("could not update balance: %v", err)
		}
		log.Printf("Account updated: %s, Balance: %f", res.Name, res.Balance)
	},
}

var GetAllAccountsCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all accounts",
	Run: func(cmd *cobra.Command, args []string) {
		// This command is not defined in the proto service
		log.Println("This service is not implemented in the proto file")
	},
}

var DeleteAccountCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete account",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			log.Println("Account name is required")
			return
		}
		client, conn := getClient()
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := &pb.DeleteRequest{Name: name}
		_, err := client.Delete(ctx, req)
		if err != nil {
			log.Fatalf("could not delete account: %v", err)
		}
		log.Printf("Account deleted: %s", name)
	},
}

var ChangeAccountNameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Change account name",
	Run: func(cmd *cobra.Command, args []string) {
		oldName, _ := cmd.Flags().GetString("name")
		newName, _ := cmd.Flags().GetString("newname")
		if oldName == "" || newName == "" {
			log.Println("Old and new names are required")
			return
		}
		client, conn := getClient()
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := &pb.UpdateNameRequest{Name: oldName, NewName: newName}
		res, err := client.UpdateName(ctx, req)
		if err != nil {
			log.Fatalf("could not update name: %v", err)
		}
		log.Printf("Account name updated: %s -> %s", res.Name, res.NewName)
	},
}

func init() {
	CreateAccountCmd.Flags().String("name", "", "Name of the account")
	GetAccountCmd.Flags().String("name", "", "Name of the account")
	UpdateAmountCmd.Flags().String("name", "", "Name of the account")
	UpdateAmountCmd.Flags().Float64("balance", 0, "Balance to update the account with")
	DeleteAccountCmd.Flags().String("name", "", "Name of the account")
	ChangeAccountNameCmd.Flags().String("name", "", "Old name of the account")
	ChangeAccountNameCmd.Flags().String("newname", "", "New name of the account")
}
