package awsservice

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

func UserHasMFA(svc *iam.IAM, user *iam.UserDetail) bool {
	input := &iam.ListMFADevicesInput{
		UserName: aws.String(*user.UserName),
	}

	result, err := svc.ListMFADevices(input)
	if err != nil {
		fmt.Sprintf("Error: %v", err)
	}

	if len(result.MFADevices) > 0 {
		return true
	}
	return false
}
