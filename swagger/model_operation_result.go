/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a relation between a message and a related entity.
type OperationResult struct {
	Message string         `json:"message,omitempty"`
	Related *RelatedEntity `json:"related,omitempty"`
}
