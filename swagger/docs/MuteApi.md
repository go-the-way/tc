# \MuteApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllMutedTests**](MuteApi.md#GetAllMutedTests) | **Get** /app/rest/mutes | Get all muted tests.
[**GetMutedTest**](MuteApi.md#GetMutedTest) | **Get** /app/rest/mutes/{muteLocator} | Get a muted test.
[**MuteMultipleTests**](MuteApi.md#MuteMultipleTests) | **Post** /app/rest/mutes/multiple | Mute multiple tests.
[**MuteTest**](MuteApi.md#MuteTest) | **Post** /app/rest/mutes | Mute a test.
[**UnmuteTest**](MuteApi.md#UnmuteTest) | **Delete** /app/rest/mutes/{muteLocator} | Unmute the matching test.


# **GetAllMutedTests**
> Mutes GetAllMutedTests(ctx, optional)
Get all muted tests.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllMutedTestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllMutedTestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Mutes**](mutes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMutedTest**
> Mute GetMutedTest(ctx, muteLocator, optional)
Get a muted test.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **muteLocator** | **string**|  | 
 **optional** | ***GetMutedTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMutedTestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Mute**](mute.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuteMultipleTests**
> Mutes MuteMultipleTests(ctx, optional)
Mute multiple tests.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MuteMultipleTestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MuteMultipleTestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Mutes**](Mutes.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Mutes**](mutes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuteTest**
> Mute MuteTest(ctx, optional)
Mute a test.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MuteTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MuteTestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Mute**](Mute.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Mute**](mute.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnmuteTest**
> UnmuteTest(ctx, muteLocator, optional)
Unmute the matching test.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **muteLocator** | **string**|  | 
 **optional** | ***UnmuteTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnmuteTestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

