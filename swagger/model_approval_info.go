/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents approval status for this build, if applicable.
type ApprovalInfo struct {
	TimeoutTimestamp           string          `json:"timeoutTimestamp,omitempty"`
	ConfigurationValid         bool            `json:"configurationValid,omitempty"`
	CanBeApprovedByCurrentUser bool            `json:"canBeApprovedByCurrentUser,omitempty"`
	UserApprovals              *UserApprovals  `json:"userApprovals,omitempty"`
	GroupApprovals             *GroupApprovals `json:"groupApprovals,omitempty"`
	Status                     string          `json:"status,omitempty"`
}
