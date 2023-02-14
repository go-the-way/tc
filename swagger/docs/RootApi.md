# \RootApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiVersion**](RootApi.md#GetApiVersion) | **Get** /app/rest/apiVersion | Get the API version.
[**GetPluginInfo**](RootApi.md#GetPluginInfo) | **Get** /app/rest/info | Get the plugin info.
[**GetRootEndpointsOfRoot**](RootApi.md#GetRootEndpointsOfRoot) | **Get** /app/rest | Get root endpoints.
[**GetVersion**](RootApi.md#GetVersion) | **Get** /app/rest/version | Get the TeamCity server version.


# **GetApiVersion**
> string GetApiVersion(ctx, )
Get the API version.



### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPluginInfo**
> Plugin GetPluginInfo(ctx, optional)
Get the plugin info.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPluginInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPluginInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 

### Return type

[**Plugin**](plugin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRootEndpointsOfRoot**
> string GetRootEndpointsOfRoot(ctx, )
Get root endpoints.



### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersion**
> string GetVersion(ctx, )
Get the TeamCity server version.



### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

