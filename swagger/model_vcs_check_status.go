/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents the last known repository check status.
type VcsCheckStatus struct {
	Status        string `json:"status,omitempty"`
	RequestorType string `json:"requestorType,omitempty"`
	Timestamp     string `json:"timestamp,omitempty"`
}
