/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a paginated list of Build entities.
type Builds struct {
	Href     string  `json:"href,omitempty"`
	Build    []Build `json:"build,omitempty"`
	NextHref string  `json:"nextHref,omitempty"`
	PrevHref string  `json:"prevHref,omitempty"`
	Count    int32   `json:"count,omitempty"`
}
