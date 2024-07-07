/*
Logto API references

Testing VerificationCodesAPIService

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

func Test_logto_VerificationCodesAPIService(t *testing.T) {

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)

	t.Run("Test VerificationCodesAPIService CreateVerificationCode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.VerificationCodesAPI.CreateVerificationCode(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VerificationCodesAPIService VerifyVerificationCode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.VerificationCodesAPI.VerifyVerificationCode(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
