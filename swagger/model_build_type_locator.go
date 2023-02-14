/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a locator string for filtering BuildType entities.
type BuildTypeLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// Build locator.
	Build string `json:"build,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count      int32  `json:"count,omitempty"`
	Id         string `json:"id,omitempty"`
	InternalId string `json:"internalId,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	Name string `json:"name,omitempty"`
	// Is paused.
	Paused bool `json:"paused,omitempty"`
	// Project (direct parent) locator.
	Project string `json:"project,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`
	// Base template locator.
	Template string `json:"template,omitempty"`
	// Is a template.
	TemplateFlag bool   `json:"templateFlag,omitempty"`
	Uuid         string `json:"uuid,omitempty"`
	// VCS root locator.
	VcsRoot string `json:"vcsRoot,omitempty"`
	// VCS root instance locator.
	VcsRootInstance string `json:"vcsRootInstance,omitempty"`
}
