/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a user.
type User struct {
	Username    string       `json:"username,omitempty"`
	Name        string       `json:"name,omitempty"`
	Id          int64        `json:"id,omitempty"`
	Email       string       `json:"email,omitempty"`
	LastLogin   string       `json:"lastLogin,omitempty"`
	Password    string       `json:"password,omitempty"`
	HasPassword bool         `json:"hasPassword,omitempty"`
	Realm       string       `json:"realm,omitempty"`
	Href        string       `json:"href,omitempty"`
	Properties  *Properties  `json:"properties,omitempty"`
	Roles       *Roles       `json:"roles,omitempty"`
	Groups      *Groups      `json:"groups,omitempty"`
	Locator     string       `json:"locator,omitempty"`
	Avatars     *UserAvatars `json:"avatars,omitempty"`
	Enabled2FA  bool         `json:"enabled2FA,omitempty"`
}
