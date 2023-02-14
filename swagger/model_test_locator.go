/*
 * TeamCity REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2018.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Represents a locator string for filtering Test entities.
type TestLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// Is currently failing.
	CurrentlyFailing bool `json:"currentlyFailing,omitempty"`
	// Is currently investigated.
	CurrentlyInvestigated bool `json:"currentlyInvestigated,omitempty"`
	// Is currently muted.
	CurrentlyMuted bool `json:"currentlyMuted,omitempty"`
	// Entity ID.
	Id int32 `json:"id,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32 `json:"lookupLimit,omitempty"`
	// Build type locator (for finding out if this test is affected by mutes in build type).
	MuteAffected string `json:"muteAffected,omitempty"`
	Name         string `json:"name,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`
}