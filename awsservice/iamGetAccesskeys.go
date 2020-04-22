package awsservice

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

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
		fmt.Println(r)
	}
	return false
}
