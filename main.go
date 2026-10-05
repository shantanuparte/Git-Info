package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter username ")
		return
	}

	username := os.Args[1]

	user, err := Get_user_info(username)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Printf("Github User: %s\n", user.Login)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Bio: %s\n", user.Bio)
	fmt.Printf("Location: %s\n", user.Location)
	fmt.Printf("Company: %s\n", user.Company)
	fmt.Printf("Public Repo's: %d\n", user.PublicRepos)
	fmt.Printf("Followers: %d\n", user.Followers)
	fmt.Printf("Following: %d\n", user.Following)
	fmt.Printf("Profile: %s\n,", user.HtmlUrll)
	
	
}