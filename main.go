package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello this is my first time in neo vim")
	fmt.Println("Welcome to git info visualizer")
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	url := "https://api.github.com/users/{shantanuparte}"

	resp, err := client.Get(url)

	if err != nil {

		fmt.Println(resp)
	}
	fmt.Println("Resp failed")

}
