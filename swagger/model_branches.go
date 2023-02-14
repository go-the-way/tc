/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a list of Branch entities.
type Branches struct {
	Count  int32    `json:"count,omitempty"`
	Href   string   `json:"href,omitempty"`
	Branch []Branch `json:"branch,omitempty"`
}