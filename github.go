package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	Location    string `json:"location"`
	Company     string `json:"company"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	HtmlUrll    string `json:"html_url"`
}

func Get_user_info(username string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println("Error in request ")
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("Error: user %q not found lol", username)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: %s", resp.Status)
	}

	var user User

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("Error in decoding")
	}

	return &user, nil

}
