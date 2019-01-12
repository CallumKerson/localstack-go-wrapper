// Package wrapper is a wrapper around LocalStack, to make it easier to
// integration test Go code that use AWS
package wrapper

import (
	// localstack package
	_ "github.com/callumkerredwards/localstack-go-wrapper/localstack"
)
