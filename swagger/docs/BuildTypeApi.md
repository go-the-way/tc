# \BuildTypeApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAgentRequirementToBuildType**](BuildTypeApi.md#AddAgentRequirementToBuildType) | **Post** /app/rest/buildTypes/{btLocator}/agent-requirements | Add an agent requirement to the matching build configuration.
[**AddArtifactDependencyToBuildType**](BuildTypeApi.md#AddArtifactDependencyToBuildType) | **Post** /app/rest/buildTypes/{btLocator}/artifact-dependencies | Add an artifact dependency to the matching build configuration.
[**AddBuildFeatureToBuildType**](BuildTypeApi.md#AddBuildFeatureToBuildType) | **Post** /app/rest/buildTypes/{btLocator}/features | Add build feature to the matching build configuration.
[**AddBuildStepToBuildType**](BuildTypeApi.md#AddBuildStepToBuildType) | **Post** /app/rest/buildTypes/{btLocator}/steps | Add a build step to the matching build configuration.
[**AddBuildTemplate**](BuildTypeApi.md#AddBuildTemplate) | **Post** /app/rest/buildTypes/{btLocator}/templates | Add a build template to the matching build configuration.
[**AddParameterToBuildFeature**](BuildTypeApi.md#AddParameterToBuildFeature) | **Put** /app/rest/buildTypes/{btLocator}/features/{featureId}/parameters/{parameterName} | Update build feature parameter for the matching build configuration.
[**AddParameterToBuildStep**](BuildTypeApi.md#AddParameterToBuildStep) | **Put** /app/rest/buildTypes/{btLocator}/steps/{stepId}/parameters/{parameterName} | Add a parameter to a build step of the matching build configuration.
[**AddSnapshotDependencyToBuildType**](BuildTypeApi.md#AddSnapshotDependencyToBuildType) | **Post** /app/rest/buildTypes/{btLocator}/snapshot-dependencies | Add a snapshot dependency to the matching build configuration.
[**AddTriggerToBuildType**](BuildTypeApi.md#AddTriggerToBuildType) | **Post** /app/rest/buildTypes/{btLocator}/triggers | Add a trigger to the matching build configuration.
[**AddVcsRootToBuildType**](BuildTypeApi.md#AddVcsRootToBuildType) | **Post** /app/rest/buildTypes/{btLocator}/vcs-root-entries | Add a VCS root to the matching build.
[**CreateBuildParameterOfBuildType**](BuildTypeApi.md#CreateBuildParameterOfBuildType) | **Post** /app/rest/buildTypes/{btLocator}/parameters | Create a build parameter.
[**CreateBuildType**](BuildTypeApi.md#CreateBuildType) | **Post** /app/rest/buildTypes | Create a new build configuration.
[**DeleteAgentRequirement**](BuildTypeApi.md#DeleteAgentRequirement) | **Delete** /app/rest/buildTypes/{btLocator}/agent-requirements/{agentRequirementLocator} | Remove an agent requirement of the matching build configuration.
[**DeleteArtifactDependency**](BuildTypeApi.md#DeleteArtifactDependency) | **Delete** /app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator} | Remove an artifact dependency from the matching build configuration.
[**DeleteBuildParameterOfBuildType**](BuildTypeApi.md#DeleteBuildParameterOfBuildType) | **Delete** /app/rest/buildTypes/{btLocator}/parameters/{name} | Delete build parameter.
[**DeleteBuildParametersOfBuildType**](BuildTypeApi.md#DeleteBuildParametersOfBuildType) | **Delete** /app/rest/buildTypes/{btLocator}/parameters | Delete all build parameters.
[**DeleteBuildStep**](BuildTypeApi.md#DeleteBuildStep) | **Delete** /app/rest/buildTypes/{btLocator}/steps/{stepId} | Delete a build step of the matching build configuration.
[**DeleteBuildStepParameters**](BuildTypeApi.md#DeleteBuildStepParameters) | **Put** /app/rest/buildTypes/{btLocator}/steps/{stepId}/parameters | Update a parameter of a build step of the matching build configuration.
[**DeleteBuildType**](BuildTypeApi.md#DeleteBuildType) | **Delete** /app/rest/buildTypes/{btLocator} | Delete build configuration matching the locator.
[**DeleteFeatureOfBuildType**](BuildTypeApi.md#DeleteFeatureOfBuildType) | **Delete** /app/rest/buildTypes/{btLocator}/features/{featureId} | Remove a build feature of the matching build configuration.
[**DeleteSnapshotDependency**](BuildTypeApi.md#DeleteSnapshotDependency) | **Delete** /app/rest/buildTypes/{btLocator}/snapshot-dependencies/{snapshotDepLocator} | Delete a snapshot dependency of the matching build configuration.
[**DeleteTrigger**](BuildTypeApi.md#DeleteTrigger) | **Delete** /app/rest/buildTypes/{btLocator}/triggers/{triggerLocator} | Delete a trigger of the matching build configuration.
[**DeleteVcsRootOfBuildType**](BuildTypeApi.md#DeleteVcsRootOfBuildType) | **Delete** /app/rest/buildTypes/{btLocator}/vcs-root-entries/{vcsRootLocator} | Remove a VCS root of the matching build configuration.
[**DownloadFileOfBuildType**](BuildTypeApi.md#DownloadFileOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/vcs/files/latest/files{path} | Download specific file.
[**GetAgentRequirement**](BuildTypeApi.md#GetAgentRequirement) | **Get** /app/rest/buildTypes/{btLocator}/agent-requirements/{agentRequirementLocator} | Get an agent requirement of the matching build configuration.
[**GetAgentRequirementParameter**](BuildTypeApi.md#GetAgentRequirementParameter) | **Get** /app/rest/buildTypes/{btLocator}/agent-requirements/{agentRequirementLocator}/{fieldName} | Get a setting of an agent requirement of the matching build configuration.
[**GetAliases**](BuildTypeApi.md#GetAliases) | **Get** /app/rest/buildTypes/{btLocator}/aliases | Get external IDs of the matching build configuration.
[**GetAllAgentRequirements**](BuildTypeApi.md#GetAllAgentRequirements) | **Get** /app/rest/buildTypes/{btLocator}/agent-requirements | Get all agent requirements of the matching build configuration.
[**GetAllArtifactDependencies**](BuildTypeApi.md#GetAllArtifactDependencies) | **Get** /app/rest/buildTypes/{btLocator}/artifact-dependencies | Get all artifact dependencies of the matching build configuration.
[**GetAllBranchesOfBuildType**](BuildTypeApi.md#GetAllBranchesOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/branches | Get all branches of the matching build configuration.
[**GetAllBuildFeatureParameters**](BuildTypeApi.md#GetAllBuildFeatureParameters) | **Get** /app/rest/buildTypes/{btLocator}/features/{featureId}/parameters | Get all parameters of a build feature of the matching build configuration.
[**GetAllBuildFeatures**](BuildTypeApi.md#GetAllBuildFeatures) | **Get** /app/rest/buildTypes/{btLocator}/features | Get all build features of the matching build configuration.
[**GetAllBuildStepParameters**](BuildTypeApi.md#GetAllBuildStepParameters) | **Get** /app/rest/buildTypes/{btLocator}/steps/{stepId}/parameters | Get all parameters of a build step of the matching build configuration.
[**GetAllBuildSteps**](BuildTypeApi.md#GetAllBuildSteps) | **Get** /app/rest/buildTypes/{btLocator}/steps | Get all build steps of the matching build configuration.
[**GetAllBuildTemplates**](BuildTypeApi.md#GetAllBuildTemplates) | **Get** /app/rest/buildTypes/{btLocator}/templates | Get all build templates of the matching build configuration.
[**GetAllBuildTypes**](BuildTypeApi.md#GetAllBuildTypes) | **Get** /app/rest/buildTypes | Get all build configurations.
[**GetAllInvestigationsOfBuildType**](BuildTypeApi.md#GetAllInvestigationsOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/investigations | Get all investigations of the matching build configuration.
[**GetAllSnapshotDependencies**](BuildTypeApi.md#GetAllSnapshotDependencies) | **Get** /app/rest/buildTypes/{btLocator}/snapshot-dependencies | Get all snapshot dependencies of the matching build configuration.
[**GetAllTriggers**](BuildTypeApi.md#GetAllTriggers) | **Get** /app/rest/buildTypes/{btLocator}/triggers | Get all triggers of the matching build configuration.
[**GetAllVcsRootsOfBuildType**](BuildTypeApi.md#GetAllVcsRootsOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/vcs-root-entries | Get all VCS roots of the matching build configuration.
[**GetArtifactDependency**](BuildTypeApi.md#GetArtifactDependency) | **Get** /app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator} | Get an artifact dependency of the matching build configuration.
[**GetArtifactDependencyParameter**](BuildTypeApi.md#GetArtifactDependencyParameter) | **Get** /app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator}/{fieldName} | Get a parameter of an artifact dependency of the matching build configuration.
[**GetBuildFeature**](BuildTypeApi.md#GetBuildFeature) | **Get** /app/rest/buildTypes/{btLocator}/features/{featureId} | Get a build feature of the matching build configuration.
[**GetBuildFeatureParameter**](BuildTypeApi.md#GetBuildFeatureParameter) | **Get** /app/rest/buildTypes/{btLocator}/features/{featureId}/parameters/{parameterName} | Get a parameter of a build feature of the matching build configuration.
[**GetBuildFeatureSetting**](BuildTypeApi.md#GetBuildFeatureSetting) | **Get** /app/rest/buildTypes/{btLocator}/features/{featureId}/{name} | Get the setting of a build feature of the matching build configuration.
[**GetBuildParameterOfBuildType**](BuildTypeApi.md#GetBuildParameterOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/parameters/{name} | Get build parameter.
[**GetBuildParameterSpecificationOfBuildType**](BuildTypeApi.md#GetBuildParameterSpecificationOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/parameters/{name}/type/rawValue | Get build parameter specification.
[**GetBuildParameterTypeOfBuildType**](BuildTypeApi.md#GetBuildParameterTypeOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/parameters/{name}/type | Get type of build parameter.
[**GetBuildParameterValueOfBuildType**](BuildTypeApi.md#GetBuildParameterValueOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/parameters/{name}/value | Get value of build parameter.
[**GetBuildParametersOfBuildType**](BuildTypeApi.md#GetBuildParametersOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/parameters | Get build parameters.
[**GetBuildStep**](BuildTypeApi.md#GetBuildStep) | **Get** /app/rest/buildTypes/{btLocator}/steps/{stepId} | Get a build step of the matching build configuration.
[**GetBuildStepParameter**](BuildTypeApi.md#GetBuildStepParameter) | **Get** /app/rest/buildTypes/{btLocator}/steps/{stepId}/parameters/{parameterName} | Get a parameter of a build step of the matching build configuration.
[**GetBuildStepSetting**](BuildTypeApi.md#GetBuildStepSetting) | **Get** /app/rest/buildTypes/{btLocator}/steps/{stepId}/{fieldName} | Get the setting of a build step of the matching build configuration.
[**GetBuildTemplate**](BuildTypeApi.md#GetBuildTemplate) | **Get** /app/rest/buildTypes/{btLocator}/templates/{templateLocator} | Get a template of the matching build configuration.
[**GetBuildType**](BuildTypeApi.md#GetBuildType) | **Get** /app/rest/buildTypes/{btLocator} | Get build configuration matching the locator.
[**GetBuildTypeBuildTags**](BuildTypeApi.md#GetBuildTypeBuildTags) | **Get** /app/rest/buildTypes/{btLocator}/buildTags | Get tags of builds of the matching build configuration.
[**GetBuildTypeBuilds**](BuildTypeApi.md#GetBuildTypeBuilds) | **Get** /app/rest/buildTypes/{btLocator}/builds | Get builds of the matching build configuration.
[**GetBuildTypeField**](BuildTypeApi.md#GetBuildTypeField) | **Get** /app/rest/buildTypes/{btLocator}/{field} | Get a field of the matching build configuration.
[**GetBuildTypeSettingsFile**](BuildTypeApi.md#GetBuildTypeSettingsFile) | **Get** /app/rest/buildTypes/{btLocator}/settingsFile | Get the settings file of the matching build configuration.
[**GetFileMetadataOfBuildType**](BuildTypeApi.md#GetFileMetadataOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/vcs/files/latest/metadata{path} | Get metadata of specific file.
[**GetFilesListForSubpathOfBuildType**](BuildTypeApi.md#GetFilesListForSubpathOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/vcs/files/latest/{path} | List files under this path.
[**GetFilesListOfBuildType**](BuildTypeApi.md#GetFilesListOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/vcs/files/latest | List all files.
[**GetSnapshotDependency**](BuildTypeApi.md#GetSnapshotDependency) | **Get** /app/rest/buildTypes/{btLocator}/snapshot-dependencies/{snapshotDepLocator} | Get a snapshot dependency of the matching build configuration.
[**GetTrigger**](BuildTypeApi.md#GetTrigger) | **Get** /app/rest/buildTypes/{btLocator}/triggers/{triggerLocator} | Get a trigger of the matching build configuration.
[**GetTriggerParameter**](BuildTypeApi.md#GetTriggerParameter) | **Get** /app/rest/buildTypes/{btLocator}/triggers/{triggerLocator}/{fieldName} | Get a parameter of a trigger of the matching build configuration.
[**GetVcsRoot**](BuildTypeApi.md#GetVcsRoot) | **Get** /app/rest/buildTypes/{btLocator}/vcs-root-entries/{vcsRootLocator} | Get a VCS root of the matching build configuration.
[**GetVcsRootCheckoutRules**](BuildTypeApi.md#GetVcsRootCheckoutRules) | **Get** /app/rest/buildTypes/{btLocator}/vcs-root-entries/{vcsRootLocator}/checkout-rules | Get checkout rules of a VCS root of the matching build configuration.
[**GetVcsRootInstancesOfBuildType**](BuildTypeApi.md#GetVcsRootInstancesOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/vcsRootInstances | Get all VCS root instances of the matching build configuration.
[**GetZippedFileOfBuildType**](BuildTypeApi.md#GetZippedFileOfBuildType) | **Get** /app/rest/buildTypes/{btLocator}/vcs/files/latest/archived{path} | Get specific file zipped.
[**RemoveAllTemplates**](BuildTypeApi.md#RemoveAllTemplates) | **Delete** /app/rest/buildTypes/{btLocator}/templates | Detach all templates from the matching build configuration.
[**RemoveTemplate**](BuildTypeApi.md#RemoveTemplate) | **Delete** /app/rest/buildTypes/{btLocator}/templates/{templateLocator} | Detach a template from the matching build configuration.
[**ReplaceAgentRequirement**](BuildTypeApi.md#ReplaceAgentRequirement) | **Put** /app/rest/buildTypes/{btLocator}/agent-requirements/{agentRequirementLocator} | Update an agent requirement of the matching build configuration.
[**ReplaceAllAgentRequirements**](BuildTypeApi.md#ReplaceAllAgentRequirements) | **Put** /app/rest/buildTypes/{btLocator}/agent-requirements | Update all agent requirements of the matching build configuration.
[**ReplaceAllArtifactDependencies**](BuildTypeApi.md#ReplaceAllArtifactDependencies) | **Put** /app/rest/buildTypes/{btLocator}/artifact-dependencies | Update all artifact dependencies of the matching build configuration.
[**ReplaceAllBuildFeatures**](BuildTypeApi.md#ReplaceAllBuildFeatures) | **Put** /app/rest/buildTypes/{btLocator}/features | Update all build features of the matching build configuration.
[**ReplaceAllBuildSteps**](BuildTypeApi.md#ReplaceAllBuildSteps) | **Put** /app/rest/buildTypes/{btLocator}/steps | Update all build steps of the matching build configuration.
[**ReplaceAllSnapshotDependencies**](BuildTypeApi.md#ReplaceAllSnapshotDependencies) | **Put** /app/rest/buildTypes/{btLocator}/snapshot-dependencies | Update all snapshot dependencies of the matching build configuration.
[**ReplaceAllTriggers**](BuildTypeApi.md#ReplaceAllTriggers) | **Put** /app/rest/buildTypes/{btLocator}/triggers | Update all triggers of the matching build configuration.
[**ReplaceAllVcsRoots**](BuildTypeApi.md#ReplaceAllVcsRoots) | **Put** /app/rest/buildTypes/{btLocator}/vcs-root-entries | Update all VCS roots of the matching build configuration.
[**ReplaceArtifactDependency**](BuildTypeApi.md#ReplaceArtifactDependency) | **Put** /app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator} | Update an artifact dependency of the matching build configuration.
[**ReplaceBuildFeature**](BuildTypeApi.md#ReplaceBuildFeature) | **Put** /app/rest/buildTypes/{btLocator}/features/{featureId} | Update a build feature of the matching build configuration.
[**ReplaceBuildFeatureParameters**](BuildTypeApi.md#ReplaceBuildFeatureParameters) | **Put** /app/rest/buildTypes/{btLocator}/features/{featureId}/parameters | Update a parameter of a build feature of the matching build configuration.
[**ReplaceBuildStep**](BuildTypeApi.md#ReplaceBuildStep) | **Put** /app/rest/buildTypes/{btLocator}/steps/{stepId} | Replace a build step of the matching build configuration.
[**ReplaceSnapshotDependency**](BuildTypeApi.md#ReplaceSnapshotDependency) | **Put** /app/rest/buildTypes/{btLocator}/snapshot-dependencies/{snapshotDepLocator} | Update a snapshot dependency of the matching build configuration.
[**ReplaceTrigger**](BuildTypeApi.md#ReplaceTrigger) | **Put** /app/rest/buildTypes/{btLocator}/triggers/{triggerLocator} | Update a trigger of the matching build configuration.
[**SetAgentRequirementParameter**](BuildTypeApi.md#SetAgentRequirementParameter) | **Put** /app/rest/buildTypes/{btLocator}/agent-requirements/{agentRequirementLocator}/{fieldName} | Update a parameter of an agent requirement of the matching build configuration.
[**SetArtifactDependencyParameter**](BuildTypeApi.md#SetArtifactDependencyParameter) | **Put** /app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator}/{fieldName} | Update a parameter of an artifact dependency of the matching build configuration.
[**SetBuildFeatureParameter**](BuildTypeApi.md#SetBuildFeatureParameter) | **Put** /app/rest/buildTypes/{btLocator}/features/{featureId}/{name} | Update a parameter of a build feature of the matching build configuration.
[**SetBuildStepParameter**](BuildTypeApi.md#SetBuildStepParameter) | **Put** /app/rest/buildTypes/{btLocator}/steps/{stepId}/{fieldName} | Update a parameter of a build step of the matching build configuration.
[**SetBuildTypeField**](BuildTypeApi.md#SetBuildTypeField) | **Put** /app/rest/buildTypes/{btLocator}/{field} | Update a field of the matching build configuration.
[**SetBuildTypeTemplates**](BuildTypeApi.md#SetBuildTypeTemplates) | **Put** /app/rest/buildTypes/{btLocator}/templates | Update all templates of the matching build configuration.
[**SetTriggerParameter**](BuildTypeApi.md#SetTriggerParameter) | **Put** /app/rest/buildTypes/{btLocator}/triggers/{triggerLocator}/{fieldName} | Update a parameter of a trigger of the matching build configuration.
[**UpdateBuildParameterOfBuildType**](BuildTypeApi.md#UpdateBuildParameterOfBuildType) | **Put** /app/rest/buildTypes/{btLocator}/parameters/{name} | Update build parameter.
[**UpdateBuildParameterSpecificationOfBuildType**](BuildTypeApi.md#UpdateBuildParameterSpecificationOfBuildType) | **Put** /app/rest/buildTypes/{btLocator}/parameters/{name}/type/rawValue | Update build parameter specification.
[**UpdateBuildParameterTypeOfBuildType**](BuildTypeApi.md#UpdateBuildParameterTypeOfBuildType) | **Put** /app/rest/buildTypes/{btLocator}/parameters/{name}/type | Update type of build parameter.
[**UpdateBuildParameterValueOfBuildType**](BuildTypeApi.md#UpdateBuildParameterValueOfBuildType) | **Put** /app/rest/buildTypes/{btLocator}/parameters/{name}/value | Update value of build parameter.
[**UpdateBuildParametersOfBuildType**](BuildTypeApi.md#UpdateBuildParametersOfBuildType) | **Put** /app/rest/buildTypes/{btLocator}/parameters | Update build parameters.
[**UpdateBuildTypeVcsRoot**](BuildTypeApi.md#UpdateBuildTypeVcsRoot) | **Put** /app/rest/buildTypes/{btLocator}/vcs-root-entries/{vcsRootLocator} | Update a VCS root of the matching build configuration.
[**UpdateBuildTypeVcsRootCheckoutRules**](BuildTypeApi.md#UpdateBuildTypeVcsRootCheckoutRules) | **Put** /app/rest/buildTypes/{btLocator}/vcs-root-entries/{vcsRootLocator}/checkout-rules | Update checkout rules of a VCS root of the matching build configuration.


# **AddAgentRequirementToBuildType**
> AgentRequirement AddAgentRequirementToBuildType(ctx, btLocator, optional)
Add an agent requirement to the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***AddAgentRequirementToBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddAgentRequirementToBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of AgentRequirement**](AgentRequirement.md)|  | 

### Return type

[**AgentRequirement**](agent-requirement.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddArtifactDependencyToBuildType**
> ArtifactDependency AddArtifactDependencyToBuildType(ctx, btLocator, optional)
Add an artifact dependency to the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***AddArtifactDependencyToBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddArtifactDependencyToBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of ArtifactDependency**](ArtifactDependency.md)|  | 

### Return type

[**ArtifactDependency**](artifact-dependency.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddBuildFeatureToBuildType**
> Feature AddBuildFeatureToBuildType(ctx, btLocator, optional)
Add build feature to the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***AddBuildFeatureToBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddBuildFeatureToBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of Feature**](Feature.md)|  | 

### Return type

[**Feature**](feature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddBuildStepToBuildType**
> Step AddBuildStepToBuildType(ctx, btLocator, optional)
Add a build step to the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***AddBuildStepToBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddBuildStepToBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of Step**](Step.md)|  | 

### Return type

[**Step**](step.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddBuildTemplate**
> BuildType AddBuildTemplate(ctx, btLocator, optional)
Add a build template to the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***AddBuildTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddBuildTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BuildType**](BuildType.md)|  | 
 **optimizeSettings** | **optional.Bool**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**BuildType**](buildType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddParameterToBuildFeature**
> string AddParameterToBuildFeature(ctx, btLocator, featureId, parameterName, optional)
Update build feature parameter for the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **featureId** | **string**|  | 
  **parameterName** | **string**|  | 
 **optional** | ***AddParameterToBuildFeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddParameterToBuildFeatureOpts struct

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

# **AddParameterToBuildStep**
> string AddParameterToBuildStep(ctx, btLocator, stepId, parameterName, optional)
Add a parameter to a build step of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **stepId** | **string**|  | 
  **parameterName** | **string**|  | 
 **optional** | ***AddParameterToBuildStepOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddParameterToBuildStepOpts struct

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

# **AddSnapshotDependencyToBuildType**
> SnapshotDependency AddSnapshotDependencyToBuildType(ctx, btLocator, optional)
Add a snapshot dependency to the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***AddSnapshotDependencyToBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddSnapshotDependencyToBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of SnapshotDependency**](SnapshotDependency.md)|  | 

### Return type

[**SnapshotDependency**](snapshot-dependency.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTriggerToBuildType**
> Trigger AddTriggerToBuildType(ctx, btLocator, optional)
Add a trigger to the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***AddTriggerToBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddTriggerToBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of Trigger**](Trigger.md)|  | 

### Return type

[**Trigger**](trigger.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVcsRootToBuildType**
> VcsRootEntry AddVcsRootToBuildType(ctx, btLocator, optional)
Add a VCS root to the matching build.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***AddVcsRootToBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddVcsRootToBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VcsRootEntry**](VcsRootEntry.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**VcsRootEntry**](vcs-root-entry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBuildParameterOfBuildType**
> Property CreateBuildParameterOfBuildType(ctx, btLocator, optional)
Create a build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***CreateBuildParameterOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateBuildParameterOfBuildTypeOpts struct

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

# **CreateBuildType**
> BuildType CreateBuildType(ctx, optional)
Create a new build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateBuildTypeOpts struct

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

# **DeleteAgentRequirement**
> DeleteAgentRequirement(ctx, btLocator, agentRequirementLocator)
Remove an agent requirement of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **agentRequirementLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArtifactDependency**
> DeleteArtifactDependency(ctx, btLocator, artifactDepLocator)
Remove an artifact dependency from the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **artifactDepLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBuildParameterOfBuildType**
> DeleteBuildParameterOfBuildType(ctx, name, btLocator)
Delete build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **btLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBuildParametersOfBuildType**
> DeleteBuildParametersOfBuildType(ctx, btLocator)
Delete all build parameters.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBuildStep**
> DeleteBuildStep(ctx, btLocator, stepId)
Delete a build step of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **stepId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBuildStepParameters**
> Properties DeleteBuildStepParameters(ctx, btLocator, stepId, optional)
Update a parameter of a build step of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **stepId** | **string**|  | 
 **optional** | ***DeleteBuildStepParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteBuildStepParametersOpts struct

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

# **DeleteBuildType**
> DeleteBuildType(ctx, btLocator)
Delete build configuration matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFeatureOfBuildType**
> DeleteFeatureOfBuildType(ctx, btLocator, featureId)
Remove a build feature of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **featureId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshotDependency**
> DeleteSnapshotDependency(ctx, btLocator, snapshotDepLocator)
Delete a snapshot dependency of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **snapshotDepLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTrigger**
> DeleteTrigger(ctx, btLocator, triggerLocator)
Delete a trigger of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **triggerLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVcsRootOfBuildType**
> DeleteVcsRootOfBuildType(ctx, btLocator, vcsRootLocator)
Remove a VCS root of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **vcsRootLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFileOfBuildType**
> DownloadFileOfBuildType(ctx, path, btLocator, optional)
Download specific file.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **btLocator** | **string**|  | 
 **optional** | ***DownloadFileOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DownloadFileOfBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resolveParameters** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgentRequirement**
> AgentRequirement GetAgentRequirement(ctx, btLocator, agentRequirementLocator, optional)
Get an agent requirement of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **agentRequirementLocator** | **string**|  | 
 **optional** | ***GetAgentRequirementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAgentRequirementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**AgentRequirement**](agent-requirement.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgentRequirementParameter**
> string GetAgentRequirementParameter(ctx, btLocator, agentRequirementLocator, fieldName)
Get a setting of an agent requirement of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **agentRequirementLocator** | **string**|  | 
  **fieldName** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAliases**
> Items GetAliases(ctx, btLocator, optional)
Get external IDs of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAliasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAliasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **field** | **optional.String**|  | 

### Return type

[**Items**](items.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAgentRequirements**
> AgentRequirements GetAllAgentRequirements(ctx, btLocator, optional)
Get all agent requirements of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAllAgentRequirementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllAgentRequirementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**AgentRequirements**](agent-requirements.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllArtifactDependencies**
> ArtifactDependencies GetAllArtifactDependencies(ctx, btLocator, optional)
Get all artifact dependencies of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAllArtifactDependenciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllArtifactDependenciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**ArtifactDependencies**](artifact-dependencies.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllBranchesOfBuildType**
> Branches GetAllBranchesOfBuildType(ctx, btLocator, optional)
Get all branches of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAllBranchesOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBranchesOfBuildTypeOpts struct

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

# **GetAllBuildFeatureParameters**
> Properties GetAllBuildFeatureParameters(ctx, btLocator, featureId, optional)
Get all parameters of a build feature of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **featureId** | **string**|  | 
 **optional** | ***GetAllBuildFeatureParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBuildFeatureParametersOpts struct

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

# **GetAllBuildFeatures**
> Features GetAllBuildFeatures(ctx, btLocator, optional)
Get all build features of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAllBuildFeaturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBuildFeaturesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Features**](features.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllBuildStepParameters**
> Properties GetAllBuildStepParameters(ctx, btLocator, stepId, optional)
Get all parameters of a build step of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **stepId** | **string**|  | 
 **optional** | ***GetAllBuildStepParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBuildStepParametersOpts struct

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

# **GetAllBuildSteps**
> Steps GetAllBuildSteps(ctx, btLocator, optional)
Get all build steps of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAllBuildStepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBuildStepsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Steps**](steps.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllBuildTemplates**
> BuildTypes GetAllBuildTemplates(ctx, btLocator, optional)
Get all build templates of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAllBuildTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBuildTemplatesOpts struct

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

# **GetAllBuildTypes**
> BuildTypes GetAllBuildTypes(ctx, optional)
Get all build configurations.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllBuildTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBuildTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**BuildTypes**](buildTypes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllInvestigationsOfBuildType**
> Investigations GetAllInvestigationsOfBuildType(ctx, btLocator, optional)
Get all investigations of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAllInvestigationsOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllInvestigationsOfBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Investigations**](investigations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSnapshotDependencies**
> SnapshotDependencies GetAllSnapshotDependencies(ctx, btLocator, optional)
Get all snapshot dependencies of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAllSnapshotDependenciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllSnapshotDependenciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**SnapshotDependencies**](snapshot-dependencies.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTriggers**
> Triggers GetAllTriggers(ctx, btLocator, optional)
Get all triggers of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAllTriggersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllTriggersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Triggers**](triggers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllVcsRootsOfBuildType**
> VcsRootEntries GetAllVcsRootsOfBuildType(ctx, btLocator, optional)
Get all VCS roots of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetAllVcsRootsOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllVcsRootsOfBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**VcsRootEntries**](vcs-root-entries.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactDependency**
> ArtifactDependency GetArtifactDependency(ctx, btLocator, artifactDepLocator, optional)
Get an artifact dependency of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **artifactDepLocator** | **string**|  | 
 **optional** | ***GetArtifactDependencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetArtifactDependencyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**ArtifactDependency**](artifact-dependency.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtifactDependencyParameter**
> string GetArtifactDependencyParameter(ctx, btLocator, artifactDepLocator, fieldName)
Get a parameter of an artifact dependency of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **artifactDepLocator** | **string**|  | 
  **fieldName** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildFeature**
> Feature GetBuildFeature(ctx, btLocator, featureId, optional)
Get a build feature of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **featureId** | **string**|  | 
 **optional** | ***GetBuildFeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildFeatureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**Feature**](feature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildFeatureParameter**
> string GetBuildFeatureParameter(ctx, btLocator, featureId, parameterName)
Get a parameter of a build feature of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **featureId** | **string**|  | 
  **parameterName** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildFeatureSetting**
> string GetBuildFeatureSetting(ctx, btLocator, featureId, name)
Get the setting of a build feature of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **featureId** | **string**|  | 
  **name** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildParameterOfBuildType**
> Property GetBuildParameterOfBuildType(ctx, name, btLocator, optional)
Get build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **btLocator** | **string**|  | 
 **optional** | ***GetBuildParameterOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildParameterOfBuildTypeOpts struct

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

# **GetBuildParameterSpecificationOfBuildType**
> string GetBuildParameterSpecificationOfBuildType(ctx, name, btLocator)
Get build parameter specification.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **btLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildParameterTypeOfBuildType**
> ModelType GetBuildParameterTypeOfBuildType(ctx, name, btLocator)
Get type of build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **btLocator** | **string**|  | 

### Return type

[**ModelType**](type.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildParameterValueOfBuildType**
> string GetBuildParameterValueOfBuildType(ctx, name, btLocator)
Get value of build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **btLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildParametersOfBuildType**
> Properties GetBuildParametersOfBuildType(ctx, btLocator, optional)
Get build parameters.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetBuildParametersOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildParametersOfBuildTypeOpts struct

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

# **GetBuildStep**
> Step GetBuildStep(ctx, btLocator, stepId, optional)
Get a build step of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **stepId** | **string**|  | 
 **optional** | ***GetBuildStepOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildStepOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**Step**](step.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildStepParameter**
> string GetBuildStepParameter(ctx, btLocator, stepId, parameterName)
Get a parameter of a build step of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **stepId** | **string**|  | 
  **parameterName** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildStepSetting**
> string GetBuildStepSetting(ctx, btLocator, stepId, fieldName)
Get the setting of a build step of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **stepId** | **string**|  | 
  **fieldName** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildTemplate**
> BuildType GetBuildTemplate(ctx, btLocator, templateLocator, optional)
Get a template of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **templateLocator** | **string**|  | 
 **optional** | ***GetBuildTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildTemplateOpts struct

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

# **GetBuildType**
> BuildType GetBuildType(ctx, btLocator, optional)
Get build configuration matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildTypeOpts struct

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

# **GetBuildTypeBuildTags**
> Tags GetBuildTypeBuildTags(ctx, btLocator, optional)
Get tags of builds of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetBuildTypeBuildTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildTypeBuildTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **field** | **optional.String**|  | 

### Return type

[**Tags**](tags.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildTypeBuilds**
> Builds GetBuildTypeBuilds(ctx, btLocator, optional)
Get builds of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetBuildTypeBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBuildTypeBuildsOpts struct

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

# **GetBuildTypeField**
> string GetBuildTypeField(ctx, btLocator, field)
Get a field of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildTypeSettingsFile**
> string GetBuildTypeSettingsFile(ctx, btLocator)
Get the settings file of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileMetadataOfBuildType**
> *os.File GetFileMetadataOfBuildType(ctx, path, btLocator, optional)
Get metadata of specific file.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **btLocator** | **string**|  | 
 **optional** | ***GetFileMetadataOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileMetadataOfBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 
 **resolveParameters** | **optional.Bool**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilesListForSubpathOfBuildType**
> Files GetFilesListForSubpathOfBuildType(ctx, path, btLocator, optional)
List files under this path.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **btLocator** | **string**|  | 
 **optional** | ***GetFilesListForSubpathOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFilesListForSubpathOfBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **basePath** | **optional.String**|  | 
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **resolveParameters** | **optional.Bool**|  | 

### Return type

[**Files**](files.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilesListOfBuildType**
> Files GetFilesListOfBuildType(ctx, btLocator, optional)
List all files.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetFilesListOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFilesListOfBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **basePath** | **optional.String**|  | 
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 
 **resolveParameters** | **optional.Bool**|  | 

### Return type

[**Files**](files.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotDependency**
> SnapshotDependency GetSnapshotDependency(ctx, btLocator, snapshotDepLocator, optional)
Get a snapshot dependency of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **snapshotDepLocator** | **string**|  | 
 **optional** | ***GetSnapshotDependencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSnapshotDependencyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**SnapshotDependency**](snapshot-dependency.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrigger**
> Trigger GetTrigger(ctx, btLocator, triggerLocator, optional)
Get a trigger of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **triggerLocator** | **string**|  | 
 **optional** | ***GetTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTriggerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**Trigger**](trigger.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTriggerParameter**
> string GetTriggerParameter(ctx, btLocator, triggerLocator, fieldName)
Get a parameter of a trigger of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **triggerLocator** | **string**|  | 
  **fieldName** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRoot**
> VcsRootEntry GetVcsRoot(ctx, btLocator, vcsRootLocator, optional)
Get a VCS root of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **vcsRootLocator** | **string**|  | 
 **optional** | ***GetVcsRootOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetVcsRootOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 

### Return type

[**VcsRootEntry**](vcs-root-entry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRootCheckoutRules**
> string GetVcsRootCheckoutRules(ctx, btLocator, vcsRootLocator)
Get checkout rules of a VCS root of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **vcsRootLocator** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcsRootInstancesOfBuildType**
> VcsRootInstances GetVcsRootInstancesOfBuildType(ctx, btLocator, optional)
Get all VCS root instances of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***GetVcsRootInstancesOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetVcsRootInstancesOfBuildTypeOpts struct

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

# **GetZippedFileOfBuildType**
> GetZippedFileOfBuildType(ctx, path, btLocator, optional)
Get specific file zipped.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **btLocator** | **string**|  | 
 **optional** | ***GetZippedFileOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetZippedFileOfBuildTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **basePath** | **optional.String**|  | 
 **locator** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **resolveParameters** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAllTemplates**
> RemoveAllTemplates(ctx, btLocator, optional)
Detach all templates from the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***RemoveAllTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoveAllTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineSettings** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTemplate**
> RemoveTemplate(ctx, btLocator, templateLocator, optional)
Detach a template from the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **templateLocator** | **string**|  | 
 **optional** | ***RemoveTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoveTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineSettings** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAgentRequirement**
> AgentRequirement ReplaceAgentRequirement(ctx, btLocator, agentRequirementLocator, optional)
Update an agent requirement of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **agentRequirementLocator** | **string**|  | 
 **optional** | ***ReplaceAgentRequirementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceAgentRequirementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of AgentRequirement**](AgentRequirement.md)|  | 

### Return type

[**AgentRequirement**](agent-requirement.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllAgentRequirements**
> AgentRequirements ReplaceAllAgentRequirements(ctx, btLocator, optional)
Update all agent requirements of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***ReplaceAllAgentRequirementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceAllAgentRequirementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of AgentRequirements**](AgentRequirements.md)|  | 

### Return type

[**AgentRequirements**](agent-requirements.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllArtifactDependencies**
> ArtifactDependencies ReplaceAllArtifactDependencies(ctx, btLocator, optional)
Update all artifact dependencies of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***ReplaceAllArtifactDependenciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceAllArtifactDependenciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of ArtifactDependencies**](ArtifactDependencies.md)|  | 

### Return type

[**ArtifactDependencies**](artifact-dependencies.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllBuildFeatures**
> Features ReplaceAllBuildFeatures(ctx, btLocator, optional)
Update all build features of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***ReplaceAllBuildFeaturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceAllBuildFeaturesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of Features**](Features.md)|  | 

### Return type

[**Features**](features.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllBuildSteps**
> Steps ReplaceAllBuildSteps(ctx, btLocator, optional)
Update all build steps of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***ReplaceAllBuildStepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceAllBuildStepsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of Steps**](Steps.md)|  | 

### Return type

[**Steps**](steps.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllSnapshotDependencies**
> SnapshotDependencies ReplaceAllSnapshotDependencies(ctx, btLocator, optional)
Update all snapshot dependencies of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***ReplaceAllSnapshotDependenciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceAllSnapshotDependenciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of SnapshotDependencies**](SnapshotDependencies.md)|  | 

### Return type

[**SnapshotDependencies**](snapshot-dependencies.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllTriggers**
> Triggers ReplaceAllTriggers(ctx, btLocator, optional)
Update all triggers of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***ReplaceAllTriggersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceAllTriggersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of Triggers**](Triggers.md)|  | 

### Return type

[**Triggers**](triggers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllVcsRoots**
> VcsRootEntries ReplaceAllVcsRoots(ctx, btLocator, optional)
Update all VCS roots of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***ReplaceAllVcsRootsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceAllVcsRootsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VcsRootEntries**](VcsRootEntries.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**VcsRootEntries**](vcs-root-entries.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceArtifactDependency**
> ArtifactDependency ReplaceArtifactDependency(ctx, btLocator, artifactDepLocator, optional)
Update an artifact dependency of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **artifactDepLocator** | **string**|  | 
 **optional** | ***ReplaceArtifactDependencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceArtifactDependencyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of ArtifactDependency**](ArtifactDependency.md)|  | 

### Return type

[**ArtifactDependency**](artifact-dependency.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceBuildFeature**
> Feature ReplaceBuildFeature(ctx, btLocator, featureId, optional)
Update a build feature of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **featureId** | **string**|  | 
 **optional** | ***ReplaceBuildFeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceBuildFeatureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of Feature**](Feature.md)|  | 

### Return type

[**Feature**](feature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceBuildFeatureParameters**
> Properties ReplaceBuildFeatureParameters(ctx, btLocator, featureId, optional)
Update a parameter of a build feature of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **featureId** | **string**|  | 
 **optional** | ***ReplaceBuildFeatureParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceBuildFeatureParametersOpts struct

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

# **ReplaceBuildStep**
> Step ReplaceBuildStep(ctx, btLocator, stepId, optional)
Replace a build step of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **stepId** | **string**|  | 
 **optional** | ***ReplaceBuildStepOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceBuildStepOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of Step**](Step.md)|  | 

### Return type

[**Step**](step.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSnapshotDependency**
> SnapshotDependency ReplaceSnapshotDependency(ctx, btLocator, snapshotDepLocator, optional)
Update a snapshot dependency of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **snapshotDepLocator** | **string**|  | 
 **optional** | ***ReplaceSnapshotDependencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceSnapshotDependencyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of SnapshotDependency**](SnapshotDependency.md)|  | 

### Return type

[**SnapshotDependency**](snapshot-dependency.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceTrigger**
> Trigger ReplaceTrigger(ctx, btLocator, triggerLocator, optional)
Update a trigger of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **triggerLocator** | **string**|  | 
 **optional** | ***ReplaceTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceTriggerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**|  | 
 **body** | [**optional.Interface of Trigger**](Trigger.md)|  | 

### Return type

[**Trigger**](trigger.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAgentRequirementParameter**
> string SetAgentRequirementParameter(ctx, btLocator, agentRequirementLocator, fieldName, optional)
Update a parameter of an agent requirement of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **agentRequirementLocator** | **string**|  | 
  **fieldName** | **string**|  | 
 **optional** | ***SetAgentRequirementParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetAgentRequirementParameterOpts struct

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

# **SetArtifactDependencyParameter**
> string SetArtifactDependencyParameter(ctx, btLocator, artifactDepLocator, fieldName, optional)
Update a parameter of an artifact dependency of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **artifactDepLocator** | **string**|  | 
  **fieldName** | **string**|  | 
 **optional** | ***SetArtifactDependencyParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetArtifactDependencyParameterOpts struct

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

# **SetBuildFeatureParameter**
> string SetBuildFeatureParameter(ctx, btLocator, featureId, name, optional)
Update a parameter of a build feature of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **featureId** | **string**|  | 
  **name** | **string**|  | 
 **optional** | ***SetBuildFeatureParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildFeatureParameterOpts struct

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

# **SetBuildStepParameter**
> string SetBuildStepParameter(ctx, btLocator, stepId, fieldName, optional)
Update a parameter of a build step of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **stepId** | **string**|  | 
  **fieldName** | **string**|  | 
 **optional** | ***SetBuildStepParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildStepParameterOpts struct

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

# **SetBuildTypeField**
> string SetBuildTypeField(ctx, btLocator, field, optional)
Update a field of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***SetBuildTypeFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildTypeFieldOpts struct

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

# **SetBuildTypeTemplates**
> BuildTypes SetBuildTypeTemplates(ctx, btLocator, optional)
Update all templates of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***SetBuildTypeTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetBuildTypeTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BuildTypes**](BuildTypes.md)|  | 
 **optimizeSettings** | **optional.Bool**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**BuildTypes**](buildTypes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetTriggerParameter**
> string SetTriggerParameter(ctx, btLocator, triggerLocator, fieldName, optional)
Update a parameter of a trigger of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **triggerLocator** | **string**|  | 
  **fieldName** | **string**|  | 
 **optional** | ***SetTriggerParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetTriggerParameterOpts struct

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

# **UpdateBuildParameterOfBuildType**
> Property UpdateBuildParameterOfBuildType(ctx, name, btLocator, optional)
Update build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **btLocator** | **string**|  | 
 **optional** | ***UpdateBuildParameterOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildParameterOfBuildTypeOpts struct

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

# **UpdateBuildParameterSpecificationOfBuildType**
> string UpdateBuildParameterSpecificationOfBuildType(ctx, name, btLocator, optional)
Update build parameter specification.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **btLocator** | **string**|  | 
 **optional** | ***UpdateBuildParameterSpecificationOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildParameterSpecificationOfBuildTypeOpts struct

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

# **UpdateBuildParameterTypeOfBuildType**
> ModelType UpdateBuildParameterTypeOfBuildType(ctx, name, btLocator, optional)
Update type of build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **btLocator** | **string**|  | 
 **optional** | ***UpdateBuildParameterTypeOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildParameterTypeOfBuildTypeOpts struct

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

# **UpdateBuildParameterValueOfBuildType**
> string UpdateBuildParameterValueOfBuildType(ctx, name, btLocator, optional)
Update value of build parameter.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **btLocator** | **string**|  | 
 **optional** | ***UpdateBuildParameterValueOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildParameterValueOfBuildTypeOpts struct

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

# **UpdateBuildParametersOfBuildType**
> Properties UpdateBuildParametersOfBuildType(ctx, btLocator, optional)
Update build parameters.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
 **optional** | ***UpdateBuildParametersOfBuildTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildParametersOfBuildTypeOpts struct

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

# **UpdateBuildTypeVcsRoot**
> VcsRootEntry UpdateBuildTypeVcsRoot(ctx, btLocator, vcsRootLocator, optional)
Update a VCS root of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **vcsRootLocator** | **string**|  | 
 **optional** | ***UpdateBuildTypeVcsRootOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildTypeVcsRootOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of VcsRootEntry**](VcsRootEntry.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**VcsRootEntry**](vcs-root-entry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBuildTypeVcsRootCheckoutRules**
> string UpdateBuildTypeVcsRootCheckoutRules(ctx, btLocator, vcsRootLocator, optional)
Update checkout rules of a VCS root of the matching build configuration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **btLocator** | **string**|  | 
  **vcsRootLocator** | **string**|  | 
 **optional** | ***UpdateBuildTypeVcsRootCheckoutRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBuildTypeVcsRootCheckoutRulesOpts struct

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

