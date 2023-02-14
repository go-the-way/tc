/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a build configuration.
type BuildType struct {
	Id                    string                `json:"id,omitempty"`
	InternalId            string                `json:"internalId,omitempty"`
	Name                  string                `json:"name,omitempty"`
	TemplateFlag          bool                  `json:"templateFlag,omitempty"`
	Type_                 string                `json:"type,omitempty"`
	Paused                bool                  `json:"paused,omitempty"`
	Uuid                  string                `json:"uuid,omitempty"`
	Description           string                `json:"description,omitempty"`
	ProjectName           string                `json:"projectName,omitempty"`
	ProjectId             string                `json:"projectId,omitempty"`
	ProjectInternalId     string                `json:"projectInternalId,omitempty"`
	Href                  string                `json:"href,omitempty"`
	WebUrl                string                `json:"webUrl,omitempty"`
	Inherited             bool                  `json:"inherited,omitempty"`
	Links                 *Links                `json:"links,omitempty"`
	Project               *Project              `json:"project,omitempty"`
	Templates             *BuildTypes           `json:"templates,omitempty"`
	Template              *BuildType            `json:"template,omitempty"`
	VcsRootEntries        *VcsRootEntries       `json:"vcs-root-entries,omitempty"`
	Settings              *Properties           `json:"settings,omitempty"`
	Parameters            *Properties           `json:"parameters,omitempty"`
	Steps                 *Steps                `json:"steps,omitempty"`
	Features              *Features             `json:"features,omitempty"`
	Triggers              *Triggers             `json:"triggers,omitempty"`
	SnapshotDependencies  *SnapshotDependencies `json:"snapshot-dependencies,omitempty"`
	ArtifactDependencies  *ArtifactDependencies `json:"artifact-dependencies,omitempty"`
	AgentRequirements     *AgentRequirements    `json:"agent-requirements,omitempty"`
	Branches              *Branches             `json:"branches,omitempty"`
	Builds                *Builds               `json:"builds,omitempty"`
	Investigations        *Investigations       `json:"investigations,omitempty"`
	CompatibleAgents      *Agents               `json:"compatibleAgents,omitempty"`
	VcsRootInstances      *VcsRootInstances     `json:"vcsRootInstances,omitempty"`
	ExternalStatusAllowed bool                  `json:"externalStatusAllowed,omitempty"`
	PauseComment          *Comment              `json:"pauseComment,omitempty"`
	Locator               string                `json:"locator,omitempty"`
}