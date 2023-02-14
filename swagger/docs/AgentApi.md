# \AgentApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAgent**](AgentApi.md#DeleteAgent) | **Delete** /app/rest/agents/{agentLocator} | Delete an inactive agent.
[**GetAgent**](AgentApi.md#GetAgent) | **Get** /app/rest/agents/{agentLocator} | Get agent matching the locator.
[**GetAgentField**](AgentApi.md#GetAgentField) | **Get** /app/rest/agents/{agentLocator}/{field} | Get a field of the matching agent.
[**GetAgentPool**](AgentApi.md#GetAgentPool) | **Get** /app/rest/agents/{agentLocator}/pool | Get the agent pool of the matching agent.
[**GetAllAgents**](AgentApi.md#GetAllAgents) | **Get** /app/rest/agents | Get all known agents.
[**GetAuthorizedInfo**](AgentApi.md#GetAuthorizedInfo) | **Get** /app/rest/agents/{agentLocator}/authorizedInfo | Get the authorization info of the matching agent.
[**GetBuildConfigurationRunPolicy**](AgentApi.md#GetBuildConfigurationRunPolicy) | **Get** /app/rest/agents/{agentLocator}/compatibilityPolicy | Get the build configuration run policy of the matching agent.
[**GetCompatibleBuildTypes**](AgentApi.md#GetCompatibleBuildTypes) | **Get** /app/rest/agents/{agentLocator}/compatibleBuildTypes | Get build types compatible with the matching agent.
[**GetEnabledInfo**](AgentApi.md#GetEnabledInfo) | **Get** /app/rest/agents/{agentLocator}/enabledInfo | Check if the matching agent is enabled.
[**GetIncompatibleBuildTypes**](AgentApi.md#GetIncompatibleBuildTypes) | **Get** /app/rest/agents/{agentLocator}/incompatibleBuildTypes | Get build types incompatible with the matching agent.
[**SetAgentField**](AgentApi.md#SetAgentField) | **Put** /app/rest/agents/{agentLocator}/{field} | Update a field of the matching agent.
[**SetAgentPool**](AgentApi.md#SetAgentPool) | **Put** /app/rest/agents/{agentLocator}/pool | Assign the matching agent to the specified agent pool.
[**SetAuthorizedInfo**](AgentApi.md#SetAuthorizedInfo) | **Put** /app/rest/agents/{agentLocator}/authorizedInfo | Update the authorization info of the matching agent.
[**SetBuildConfigurationRunPolicy**](AgentApi.md#SetBuildConfigurationRunPolicy) | **Put** /app/rest/agents/{agentLocator}/compatibilityPolicy | Update build configuration run policy of agent matching locator.
[**SetEnabledInfo**](AgentApi.md#SetEnabledInfo) | **Put** /app/rest/agents/{agentLocator}/enabledInfo | Update the enablement status of the matching agent.


# **DeleteAgent**
> DeleteAgent(ctx, agentLocator)
Delete an inactive agent.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgent**
> Agent GetAgent(ctx, agentLocator, optional)
Get agent matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***GetAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Agent**](agent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgentField**
> string GetAgentField(ctx, agentLocator, field)
Get a field of the matching agent.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgentPool**
> AgentPool GetAgentPool(ctx, agentLocator, optional)
Get the agent pool of the matching agent.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***GetAgentPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAgentPoolOpts struct

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

# **GetAllAgents**
> Agents GetAllAgents(ctx, optional)
Get all known agents.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllAgentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllAgentsOpts struct

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

# **GetAuthorizedInfo**
> AuthorizedInfo GetAuthorizedInfo(ctx, agentLocator, optional)
Get the authorization info of the matching agent.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***GetAuthorizedInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAuthorizedInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**AuthorizedInfo**](authorizedInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildConfigurationRunPolicy**
> CompatibilityPolicy GetBuildConfigurationRunPolicy(ctx, agentLocator, optional)
Get the build configuration run policy of the matching agent.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***GetBuildConfigurationRunPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildConfigurationRunPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**CompatibilityPolicy**](compatibilityPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompatibleBuildTypes**
> BuildTypes GetCompatibleBuildTypes(ctx, agentLocator, optional)
Get build types compatible with the matching agent.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***GetCompatibleBuildTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCompatibleBuildTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**BuildTypes**](buildTypes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnabledInfo**
> EnabledInfo GetEnabledInfo(ctx, agentLocator, optional)
Check if the matching agent is enabled.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***GetEnabledInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEnabledInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**EnabledInfo**](enabledInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIncompatibleBuildTypes**
> Compatibilities GetIncompatibleBuildTypes(ctx, agentLocator, optional)
Get build types incompatible with the matching agent.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***GetIncompatibleBuildTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetIncompatibleBuildTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Compatibilities**](compatibilities.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAgentField**
> string SetAgentField(ctx, agentLocator, field, optional)
Update a field of the matching agent.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***SetAgentFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetAgentFieldOpts struct

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

# **SetAgentPool**
> AgentPool SetAgentPool(ctx, agentLocator, optional)
Assign the matching agent to the specified agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***SetAgentPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetAgentPoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AgentPool**](AgentPool.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**AgentPool**](agentPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAuthorizedInfo**
> AuthorizedInfo SetAuthorizedInfo(ctx, agentLocator, optional)
Update the authorization info of the matching agent.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***SetAuthorizedInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetAuthorizedInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AuthorizedInfo**](AuthorizedInfo.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**AuthorizedInfo**](authorizedInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBuildConfigurationRunPolicy**
> CompatibilityPolicy SetBuildConfigurationRunPolicy(ctx, agentLocator, optional)
Update build configuration run policy of agent matching locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***SetBuildConfigurationRunPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildConfigurationRunPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CompatibilityPolicy**](CompatibilityPolicy.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**CompatibilityPolicy**](compatibilityPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetEnabledInfo**
> EnabledInfo SetEnabledInfo(ctx, agentLocator, optional)
Update the enablement status of the matching agent.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentLocator** | **string**|  | 
 **optional** | ***SetEnabledInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetEnabledInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EnabledInfo**](EnabledInfo.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**EnabledInfo**](enabledInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

