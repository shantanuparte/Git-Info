package main

import (
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
	HTMLURL     string `json:"html_url"`
}

func Get_user_info(username string) (*User, error) {
	url := fmt.Sprintf("https://github.com/user/users/%s",username)

	

}
