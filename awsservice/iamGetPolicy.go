package awsservice

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

type Policy struct {
	Id              string
	Name            string
	AttachmentCount int
	IsUsed          bool
}

// Func return a list of CustomerPolicy
func ListCustomerPolicies(svc *iam.IAM) (policies []Policy) {
	input := &iam.GetAccountAuthorizationDetailsInput{
		Filter: []*string{
			aws.String("LocalManagedPolicy"),
		},
	}
	result, err := svc.GetAccountAuthorizationDetails(input)
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range result.Policies {
		policies = append(policies, Policy{
			Id:              *p.PolicyId,
			AttachmentCount: int(*p.AttachmentCount),
			Name:            *p.PolicyName,
			IsUsed:          isUsed(*p.AttachmentCount),
		})
	}
	return policies
}

// Func return true if AttachementCount is != zero
func isUsed(c int64) bool {
	return c != *aws.Int64(0)
}
