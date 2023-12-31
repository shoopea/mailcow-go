/*
mailcow API

Testing SingleSignOnApiService

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

func Test_mailcow_SingleSignOnApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SingleSignOnApiService IssueDomainAdminSSOToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SingleSignOnApi.IssueDomainAdminSSOToken(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
