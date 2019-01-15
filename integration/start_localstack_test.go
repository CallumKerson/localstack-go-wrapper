package integration

import (
	"log"
	"os"
	"testing"

	"github.com/callumkerredwards/localstack-go-wrapper/localstack"
	"github.com/callumkerredwards/localstack-go-wrapper/localstack/services"
)

// s3Port is a custom port
const s3Port int = 4590

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	if !testing.Short() {
		s3Config := &services.ServiceConfig{
			Service: services.S3,
			Port:    s3Port,
		}
		log.Printf("Creating localstack S3 config with port %v", s3Port)
		sqsConfig := &services.ServiceConfig{
			Service: services.SQS,
		}
		log.Print("Creating localstack SQS config with port default port")
		snsConfig := &services.ServiceConfig{
			Service: services.SNS,
		}
		log.Print("Creating localstack SNS config with port default port")
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
