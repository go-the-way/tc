# \VcsRootApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVcsRoot**](VcsRootApi.md#AddVcsRoot) | **Post** /app/rest/vcs-roots | Add a new VCS root.
[**DeleteAllVcsRootProperties**](VcsRootApi.md#DeleteAllVcsRootProperties) | **Delete** /app/rest/vcs-roots/{vcsRootLocator}/properties | Delete all properties of the matching VCS root.
[**DeleteVcsRoot**](VcsRootApi.md#DeleteVcsRoot) | **Delete** /app/rest/vcs-roots/{vcsRootLocator} | Remove VCS root matching the locator.
[**DeleteVcsRootProperty**](VcsRootApi.md#DeleteVcsRootProperty) | **Delete** /app/rest/vcs-roots/{vcsRootLocator}/properties/{name} | Delete a property of the matching VCS root.
[**GetAllVcsRootProperties**](VcsRootApi.md#GetAllVcsRootProperties) | **Get** /app/rest/vcs-roots/{vcsRootLocator}/properties | Get all properties of the matching VCS root.
[**GetAllVcsRoots**](VcsRootApi.md#GetAllVcsRoots) | **Get** /app/rest/vcs-roots | Get all VCS roots.
[**GetRootEndpoints**](VcsRootApi.md#GetRootEndpoints) | **Get** /app/rest/vcs-roots/{vcsRootLocator} | Get root endpoints.
[**GetVcsRootField**](VcsRootApi.md#GetVcsRootField) | **Get** /app/rest/vcs-roots/{vcsRootLocator}/{field} | Get a field of the matching VCS root.
[**GetVcsRootInstances**](VcsRootApi.md#GetVcsRootInstances) | **Get** /app/rest/vcs-roots/{vcsRootLocator}/instances | Get all VCS root instances of the matching VCS root.
[**GetVcsRootProperty**](VcsRootApi.md#GetVcsRootProperty) | **Get** /app/rest/vcs-roots/{vcsRootLocator}/properties/{name} | Get a property on the matching VCS root.
[**GetVcsRootSettingsFile**](VcsRootApi.md#GetVcsRootSettingsFile) | **Get** /app/rest/vcs-roots/{vcsRootLocator}/settingsFile | Get the settings file of the matching VCS root.
[**SetVcsRootField**](VcsRootApi.md#SetVcsRootField) | **Put** /app/rest/vcs-roots/{vcsRootLocator}/{field} | Update a field of the matching VCS root.
[**SetVcsRootProperties**](VcsRootApi.md#SetVcsRootProperties) | **Put** /app/rest/vcs-roots/{vcsRootLocator}/properties | Update all properties of the matching VCS root.
[**SetVcsRootProperty**](VcsRootApi.md#SetVcsRootProperty) | **Put** /app/rest/vcs-roots/{vcsRootLocator}/properties/{name} | Update a property of the matching VCS root.


# **AddVcsRoot**
> VcsRoot AddVcsRoot(ctx, optional)
Add a new VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddVcsRootOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddVcsRootOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VcsRoot**](VcsRoot.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**VcsRoot**](vcs-root.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllVcsRootProperties**
> DeleteAllVcsRootProperties(ctx, vcsRootLocator)
Delete all properties of the matching VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVcsRoot**
> DeleteVcsRoot(ctx, vcsRootLocator)
Remove VCS root matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVcsRootProperty**
> DeleteVcsRootProperty(ctx, vcsRootLocator, name)
Delete a property of the matching VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllVcsRootProperties**
> Properties GetAllVcsRootProperties(ctx, vcsRootLocator, optional)
Get all properties of the matching VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 
 **optional** | ***GetAllVcsRootPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllVcsRootPropertiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Properties**](properties.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllVcsRoots**
> VcsRoots GetAllVcsRoots(ctx, optional)
Get all VCS roots.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllVcsRootsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllVcsRootsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**VcsRoots**](vcs-roots.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRootEndpoints**
> VcsRoot GetRootEndpoints(ctx, vcsRootLocator, optional)
Get root endpoints.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 
 **optional** | ***GetRootEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRootEndpointsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**VcsRoot**](vcs-root.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRootField**
> string GetVcsRootField(ctx, vcsRootLocator, field)
Get a field of the matching VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRootInstances**
> VcsRootInstances GetVcsRootInstances(ctx, vcsRootLocator, optional)
Get all VCS root instances of the matching VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 
 **optional** | ***GetVcsRootInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetVcsRootInstancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**VcsRootInstances**](vcs-root-instances.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRootProperty**
> string GetVcsRootProperty(ctx, vcsRootLocator, name)
Get a property on the matching VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 
  **name** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRootSettingsFile**
> string GetVcsRootSettingsFile(ctx, vcsRootLocator)
Get the settings file of the matching VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetVcsRootField**
> string SetVcsRootField(ctx, vcsRootLocator, field, optional)
Update a field of the matching VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***SetVcsRootFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetVcsRootFieldOpts struct

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

# **SetVcsRootProperties**
> Properties SetVcsRootProperties(ctx, vcsRootLocator, optional)
Update all properties of the matching VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 
 **optional** | ***SetVcsRootPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetVcsRootPropertiesOpts struct

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

# **SetVcsRootProperty**
> string SetVcsRootProperty(ctx, vcsRootLocator, name, optional)
Update a property of the matching VCS root.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootLocator** | **string**|  | 
  **name** | **string**|  | 
 **optional** | ***SetVcsRootPropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetVcsRootPropertyOpts struct

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

