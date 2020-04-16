package main

import (
	"cisgo/awssession"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// account := flag.String("account", "", "AWS Account number")
	region := flag.String("region", "eu-west-1", "Name of the region")
	profile := flag.String("profile", "default", "AWS profile name")

	flag.Parse()

	// If AWS account not 12 digits
	// if len(*account) != 12 {
	// 	fmt.Printf("Invalid account. got='%v'\n", *account)
	// 	os.Exit(1)
	// }

	// p := session.SessionConfig.Profile
	// r := session.SessionConfig.Region
	// region = append(r, region)
	// profile = append(p, profile)

	var config awssession.SessionConfig
	config.Profile = *profile
	config.Region = *region

	sess, err := awssession.CreateSession(config)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	svc := ec2.New(sess)
	input := &ec2.DescribeVpcsInput{}

	result, err := svc.DescribeVpcs(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
