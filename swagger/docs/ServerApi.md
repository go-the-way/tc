# \ServerApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLicenseKeys**](ServerApi.md#AddLicenseKeys) | **Post** /app/rest/server/licensingData/licenseKeys | Add license keys.
[**DeleteLicenseKey**](ServerApi.md#DeleteLicenseKey) | **Delete** /app/rest/server/licensingData/licenseKeys/{licenseKey} | Delete a license key.
[**DownloadFileOfServer**](ServerApi.md#DownloadFileOfServer) | **Get** /app/rest/server/files/{areaId}/files{path} | Download specific file.
[**GetAllMetrics**](ServerApi.md#GetAllMetrics) | **Get** /app/rest/server/metrics | Get metrics.
[**GetAllPlugins**](ServerApi.md#GetAllPlugins) | **Get** /app/rest/server/plugins | Get all plugins.
[**GetBackupStatus**](ServerApi.md#GetBackupStatus) | **Get** /app/rest/server/backup | Get the latest backup status.
[**GetFileMetadataOfServer**](ServerApi.md#GetFileMetadataOfServer) | **Get** /app/rest/server/files/{areaId}/metadata{path} | Get metadata of specific file.
[**GetFilesListForSubpathOfServer**](ServerApi.md#GetFilesListForSubpathOfServer) | **Get** /app/rest/server/files/{areaId}/{path} | List files under this path.
[**GetFilesListOfServer**](ServerApi.md#GetFilesListOfServer) | **Get** /app/rest/server/files/{areaId} | List all files.
[**GetLicenseKey**](ServerApi.md#GetLicenseKey) | **Get** /app/rest/server/licensingData/licenseKeys/{licenseKey} | Get a license key.
[**GetLicenseKeys**](ServerApi.md#GetLicenseKeys) | **Get** /app/rest/server/licensingData/licenseKeys | Get all license keys.
[**GetLicensingData**](ServerApi.md#GetLicensingData) | **Get** /app/rest/server/licensingData | Get the licensing data.
[**GetServerField**](ServerApi.md#GetServerField) | **Get** /app/rest/server/{field} | Get a field of the server info.
[**GetServerInfo**](ServerApi.md#GetServerInfo) | **Get** /app/rest/server | Get the server info.
[**GetZippedFileOfServer**](ServerApi.md#GetZippedFileOfServer) | **Get** /app/rest/server/files/{areaId}/archived{path} | Get specific file zipped.
[**StartBackup**](ServerApi.md#StartBackup) | **Post** /app/rest/server/backup | Start a new backup.


# **AddLicenseKeys**
> LicenseKeys AddLicenseKeys(ctx, optional)
Add license keys.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddLicenseKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddLicenseKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**LicenseKeys**](licenseKeys.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLicenseKey**
> DeleteLicenseKey(ctx, licenseKey)
Delete a license key.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **licenseKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFileOfServer**
> DownloadFileOfServer(ctx, path, areaId)
Download specific file.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **areaId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllMetrics**
> Metrics GetAllMetrics(ctx, optional)
Get metrics.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllMetricsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 

### Return type

[**Metrics**](metrics.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPlugins**
> Plugins GetAllPlugins(ctx, optional)
Get all plugins.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllPluginsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllPluginsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 

### Return type

[**Plugins**](plugins.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupStatus**
> string GetBackupStatus(ctx, )
Get the latest backup status.



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

# **GetFileMetadataOfServer**
> *os.File GetFileMetadataOfServer(ctx, path, areaId, optional)
Get metadata of specific file.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **areaId** | **string**|  | 
 **optional** | ***GetFileMetadataOfServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileMetadataOfServerOpts struct

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

# **GetFilesListForSubpathOfServer**
> Files GetFilesListForSubpathOfServer(ctx, path, areaId, optional)
List files under this path.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **areaId** | **string**|  | 
 **optional** | ***GetFilesListForSubpathOfServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFilesListForSubpathOfServerOpts struct

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

# **GetFilesListOfServer**
> Files GetFilesListOfServer(ctx, areaId, optional)
List all files.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **areaId** | **string**|  | 
 **optional** | ***GetFilesListOfServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFilesListOfServerOpts struct

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

# **GetLicenseKey**
> LicenseKey GetLicenseKey(ctx, licenseKey, optional)
Get a license key.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **licenseKey** | **string**|  | 
 **optional** | ***GetLicenseKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetLicenseKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**LicenseKey**](licenseKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseKeys**
> LicenseKeys GetLicenseKeys(ctx, optional)
Get all license keys.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLicenseKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetLicenseKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 

### Return type

[**LicenseKeys**](licenseKeys.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicensingData**
> LicensingData GetLicensingData(ctx, optional)
Get the licensing data.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLicensingDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetLicensingDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 

### Return type

[**LicensingData**](licensingData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerField**
> string GetServerField(ctx, field)
Get a field of the server info.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **field** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerInfo**
> Server GetServerInfo(ctx, optional)
Get the server info.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetServerInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetServerInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 

### Return type

[**Server**](server.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetZippedFileOfServer**
> GetZippedFileOfServer(ctx, path, areaId, optional)
Get specific file zipped.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **areaId** | **string**|  | 
 **optional** | ***GetZippedFileOfServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetZippedFileOfServerOpts struct

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

# **StartBackup**
> string StartBackup(ctx, optional)
Start a new backup.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StartBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StartBackupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileName** | **optional.String**|  | 
 **addTimestamp** | **optional.Bool**|  | 
 **includeConfigs** | **optional.Bool**|  | 
 **includeDatabase** | **optional.Bool**|  | 
 **includeBuildLogs** | **optional.Bool**|  | 
 **includePersonalChanges** | **optional.Bool**|  | 
 **includeRunningBuilds** | **optional.Bool**|  | 
 **includeSupplimentaryData** | **optional.Bool**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

