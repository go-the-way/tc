# \ChangeApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllChanges**](ChangeApi.md#GetAllChanges) | **Get** /app/rest/changes | Get all changes.
[**GetChange**](ChangeApi.md#GetChange) | **Get** /app/rest/changes/{changeLocator} | Get change matching the locator.
[**GetChangeAttributes**](ChangeApi.md#GetChangeAttributes) | **Get** /app/rest/changes/{changeLocator}/attributes | Get attributes of the matching change.
[**GetChangeDuplicates**](ChangeApi.md#GetChangeDuplicates) | **Get** /app/rest/changes/{changeLocator}/duplicates | Get duplicates of the matching change.
[**GetChangeField**](ChangeApi.md#GetChangeField) | **Get** /app/rest/changes/{changeLocator}/{field} | Get a field of the matching change.
[**GetChangeIssue**](ChangeApi.md#GetChangeIssue) | **Get** /app/rest/changes/{changeLocator}/issues | Get issues of the matching change.
[**GetChangeParentChanges**](ChangeApi.md#GetChangeParentChanges) | **Get** /app/rest/changes/{changeLocator}/parentChanges | Get parent changes of the matching change.
[**GetChangeParentRevisions**](ChangeApi.md#GetChangeParentRevisions) | **Get** /app/rest/changes/{changeLocator}/parentRevisions | Get parent revisions of the matching change.
[**GetChangeVcsRoot**](ChangeApi.md#GetChangeVcsRoot) | **Get** /app/rest/changes/{changeLocator}/vcsRootInstance | Get a VCS root instance of the matching change.


# **GetAllChanges**
> Changes GetAllChanges(ctx, optional)
Get all changes.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllChangesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Changes**](changes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChange**
> Change GetChange(ctx, changeLocator, optional)
Get change matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeLocator** | **string**|  | 
 **optional** | ***GetChangeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChangeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Change**](change.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChangeAttributes**
> Entries GetChangeAttributes(ctx, changeLocator, optional)
Get attributes of the matching change.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeLocator** | **string**|  | 
 **optional** | ***GetChangeAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChangeAttributesOpts struct

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

# **GetChangeDuplicates**
> Changes GetChangeDuplicates(ctx, changeLocator, optional)
Get duplicates of the matching change.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeLocator** | **string**|  | 
 **optional** | ***GetChangeDuplicatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChangeDuplicatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Changes**](changes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChangeField**
> string GetChangeField(ctx, changeLocator, field)
Get a field of the matching change.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChangeIssue**
> Issues GetChangeIssue(ctx, changeLocator)
Get issues of the matching change.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeLocator** | **string**|  | 

### Return type

[**Issues**](issues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChangeParentChanges**
> Changes GetChangeParentChanges(ctx, changeLocator, optional)
Get parent changes of the matching change.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeLocator** | **string**|  | 
 **optional** | ***GetChangeParentChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChangeParentChangesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Changes**](changes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChangeParentRevisions**
> Items GetChangeParentRevisions(ctx, changeLocator)
Get parent revisions of the matching change.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeLocator** | **string**|  | 

### Return type

[**Items**](items.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChangeVcsRoot**
> VcsRootInstance GetChangeVcsRoot(ctx, changeLocator, optional)
Get a VCS root instance of the matching change.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeLocator** | **string**|  | 
 **optional** | ***GetChangeVcsRootOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChangeVcsRootOpts struct

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

