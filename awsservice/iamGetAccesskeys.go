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
	LastUsed   *time.Time
}

func ListAccessKeysInfo(a []AccessKey) string {
	for _, a := range a {
		fmt.Printf("AccessKeyId: %v - CreateDate: %v - Status: %v - LastUsed: %v\n", a.Id, a.CreateDate, a.Status, a.LastUsed)
	}
	return ""
}

// Func return bool if user has access keys
func UserHasAccessKey(svc *iam.IAM, user *iam.UserDetail) bool {
	input := &iam.ListAccessKeysInput{
		UserName: aws.String(*user.UserName),
	}

	result, err := svc.ListAccessKeys(input)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, r := range result.AccessKeyMetadata {
		if r.Status == aws.String("Active") {
			return true
		}
	}
	return false
}

// Func to list accesskeys of a user
// Return a slice of AccessKey struct
func ListAccessKeys(svc *iam.IAM, user *string) (keys []AccessKey) {
	input := &iam.ListAccessKeysInput{
		UserName: aws.String(*user),
	}
	result, err := svc.ListAccessKeys(input)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, r := range result.AccessKeyMetadata {
		keys = append(keys, AccessKey{
			Id:         *r.AccessKeyId,
			CreateDate: r.CreateDate,
			Status:     *r.Status,
			LastUsed:   getAccessKeyLastUsed(svc, *r.AccessKeyId),
		})
	}
	return keys
}

// Func to retrieve the lastUsedDate of an AccessKey
// Param: key string
// return timestamp
func getAccessKeyLastUsed(svc *iam.IAM, key string) *time.Time {
	result, err := svc.GetAccessKeyLastUsed(&iam.GetAccessKeyLastUsedInput{
		AccessKeyId: aws.String(key),
	})
	if err != nil {
		fmt.Println("Error", err)
	}
	return result.AccessKeyLastUsed.LastUsedDate
}
