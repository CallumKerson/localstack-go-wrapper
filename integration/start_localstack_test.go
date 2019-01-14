package integration

import (
	"log"
	"os"
	"testing"

	"github.com/callumkerredwards/localstack-go-wrapper/localstack"
)

// s3Port is a custom port
const s3Port int = 33000

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	if !testing.Short() {
		s3Config := &localstack.ServiceConfig{
			Serv: localstack.S3,
			Port: s3Port,
		}
		sqsConfig := &localstack.ServiceConfig{
			Serv: localstack.SQS,
		}
		snsConfig := &localstack.ServiceConfig{
			Serv: localstack.SNS,
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
