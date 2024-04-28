package main

import (
	"fmt"
	"os"

	http_client "github.com/leetcode-golang-classroom/go-with-testcontainer/internal/http-client"
)

func main() {
	client, err := http_client.NewHttpClient("https://api.agify.io")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res, err := client.GetAge("Sig")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)

}
