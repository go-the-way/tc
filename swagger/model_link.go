/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a list of URLs.
type Link struct {
	Type_       string `json:"type,omitempty"`
	Url         string `json:"url,omitempty"`
	RelativeUrl string `json:"relativeUrl,omitempty"`
}