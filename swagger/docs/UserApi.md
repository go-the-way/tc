# \UserApi

All URIs are relative to *https://teamcity.jetbrains.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoleToUser**](UserApi.md#AddRoleToUser) | **Post** /app/rest/users/{userLocator}/roles | Add a role to the matching user.
[**AddRoleToUserAtScope**](UserApi.md#AddRoleToUserAtScope) | **Put** /app/rest/users/{userLocator}/roles/{roleId}/{scope} | Add a role with the specific scope to the matching user.
[**AddUser**](UserApi.md#AddUser) | **Post** /app/rest/users | Create a new user.
[**AddUserToken**](UserApi.md#AddUserToken) | **Post** /app/rest/users/{userLocator}/tokens | Create a new authentication token for the matching user.
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /app/rest/users/{userLocator} | Delete user matching the locator.
[**DeleteUserField**](UserApi.md#DeleteUserField) | **Delete** /app/rest/users/{userLocator}/{field} | Remove a property of the matching user.
[**DeleteUserToken**](UserApi.md#DeleteUserToken) | **Delete** /app/rest/users/{userLocator}/tokens/{name} | Remove an authentication token from the matching user.
[**GetAllUserGroups**](UserApi.md#GetAllUserGroups) | **Get** /app/rest/users/{userLocator}/groups | Get all groups of the matching user.
[**GetAllUserRoles**](UserApi.md#GetAllUserRoles) | **Get** /app/rest/users/{userLocator}/roles | Get all user roles of the matching user.
[**GetAllUsers**](UserApi.md#GetAllUsers) | **Get** /app/rest/users | Get all users.
[**GetUser**](UserApi.md#GetUser) | **Get** /app/rest/users/{userLocator} | Get user matching the locator.
[**GetUserField**](UserApi.md#GetUserField) | **Get** /app/rest/users/{userLocator}/{field} | Get a field of the matching user.
[**GetUserGroup**](UserApi.md#GetUserGroup) | **Get** /app/rest/users/{userLocator}/groups/{groupLocator} | Get a user group of the matching user.
[**GetUserPermissions**](UserApi.md#GetUserPermissions) | **Get** /app/rest/users/{userLocator}/permissions | Get all permissions effective for the matching user.
[**GetUserProperties**](UserApi.md#GetUserProperties) | **Get** /app/rest/users/{userLocator}/properties | Get all properties of the matching user.
[**GetUserProperty**](UserApi.md#GetUserProperty) | **Get** /app/rest/users/{userLocator}/properties/{name} | Get a property of the matching user.
[**GetUserRolesAtScope**](UserApi.md#GetUserRolesAtScope) | **Get** /app/rest/users/{userLocator}/roles/{roleId}/{scope} | Get a user role with the specific scope from the matching user.
[**GetUserTokens**](UserApi.md#GetUserTokens) | **Get** /app/rest/users/{userLocator}/tokens | Get all authentication tokens of the matching user.
[**RemoveUserFromGroup**](UserApi.md#RemoveUserFromGroup) | **Delete** /app/rest/users/{userLocator}/groups/{groupLocator} | Remove the matching user from the specific group.
[**RemoveUserProperty**](UserApi.md#RemoveUserProperty) | **Delete** /app/rest/users/{userLocator}/properties/{name} | Remove a property of the matching user.
[**RemoveUserRememberMe**](UserApi.md#RemoveUserRememberMe) | **Delete** /app/rest/users/{userLocator}/debug/rememberMe | Remove the RememberMe data of the matching user.
[**RemoveUserRoleAtScope**](UserApi.md#RemoveUserRoleAtScope) | **Delete** /app/rest/users/{userLocator}/roles/{roleId}/{scope} | Remove a role with the specific scope from the matching user.
[**ReplaceUser**](UserApi.md#ReplaceUser) | **Put** /app/rest/users/{userLocator} | Update user matching the locator.
[**SetUserField**](UserApi.md#SetUserField) | **Put** /app/rest/users/{userLocator}/{field} | Update a field of the matching user.
[**SetUserGroups**](UserApi.md#SetUserGroups) | **Put** /app/rest/users/{userLocator}/groups | Update groups of the matching user.
[**SetUserProperty**](UserApi.md#SetUserProperty) | **Put** /app/rest/users/{userLocator}/properties/{name} | Update a property of the matching user.
[**SetUserRoles**](UserApi.md#SetUserRoles) | **Put** /app/rest/users/{userLocator}/roles | Update user roles of the matching user.


# **AddRoleToUser**
> Role AddRoleToUser(ctx, userLocator, optional)
Add a role to the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***AddRoleToUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddRoleToUserOpts struct

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

# **AddRoleToUserAtScope**
> Role AddRoleToUserAtScope(ctx, userLocator, roleId, scope)
Add a role with the specific scope to the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
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

# **AddUser**
> User AddUser(ctx, optional)
Create a new user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of User**](User.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**User**](user.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUserToken**
> Token AddUserToken(ctx, userLocator, optional)
Create a new authentication token for the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***AddUserTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddUserTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Token**](Token.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Token**](token.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, userLocator)
Delete user matching the locator.



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

# **DeleteUserField**
> DeleteUserField(ctx, userLocator, field)
Remove a property of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserToken**
> DeleteUserToken(ctx, userLocator, name)
Remove an authentication token from the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllUserGroups**
> Groups GetAllUserGroups(ctx, userLocator, optional)
Get all groups of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***GetAllUserGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllUserGroupsOpts struct

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

# **GetAllUserRoles**
> Roles GetAllUserRoles(ctx, userLocator)
Get all user roles of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 

### Return type

[**Roles**](roles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllUsers**
> Users GetAllUsers(ctx, optional)
Get all users.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**Users**](users.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser(ctx, userLocator, optional)
Get user matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***GetUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**User**](user.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserField**
> string GetUserField(ctx, userLocator, field)
Get a field of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **field** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroup**
> Group GetUserGroup(ctx, userLocator, groupLocator, optional)
Get a user group of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **groupLocator** | **string**|  | 
 **optional** | ***GetUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUserGroupOpts struct

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

# **GetUserPermissions**
> PermissionAssignments GetUserPermissions(ctx, userLocator, optional)
Get all permissions effective for the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***GetUserPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUserPermissionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locator** | **optional.String**|  | 
 **fields** | **optional.String**|  | 

### Return type

[**PermissionAssignments**](permissionAssignments.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserProperties**
> Properties GetUserProperties(ctx, userLocator, optional)
Get all properties of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***GetUserPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUserPropertiesOpts struct

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

# **GetUserProperty**
> string GetUserProperty(ctx, userLocator, name)
Get a property of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **name** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserRolesAtScope**
> Role GetUserRolesAtScope(ctx, userLocator, roleId, scope)
Get a user role with the specific scope from the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
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

# **GetUserTokens**
> Tokens GetUserTokens(ctx, userLocator, optional)
Get all authentication tokens of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***GetUserTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUserTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**|  | 

### Return type

[**Tokens**](tokens.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserFromGroup**
> RemoveUserFromGroup(ctx, userLocator, groupLocator, optional)
Remove the matching user from the specific group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **groupLocator** | **string**|  | 
 **optional** | ***RemoveUserFromGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoveUserFromGroupOpts struct

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

# **RemoveUserProperty**
> RemoveUserProperty(ctx, userLocator, name)
Remove a property of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserRememberMe**
> RemoveUserRememberMe(ctx, userLocator)
Remove the RememberMe data of the matching user.



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
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserRoleAtScope**
> RemoveUserRoleAtScope(ctx, userLocator, roleId, scope)
Remove a role with the specific scope from the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
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

# **ReplaceUser**
> User ReplaceUser(ctx, userLocator, optional)
Update user matching the locator.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***ReplaceUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplaceUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of User**](User.md)|  | 
 **fields** | **optional.String**|  | 

### Return type

[**User**](user.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetUserField**
> string SetUserField(ctx, userLocator, field, optional)
Update a field of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***SetUserFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetUserFieldOpts struct

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

# **SetUserGroups**
> Groups SetUserGroups(ctx, userLocator, optional)
Update groups of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***SetUserGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetUserGroupsOpts struct

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

# **SetUserProperty**
> string SetUserProperty(ctx, userLocator, name, optional)
Update a property of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
  **name** | **string**|  | 
 **optional** | ***SetUserPropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetUserPropertyOpts struct

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

# **SetUserRoles**
> Roles SetUserRoles(ctx, userLocator, optional)
Update user roles of the matching user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userLocator** | **string**|  | 
 **optional** | ***SetUserRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetUserRolesOpts struct

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

