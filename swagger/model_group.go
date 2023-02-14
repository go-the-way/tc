/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a user group.
type Group struct {
	Key          string      `json:"key,omitempty"`
	Name         string      `json:"name,omitempty"`
	Href         string      `json:"href,omitempty"`
	Description  string      `json:"description,omitempty"`
	ParentGroups *Groups     `json:"parent-groups,omitempty"`
	ChildGroups  *Groups     `json:"child-groups,omitempty"`
	Users        *Users      `json:"users,omitempty"`
	Roles        *Roles      `json:"roles,omitempty"`
	Properties   *Properties `json:"properties,omitempty"`
}