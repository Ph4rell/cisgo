package awsservice

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

type Policy struct {
	Id              string
	AttachmentCount int
}

// TODO add filter to input
func ListPolicies(svc *iam.IAM) (policies []Policy) {
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
		})
	}
	return policies
}
