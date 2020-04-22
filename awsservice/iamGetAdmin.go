package awsservice

import "github.com/aws/aws-sdk-go/service/iam"

func IsUserAdmin(svc *iam.IAM, user *iam.UserDetail, admin string) bool {
	// Check policy, attached policy, and groups (policy and attached policy)
	policyHasAdmin := UserPolicyHasAdmin(user, admin)
	if policyHasAdmin {
		return true
	}

	attachedPolicyHasAdmin := AttachedUserPolicyHasAdmin(user, admin)
	if attachedPolicyHasAdmin {
		return true
	}

	// userGroupsHaveAdmin := UsersGroupsHaveAdmin(svc, user, admin)
	// if userGroupsHaveAdmin {
	// 	return true
	// }

	return false
}

func UserPolicyHasAdmin(user *iam.UserDetail, admin string) bool {
	for _, policy := range user.UserPolicyList {
		if *policy.PolicyName == admin {
			return true
		}
	}

	return false
}

func AttachedUserPolicyHasAdmin(user *iam.UserDetail, admin string) bool {
	for _, policy := range user.AttachedManagedPolicies {
		if *policy.PolicyName == admin {
			return true
		}
	}

	return false
}
