/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a list of MetricValue entities.
type MetricValues struct {
	Count       int32         `json:"count,omitempty"`
	MetricValue []MetricValue `json:"metricValue,omitempty"`
}