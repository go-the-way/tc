/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents build customization settings of a trigger
type BuildTriggerCustomization struct {
	EnforceCleanCheckout                bool        `json:"enforceCleanCheckout,omitempty"`
	EnforceCleanCheckoutForDependencies bool        `json:"enforceCleanCheckoutForDependencies,omitempty"`
	Parameters                          *Properties `json:"parameters,omitempty"`
}