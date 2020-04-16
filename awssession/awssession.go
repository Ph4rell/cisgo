package awssession

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type SessionConfig struct {
	Profile string
	Region  string
}

func CreateSession(config SessionConfig) (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(
		session.Options{
			Profile: config.Profile,
			Config: aws.Config{
				Region: &config.Region,
			},
		})
	return sess, err
}
