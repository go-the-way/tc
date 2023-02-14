# \AgentPoolApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAgentToAgentPool**](AgentPoolApi.md#AddAgentToAgentPool) | **Post** /app/rest/agentPools/{agentPoolLocator}/agents | Assign the agent to the matching agent pool.
[**AddProjectToAgentPool**](AgentPoolApi.md#AddProjectToAgentPool) | **Post** /app/rest/agentPools/{agentPoolLocator}/projects | Assign the project to the matching agent pool.
[**CreateAgentPool**](AgentPoolApi.md#CreateAgentPool) | **Post** /app/rest/agentPools | Create a new agent pool.
[**DeleteAgentPool**](AgentPoolApi.md#DeleteAgentPool) | **Delete** /app/rest/agentPools/{agentPoolLocator} | Delete the agent pool matching the locator.
[**DeleteAllProjectsFromAgentPool**](AgentPoolApi.md#DeleteAllProjectsFromAgentPool) | **Delete** /app/rest/agentPools/{agentPoolLocator}/projects | Unassign all projects from the matching agent pool.
[**DeleteProjectFromAgentPool**](AgentPoolApi.md#DeleteProjectFromAgentPool) | **Delete** /app/rest/agentPools/{agentPoolLocator}/projects/{projectLocator} | Unassign the project from the matching agent pool.
[**GetAgentPoolOfAgentPool**](AgentPoolApi.md#GetAgentPoolOfAgentPool) | **Get** /app/rest/agentPools/{agentPoolLocator} | Get the agent pool matching the locator.
[**GetAllAgentPools**](AgentPoolApi.md#GetAllAgentPools) | **Get** /app/rest/agentPools | Get all agent pools.
[**GetAllAgentsFromAgentPool**](AgentPoolApi.md#GetAllAgentsFromAgentPool) | **Get** /app/rest/agentPools/{agentPoolLocator}/agents | Get the agent of the matching agent pool.
[**GetAllProjectsFromAgentPool**](AgentPoolApi.md#GetAllProjectsFromAgentPool) | **Get** /app/rest/agentPools/{agentPoolLocator}/projects | Get all projects of the matching agent pool.
[**GetFieldFromAgentPool**](AgentPoolApi.md#GetFieldFromAgentPool) | **Get** /app/rest/agentPools/{agentPoolLocator}/{field} | Get a field of the matching agent pool.
[**SetAgentPoolField**](AgentPoolApi.md#SetAgentPoolField) | **Put** /app/rest/agentPools/{agentPoolLocator}/{field} | Update a field of the matching agent pool.
[**SetAgentPoolProjects**](AgentPoolApi.md#SetAgentPoolProjects) | **Put** /app/rest/agentPools/{agentPoolLocator}/projects | Update projects of the matching agent pool.


# **AddAgentToAgentPool**
> Agent AddAgentToAgentPool(ctx, agentPoolLocator, optional)
Assign the agent to the matching agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 
 **optional** | ***AddAgentToAgentPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddAgentToAgentPoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Agent**](Agent.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Agent**](agent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddProjectToAgentPool**
> Project AddProjectToAgentPool(ctx, agentPoolLocator, optional)
Assign the project to the matching agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 
 **optional** | ***AddProjectToAgentPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddProjectToAgentPoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Project**](Project.md)|  | 

### Return type

[**Project**](project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAgentPool**
> AgentPool CreateAgentPool(ctx, optional)
Create a new agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAgentPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateAgentPoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AgentPool**](AgentPool.md)|  | 

### Return type

[**AgentPool**](agentPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAgentPool**
> DeleteAgentPool(ctx, agentPoolLocator)
Delete the agent pool matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllProjectsFromAgentPool**
> DeleteAllProjectsFromAgentPool(ctx, agentPoolLocator)
Unassign all projects from the matching agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectFromAgentPool**
> DeleteProjectFromAgentPool(ctx, agentPoolLocator, projectLocator)
Unassign the project from the matching agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 
  **projectLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgentPoolOfAgentPool**
> AgentPool GetAgentPoolOfAgentPool(ctx, agentPoolLocator, optional)
Get the agent pool matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 
 **optional** | ***GetAgentPoolOfAgentPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAgentPoolOfAgentPoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**AgentPool**](agentPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAgentPools**
> AgentPools GetAllAgentPools(ctx, optional)
Get all agent pools.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllAgentPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllAgentPoolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**AgentPools**](agentPools.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAgentsFromAgentPool**
> Agents GetAllAgentsFromAgentPool(ctx, agentPoolLocator, optional)
Get the agent of the matching agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 
 **optional** | ***GetAllAgentsFromAgentPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllAgentsFromAgentPoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Agents**](agents.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProjectsFromAgentPool**
> Projects GetAllProjectsFromAgentPool(ctx, agentPoolLocator, optional)
Get all projects of the matching agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 
 **optional** | ***GetAllProjectsFromAgentPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllProjectsFromAgentPoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Projects**](projects.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFieldFromAgentPool**
> string GetFieldFromAgentPool(ctx, agentPoolLocator, field)
Get a field of the matching agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAgentPoolField**
> string SetAgentPoolField(ctx, agentPoolLocator, field, optional)
Update a field of the matching agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***SetAgentPoolFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetAgentPoolFieldOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAgentPoolProjects**
> Projects SetAgentPoolProjects(ctx, agentPoolLocator, optional)
Update projects of the matching agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentPoolLocator** | **string**|  | 
 **optional** | ***SetAgentPoolProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetAgentPoolProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Projects**](Projects.md)|  | 

### Return type

[**Projects**](projects.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

