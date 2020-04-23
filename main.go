package main

import (
	"cisgo/awsservice"
	"cisgo/awssession"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/iam"
)

func main() {
	// account := flag.String("account", "", "AWS Account number")
	region := flag.String("region", "eu-west-1", "Name of the region")
	profile := flag.String("profile", "default", "AWS profile name")

	flag.Parse()

	// If AWS account not 12 digits
	// if len(*account) != 12 {
	// 	fmt.Printf("Invalid account. got='%v'\n", *account)
	// 	os.Exit(1)
	// }

	// create the config
	var config awssession.SessionConfig
	// assign values
	config.Profile = *profile
	config.Region = *region

	sess, err := awssession.CreateSession(config)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	svc := iam.New(sess)
	input := &iam.GetAccountAuthorizationDetailsInput{}

	result, err := svc.GetAccountAuthorizationDetails(input)
	if err != nil {
		fmt.Println("Got error getting account details")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	type User struct {
		Id         string
		Arn        string
		Name       string
		MFA        bool
		IsAdmin    bool
		AccessKey  bool
		AccessKeys []awsservice.AccessKey
	}

	var users []User

	usercount := 0
	admincount := 0

	for _, u := range result.UserDetailList {
		users = append(users, User{
			Id:         *u.UserId,
			Arn:        *u.Arn,
			Name:       *u.UserName,
			MFA:        awsservice.UserHasMFA(svc, u),
			IsAdmin:    awsservice.IsUserAdmin(svc, u, "AdministratorAccess"),
			AccessKey:  awsservice.UserHasAccessKey(svc, u),
			AccessKeys: awsservice.ListAccessKeys(svc, u.UserName),
		})
	}

	for _, u := range users {
		fmt.Println(u)
	}
	fmt.Printf("Number of users: %v\n", usercount)
	fmt.Printf("Number of admins: %v\n", admincount)

}

// func String(user *iam.UserDetail) string {
// 	fmt.Printf("User: %v - MFA: %v - AccessKey: %v\n", user.UserName,
// }
