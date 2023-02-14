/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a test results counter (how many times this test was successful/failed/muted/ignored).
type TestCounters struct {
	Ignored   int32 `json:"ignored,omitempty"`
	Failed    int32 `json:"failed,omitempty"`
	Muted     int32 `json:"muted,omitempty"`
	Success   int32 `json:"success,omitempty"`
	All       int32 `json:"all,omitempty"`
	NewFailed int32 `json:"newFailed,omitempty"`
	Duration  int64 `json:"duration,omitempty"`
}