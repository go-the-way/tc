/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a paginated list of CloudProfile entities.
type CloudProfiles struct {
	Count        int32          `json:"count,omitempty"`
	NextHref     string         `json:"nextHref,omitempty"`
	PrevHref     string         `json:"prevHref,omitempty"`
	Href         string         `json:"href,omitempty"`
	CloudProfile []CloudProfile `json:"cloudProfile,omitempty"`
}
