/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents build customizations (artifact dependency overrides, custom parameters or changesets).
type Customizations struct {
	Parameters           map[string]string `json:"parameters,omitempty"`
	Changes              map[string]string `json:"changes,omitempty"`
	ArtifactDependencies map[string]string `json:"artifactDependencies,omitempty"`
}
