package awsservice

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

type AccessKey struct {
	Id         string
	CreateDate *time.Time
	Status     string
}

func ListAccessKeysInfo(a []AccessKey) string {
	for _, a := range a {
		fmt.Printf("AccessKeyId: %v - CreateDate: %v - Status: %v\n", a.Id, a.CreateDate, a.Status)
	}
	return ""
}

func UserHasAccessKey(svc *iam.IAM, user *iam.UserDetail) bool {
	input := &iam.ListAccessKeysInput{
		UserName: aws.String(*user.UserName),
	}

	result, err := svc.ListAccessKeys(input)
	if err != nil {
		fmt.Sprintf("Error: %v", err)
	}
	for _, r := range result.AccessKeyMetadata {
		if r.Status == aws.String("Active") {
			return true
		}
	}
	//fmt.Println(result)
	return false
}

func ListAccessKeys(svc *iam.IAM, user *string) (keys []AccessKey) {
	input := &iam.ListAccessKeysInput{
		UserName: aws.String(*user),
	}
	result, err := svc.ListAccessKeys(input)
	if err != nil {
		fmt.Sprintf("Error: %v", err)
	}
	for _, r := range result.AccessKeyMetadata {
		keys = append(keys, AccessKey{
			Id:         *r.AccessKeyId,
			CreateDate: r.CreateDate,
			Status:     *r.Status,
		})
	}
	//fmt.Println(keys)
	return keys
}
