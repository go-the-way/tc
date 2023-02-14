# BuildLocator

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProject** | **string** | Project (direct or indirect parent) locator. | [optional] [default to null]
**Agent** | **string** | Agent locator. | [optional] [default to null]
**AgentTypeId** | **int32** | typeId of agent used to execute build. | [optional] [default to null]
**Any** | **bool** | State can be any. | [optional] [default to null]
**ArtifactDependency** | **string** |  | [optional] [default to null]
**Branch** | **string** | Branch locator. | [optional] [default to null]
**BuildType** | **string** | Build type locator. | [optional] [default to null]
**Canceled** | **bool** | Is canceled. | [optional] [default to null]
**CompatibleAgent** | **string** | Agent locator. | [optional] [default to null]
**Composite** | **bool** | Is composite. | [optional] [default to null]
**Count** | **int32** | For paginated calls, how many entities to return per page. | [optional] [default to null]
**DefaultFilter** | **bool** | If true, applies default filter which returns only \&quot;normal\&quot; builds (finished builds which are not canceled, not failed-to-start, not personal, and on default branch (in branched build configurations)). | [optional] [default to null]
**FailedToStart** | **bool** | Is failed to start. | [optional] [default to null]
**FinishDate** | **string** | Requires either date or build dimension. | [optional] [default to null]
**Finished** | **bool** | Is finished. | [optional] [default to null]
**Hanging** | **bool** | Is hanging. | [optional] [default to null]
**History** | **bool** | Is history build. | [optional] [default to null]
**Id** | **int32** | Entity ID. | [optional] [default to null]
**Item** | **string** | Supply multiple locators and return a union of the results. | [optional] [default to null]
**LookupLimit** | **int32** | Limit processing to the latest &#x60;&lt;lookupLimit&gt;&#x60; entities. | [optional] [default to null]
**Number** | **string** |  | [optional] [default to null]
**Personal** | **bool** | Is a personal build. | [optional] [default to null]
**Pinned** | **bool** | Is pinned. | [optional] [default to null]
**Project** | **string** | Project (direct parent) locator. | [optional] [default to null]
**Property** | **string** |  | [optional] [default to null]
**Queued** | **bool** | Is queued. | [optional] [default to null]
**QueuedDate** | **string** | Requires either date or build dimension. | [optional] [default to null]
**Revision** | **string** | Build revision. | [optional] [default to null]
**Running** | **bool** | Is running. | [optional] [default to null]
**SnapshotDependency** | **string** |  | [optional] [default to null]
**Start** | **int32** | For paginated calls, from which entity to start rendering the page. | [optional] [default to null]
**StartDate** | **string** | Requires either date or build dimension. | [optional] [default to null]
**State** | **string** |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Tag** | **string** | Tag locator. | [optional] [default to null]
**TaskId** | **int32** | ID of a build or build promotion. | [optional] [default to null]
**User** | **string** | For personal builds checks the owner of the build, triggerring user in other cases. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


