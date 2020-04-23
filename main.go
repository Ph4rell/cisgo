package main

import (
	"cisgo/awsservice"
	"cisgo/awssession"
	"flag"
	"fmt"

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

	users := awsservice.ListUsers(svc)

	usercount := 0
	admincount := 0

	for _, u := range users {
		usercount += 1
		if u.IsAdmin {
			admincount += 1
		}
		awsservice.ListUserInfo(u)
	}
	fmt.Printf("Number of users: %v\n", usercount)
	fmt.Printf("Number of admins: %v\n", admincount)

}
