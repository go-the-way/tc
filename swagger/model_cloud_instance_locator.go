/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a locator string for filtering CloudInstance entities.
type CloudInstanceLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// Agent locator.
	Agent string `json:"agent,omitempty"`
	Id    string `json:"id,omitempty"`
	// Cloud image locator.
	Instance string `json:"instance,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item           string `json:"item,omitempty"`
	NetworkAddress string `json:"networkAddress,omitempty"`
	// Cloud profile locator.
	Profile string `json:"profile,omitempty"`
	// Project locator.
	Project  string `json:"project,omitempty"`
	Property string `json:"property,omitempty"`
}
