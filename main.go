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

	//List all the users
	// users := awsservice.ListUsers(sess)
	// for _, u := range users {
	// 	awsservice.GetUser(sess, u)
	// 	u.Mfa.Serial = awsservice.ListMFA(sess, u)
	// 	// if there is no MFA
	// 	if u.Mfa.Serial == "" {
	// 		u.Mfa.Serial = "MFA Missing"
	// 	}
	// 	awsservice.ListUserInfo(u)
	// }
	svc := iam.New(sess)
	input := &iam.GetAccountAuthorizationDetailsInput{}

	result, err := svc.GetAccountAuthorizationDetails(input)
	if err != nil {
		fmt.Println("Got error getting account details")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	users := 0
	admins := 0

	for _, u := range result.UserDetailList {
		users += 1
		fmt.Println(*u.UserName)
		// check if user got MFA
		if awsservice.UserHasMFA(svc, u) {
			fmt.Println("MFA found")
		}
		// check if user got Admin rights
		if awsservice.IsUserAdmin(svc, u, "AdministratorAccess") {
			admins += 1
		}
	}
	fmt.Printf("Number of user: %v\n", users)
	fmt.Printf("Number of admin: %v\n", admins)
}
