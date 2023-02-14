/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a progress estimate of this build.
type ProgressInfo struct {
	PercentageComplete    int32  `json:"percentageComplete,omitempty"`
	ElapsedSeconds        int64  `json:"elapsedSeconds,omitempty"`
	EstimatedTotalSeconds int64  `json:"estimatedTotalSeconds,omitempty"`
	LeftSeconds           int64  `json:"leftSeconds,omitempty"`
	CurrentStageText      string `json:"currentStageText,omitempty"`
	Outdated              bool   `json:"outdated,omitempty"`
	ProbablyHanging       bool   `json:"probablyHanging,omitempty"`
	LastActivityTime      string `json:"lastActivityTime,omitempty"`
	OutdatedReasonBuild   *Build `json:"outdatedReasonBuild,omitempty"`
}