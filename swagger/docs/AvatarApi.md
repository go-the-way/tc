# \AvatarApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAvatar**](AvatarApi.md#DeleteAvatar) | **Delete** /app/rest/avatars/{userLocator} | Delete a users avatar
[**GetAvatar**](AvatarApi.md#GetAvatar) | **Get** /app/rest/avatars/{userLocator}/{size}/avatar.png | Get a users avatar
[**GetAvatarWithHash**](AvatarApi.md#GetAvatarWithHash) | **Get** /app/rest/avatars/{userLocator}/{size}/avatar.{hash}.png | Get a users avatar
[**PutAvatar**](AvatarApi.md#PutAvatar) | **Put** /app/rest/avatars/{userLocator} | Update a users avatar


# **DeleteAvatar**
> DeleteAvatar(ctx, userLocator)
Delete a users avatar



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvatar**
> GetAvatar(ctx, userLocator, size)
Get a users avatar



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **size** | **int32**| avatar&#39;s size | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvatarWithHash**
> GetAvatarWithHash(ctx, userLocator, size, hash)
Get a users avatar



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **size** | **int32**| avatar&#39;s size | 
  **hash** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAvatar**
> PutAvatar(ctx, userLocator, optional)
Update a users avatar



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***PutAvatarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutAvatarOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **avatar** | **optional.Interface of *os.File**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

