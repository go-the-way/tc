/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents agent authorization data.
type AuthorizedInfo struct {
	Status  bool     `json:"status,omitempty"`
	Comment *Comment `json:"comment,omitempty"`
}
