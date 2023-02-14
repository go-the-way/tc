# \BuildApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBuildVcsLabel**](BuildApi.md#AddBuildVcsLabel) | **Post** /app/rest/builds/{buildLocator}/vcsLabels | Add a VCS label to the matching build.
[**AddLogMessageToBuild**](BuildApi.md#AddLogMessageToBuild) | **Post** /app/rest/builds/{buildLocator}/log | Adds a message to the build log. Service messages are accepted.
[**AddProblemToBuild**](BuildApi.md#AddProblemToBuild) | **Post** /app/rest/builds/{buildLocator}/problemOccurrences | Add a build problem to the matching build.
[**AddTagsToBuild**](BuildApi.md#AddTagsToBuild) | **Post** /app/rest/builds/{buildLocator}/tags | Add tags to the matching build.
[**AddTagsToMultipleBuilds**](BuildApi.md#AddTagsToMultipleBuilds) | **Post** /app/rest/builds/multiple/{buildLocator}/tags | Add tags to multiple matching builds.
[**CancelBuild**](BuildApi.md#CancelBuild) | **Post** /app/rest/builds/{buildLocator} | cancelBuild
[**CancelMultiple**](BuildApi.md#CancelMultiple) | **Post** /app/rest/builds/multiple/{buildLocator} | cancelMultipleBuilds
[**DeleteBuild**](BuildApi.md#DeleteBuild) | **Delete** /app/rest/builds/{buildLocator} | Delete build matching the locator.
[**DeleteBuildComment**](BuildApi.md#DeleteBuildComment) | **Delete** /app/rest/builds/{buildLocator}/comment | Remove the build comment matching the locator.
[**DeleteMultipleBuildComments**](BuildApi.md#DeleteMultipleBuildComments) | **Delete** /app/rest/builds/multiple/{buildLocator}/comment | Delete comments of multiple matching builds.
[**DeleteMultipleBuilds**](BuildApi.md#DeleteMultipleBuilds) | **Delete** /app/rest/builds/multiple/{buildLocator} | Delete multiple builds matching the locator.
[**DownloadFileOfBuild**](BuildApi.md#DownloadFileOfBuild) | **Get** /app/rest/builds/{buildLocator}/artifacts/files{path} | Download specific file.
[**GetAggregatedBuildStatus**](BuildApi.md#GetAggregatedBuildStatus) | **Get** /app/rest/builds/aggregated/{buildLocator}/status | Get the build status of aggregated matching builds.
[**GetAggregatedBuildStatusIcon**](BuildApi.md#GetAggregatedBuildStatusIcon) | **Get** /app/rest/builds/aggregated/{buildLocator}/statusIcon{suffix} | Get the status icon (in specified format) of aggregated matching builds.
[**GetAllBuilds**](BuildApi.md#GetAllBuilds) | **Get** /app/rest/builds | Get all builds.
[**GetArtifactDependencyChanges**](BuildApi.md#GetArtifactDependencyChanges) | **Get** /app/rest/builds/{buildLocator}/artifactDependencyChanges | Get artifact dependency changes of the matching build.
[**GetArtifactsDirectory**](BuildApi.md#GetArtifactsDirectory) | **Get** /app/rest/builds/{buildLocator}/artifactsDirectory | Get the artifacts&#39; directory of the matching build.
[**GetBuild**](BuildApi.md#GetBuild) | **Get** /app/rest/builds/{buildLocator} | Get build matching the locator.
[**GetBuildActualParameters**](BuildApi.md#GetBuildActualParameters) | **Get** /app/rest/builds/{buildLocator}/resulting-properties | Get actual build parameters of the matching build.
[**GetBuildField**](BuildApi.md#GetBuildField) | **Get** /app/rest/builds/{buildLocator}/{field} | Get a field of the matching build.
[**GetBuildFinishDate**](BuildApi.md#GetBuildFinishDate) | **Get** /app/rest/builds/{buildLocator}/finishDate | Get the finish date of the matching build.
[**GetBuildNumber**](BuildApi.md#GetBuildNumber) | **Get** /app/rest/builds/{buildLocator}/number | Get the number of the matching build.
[**GetBuildPinInfo**](BuildApi.md#GetBuildPinInfo) | **Get** /app/rest/builds/{buildLocator}/pinInfo | Check if the matching build is pinned.
[**GetBuildProblems**](BuildApi.md#GetBuildProblems) | **Get** /app/rest/builds/{buildLocator}/problemOccurrences | Get build problems of the matching build.
[**GetBuildRelatedIssues**](BuildApi.md#GetBuildRelatedIssues) | **Get** /app/rest/builds/{buildLocator}/relatedIssues | Get related issues of the matching build.
[**GetBuildResolved**](BuildApi.md#GetBuildResolved) | **Get** /app/rest/builds/{buildLocator}/resolved/{value} | Get the resolvement status of the matching build.
[**GetBuildResultingProperties**](BuildApi.md#GetBuildResultingProperties) | **Get** /app/rest/builds/{buildLocator}/resulting-properties/{propertyName} | Update a build parameter of the matching build.
[**GetBuildSourceFile**](BuildApi.md#GetBuildSourceFile) | **Get** /app/rest/builds/{buildLocator}/sources/files/{fileName} | Get a source file of the matching build.
[**GetBuildStatisticValue**](BuildApi.md#GetBuildStatisticValue) | **Get** /app/rest/builds/{buildLocator}/statistics/{name} | Get a statistical value of the matching build.
[**GetBuildStatisticValues**](BuildApi.md#GetBuildStatisticValues) | **Get** /app/rest/builds/{buildLocator}/statistics | Get all statistical values of the matching build.
[**GetBuildStatus**](BuildApi.md#GetBuildStatus) | **Get** /app/rest/builds/{buildLocator}/status | Get status of the matching build.
[**GetBuildStatusIcon**](BuildApi.md#GetBuildStatusIcon) | **Get** /app/rest/builds/{buildLocator}/statusIcon{suffix} | Get the status icon (in specified format) of the matching build.
[**GetBuildStatusText**](BuildApi.md#GetBuildStatusText) | **Get** /app/rest/builds/{buildLocator}/statusText | Get the build status text of the matching build.
[**GetBuildTags**](BuildApi.md#GetBuildTags) | **Get** /app/rest/builds/{buildLocator}/tags | Get tags of the matching build.
[**GetBuildTestOccurrences**](BuildApi.md#GetBuildTestOccurrences) | **Get** /app/rest/builds/{buildLocator}/testOccurrences | Get test occurrences of the matching build.
[**GetBuildVcsLabels**](BuildApi.md#GetBuildVcsLabels) | **Get** /app/rest/builds/{buildLocator}/vcsLabels | Get VCS labels of the matching build.
[**GetCanceledInfo**](BuildApi.md#GetCanceledInfo) | **Get** /app/rest/builds/{buildLocator}/canceledInfo | Check if the matching build is canceled.
[**GetFileMetadataOfBuild**](BuildApi.md#GetFileMetadataOfBuild) | **Get** /app/rest/builds/{buildLocator}/artifacts/metadata{path} | Get metadata of specific file.
[**GetFilesListForSubpathOfBuild**](BuildApi.md#GetFilesListForSubpathOfBuild) | **Get** /app/rest/builds/{buildLocator}/artifacts/{path} | List files under this path.
[**GetFilesListOfBuild**](BuildApi.md#GetFilesListOfBuild) | **Get** /app/rest/builds/{buildLocator}/artifacts | List all files.
[**GetMultipleBuilds**](BuildApi.md#GetMultipleBuilds) | **Get** /app/rest/builds/multiple/{buildLocator} | Get multiple builds matching the locator.
[**GetZippedFileOfBuild**](BuildApi.md#GetZippedFileOfBuild) | **Get** /app/rest/builds/{buildLocator}/artifacts/archived{path} | Get specific file zipped.
[**MarkBuildAsRunning**](BuildApi.md#MarkBuildAsRunning) | **Put** /app/rest/builds/{buildLocator}/runningData | Starts the queued build as an agent-less build and returns the corresponding running build.
[**PinMultipleBuilds**](BuildApi.md#PinMultipleBuilds) | **Put** /app/rest/builds/multiple/{buildLocator}/pinInfo | Pin multiple matching builds.
[**RemoveMultipleBuildTags**](BuildApi.md#RemoveMultipleBuildTags) | **Delete** /app/rest/builds/multiple/{buildLocator}/tags | Remove tags from multiple matching builds.
[**ResetBuildFinishProperties**](BuildApi.md#ResetBuildFinishProperties) | **Delete** /app/rest/builds/{buildLocator}/caches/finishProperties | Remove build parameters from the matching build.
[**SetBuildComment**](BuildApi.md#SetBuildComment) | **Put** /app/rest/builds/{buildLocator}/comment | Update the comment on the matching build.
[**SetBuildFinishDate**](BuildApi.md#SetBuildFinishDate) | **Put** /app/rest/builds/{buildLocator}/finishDate | Marks the running build as finished by passing agent the current time of the build to finish.
[**SetBuildNumber**](BuildApi.md#SetBuildNumber) | **Put** /app/rest/builds/{buildLocator}/number | Update the number of the matching build.
[**SetBuildPinInfo**](BuildApi.md#SetBuildPinInfo) | **Put** /app/rest/builds/{buildLocator}/pinInfo | Update the pin info of the matching build.
[**SetBuildStatus**](BuildApi.md#SetBuildStatus) | **Post** /app/rest/builds/{buildLocator}/status | Change status of the build.
[**SetBuildStatusText**](BuildApi.md#SetBuildStatusText) | **Put** /app/rest/builds/{buildLocator}/statusText | Update the build status of the matching build.
[**SetBuildTags**](BuildApi.md#SetBuildTags) | **Put** /app/rest/builds/{buildLocator}/tags | Update tags of the matching build.
[**SetFinishedTime**](BuildApi.md#SetFinishedTime) | **Put** /app/rest/builds/{buildLocator}/finish | Marks the running build as finished by passing agent the current time of the build to finish.
[**SetMultipleBuildComments**](BuildApi.md#SetMultipleBuildComments) | **Put** /app/rest/builds/multiple/{buildLocator}/comment | Update comments in multiple matching builds.


# **AddBuildVcsLabel**
> VcsLabels AddBuildVcsLabel(ctx, buildLocator, optional)
Add a VCS label to the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***AddBuildVcsLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddBuildVcsLabelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **body** | **optional.String**|  | 

### Return type

[**VcsLabels**](vcsLabels.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddLogMessageToBuild**
> AddLogMessageToBuild(ctx, buildLocator, optional)
Adds a message to the build log. Service messages are accepted.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***AddLogMessageToBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddLogMessageToBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddProblemToBuild**
> ProblemOccurrence AddProblemToBuild(ctx, buildLocator, optional)
Add a build problem to the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***AddProblemToBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddProblemToBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**ProblemOccurrence**](problemOccurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTagsToBuild**
> Tags AddTagsToBuild(ctx, buildLocator, optional)
Add tags to the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***AddTagsToBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddTagsToBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Tags**](Tags.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Tags**](tags.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTagsToMultipleBuilds**
> MultipleOperationResult AddTagsToMultipleBuilds(ctx, buildLocator, optional)
Add tags to multiple matching builds.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***AddTagsToMultipleBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddTagsToMultipleBuildsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Tags**](Tags.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**MultipleOperationResult**](multipleOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelBuild**
> Build CancelBuild(ctx, buildLocator, optional)
cancelBuild



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***CancelBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CancelBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BuildCancelRequest**](BuildCancelRequest.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Build**](build.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelMultiple**
> MultipleOperationResult CancelMultiple(ctx, buildLocator, optional)
cancelMultipleBuilds



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***CancelMultipleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CancelMultipleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BuildCancelRequest**](BuildCancelRequest.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**MultipleOperationResult**](multipleOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBuild**
> DeleteBuild(ctx, buildLocator)
Delete build matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBuildComment**
> DeleteBuildComment(ctx, buildLocator)
Remove the build comment matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMultipleBuildComments**
> MultipleOperationResult DeleteMultipleBuildComments(ctx, buildLocator, optional)
Delete comments of multiple matching builds.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***DeleteMultipleBuildCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteMultipleBuildCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**MultipleOperationResult**](multipleOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMultipleBuilds**
> MultipleOperationResult DeleteMultipleBuilds(ctx, buildLocator, optional)
Delete multiple builds matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***DeleteMultipleBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteMultipleBuildsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**MultipleOperationResult**](multipleOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFileOfBuild**
> DownloadFileOfBuild(ctx, path, buildLocator, optional)
Download specific file.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **buildLocator** | **string**|  | 
 **optional** | ***DownloadFileOfBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DownloadFileOfBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resolveParameters** | **optional.Bool**|  | 
 **logBuildUsage** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedBuildStatus**
> string GetAggregatedBuildStatus(ctx, buildLocator)
Get the build status of aggregated matching builds.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedBuildStatusIcon**
> GetAggregatedBuildStatusIcon(ctx, buildLocator, suffix)
Get the status icon (in specified format) of aggregated matching builds.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
  **suffix** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllBuilds**
> Builds GetAllBuilds(ctx, optional)
Get all builds.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBuildsOpts struct

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

# **GetArtifactDependencyChanges**
> BuildChanges GetArtifactDependencyChanges(ctx, buildLocator, optional)
Get artifact dependency changes of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetArtifactDependencyChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetArtifactDependencyChangesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**BuildChanges**](buildChanges.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactsDirectory**
> string GetArtifactsDirectory(ctx, buildLocator)
Get the artifacts' directory of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuild**
> Build GetBuild(ctx, buildLocator, optional)
Get build matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildOpts struct

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

# **GetBuildActualParameters**
> Properties GetBuildActualParameters(ctx, buildLocator, optional)
Get actual build parameters of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetBuildActualParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildActualParametersOpts struct

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

# **GetBuildField**
> string GetBuildField(ctx, buildLocator, field)
Get a field of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildFinishDate**
> string GetBuildFinishDate(ctx, buildLocator)
Get the finish date of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildNumber**
> string GetBuildNumber(ctx, buildLocator)
Get the number of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildPinInfo**
> PinInfo GetBuildPinInfo(ctx, buildLocator, optional)
Check if the matching build is pinned.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetBuildPinInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildPinInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**PinInfo**](pinInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildProblems**
> ProblemOccurrences GetBuildProblems(ctx, buildLocator, optional)
Get build problems of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetBuildProblemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildProblemsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**ProblemOccurrences**](problemOccurrences.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildRelatedIssues**
> IssuesUsages GetBuildRelatedIssues(ctx, buildLocator, optional)
Get related issues of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetBuildRelatedIssuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildRelatedIssuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**IssuesUsages**](issuesUsages.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildResolved**
> string GetBuildResolved(ctx, buildLocator, value)
Get the resolvement status of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
  **value** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildResultingProperties**
> string GetBuildResultingProperties(ctx, buildLocator, propertyName)
Update a build parameter of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
  **propertyName** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildSourceFile**
> GetBuildSourceFile(ctx, buildLocator, fileName)
Get a source file of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
  **fileName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildStatisticValue**
> string GetBuildStatisticValue(ctx, buildLocator, name)
Get a statistical value of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
  **name** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildStatisticValues**
> Properties GetBuildStatisticValues(ctx, buildLocator, optional)
Get all statistical values of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetBuildStatisticValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildStatisticValuesOpts struct

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

# **GetBuildStatus**
> string GetBuildStatus(ctx, buildLocator)
Get status of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildStatusIcon**
> GetBuildStatusIcon(ctx, buildLocator, suffix)
Get the status icon (in specified format) of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
  **suffix** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildStatusText**
> string GetBuildStatusText(ctx, buildLocator)
Get the build status text of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildTags**
> Tags GetBuildTags(ctx, buildLocator, optional)
Get tags of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetBuildTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildTagsOpts struct

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

# **GetBuildTestOccurrences**
> TestOccurrences GetBuildTestOccurrences(ctx, buildLocator, optional)
Get test occurrences of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetBuildTestOccurrencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildTestOccurrencesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**TestOccurrences**](testOccurrences.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildVcsLabels**
> VcsLabels GetBuildVcsLabels(ctx, buildLocator, optional)
Get VCS labels of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetBuildVcsLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildVcsLabelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**VcsLabels**](vcsLabels.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCanceledInfo**
> Comment GetCanceledInfo(ctx, buildLocator, optional)
Check if the matching build is canceled.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetCanceledInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCanceledInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Comment**](comment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileMetadataOfBuild**
> *os.File GetFileMetadataOfBuild(ctx, path, buildLocator, optional)
Get metadata of specific file.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **buildLocator** | **string**|  | 
 **optional** | ***GetFileMetadataOfBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileMetadataOfBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 
 **resolveParameters** | **optional.Bool**|  | 
 **logBuildUsage** | **optional.Bool**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilesListForSubpathOfBuild**
> Files GetFilesListForSubpathOfBuild(ctx, path, buildLocator, optional)
List files under this path.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **buildLocator** | **string**|  | 
 **optional** | ***GetFilesListForSubpathOfBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFilesListForSubpathOfBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **basePath** | **optional.String**|  | 
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **resolveParameters** | **optional.Bool**|  | 
 **logBuildUsage** | **optional.Bool**|  | 

### Return type

[**Files**](files.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilesListOfBuild**
> Files GetFilesListOfBuild(ctx, buildLocator, optional)
List all files.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetFilesListOfBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFilesListOfBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **basePath** | **optional.String**|  | 
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **resolveParameters** | **optional.Bool**|  | 
 **logBuildUsage** | **optional.Bool**|  | 

### Return type

[**Files**](files.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMultipleBuilds**
> Builds GetMultipleBuilds(ctx, buildLocator, optional)
Get multiple builds matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***GetMultipleBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMultipleBuildsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Builds**](builds.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetZippedFileOfBuild**
> GetZippedFileOfBuild(ctx, path, buildLocator, optional)
Get specific file zipped.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **buildLocator** | **string**|  | 
 **optional** | ***GetZippedFileOfBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetZippedFileOfBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **basePath** | **optional.String**|  | 
 **locator** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **resolveParameters** | **optional.Bool**|  | 
 **logBuildUsage** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkBuildAsRunning**
> Build MarkBuildAsRunning(ctx, buildLocator, optional)
Starts the queued build as an agent-less build and returns the corresponding running build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***MarkBuildAsRunningOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarkBuildAsRunningOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Build**](build.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PinMultipleBuilds**
> MultipleOperationResult PinMultipleBuilds(ctx, buildLocator, optional)
Pin multiple matching builds.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***PinMultipleBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PinMultipleBuildsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PinInfo**](PinInfo.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**MultipleOperationResult**](multipleOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveMultipleBuildTags**
> MultipleOperationResult RemoveMultipleBuildTags(ctx, buildLocator, optional)
Remove tags from multiple matching builds.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***RemoveMultipleBuildTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoveMultipleBuildTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Tags**](Tags.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**MultipleOperationResult**](multipleOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetBuildFinishProperties**
> ResetBuildFinishProperties(ctx, buildLocator)
Remove build parameters from the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBuildComment**
> SetBuildComment(ctx, buildLocator, optional)
Update the comment on the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***SetBuildCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildCommentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBuildFinishDate**
> string SetBuildFinishDate(ctx, buildLocator, optional)
Marks the running build as finished by passing agent the current time of the build to finish.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***SetBuildFinishDateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildFinishDateOpts struct

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

# **SetBuildNumber**
> string SetBuildNumber(ctx, buildLocator, optional)
Update the number of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***SetBuildNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildNumberOpts struct

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

# **SetBuildPinInfo**
> PinInfo SetBuildPinInfo(ctx, buildLocator, optional)
Update the pin info of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***SetBuildPinInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildPinInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PinInfo**](PinInfo.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**PinInfo**](pinInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBuildStatus**
> SetBuildStatus(ctx, buildLocator, optional)
Change status of the build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***SetBuildStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of BuildStatusUpdate**](BuildStatusUpdate.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBuildStatusText**
> string SetBuildStatusText(ctx, buildLocator, optional)
Update the build status of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***SetBuildStatusTextOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildStatusTextOpts struct

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

# **SetBuildTags**
> Tags SetBuildTags(ctx, buildLocator, optional)
Update tags of the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***SetBuildTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locator** | **optional.String**|  | 
 **body** | [**optional.Interface of Tags**](Tags.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Tags**](tags.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetFinishedTime**
> string SetFinishedTime(ctx, buildLocator)
Marks the running build as finished by passing agent the current time of the build to finish.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMultipleBuildComments**
> MultipleOperationResult SetMultipleBuildComments(ctx, buildLocator, optional)
Update comments in multiple matching builds.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildLocator** | **string**|  | 
 **optional** | ***SetMultipleBuildCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetMultipleBuildCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**MultipleOperationResult**](multipleOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

