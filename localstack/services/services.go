package services

import "fmt"

// ServiceConfig allows the configuration of LocalStack Service to specify a custom port
type ServiceConfig struct {
	Service Service
	Port    int
}

// Service represents a LocalStack service
type Service int

// The supported LocalStack services
const (
	APIGateway Service = iota + 1
	Kinesis
	DynamoDB
	DynamoDBStreams
	S3
	Firehose
	Lambda
	SNS
	SQS
	Redshift
	ES
	SES
	Route53
	CloudFormation
	CloudWatch
	SSM
	SecretsManager
)

// GetDefaultPort gets the default LocalStack port for the given service
func GetDefaultPort(ser Service) (int, error) {
	var port int
	switch ser {
	case APIGateway:
		port = 4567
	case Kinesis:
		port = 4568
	case DynamoDB:
		port = 4569
	case DynamoDBStreams:
		port = 4570
	case S3:
		port = 4572
	case Firehose:
		port = 4573
	case Lambda:
		port = 4574
	case SNS:
		port = 4575
	case SQS:
		port = 4576
	case Redshift:
		port = 4577
	case ES:
		port = 4578
	case SES:
		port = 4579
	case Route53:
		port = 4580
	case CloudFormation:
		port = 4581
	case CloudWatch:
		port = 4582
	case SSM:
		port = 4583
	case SecretsManager:
		port = 4584
	default:
		return 0, fmt.Errorf("Unrecognised Service %s", ser)
	}
	return port, nil
}
