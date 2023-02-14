# \BuildQueueApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBuildToQueue**](BuildQueueApi.md#AddBuildToQueue) | **Post** /app/rest/buildQueue | Add a new build to the queue.
[**AddTagsToBuildOfBuildQueue**](BuildQueueApi.md#AddTagsToBuildOfBuildQueue) | **Post** /app/rest/buildQueue/{buildLocator}/tags | Add tags to the matching build.
[**ApproveQueuedBuild**](BuildQueueApi.md#ApproveQueuedBuild) | **Post** /app/rest/buildQueue/{buildLocator}/approve | Approve queued build with approval feature enabled.
[**CancelQueuedBuild**](BuildQueueApi.md#CancelQueuedBuild) | **Post** /app/rest/buildQueue/{queuedBuildLocator} | Cancel a queued matching build.
[**DeleteAllQueuedBuilds**](BuildQueueApi.md#DeleteAllQueuedBuilds) | **Delete** /app/rest/buildQueue | Delete all queued builds.
[**DeleteQueuedBuild**](BuildQueueApi.md#DeleteQueuedBuild) | **Delete** /app/rest/buildQueue/{queuedBuildLocator} | Delete a queued matching build.
[**GetAllQueuedBuilds**](BuildQueueApi.md#GetAllQueuedBuilds) | **Get** /app/rest/buildQueue | Get all queued builds.
[**GetApprovalInfo**](BuildQueueApi.md#GetApprovalInfo) | **Get** /app/rest/buildQueue/{buildLocator}/approvalInfo | Get approval info of a queued matching build.
[**GetCompatibleAgentsForBuild**](BuildQueueApi.md#GetCompatibleAgentsForBuild) | **Get** /app/rest/buildQueue/{queuedBuildLocator}/compatibleAgents | Get compatible agents for a queued matching build.
[**GetQueuedBuild**](BuildQueueApi.md#GetQueuedBuild) | **Get** /app/rest/buildQueue/{queuedBuildLocator} | Get a queued matching build.
[**GetQueuedBuildPosition**](BuildQueueApi.md#GetQueuedBuildPosition) | **Get** /app/rest/buildQueue/order/{queuePosition} | Get the queue position of a queued matching build.
[**GetQueuedBuildTags**](BuildQueueApi.md#GetQueuedBuildTags) | **Get** /app/rest/buildQueue/{buildLocator}/tags | Get tags of the queued matching build.
[**SetQueuedBuildPosition**](BuildQueueApi.md#SetQueuedBuildPosition) | **Put** /app/rest/buildQueue/order/{queuePosition} | Update the queue position of a queued matching build.
[**SetQueuedBuildsOrder**](BuildQueueApi.md#SetQueuedBuildsOrder) | **Put** /app/rest/buildQueue/order | Update the build queue order.


# **AddBuildToQueue**
> Build AddBuildToQueue(ctx, optional)
Add a new build to the queue.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddBuildToQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddBuildToQueueOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Build**](Build.md)|  | 
 **moveToTop** | **optional.Bool**|  | 

### Return type

[**Build**](build.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTagsToBuildOfBuildQueue**
> AddTagsToBuildOfBuildQueue(ctx, buildLocator, optional)
Add tags to the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***AddTagsToBuildOfBuildQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddTagsToBuildOfBuildQueueOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Tags**](Tags.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApproveQueuedBuild**
> ApprovalInfo ApproveQueuedBuild(ctx, buildLocator, optional)
Approve queued build with approval feature enabled.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***ApproveQueuedBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApproveQueuedBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**ApprovalInfo**](approvalInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelQueuedBuild**
> Build CancelQueuedBuild(ctx, queuedBuildLocator, optional)
Cancel a queued matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queuedBuildLocator** | **string**|  | 
 **optional** | ***CancelQueuedBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CancelQueuedBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BuildCancelRequest**](BuildCancelRequest.md)|  | 

### Return type

[**Build**](build.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllQueuedBuilds**
> DeleteAllQueuedBuilds(ctx, optional)
Delete all queued builds.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteAllQueuedBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteAllQueuedBuildsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQueuedBuild**
> DeleteQueuedBuild(ctx, queuedBuildLocator)
Delete a queued matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queuedBuildLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllQueuedBuilds**
> Builds GetAllQueuedBuilds(ctx, optional)
Get all queued builds.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllQueuedBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllQueuedBuildsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Builds**](builds.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApprovalInfo**
> ApprovalInfo GetApprovalInfo(ctx, buildLocator, optional)
Get approval info of a queued matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetApprovalInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApprovalInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**ApprovalInfo**](approvalInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompatibleAgentsForBuild**
> Agents GetCompatibleAgentsForBuild(ctx, queuedBuildLocator, optional)
Get compatible agents for a queued matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queuedBuildLocator** | **string**|  | 
 **optional** | ***GetCompatibleAgentsForBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCompatibleAgentsForBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Agents**](agents.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQueuedBuild**
> Build GetQueuedBuild(ctx, queuedBuildLocator, optional)
Get a queued matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queuedBuildLocator** | **string**|  | 
 **optional** | ***GetQueuedBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetQueuedBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Build**](build.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQueuedBuildPosition**
> Build GetQueuedBuildPosition(ctx, queuePosition, optional)
Get the queue position of a queued matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queuePosition** | **string**|  | 
 **optional** | ***GetQueuedBuildPositionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetQueuedBuildPositionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Build**](build.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQueuedBuildTags**
> Tags GetQueuedBuildTags(ctx, buildLocator, optional)
Get tags of the queued matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetQueuedBuildTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetQueuedBuildTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Tags**](tags.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetQueuedBuildPosition**
> Build SetQueuedBuildPosition(ctx, queuePosition, optional)
Update the queue position of a queued matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queuePosition** | **string**|  | 
 **optional** | ***SetQueuedBuildPositionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetQueuedBuildPositionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Build**](Build.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Build**](build.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetQueuedBuildsOrder**
> Builds SetQueuedBuildsOrder(ctx, optional)
Update the build queue order.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SetQueuedBuildsOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetQueuedBuildsOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Builds**](Builds.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Builds**](builds.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

