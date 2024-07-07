/*
Logto API references

Testing StatusAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package logto

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	logto "github.com/mostafa/go-api-client"
)

func Test_logto_StatusAPIService(t *testing.T) {

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)

	t.Run("Test StatusAPIService GetStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StatusAPI.GetStatus(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
