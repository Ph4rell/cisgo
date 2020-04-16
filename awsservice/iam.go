package awsservice

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

// Ensure multi-factor authentication (MFA) is enabled
// for all IAM users that have a console password

type User struct {
	Id   string
	Name string
	MFA  bool
}

func ListUsers(sess *session.Session) (users []User) {
	svc := iam.New(sess)
	input := &iam.ListUsersInput{}

	result, err := svc.ListUsers(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}
	for _, u := range result.Users {
		users = append(users, User{
			Id:   *u.UserId,
			Name: *u.UserName,
		})
	}
	return users
}

func ListMFA(sess *session.Session, username string) (user []User) {
	svc := iam.New(sess)
	input := &iam.ListMFADevicesInput{
		UserName: aws.String(username),
	}

	result, err := svc.ListMFADevices(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	if len(result.MFADevices) > 0 {
		user = append(user, User{
			MFA: true,
		})
	}
	return user
}
