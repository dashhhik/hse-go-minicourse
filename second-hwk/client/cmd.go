package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

const url = "http://localhost:8080/account"

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(CreateAccountCmd)
	rootCmd.AddCommand(GetAccountCmd)
	rootCmd.AddCommand(UpdateAmountCmd)
	rootCmd.AddCommand(GetAllAccountsCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

var CreateAccountCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new account",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Account name is required")
			return
		}
		data := map[string]string{"name": name}
		jsonData, _ := json.Marshal(data)
		resp, err := http.Post(url+"/create", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body: %v\n", err)
			return
		}

		fmt.Printf("Response: %s\n", body)
	},
}

var GetAccountCmd = &cobra.Command{
	Use:   "get",
	Short: "Get account by name",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Account name is required")
			return
		}
		resp, err := http.Get(url + "/" + name)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body: %v\n", err)
			return
		}

		fmt.Printf("Response: %s\n", body)
	},
}

var UpdateAmountCmd = &cobra.Command{
	Use:   "update",
	Short: "Update account amount",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		amount, _ := cmd.Flags().GetFloat64("amount")
		if name == "" {
			fmt.Println("Account name is required")
			return
		}
		if amount < 0 {
			fmt.Println("Amount must be positive")
			return
		}
		data := map[string]interface{}{"name": name, "amount": amount}
		jsonData, _ := json.Marshal(data)
		resp, err := http.Post(url+"/update", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body: %v\n", err)
			return
		}

		fmt.Printf("Response: %s\n", body)
	},
}

var GetAllAccountsCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all accounts",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body: %v\n", err)
			return
		}

		fmt.Printf("Response: %s\n", body)
	},
}

func init() {
	CreateAccountCmd.Flags().String("name", "", "Name of the account")
	GetAccountCmd.Flags().String("name", "", "Name of the account")
	UpdateAmountCmd.Flags().String("name", "", "Name of the account")
	UpdateAmountCmd.Flags().Float64("amount", 0, "Amount to update the account with")
}
