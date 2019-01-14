package integration

import (
	"log"
	"os"
	"testing"

	"github.com/callumkerredwards/localstack-go-wrapper/localstack"
)

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	if !testing.Short() {
		container, err := localstack.Start()
		if err != nil {
			log.Printf("Cannot start localstack, %v", err)
			return 1
		}
		defer localstack.Stop(container)
	}
	return m.Run()
}
