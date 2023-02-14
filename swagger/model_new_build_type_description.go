/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a build configuration stub.
type NewBuildTypeDescription struct {
	CopyAllAssociatedSettings bool        `json:"copyAllAssociatedSettings,omitempty"`
	ProjectsIdsMap            *Properties `json:"projectsIdsMap,omitempty"`
	BuildTypesIdsMap          *Properties `json:"buildTypesIdsMap,omitempty"`
	VcsRootsIdsMap            *Properties `json:"vcsRootsIdsMap,omitempty"`
	Name                      string      `json:"name,omitempty"`
	Id                        string      `json:"id,omitempty"`
	SourceBuildTypeLocator    string      `json:"sourceBuildTypeLocator,omitempty"`
	SourceBuildType           *BuildType  `json:"sourceBuildType,omitempty"`
}
