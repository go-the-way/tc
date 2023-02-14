# \GroupApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroup**](GroupApi.md#AddGroup) | **Post** /app/rest/userGroups | Add a new user group.
[**AddRoleAtScopeToGroup**](GroupApi.md#AddRoleAtScopeToGroup) | **Post** /app/rest/userGroups/{groupLocator}/roles/{roleId}/{scope} | Add a role with the specific scope to the matching user group.
[**AddRoleToGroup**](GroupApi.md#AddRoleToGroup) | **Post** /app/rest/userGroups/{groupLocator}/roles | Add a role to the matching user group.
[**DeleteGroup**](GroupApi.md#DeleteGroup) | **Delete** /app/rest/userGroups/{groupLocator} | Delete user group matching the locator.
[**GetAllGroups**](GroupApi.md#GetAllGroups) | **Get** /app/rest/userGroups | Get all user groups.
[**GetGroupParentGroups**](GroupApi.md#GetGroupParentGroups) | **Get** /app/rest/userGroups/{groupLocator}/parent-groups | Get parent groups of the matching user group.
[**GetGroupProperties**](GroupApi.md#GetGroupProperties) | **Get** /app/rest/userGroups/{groupLocator}/properties | Get properties of the matching user group.
[**GetGroupProperty**](GroupApi.md#GetGroupProperty) | **Get** /app/rest/userGroups/{groupLocator}/properties/{name} | Get a property of the matching user group.
[**GetGroupRoleAtScope**](GroupApi.md#GetGroupRoleAtScope) | **Get** /app/rest/userGroups/{groupLocator}/roles/{roleId}/{scope} | Get a role with the specific scope of the matching user group.
[**GetGroupRoles**](GroupApi.md#GetGroupRoles) | **Get** /app/rest/userGroups/{groupLocator}/roles | Get all roles of the matching user group.
[**GetUserGroupOfGroup**](GroupApi.md#GetUserGroupOfGroup) | **Get** /app/rest/userGroups/{groupLocator} | Get user group matching the locator.
[**RemoveGroupProperty**](GroupApi.md#RemoveGroupProperty) | **Delete** /app/rest/userGroups/{groupLocator}/properties/{name} | Remove a property of the matching user group.
[**RemoveRoleAtScopeFromGroup**](GroupApi.md#RemoveRoleAtScopeFromGroup) | **Delete** /app/rest/userGroups/{groupLocator}/roles/{roleId}/{scope} | Remove a role with the specific scope from the matching user group.
[**SetGroupParentGroups**](GroupApi.md#SetGroupParentGroups) | **Put** /app/rest/userGroups/{groupLocator}/parent-groups | Update parent groups of the matching user group.
[**SetGroupProperty**](GroupApi.md#SetGroupProperty) | **Put** /app/rest/userGroups/{groupLocator}/properties/{name} | Update a property of the matching user group.
[**SetGroupRoles**](GroupApi.md#SetGroupRoles) | **Put** /app/rest/userGroups/{groupLocator}/roles | Update roles of the matching user group.


# **AddGroup**
> Group AddGroup(ctx, optional)
Add a new user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Group**](Group.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Group**](group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRoleAtScopeToGroup**
> Role AddRoleAtScopeToGroup(ctx, groupLocator, roleId, scope)
Add a role with the specific scope to the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
  **roleId** | **string**|  | 
  **scope** | **string**|  | 

### Return type

[**Role**](role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRoleToGroup**
> Role AddRoleToGroup(ctx, groupLocator, optional)
Add a role to the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
 **optional** | ***AddRoleToGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddRoleToGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Role**](Role.md)|  | 

### Return type

[**Role**](role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DeleteGroup(ctx, groupLocator)
Delete user group matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllGroups**
> Groups GetAllGroups(ctx, optional)
Get all user groups.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**|  | 

### Return type

[**Groups**](groups.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupParentGroups**
> Groups GetGroupParentGroups(ctx, groupLocator, optional)
Get parent groups of the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
 **optional** | ***GetGroupParentGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetGroupParentGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Groups**](groups.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupProperties**
> Properties GetGroupProperties(ctx, groupLocator, optional)
Get properties of the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
 **optional** | ***GetGroupPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetGroupPropertiesOpts struct

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

# **GetGroupProperty**
> string GetGroupProperty(ctx, groupLocator, name)
Get a property of the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
  **name** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupRoleAtScope**
> Role GetGroupRoleAtScope(ctx, groupLocator, roleId, scope)
Get a role with the specific scope of the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
  **roleId** | **string**|  | 
  **scope** | **string**|  | 

### Return type

[**Role**](role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupRoles**
> Roles GetGroupRoles(ctx, groupLocator)
Get all roles of the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 

### Return type

[**Roles**](roles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroupOfGroup**
> Group GetUserGroupOfGroup(ctx, groupLocator, optional)
Get user group matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
 **optional** | ***GetUserGroupOfGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUserGroupOfGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Group**](group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveGroupProperty**
> RemoveGroupProperty(ctx, groupLocator, name)
Remove a property of the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoleAtScopeFromGroup**
> RemoveRoleAtScopeFromGroup(ctx, groupLocator, roleId, scope)
Remove a role with the specific scope from the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
  **roleId** | **string**|  | 
  **scope** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetGroupParentGroups**
> Groups SetGroupParentGroups(ctx, groupLocator, optional)
Update parent groups of the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
 **optional** | ***SetGroupParentGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetGroupParentGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Groups**](Groups.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Groups**](groups.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetGroupProperty**
> string SetGroupProperty(ctx, groupLocator, name, optional)
Update a property of the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
  **name** | **string**|  | 
 **optional** | ***SetGroupPropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetGroupPropertyOpts struct

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

# **SetGroupRoles**
> Roles SetGroupRoles(ctx, groupLocator, optional)
Update roles of the matching user group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupLocator** | **string**|  | 
 **optional** | ***SetGroupRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetGroupRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Roles**](Roles.md)|  | 

### Return type

[**Roles**](roles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

