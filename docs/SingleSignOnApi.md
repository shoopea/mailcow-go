# \SingleSignOnApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssueDomainAdminSSOToken**](SingleSignOnApi.md#IssueDomainAdminSSOToken) | **Post** /api/v1/add/sso/domain-admin | Issue Domain Admin SSO token



## IssueDomainAdminSSOToken

> IssueDomainAdminSSOToken(ctx).IssueDomainAdminSSOTokenRequest(issueDomainAdminSSOTokenRequest).Execute()

Issue Domain Admin SSO token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shoopea/mailcow-go"
)

func main() {
    issueDomainAdminSSOTokenRequest := *openapiclient.NewIssueDomainAdminSSOTokenRequest() // IssueDomainAdminSSOTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SingleSignOnApi.IssueDomainAdminSSOToken(context.Background()).IssueDomainAdminSSOTokenRequest(issueDomainAdminSSOTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SingleSignOnApi.IssueDomainAdminSSOToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueDomainAdminSSOTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueDomainAdminSSOTokenRequest** | [**IssueDomainAdminSSOTokenRequest**](IssueDomainAdminSSOTokenRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

