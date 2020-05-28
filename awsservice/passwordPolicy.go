package awsservice

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/iam"
)

type PasswordPolicy struct {
	ExpirePasswpord            bool
	MaxPasswordAge             int64
	MinimumPasswordLength      int64
	PasswordReusePrevention    int64
	RequireLowercaseCharacters bool
	RequireNumbers             bool
	RequireSymbols             bool
	RequireUppercaseCharacters bool
}

func ListPasswordPolicy(p []PasswordPolicy) string {
	for _, p := range p {
		fmt.Printf("ExpirePassword: %v\n", p.ExpirePasswpord)
		fmt.Printf("MaxPasswordAge: %v\n", p.MaxPasswordAge)
		fmt.Printf("MinimumPasswordLength: %v\n", p.MinimumPasswordLength)
		fmt.Printf("PasswordReusePrevention: %v\n", p.PasswordReusePrevention)
		fmt.Printf("RequireLowercaseCharacters: %v\n", p.RequireLowercaseCharacters)
		fmt.Printf("RequireNumbers: %v\n", p.RequireNumbers)
		fmt.Printf("RequireSymbols: %v\n", p.RequireSymbols)
		fmt.Printf("RequireUppercaseCharacters: %v\n", p.RequireUppercaseCharacters)
	}
	return "-------------"
}

func GetAccountPasswordPolicy(svc *iam.IAM) (policy []PasswordPolicy) {
	input := &iam.GetAccountPasswordPolicyInput{}

	result, err := svc.GetAccountPasswordPolicy(input)
	if err != nil {
		fmt.Println("This AWS account does not have a password policy set.")

	} else {
		policy = append(policy, PasswordPolicy{
			ExpirePasswpord:            *result.PasswordPolicy.ExpirePasswords,
			MaxPasswordAge:             *result.PasswordPolicy.MaxPasswordAge,
			MinimumPasswordLength:      *result.PasswordPolicy.MinimumPasswordLength,
			PasswordReusePrevention:    *result.PasswordPolicy.PasswordReusePrevention,
			RequireLowercaseCharacters: *result.PasswordPolicy.RequireLowercaseCharacters,
			RequireNumbers:             *result.PasswordPolicy.RequireNumbers,
			RequireSymbols:             *result.PasswordPolicy.RequireSymbols,
			RequireUppercaseCharacters: *result.PasswordPolicy.RequireUppercaseCharacters,
		})
	}
	return policy
}
