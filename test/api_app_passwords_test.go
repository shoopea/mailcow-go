/*
mailcow API

Testing AppPasswordsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mailcow

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/shoopea/mailcow-go"
)

func Test_mailcow_AppPasswordsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppPasswordsApiService CreateAppPassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppPasswordsApi.CreateAppPassword(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppPasswordsApiService DeleteAppPassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppPasswordsApi.DeleteAppPassword(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppPasswordsApiService GetAppPassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mailbox string

		httpRes, err := apiClient.AppPasswordsApi.GetAppPassword(context.Background(), mailbox).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
