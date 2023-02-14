/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a project.
type Project struct {
	Id                      string           `json:"id,omitempty"`
	InternalId              string           `json:"internalId,omitempty"`
	Uuid                    string           `json:"uuid,omitempty"`
	Name                    string           `json:"name,omitempty"`
	ParentProjectId         string           `json:"parentProjectId,omitempty"`
	ParentProjectInternalId string           `json:"parentProjectInternalId,omitempty"`
	ParentProjectName       string           `json:"parentProjectName,omitempty"`
	Archived                bool             `json:"archived,omitempty"`
	Virtual                 bool             `json:"virtual,omitempty"`
	Description             string           `json:"description,omitempty"`
	Href                    string           `json:"href,omitempty"`
	WebUrl                  string           `json:"webUrl,omitempty"`
	Links                   *Links           `json:"links,omitempty"`
	ParentProject           *Project         `json:"parentProject,omitempty"`
	ReadOnlyUI              *StateField      `json:"readOnlyUI,omitempty"`
	DefaultTemplate         *BuildType       `json:"defaultTemplate,omitempty"`
	BuildTypes              *BuildTypes      `json:"buildTypes,omitempty"`
	Templates               *BuildTypes      `json:"templates,omitempty"`
	Parameters              *Properties      `json:"parameters,omitempty"`
	VcsRoots                *VcsRoots        `json:"vcsRoots,omitempty"`
	ProjectFeatures         *ProjectFeatures `json:"projectFeatures,omitempty"`
	Projects                *Projects        `json:"projects,omitempty"`
	CloudProfiles           *CloudProfiles   `json:"cloudProfiles,omitempty"`
	AncestorProjects        *Projects        `json:"ancestorProjects,omitempty"`
	Locator                 string           `json:"locator,omitempty"`
}
