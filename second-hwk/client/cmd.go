package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "http://localhost:8080/account"

var httpClient = &http.Client{}

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(CreateAccountCmd, GetAccountCmd, UpdateAmountCmd, GetAllAccountsCmd, DeleteAccountCmd, ChangeAccountNameCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
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
		data := map[string]string{"name": name}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshalling data: %v\n", err)
			return
		}
		resp, err := httpClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v\n", err)
			return
		}
		log.Printf("Response: %s\n", body)
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
		resp, err := httpClient.Get(url + "/" + name)
		if err != nil {
			log.Printf("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v\n", err)
			return
		}
		log.Printf("Response: %s\n", body)
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
		data := map[string]interface{}{"balance": balance}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshalling data: %v\n", err)
			return
		}
		req, err := http.NewRequest(http.MethodPatch, url+"/"+name, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Error making request: %v\n", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v\n", err)
			return
		}
		log.Printf("Response: %s\n", body)
	},
}

var GetAllAccountsCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all accounts",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := httpClient.Get(url)
		if err != nil {
			log.Printf("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v\n", err)
			return
		}
		log.Printf("Response: %s\n", body)
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
		req, err := http.NewRequest(http.MethodDelete, url+"/"+name, nil)
		if err != nil {
			log.Printf("Error creating request: %v\n", err)
			return
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Printf("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v\n", err)
			return
		}
		log.Printf("Response: %s\n", body)
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
		data := map[string]string{"new_name": newName}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshalling data: %v\n", err)
			return
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", url, oldName), bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Error making request: %v\n", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v\n", err)
			return
		}
		log.Printf("Response: %s\n", body)
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
