# \ProjectApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAgentPoolsProject**](ProjectApi.md#AddAgentPoolsProject) | **Post** /app/rest/projects/{projectLocator}/agentPools | Assign the matching project to the agent pool.
[**AddBuildType**](ProjectApi.md#AddBuildType) | **Post** /app/rest/projects/{projectLocator}/buildTypes | Add a build configuration to the matching project.
[**AddFeature**](ProjectApi.md#AddFeature) | **Post** /app/rest/projects/{projectLocator}/projectFeatures | Add a feature.
[**AddProject**](ProjectApi.md#AddProject) | **Post** /app/rest/projects | Create a new project.
[**AddSecureToken**](ProjectApi.md#AddSecureToken) | **Post** /app/rest/projects/{projectLocator}/secure/tokens | Create a new secure token for the matching project.
[**AddTemplate**](ProjectApi.md#AddTemplate) | **Post** /app/rest/projects/{projectLocator}/templates | Add a build configuration template to the matching project.
[**CreateBuildParameter**](ProjectApi.md#CreateBuildParameter) | **Post** /app/rest/projects/{projectLocator}/parameters | Create a build parameter.
[**DeleteBuildParameter**](ProjectApi.md#DeleteBuildParameter) | **Delete** /app/rest/projects/{projectLocator}/parameters/{name} | Delete build parameter.
[**DeleteBuildParameters**](ProjectApi.md#DeleteBuildParameters) | **Delete** /app/rest/projects/{projectLocator}/parameters | Delete all build parameters.
[**DeleteFeature**](ProjectApi.md#DeleteFeature) | **Delete** /app/rest/projects/{projectLocator}/projectFeatures/{featureLocator} | Delete a matching feature.
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /app/rest/projects/{projectLocator} | Delete project matching the locator.
[**GetAgentPoolsProject**](ProjectApi.md#GetAgentPoolsProject) | **Get** /app/rest/projects/{projectLocator}/agentPools | Get agent pools appointed to the matching project.
[**GetAllBranches**](ProjectApi.md#GetAllBranches) | **Get** /app/rest/projects/{projectLocator}/branches | Get all branches of the matching project.
[**GetAllBuildTypesOrdered**](ProjectApi.md#GetAllBuildTypesOrdered) | **Get** /app/rest/projects/{projectLocator}/order/buildTypes | Get all build configurations from the matching project, with custom ordering applied.
[**GetAllProjects**](ProjectApi.md#GetAllProjects) | **Get** /app/rest/projects | Get all projects.
[**GetAllSubprojectsOrdered**](ProjectApi.md#GetAllSubprojectsOrdered) | **Get** /app/rest/projects/{projectLocator}/order/projects | Get all subprojects of the matching project, with custom ordering applied.
[**GetBuildParameter**](ProjectApi.md#GetBuildParameter) | **Get** /app/rest/projects/{projectLocator}/parameters/{name} | Get build parameter.
[**GetBuildParameterSpecification**](ProjectApi.md#GetBuildParameterSpecification) | **Get** /app/rest/projects/{projectLocator}/parameters/{name}/type/rawValue | Get build parameter specification.
[**GetBuildParameterType**](ProjectApi.md#GetBuildParameterType) | **Get** /app/rest/projects/{projectLocator}/parameters/{name}/type | Get type of build parameter.
[**GetBuildParameterValue**](ProjectApi.md#GetBuildParameterValue) | **Get** /app/rest/projects/{projectLocator}/parameters/{name}/value | Get value of build parameter.
[**GetBuildParameters**](ProjectApi.md#GetBuildParameters) | **Get** /app/rest/projects/{projectLocator}/parameters | Get build parameters.
[**GetDefaultTemplate**](ProjectApi.md#GetDefaultTemplate) | **Get** /app/rest/projects/{projectLocator}/defaultTemplate | Get the default template of the matching project.
[**GetFeature**](ProjectApi.md#GetFeature) | **Get** /app/rest/projects/{projectLocator}/projectFeatures/{featureLocator} | Get a matching feature.
[**GetFeatures**](ProjectApi.md#GetFeatures) | **Get** /app/rest/projects/{projectLocator}/projectFeatures | Get all features.
[**GetProject**](ProjectApi.md#GetProject) | **Get** /app/rest/projects/{projectLocator} | Get project matching the locator.
[**GetProjectField**](ProjectApi.md#GetProjectField) | **Get** /app/rest/projects/{projectLocator}/{field} | Get a field of the matching project.
[**GetProjectParentProject**](ProjectApi.md#GetProjectParentProject) | **Get** /app/rest/projects/{projectLocator}/parentProject | Get the parent project of the matching project.
[**GetProjectSettingsFile**](ProjectApi.md#GetProjectSettingsFile) | **Get** /app/rest/projects/{projectLocator}/settingsFile | Get the settings file of the matching build configuration.
[**GetProjectTemplates**](ProjectApi.md#GetProjectTemplates) | **Get** /app/rest/projects/{projectLocator}/templates | Get all templates of the matching project.
[**GetSecureValue**](ProjectApi.md#GetSecureValue) | **Get** /app/rest/projects/{projectLocator}/secure/values/{token} | Get a secure token of the matching project.
[**RemoveDefaultTemplate**](ProjectApi.md#RemoveDefaultTemplate) | **Delete** /app/rest/projects/{projectLocator}/defaultTemplate | Remove the default template from the matching project.
[**RemoveProjectFromAgentPool**](ProjectApi.md#RemoveProjectFromAgentPool) | **Delete** /app/rest/projects/{projectLocator}/agentPools/{agentPoolLocator} | Unassign a project from the matching agent pool.
[**SetAgentPoolsProject**](ProjectApi.md#SetAgentPoolsProject) | **Put** /app/rest/projects/{projectLocator}/agentPools | Update agent pools apppointed to the matching project.
[**SetBuildTypesOrder**](ProjectApi.md#SetBuildTypesOrder) | **Put** /app/rest/projects/{projectLocator}/order/buildTypes | Update custom ordering of build configurations of the matching project.
[**SetDefaultTemplate**](ProjectApi.md#SetDefaultTemplate) | **Put** /app/rest/projects/{projectLocator}/defaultTemplate | Update the default template of the matching project.
[**SetParentProject**](ProjectApi.md#SetParentProject) | **Put** /app/rest/projects/{projectLocator}/parentProject | Update the parent project of the matching project.
[**SetProjectField**](ProjectApi.md#SetProjectField) | **Put** /app/rest/projects/{projectLocator}/{field} | Update a field of the matching project.
[**SetSubprojectsOrder**](ProjectApi.md#SetSubprojectsOrder) | **Put** /app/rest/projects/{projectLocator}/order/projects | Update custom ordering of subprojects of the matching project.
[**UpdateBuildParameter**](ProjectApi.md#UpdateBuildParameter) | **Put** /app/rest/projects/{projectLocator}/parameters/{name} | Update build parameter.
[**UpdateBuildParameterSpecification**](ProjectApi.md#UpdateBuildParameterSpecification) | **Put** /app/rest/projects/{projectLocator}/parameters/{name}/type/rawValue | Update build parameter specification.
[**UpdateBuildParameterType**](ProjectApi.md#UpdateBuildParameterType) | **Put** /app/rest/projects/{projectLocator}/parameters/{name}/type | Update type of build parameter.
[**UpdateBuildParameterValue**](ProjectApi.md#UpdateBuildParameterValue) | **Put** /app/rest/projects/{projectLocator}/parameters/{name}/value | Update value of build parameter.
[**UpdateBuildParameters**](ProjectApi.md#UpdateBuildParameters) | **Put** /app/rest/projects/{projectLocator}/parameters | Update build parameters.
[**UpdateFeature**](ProjectApi.md#UpdateFeature) | **Put** /app/rest/projects/{projectLocator}/projectFeatures/{featureLocator} | Update a matching feature.
[**UpdateFeatures**](ProjectApi.md#UpdateFeatures) | **Put** /app/rest/projects/{projectLocator}/projectFeatures | Update all features.


# **AddAgentPoolsProject**
> AgentPool AddAgentPoolsProject(ctx, projectLocator, optional)
Assign the matching project to the agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***AddAgentPoolsProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddAgentPoolsProjectOpts struct

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

# **AddBuildType**
> BuildType AddBuildType(ctx, projectLocator, optional)
Add a build configuration to the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***AddBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NewBuildTypeDescription**](NewBuildTypeDescription.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**BuildType**](buildType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddFeature**
> interface{} AddFeature(ctx, projectLocator, optional)
Add a feature.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***AddFeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddFeatureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProjectFeature**](ProjectFeature.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddProject**
> Project AddProject(ctx, optional)
Create a new project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NewProjectDescription**](NewProjectDescription.md)|  | 

### Return type

[**Project**](project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSecureToken**
> string AddSecureToken(ctx, projectLocator, optional)
Create a new secure token for the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***AddSecureTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddSecureTokenOpts struct

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

# **AddTemplate**
> BuildType AddTemplate(ctx, projectLocator, optional)
Add a build configuration template to the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***AddTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NewBuildTypeDescription**](NewBuildTypeDescription.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**BuildType**](buildType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBuildParameter**
> Property CreateBuildParameter(ctx, projectLocator, optional)
Create a build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***CreateBuildParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateBuildParameterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Property**](Property.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Property**](property.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBuildParameter**
> DeleteBuildParameter(ctx, name, projectLocator)
Delete build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **projectLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBuildParameters**
> DeleteBuildParameters(ctx, projectLocator)
Delete all build parameters.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFeature**
> DeleteFeature(ctx, featureLocator, projectLocator)
Delete a matching feature.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureLocator** | **string**|  | 
  **projectLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject(ctx, projectLocator)
Delete project matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgentPoolsProject**
> AgentPools GetAgentPoolsProject(ctx, projectLocator, optional)
Get agent pools appointed to the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***GetAgentPoolsProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAgentPoolsProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**AgentPools**](agentPools.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllBranches**
> Branches GetAllBranches(ctx, projectLocator, optional)
Get all branches of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***GetAllBranchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBranchesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Branches**](branches.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllBuildTypesOrdered**
> BuildTypes GetAllBuildTypesOrdered(ctx, projectLocator, optional)
Get all build configurations from the matching project, with custom ordering applied.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***GetAllBuildTypesOrderedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBuildTypesOrderedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **field** | **optional.String**|  | 

### Return type

[**BuildTypes**](buildTypes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProjects**
> Projects GetAllProjects(ctx, optional)
Get all projects.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Projects**](projects.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSubprojectsOrdered**
> Projects GetAllSubprojectsOrdered(ctx, projectLocator, optional)
Get all subprojects of the matching project, with custom ordering applied.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***GetAllSubprojectsOrderedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllSubprojectsOrderedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **field** | **optional.String**|  | 

### Return type

[**Projects**](projects.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildParameter**
> Property GetBuildParameter(ctx, name, projectLocator, optional)
Get build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **projectLocator** | **string**|  | 
 **optional** | ***GetBuildParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildParameterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**Property**](property.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildParameterSpecification**
> string GetBuildParameterSpecification(ctx, name, projectLocator)
Get build parameter specification.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **projectLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildParameterType**
> ModelType GetBuildParameterType(ctx, name, projectLocator)
Get type of build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **projectLocator** | **string**|  | 

### Return type

[**ModelType**](type.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildParameterValue**
> string GetBuildParameterValue(ctx, name, projectLocator)
Get value of build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **projectLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildParameters**
> Properties GetBuildParameters(ctx, projectLocator, optional)
Get build parameters.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***GetBuildParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildParametersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Properties**](properties.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultTemplate**
> BuildType GetDefaultTemplate(ctx, projectLocator, optional)
Get the default template of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***GetDefaultTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDefaultTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**BuildType**](buildType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeature**
> interface{} GetFeature(ctx, featureLocator, projectLocator, optional)
Get a matching feature.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureLocator** | **string**|  | 
  **projectLocator** | **string**|  | 
 **optional** | ***GetFeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFeatureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatures**
> interface{} GetFeatures(ctx, projectLocator, optional)
Get all features.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***GetFeaturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFeaturesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> Project GetProject(ctx, projectLocator, optional)
Get project matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***GetProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Project**](project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectField**
> string GetProjectField(ctx, projectLocator, field)
Get a field of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectParentProject**
> Project GetProjectParentProject(ctx, projectLocator, optional)
Get the parent project of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***GetProjectParentProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetProjectParentProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Project**](project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectSettingsFile**
> string GetProjectSettingsFile(ctx, projectLocator)
Get the settings file of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectTemplates**
> BuildTypes GetProjectTemplates(ctx, projectLocator, optional)
Get all templates of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***GetProjectTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetProjectTemplatesOpts struct

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

# **GetSecureValue**
> string GetSecureValue(ctx, projectLocator, token)
Get a secure token of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
  **token** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveDefaultTemplate**
> RemoveDefaultTemplate(ctx, projectLocator, optional)
Remove the default template from the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***RemoveDefaultTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoveDefaultTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveProjectFromAgentPool**
> RemoveProjectFromAgentPool(ctx, projectLocator, agentPoolLocator)
Unassign a project from the matching agent pool.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
  **agentPoolLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAgentPoolsProject**
> AgentPools SetAgentPoolsProject(ctx, projectLocator, optional)
Update agent pools apppointed to the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***SetAgentPoolsProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetAgentPoolsProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AgentPools**](AgentPools.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**AgentPools**](agentPools.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBuildTypesOrder**
> BuildTypes SetBuildTypesOrder(ctx, projectLocator, optional)
Update custom ordering of build configurations of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***SetBuildTypesOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildTypesOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BuildTypes**](BuildTypes.md)|  | 
 **field** | **optional.String**|  | 

### Return type

[**BuildTypes**](buildTypes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDefaultTemplate**
> BuildType SetDefaultTemplate(ctx, projectLocator, optional)
Update the default template of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***SetDefaultTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetDefaultTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BuildType**](BuildType.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**BuildType**](buildType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetParentProject**
> Project SetParentProject(ctx, projectLocator, optional)
Update the parent project of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***SetParentProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetParentProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Project**](Project.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Project**](project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetProjectField**
> string SetProjectField(ctx, projectLocator, field, optional)
Update a field of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***SetProjectFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetProjectFieldOpts struct

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

# **SetSubprojectsOrder**
> Projects SetSubprojectsOrder(ctx, projectLocator, optional)
Update custom ordering of subprojects of the matching project.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***SetSubprojectsOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetSubprojectsOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Projects**](Projects.md)|  | 
 **field** | **optional.String**|  | 

### Return type

[**Projects**](projects.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBuildParameter**
> Property UpdateBuildParameter(ctx, name, projectLocator, optional)
Update build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **projectLocator** | **string**|  | 
 **optional** | ***UpdateBuildParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildParameterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Property**](Property.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Property**](property.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBuildParameterSpecification**
> string UpdateBuildParameterSpecification(ctx, name, projectLocator, optional)
Update build parameter specification.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **projectLocator** | **string**|  | 
 **optional** | ***UpdateBuildParameterSpecificationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildParameterSpecificationOpts struct

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

# **UpdateBuildParameterType**
> ModelType UpdateBuildParameterType(ctx, name, projectLocator, optional)
Update type of build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **projectLocator** | **string**|  | 
 **optional** | ***UpdateBuildParameterTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildParameterTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ModelType**](ModelType.md)|  | 

### Return type

[**ModelType**](type.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBuildParameterValue**
> string UpdateBuildParameterValue(ctx, name, projectLocator, optional)
Update value of build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **projectLocator** | **string**|  | 
 **optional** | ***UpdateBuildParameterValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildParameterValueOpts struct

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

# **UpdateBuildParameters**
> Properties UpdateBuildParameters(ctx, projectLocator, optional)
Update build parameters.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***UpdateBuildParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildParametersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Properties**](Properties.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Properties**](properties.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeature**
> interface{} UpdateFeature(ctx, featureLocator, projectLocator, optional)
Update a matching feature.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureLocator** | **string**|  | 
  **projectLocator** | **string**|  | 
 **optional** | ***UpdateFeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateFeatureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ProjectFeature**](ProjectFeature.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeatures**
> interface{} UpdateFeatures(ctx, projectLocator, optional)
Update all features.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectLocator** | **string**|  | 
 **optional** | ***UpdateFeaturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateFeaturesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProjectFeatures**](ProjectFeatures.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

