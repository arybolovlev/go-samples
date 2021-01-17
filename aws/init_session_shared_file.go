package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	// Obtaining AWS Access and Secret keys from a shared file
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
		// First parameter for NewSharedCredentials is a shared file path
		// By default it is ~/.aws/credentials
		// Second parameter is a profile name
		// By default it is "default"
		Credentials: credentials.NewSharedCredentials("", "profile-name"),
	})

	if err != nil {
		fmt.Println("The following error occurred during setting up an AWS session:", err)
	}
}
