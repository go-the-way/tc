/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents permission restrictions of an authentication token.
type PermissionRestriction struct {
	IsGlobalScope bool        `json:"isGlobalScope,omitempty"`
	Project       *Project    `json:"project,omitempty"`
	Permission    *Permission `json:"permission,omitempty"`
}
