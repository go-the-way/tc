/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a set of enabled responsibilities of a TeamCity node.
type EnabledResponsibilities struct {
	Count          int32            `json:"count,omitempty"`
	Responsibility []Responsibility `json:"responsibility,omitempty"`
}
