# \ProblemOccurrenceApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllBuildProblemOccurrences**](ProblemOccurrenceApi.md#GetAllBuildProblemOccurrences) | **Get** /app/rest/problemOccurrences | Get all build problem occurrences.
[**GetBuildProblemOccurrence**](ProblemOccurrenceApi.md#GetBuildProblemOccurrence) | **Get** /app/rest/problemOccurrences/{problemLocator} | Get a matching build problem occurrence.


# **GetAllBuildProblemOccurrences**
> ProblemOccurrences GetAllBuildProblemOccurrences(ctx, optional)
Get all build problem occurrences.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllBuildProblemOccurrencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBuildProblemOccurrencesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**ProblemOccurrences**](problemOccurrences.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildProblemOccurrence**
> ProblemOccurrence GetBuildProblemOccurrence(ctx, problemLocator, optional)
Get a matching build problem occurrence.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **problemLocator** | **string**|  | 
 **optional** | ***GetBuildProblemOccurrenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildProblemOccurrenceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**ProblemOccurrence**](problemOccurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

