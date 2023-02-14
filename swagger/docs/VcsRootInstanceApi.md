# \VcsRootInstanceApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVcsRootInstanceField**](VcsRootInstanceApi.md#DeleteVcsRootInstanceField) | **Delete** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/{field} | Remove a field of the matching VCS root instance.
[**DeleteVcsRootInstanceRepositoryState**](VcsRootInstanceApi.md#DeleteVcsRootInstanceRepositoryState) | **Delete** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/repositoryState | Delete the last repository state of the matching VCS root instance.
[**DownloadFile**](VcsRootInstanceApi.md#DownloadFile) | **Get** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/files/latest/files{path} | Download specific file.
[**GetAllVcsRootInstances**](VcsRootInstanceApi.md#GetAllVcsRootInstances) | **Get** /app/rest/vcs-root-instances | Get all VCS root instances.
[**GetFileMetadata**](VcsRootInstanceApi.md#GetFileMetadata) | **Get** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/files/latest/metadata{path} | Get metadata of specific file.
[**GetFilesList**](VcsRootInstanceApi.md#GetFilesList) | **Get** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/files/latest | List all files.
[**GetFilesListForSubpath**](VcsRootInstanceApi.md#GetFilesListForSubpath) | **Get** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/files/latest/{path} | List files under this path.
[**GetVcsRootInstance**](VcsRootInstanceApi.md#GetVcsRootInstance) | **Get** /app/rest/vcs-root-instances/{vcsRootInstanceLocator} | Get VCS root instance matching the locator.
[**GetVcsRootInstanceCreationDate**](VcsRootInstanceApi.md#GetVcsRootInstanceCreationDate) | **Get** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/repositoryState/creationDate | Get the creation date of the matching VCS root instance.
[**GetVcsRootInstanceField**](VcsRootInstanceApi.md#GetVcsRootInstanceField) | **Get** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/{field} | Get a field of the matching VCS root instance.
[**GetVcsRootInstanceProperties**](VcsRootInstanceApi.md#GetVcsRootInstanceProperties) | **Get** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/properties | Get all properties of the matching VCS root instance.
[**GetVcsRootInstanceRepositoryState**](VcsRootInstanceApi.md#GetVcsRootInstanceRepositoryState) | **Get** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/repositoryState | Get the repository state of the matching VCS root instance.
[**GetZippedFile**](VcsRootInstanceApi.md#GetZippedFile) | **Get** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/files/latest/archived{path} | Get specific file zipped.
[**RequestPendingChangesCheck**](VcsRootInstanceApi.md#RequestPendingChangesCheck) | **Post** /app/rest/vcs-root-instances/checkingForChangesQueue | Check for the pending changes for all VCS root instances.
[**SetVcsRootInstanceField**](VcsRootInstanceApi.md#SetVcsRootInstanceField) | **Put** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/{field} | Get a field of the matching VCS root instance.
[**SetVcsRootInstanceRepositoryState**](VcsRootInstanceApi.md#SetVcsRootInstanceRepositoryState) | **Put** /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/repositoryState | Update the repository state of the matching VCS root instance.
[**TriggerCommitHookNotification**](VcsRootInstanceApi.md#TriggerCommitHookNotification) | **Post** /app/rest/vcs-root-instances/commitHookNotification | Send the commit hook notification.


# **DeleteVcsRootInstanceField**
> DeleteVcsRootInstanceField(ctx, vcsRootInstanceLocator, field)
Remove a field of the matching VCS root instance.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootInstanceLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVcsRootInstanceRepositoryState**
> DeleteVcsRootInstanceRepositoryState(ctx, vcsRootInstanceLocator)
Delete the last repository state of the matching VCS root instance.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootInstanceLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFile**
> DownloadFile(ctx, path, vcsRootInstanceLocator)
Download specific file.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **vcsRootInstanceLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllVcsRootInstances**
> VcsRootInstances GetAllVcsRootInstances(ctx, optional)
Get all VCS root instances.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllVcsRootInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllVcsRootInstancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**VcsRootInstances**](vcs-root-instances.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileMetadata**
> *os.File GetFileMetadata(ctx, path, vcsRootInstanceLocator, optional)
Get metadata of specific file.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **vcsRootInstanceLocator** | **string**|  | 
 **optional** | ***GetFileMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilesList**
> Files GetFilesList(ctx, vcsRootInstanceLocator, optional)
List all files.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootInstanceLocator** | **string**|  | 
 **optional** | ***GetFilesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFilesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **basePath** | **optional.String**|  | 
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Files**](files.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilesListForSubpath**
> Files GetFilesListForSubpath(ctx, path, vcsRootInstanceLocator, optional)
List files under this path.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **vcsRootInstanceLocator** | **string**|  | 
 **optional** | ***GetFilesListForSubpathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFilesListForSubpathOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **basePath** | **optional.String**|  | 
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Files**](files.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRootInstance**
> VcsRootInstance GetVcsRootInstance(ctx, vcsRootInstanceLocator, optional)
Get VCS root instance matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootInstanceLocator** | **string**|  | 
 **optional** | ***GetVcsRootInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetVcsRootInstanceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**VcsRootInstance**](vcs-root-instance.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRootInstanceCreationDate**
> string GetVcsRootInstanceCreationDate(ctx, vcsRootInstanceLocator)
Get the creation date of the matching VCS root instance.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootInstanceLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRootInstanceField**
> string GetVcsRootInstanceField(ctx, vcsRootInstanceLocator, field)
Get a field of the matching VCS root instance.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootInstanceLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRootInstanceProperties**
> Properties GetVcsRootInstanceProperties(ctx, vcsRootInstanceLocator, optional)
Get all properties of the matching VCS root instance.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootInstanceLocator** | **string**|  | 
 **optional** | ***GetVcsRootInstancePropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetVcsRootInstancePropertiesOpts struct

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

# **GetVcsRootInstanceRepositoryState**
> Entries GetVcsRootInstanceRepositoryState(ctx, vcsRootInstanceLocator, optional)
Get the repository state of the matching VCS root instance.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootInstanceLocator** | **string**|  | 
 **optional** | ***GetVcsRootInstanceRepositoryStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetVcsRootInstanceRepositoryStateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Entries**](entries.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetZippedFile**
> GetZippedFile(ctx, path, vcsRootInstanceLocator, optional)
Get specific file zipped.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **vcsRootInstanceLocator** | **string**|  | 
 **optional** | ***GetZippedFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetZippedFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **basePath** | **optional.String**|  | 
 **locator** | **optional.String**|  | 
 **name** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestPendingChangesCheck**
> VcsRootInstances RequestPendingChangesCheck(ctx, optional)
Check for the pending changes for all VCS root instances.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RequestPendingChangesCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequestPendingChangesCheckOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **requestor** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**VcsRootInstances**](vcs-root-instances.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetVcsRootInstanceField**
> string SetVcsRootInstanceField(ctx, vcsRootInstanceLocator, field, optional)
Get a field of the matching VCS root instance.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootInstanceLocator** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***SetVcsRootInstanceFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetVcsRootInstanceFieldOpts struct

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

# **SetVcsRootInstanceRepositoryState**
> Entries SetVcsRootInstanceRepositoryState(ctx, vcsRootInstanceLocator, optional)
Update the repository state of the matching VCS root instance.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcsRootInstanceLocator** | **string**|  | 
 **optional** | ***SetVcsRootInstanceRepositoryStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetVcsRootInstanceRepositoryStateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Entries**](Entries.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Entries**](entries.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerCommitHookNotification**
> TriggerCommitHookNotification(ctx, optional)
Send the commit hook notification.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TriggerCommitHookNotificationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerCommitHookNotificationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **okOnNothingFound** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

