package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/joho/godotenv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func encodePassword(website string, userInput string) {
	API_KEY := os.Getenv("apiKey")
	API_EMAIL := os.Getenv("apiMail")
	ACCOUNT_ID := os.Getenv("accountID")
	KV_NAMESPACE := os.Getenv("kvNamespace")
	ctx := context.Background()
	client, err := cloudflare.New(API_KEY, API_EMAIL, cloudflare.UsingAccount(ACCOUNT_ID))
	check(err)
	key := website
	payload := []byte(base64.StdEncoding.EncodeToString([]byte(userInput)))
	resp, err := client.WriteWorkersKV(ctx, KV_NAMESPACE, key, payload)
	check(err)
	fmt.Println(resp)
}

func getPassword(website string) {
	API_KEY := os.Getenv("apiKey")
	API_EMAIL := os.Getenv("apiMail")
	ACCOUNT_ID := os.Getenv("accountID")
	KV_NAMESPACE := os.Getenv("kvNamespace")
	ctx := context.Background()
	client, err := cloudflare.New(API_KEY, API_EMAIL, cloudflare.UsingAccount(ACCOUNT_ID))
	check(err)
	key := website
	resp, err := client.ReadWorkersKV(ctx, KV_NAMESPACE, key)
	check(err)
	password, err := base64.StdEncoding.DecodeString(string(resp))
	check(err)
	fmt.Println("your password is", string(password))
}

func deletePassword(website string) {
	API_KEY := os.Getenv("apiKey")
	API_EMAIL := os.Getenv("apiMail")
	ACCOUNT_ID := os.Getenv("accountID")
	KV_NAMESPACE := os.Getenv("kvNamespace")
	ctx := context.Background()
	client, err := cloudflare.New(API_KEY, API_EMAIL, cloudflare.UsingAccount(ACCOUNT_ID))
	check(err)
	key := website
	resp, err := client.DeleteWorkersKV(ctx, KV_NAMESPACE, key)
	check(err)
	fmt.Println(resp)
}

func main() {
	err := godotenv.Load(".env")
	check(err)
	fmt.Println("Welcome to went. Press 1 for saving a password, 2 for retrieving a password and 3 for deleting a password")
	var option int
	fmt.Scanln(&option)
	if option == 1 {
		var userInput string
		var website string
		fmt.Println("website you use with the password: ")
		fmt.Scan(&website)
		fmt.Println("enter password: ")
		fmt.Scan(&userInput)
		encodePassword(website, userInput)
	} else if option == 2 {
		var website string
		fmt.Println("website you use the password with: ")
		fmt.Scan(&website)
		getPassword(website)
	} else if option == 3 {
		var website string
		fmt.Println("website who's password you want to delete: ")
		fmt.Scan(&website)
		deletePassword(website)
	}
}
