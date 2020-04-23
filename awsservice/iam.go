package awsservice

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/iam"
)

type User struct {
	Id         string
	Arn        string
	Name       string
	MFA        bool
	IsAdmin    bool
	AccessKey  bool
	AccessKeys []AccessKey
}

func ListUserInfo(u User) string {
	fmt.Printf("Name: %v - ID: %v\n", u.Name, u.Id)
	fmt.Printf("IsAdmin: %v - MFA: %v\n", u.IsAdmin, u.MFA)
	ListAccessKeysInfo(u.AccessKeys)
	fmt.Println("----------")
	return ""
}

func ListUsers(svc *iam.IAM) (users []User) {
	input := &iam.GetAccountAuthorizationDetailsInput{}

	result, err := svc.GetAccountAuthorizationDetails(input)
	if err != nil {
		fmt.Println("Got error getting account details")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, u := range result.UserDetailList {
		users = append(users, User{
			Id:         *u.UserId,
			Arn:        *u.Arn,
			Name:       *u.UserName,
			MFA:        UserHasMFA(svc, u),
			IsAdmin:    IsUserAdmin(svc, u, "AdministratorAccess"),
			AccessKey:  UserHasAccessKey(svc, u),
			AccessKeys: ListAccessKeys(svc, u.UserName),
		})
	}
	return users
}
