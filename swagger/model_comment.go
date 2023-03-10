/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a dated comment of the specific user.
type Comment struct {
	Timestamp string `json:"timestamp,omitempty"`
	Text      string `json:"text,omitempty"`
	User      *User  `json:"user,omitempty"`
}
