/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents license state details (available build configurations, agents, etc.).
type LicensingData struct {
	LicenseUseExceeded         bool         `json:"licenseUseExceeded,omitempty"`
	MaxAgents                  int32        `json:"maxAgents,omitempty"`
	UnlimitedAgents            bool         `json:"unlimitedAgents,omitempty"`
	MaxBuildTypes              int32        `json:"maxBuildTypes,omitempty"`
	UnlimitedBuildTypes        bool         `json:"unlimitedBuildTypes,omitempty"`
	BuildTypesLeft             int32        `json:"buildTypesLeft,omitempty"`
	ServerLicenseType          string       `json:"serverLicenseType,omitempty"`
	ServerEffectiveReleaseDate string       `json:"serverEffectiveReleaseDate,omitempty"`
	AgentsLeft                 int32        `json:"agentsLeft,omitempty"`
	LicenseKeys                *LicenseKeys `json:"licenseKeys,omitempty"`
}