/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a list of MetaData entities.
type Datas struct {
	Count int32      `json:"count,omitempty"`
	Data  []MetaData `json:"data,omitempty"`
}
