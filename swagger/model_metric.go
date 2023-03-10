/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents the specific server metric.
type Metric struct {
	Name           string        `json:"name,omitempty"`
	Description    string        `json:"description,omitempty"`
	PrometheusName string        `json:"prometheusName,omitempty"`
	MetricValues   *MetricValues `json:"metricValues,omitempty"`
	MetricTags     *MetricTags   `json:"metricTags,omitempty"`
}
