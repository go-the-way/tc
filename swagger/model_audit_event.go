/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents an audit event including a user and affected entities.
type AuditEvent struct {
	Id              string           `json:"id,omitempty"`
	Timestamp       string           `json:"timestamp,omitempty"`
	Comment         string           `json:"comment,omitempty"`
	Action          *AuditAction     `json:"action,omitempty"`
	RelatedEntities *RelatedEntities `json:"relatedEntities,omitempty"`
	User            *User            `json:"user,omitempty"`
}
