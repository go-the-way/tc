/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a relation between a VCS root and unique settings set for this root.
type VcsRootInstance struct {
	Id                        string           `json:"id,omitempty"`
	VcsRootId                 string           `json:"vcs-root-id,omitempty"`
	VcsRootInternalId         string           `json:"vcsRootInternalId,omitempty"`
	Name                      string           `json:"name,omitempty"`
	VcsName                   string           `json:"vcsName,omitempty"`
	ModificationCheckInterval int32            `json:"modificationCheckInterval,omitempty"`
	CommitHookMode            bool             `json:"commitHookMode,omitempty"`
	LastVersion               string           `json:"lastVersion,omitempty"`
	LastVersionInternal       string           `json:"lastVersionInternal,omitempty"`
	Href                      string           `json:"href,omitempty"`
	VcsRoot                   *VcsRoot         `json:"vcs-root,omitempty"`
	Status                    *VcsStatus       `json:"status,omitempty"`
	RepositoryState           *RepositoryState `json:"repositoryState,omitempty"`
	Properties                *Properties      `json:"properties,omitempty"`
	RepositoryIdStrings       *Items           `json:"repositoryIdStrings,omitempty"`
	ProjectLocator            string           `json:"projectLocator,omitempty"`
}
