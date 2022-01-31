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
	err := godotenv.Load(".env")
	check(err)
	API_KEY := os.Getenv("apiKey")
	API_EMAIL := os.Getenv("apiMail")
	ACCOUNT_ID := os.Getenv("accountID")
	KV_NAMESPACE := os.Getenv("kvNamespace")
	ctx := context.Background()
	api, err := cloudflare.New(API_KEY, API_EMAIL, cloudflare.UsingAccount(ACCOUNT_ID))
	check(err)
	key := website
	payload := []byte(base64.StdEncoding.EncodeToString([]byte(userInput)))
	resp, err := api.WriteWorkersKV(ctx, KV_NAMESPACE, key, payload)
	check(err)
	fmt.Println(resp)
}

func main() {
	var userInput string
	var website string
	fmt.Println("website you use with the password: ")
	fmt.Scan(&website)
	fmt.Println("enter password: ")
	fmt.Scan(&userInput)
	encodePassword(website, userInput)
}
