package main

import (
	"cisgo/awsservice"
	"cisgo/awssession"
	"flag"
	"fmt"
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

	// user1 := &awsservice.User{
	// 	Id:   "3456789",
	// 	Name: "Pierre",
	// }
	// user1.Mfa.Serial = awsservice.ListMFA(sess, *user1)

	// user2 := awsservice.User{
	// 	Id:   "23456789",
	// 	Name: "Bob",
	// 	Mfa:  awsservice.Mfa{"serial"},
	// }

	// awsservice.ListUserInfo(*user1)
	// awsservice.ListUserInfo(user2)
	users := awsservice.ListUsers(sess)
	for _, u := range users {
		u.Mfa.Serial = awsservice.ListMFA(sess, u)
		awsservice.ListUserInfo(u)
	}
	// mfa := &awsservice.Mfa{}
	// mfa.ListMFA(sess, users.Name)

	// for _, u := range users {
	// 	fmt.Println("User:", u.Name)
	// 	u.Mfa.ListMFA(sess, u)
	// 	fmt.Println(u)
	// 	awsservice.ListMFA(sess, u)
	// 	fmt.Printf("User: %v - MFA: %v\n", u.Name, u.Mfa.Serial)
	// 	for _, m := range mfa {
	// 		fmt.Println("MFA:", m.Serial)
	// 	}
	// }
}
