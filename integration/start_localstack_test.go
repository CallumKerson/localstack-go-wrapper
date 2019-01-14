package integration

import (
	"log"
	"os"
	"testing"

	"github.com/callumkerredwards/localstack-go-wrapper/localstack"
)

// s3Port is a custom port
const s3Port int = 3572

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	if !testing.Short() {
		s3Config := &localstack.ServiceConfig{
			Service: localstack.S3,
			Port:    s3Port,
		}
		sqsConfig := &localstack.ServiceConfig{
			Service: localstack.SQS,
		}
		snsConfig := &localstack.ServiceConfig{
			Service: localstack.SNS,
		}
		container, err := localstack.New(s3Config, sqsConfig, snsConfig)
		if err != nil {
			log.Printf("Cannot create localstack, %v", err)
			return 1
		}
		err = container.Start()
		if err != nil {
			log.Printf("Cannot start localstack, %v", err)
			return 1
		}
		defer container.Stop()
	}
	return m.Run()
}
