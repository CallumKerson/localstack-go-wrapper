package localstack

import "fmt"

// ServiceConfig allows the configuration of LocalStack Service to specify a custom port
type ServiceConfig struct {
	Serv Service
	Port int
}

// Service represents a LocalStack service
type Service int

// S3, SNS and SQS represent the supported LocalStack services
const (
	S3 Service = iota + 1
	SNS
	SQS
)

func getDefaultPort(ser Service) (int, error) {
	var port int
	switch ser {
	case S3:
		port = 4572
	case SNS:
		port = 4575
	case SQS:
		port = 4576
	default:
		return 0, fmt.Errorf("Unrecognised Service %s", ser)
	}
	return port, nil
}
