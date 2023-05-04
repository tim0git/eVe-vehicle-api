package clients_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"timdevs.rest.api.com/m/v2/clients"
)

var mockRegion string = "us-east-1"

func TestReturnsDynamoDbClient(t *testing.T) {
	t.Parallel()
	actual := clients.DynamoDbClient()
	assert.NotNil(t, actual)
}

func TestReturnsDynamoDbClientWithConfig(t *testing.T) {
	t.Parallel()
	actual := clients.DynamoDbClient()
	assert.NotNil(t, actual.Config)
}

func TestReturnsDynamoDbClientWithRegion(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("AWS_REGION", mockRegion)
	actual := clients.DynamoDbClient()
	assert.Equal(t, mockRegion, *actual.Config.Region)
}
