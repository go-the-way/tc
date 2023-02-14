/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a paginated list of Project entities.
type Projects struct {
	Count    int32     `json:"count,omitempty"`
	Href     string    `json:"href,omitempty"`
	NextHref string    `json:"nextHref,omitempty"`
	PrevHref string    `json:"prevHref,omitempty"`
	Project  []Project `json:"project,omitempty"`
}
