package main

import (
	"context"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"os"

	logto "github.com/mostafa/go-api-client"
)

type BearerToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func main() {
	// Create a new Logto API client using the token from the token file
	tokenPath := "./token.json"
	// Read the token from the file
	token, err := os.ReadFile(tokenPath)
	if err != nil {
		panic(err)
	}

	// Unmarshal the token
	var bearerToken BearerToken
	err = json.Unmarshal(token, &bearerToken)
	if err != nil {
		panic(err)
	}

	// Create a new Logto API client
	config := logto.NewConfiguration()
	config.Debug = false // Enable debug mode to see the request and response details
	config.Servers = logto.ServerConfigurations{
		{
			URL: "http://localhost:3001",
		},
	}
	// Add the authorization header with the bearer token
	config.AddDefaultHeader("Authorization", bearerToken.TokenType+" "+bearerToken.AccessToken)
	// Create the client
	client := logto.NewAPIClient(config)

	// Get the organization ID from the command line arguments
	organizationID := os.Args[1]

	// Get the organization details
	org, _, err := client.OrganizationsAPI.GetOrganization(
		context.Background(), organizationID).Execute()
	if err != nil {
		panic(err)
	}

	// Print the organization details
	spew.Dump(org)

	// Get the list of users in the organization
	users, _, err := client.OrganizationsAPI.ListOrganizationUsers(
		context.Background(), organizationID).Execute()
	if err != nil {
		panic(err)
	}

	// Print the list of users
	spew.Dump(users)
}
