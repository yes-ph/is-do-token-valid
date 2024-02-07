package main

import (
	"context"
	"fmt"
	"os"

	"github.com/digitalocean/godo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Token required\n")
		return
	}

	token := os.Args[1]

	client := godo.NewFromToken(token)

	account, _, err := client.Account.Get(context.TODO())

	if err != nil {
		fmt.Printf("\033[1;31m%s\033[0m", "Invalid token\n")
		return
	}

	if account != nil {
		fmt.Printf("\033[1;32m%s\033[0m", "Valid\n")
		return
	}

	fmt.Printf("\033[1;33m%s\033[0m", "Account not found\n")
}
