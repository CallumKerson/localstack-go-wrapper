# localstack-go-wrapper
A wrapper around [LocalStack](https://github.com/localstack/localstack) for integration testing in Go.

### Prerequesites

- [Go 1.11](https://golang.org/doc/go1.11) for module support `$GO111MODULE=on`
- [Docker](https://www.docker.com) installed and daemon running

### Usage

```
package demo

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
    s3Config := &localstack.ServiceConfig{
        Service: localstack.S3,
        Port: 33000,
    }
    container, err := localstack.New(s3Config)
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
	return m.Run()
}
```

If the port is not specified, it will use the [LocalStack](https://github.com/localstack/localstack) 
defaults. Remember to ensure that LocalStack is not running elsewhere if using default ports.

### Make

- `make` can be used to format, lint, install and test this library.
