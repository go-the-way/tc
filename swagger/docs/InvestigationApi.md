# \InvestigationApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInvestigation**](InvestigationApi.md#AddInvestigation) | **Post** /app/rest/investigations | Create a new investigation.
[**AddMultipleInvestigations**](InvestigationApi.md#AddMultipleInvestigations) | **Post** /app/rest/investigations/multiple | Create multiple new investigations.
[**DeleteInvestigation**](InvestigationApi.md#DeleteInvestigation) | **Delete** /app/rest/investigations/{investigationLocator} | Delete investigation matching the locator.
[**GetAllInvestigations**](InvestigationApi.md#GetAllInvestigations) | **Get** /app/rest/investigations | Get all investigations.
[**GetInvestigation**](InvestigationApi.md#GetInvestigation) | **Get** /app/rest/investigations/{investigationLocator} | Get investigation matching the locator.
[**ReplaceInvestigation**](InvestigationApi.md#ReplaceInvestigation) | **Put** /app/rest/investigations/{investigationLocator} | Update investigation matching the locator.


# **AddInvestigation**
> Investigation AddInvestigation(ctx, optional)
Create a new investigation.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddInvestigationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddInvestigationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Investigation**](Investigation.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Investigation**](investigation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMultipleInvestigations**
> Investigations AddMultipleInvestigations(ctx, optional)
Create multiple new investigations.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddMultipleInvestigationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddMultipleInvestigationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Investigations**](Investigations.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Investigations**](investigations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInvestigation**
> DeleteInvestigation(ctx, investigationLocator)
Delete investigation matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **investigationLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllInvestigations**
> Investigations GetAllInvestigations(ctx, optional)
Get all investigations.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInvestigationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllInvestigationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Investigations**](investigations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvestigation**
> Investigation GetInvestigation(ctx, investigationLocator, optional)
Get investigation matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **investigationLocator** | **string**|  | 
 **optional** | ***GetInvestigationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetInvestigationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Investigation**](investigation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceInvestigation**
> Investigation ReplaceInvestigation(ctx, investigationLocator, optional)
Update investigation matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **investigationLocator** | **string**|  | 
 **optional** | ***ReplaceInvestigationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceInvestigationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Investigation**](Investigation.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Investigation**](investigation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

