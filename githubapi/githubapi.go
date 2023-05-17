package main

import (
	"context"
	"fmt"
	"strconv"

	"example.com/getpass"

	// "github.com/google/go-github/v52/github"	// with go modules enabled (GO111MODULE=on or outside GOPATH)

	"github.com/google/go-github/github" // with go modules disabled

	"golang.org/x/oauth2"
)

func main() {


	//getting personal access token from user
	pst, err := getpass.GetPassword("Enter your personal access token: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//creating  a github client with personal access token
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pst},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	
	//getting data from user
	var name string
	var isPrivate string
	var description string
	fmt.Print(`Enter your repo name:`)
	fmt.Scan(&name)
	fmt.Print(`Enter your repo's description:`)
	fmt.Scan(&description)
	fmt.Print(`Is your repo private?("true"/"false"):`)
	fmt.Scan(&isPrivate)
	isPrivateCheck,err := strconv.ParseBool(isPrivate)
	
	
	// create a new private repository with user data
	repo := &github.Repository{
		Name:    github.String(name),
		Description: github.String(description),
		Private: github.Bool(isPrivateCheck),
	}
	client.Repositories.Create(ctx, "", repo)


}
