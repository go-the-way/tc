# \NodeApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeNodeResponsibility**](NodeApi.md#ChangeNodeResponsibility) | **Put** /app/rest/server/nodes/{nodeLocator}/enabledResponsibilities/{name} | Enables or disables responsibility for a node.
[**GetAllNodes**](NodeApi.md#GetAllNodes) | **Get** /app/rest/server/nodes | Get all TeamCity nodes.
[**GetDisabledResponsibilities**](NodeApi.md#GetDisabledResponsibilities) | **Get** /app/rest/server/nodes/{nodeLocator}/disabledResponsibilities | Get all effective responsibilities for a node matching the locator.
[**GetEffectiveResponsibilities**](NodeApi.md#GetEffectiveResponsibilities) | **Get** /app/rest/server/nodes/{nodeLocator}/effectiveResponsibilities | Get all effective responsibilities for a node matching the locator.
[**GetEnabledResponsibilities**](NodeApi.md#GetEnabledResponsibilities) | **Get** /app/rest/server/nodes/{nodeLocator}/enabledResponsibilities | Get all enabled responsibilities for a node matching the locator.
[**GetNode**](NodeApi.md#GetNode) | **Get** /app/rest/server/nodes/{nodeLocator} | Get a node matching the locator.


# **ChangeNodeResponsibility**
> EnabledResponsibilities ChangeNodeResponsibility(ctx, nodeLocator, name, optional)
Enables or disables responsibility for a node.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeLocator** | **string**|  | 
  **name** | **string**|  | 
 **optional** | ***ChangeNodeResponsibilityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChangeNodeResponsibilityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **optional.String**|  | 

### Return type

[**EnabledResponsibilities**](enabledResponsibilities.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllNodes**
> Nodes GetAllNodes(ctx, optional)
Get all TeamCity nodes.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllNodesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Nodes**](nodes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDisabledResponsibilities**
> DisabledResponsibilities GetDisabledResponsibilities(ctx, nodeLocator, optional)
Get all effective responsibilities for a node matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeLocator** | **string**|  | 
 **optional** | ***GetDisabledResponsibilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDisabledResponsibilitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**DisabledResponsibilities**](disabledResponsibilities.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveResponsibilities**
> EffectiveResponsibilities GetEffectiveResponsibilities(ctx, nodeLocator, optional)
Get all effective responsibilities for a node matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeLocator** | **string**|  | 
 **optional** | ***GetEffectiveResponsibilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEffectiveResponsibilitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**EffectiveResponsibilities**](effectiveResponsibilities.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnabledResponsibilities**
> EnabledResponsibilities GetEnabledResponsibilities(ctx, nodeLocator, optional)
Get all enabled responsibilities for a node matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeLocator** | **string**|  | 
 **optional** | ***GetEnabledResponsibilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEnabledResponsibilitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**EnabledResponsibilities**](enabledResponsibilities.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNode**
> Node GetNode(ctx, nodeLocator, optional)
Get a node matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeLocator** | **string**|  | 
 **optional** | ***GetNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetNodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Node**](node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

