/*
mailcow API

Testing MailboxesApiService

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

func Test_mailcow_MailboxesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MailboxesApiService CreateMailbox", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MailboxesApi.CreateMailbox(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailboxesApiService DeleteMailbox", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MailboxesApi.DeleteMailbox(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailboxesApiService DeleteMailboxTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mailbox string

		resp, httpRes, err := apiClient.MailboxesApi.DeleteMailboxTags(context.Background(), mailbox).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailboxesApiService EditMailboxSpamFilterScore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MailboxesApi.EditMailboxSpamFilterScore(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailboxesApiService GetMailboxes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.MailboxesApi.GetMailboxes(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailboxesApiService GetMailboxesOfADomain", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		httpRes, err := apiClient.MailboxesApi.GetMailboxesOfADomain(context.Background(), domain).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailboxesApiService QuarantineNotifications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.MailboxesApi.QuarantineNotifications(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailboxesApiService UpdateMailbox", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MailboxesApi.UpdateMailbox(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailboxesApiService UpdateMailboxACL", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MailboxesApi.UpdateMailboxACL(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailboxesApiService UpdatePushoverSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MailboxesApi.UpdatePushoverSettings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}