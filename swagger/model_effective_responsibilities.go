/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a set of effective responsibilities of a TeamCity node. A responsibility becomes effective once it is enabled in the server configuration and a responsible node detects this change and updates its state.
type EffectiveResponsibilities struct {
	Count          int32            `json:"count,omitempty"`
	Responsibility []Responsibility `json:"responsibility,omitempty"`
}
