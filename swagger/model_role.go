/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a user/group role.
type Role struct {
	RoleId string `json:"roleId,omitempty"`
	Scope  string `json:"scope,omitempty"`
	Href   string `json:"href,omitempty"`
}
