/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a commiter to a VCS.
type Commiter struct {
	VcsUsername string `json:"vcsUsername,omitempty"`
	Users       *Users `json:"users,omitempty"`
}
