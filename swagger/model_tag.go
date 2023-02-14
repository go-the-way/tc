/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a single build tag.
type Tag struct {
	Name    string `json:"name,omitempty"`
	Owner   *User  `json:"owner,omitempty"`
	Private bool   `json:"private,omitempty"`
}
