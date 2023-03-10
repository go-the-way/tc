/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a cancel request for the specific build.
type BuildCancelRequest struct {
	Comment        string `json:"comment,omitempty"`
	ReaddIntoQueue bool   `json:"readdIntoQueue,omitempty"`
}
