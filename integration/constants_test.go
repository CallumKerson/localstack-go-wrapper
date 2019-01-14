package integration

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var testCredentials = credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")
var testRegion = aws.String("eu-west-1")
var disableSSL = aws.Bool(true)
