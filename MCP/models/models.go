package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// IssueTypeSchemeID represents the IssueTypeSchemeID schema from the OpenAPI specification
type IssueTypeSchemeID struct {
	Issuetypeschemeid string `json:"issueTypeSchemeId"` // The ID of the issue type scheme.
}

// PageBeanIssueSecurityLevelMember represents the PageBeanIssueSecurityLevelMember schema from the OpenAPI specification
type PageBeanIssueSecurityLevelMember struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []IssueSecurityLevelMember `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// FieldConfigurationSchemeProjects represents the FieldConfigurationSchemeProjects schema from the OpenAPI specification
type FieldConfigurationSchemeProjects struct {
	Fieldconfigurationscheme FieldConfigurationScheme `json:"fieldConfigurationScheme,omitempty"` // Details of a field configuration scheme.
	Projectids []string `json:"projectIds"` // The IDs of projects using the field configuration scheme.
}

// PageBeanUiModificationDetails represents the PageBeanUiModificationDetails schema from the OpenAPI specification
type PageBeanUiModificationDetails struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []UiModificationDetails `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// ReorderIssueResolutionsRequest represents the ReorderIssueResolutionsRequest schema from the OpenAPI specification
type ReorderIssueResolutionsRequest struct {
	Position string `json:"position,omitempty"` // The position for issue resolutions to be moved to. Required if `after` isn't provided.
	After string `json:"after,omitempty"` // The ID of the resolution. Required if `position` isn't provided.
	Ids []string `json:"ids"` // The list of resolution IDs to be reordered. Cannot contain duplicates nor after ID.
}

// IssueCommentListRequestBean represents the IssueCommentListRequestBean schema from the OpenAPI specification
type IssueCommentListRequestBean struct {
	Ids []int64 `json:"ids"` // The list of comment IDs. A maximum of 1000 IDs can be specified.
}

// VersionMoveBean represents the VersionMoveBean schema from the OpenAPI specification
type VersionMoveBean struct {
	After string `json:"after,omitempty"` // The URL (self link) of the version after which to place the moved version. Cannot be used with `position`.
	Position string `json:"position,omitempty"` // An absolute position in which to place the moved version. Cannot be used with `after`.
}

// IssueTypeSchemeProjectAssociation represents the IssueTypeSchemeProjectAssociation schema from the OpenAPI specification
type IssueTypeSchemeProjectAssociation struct {
	Issuetypeschemeid string `json:"issueTypeSchemeId"` // The ID of the issue type scheme.
	Projectid string `json:"projectId"` // The ID of the project.
}

// JiraExpressionEvalRequestBean represents the JiraExpressionEvalRequestBean schema from the OpenAPI specification
type JiraExpressionEvalRequestBean struct {
	Context interface{} `json:"context,omitempty"` // The context in which the Jira expression is evaluated.
	Expression string `json:"expression"` // The Jira expression to evaluate.
}

// IssueTypeSchemeUpdateDetails represents the IssueTypeSchemeUpdateDetails schema from the OpenAPI specification
type IssueTypeSchemeUpdateDetails struct {
	Defaultissuetypeid string `json:"defaultIssueTypeId,omitempty"` // The ID of the default issue type of the issue type scheme.
	Description string `json:"description,omitempty"` // The description of the issue type scheme. The maximum length is 4000 characters.
	Name string `json:"name,omitempty"` // The name of the issue type scheme. The name must be unique. The maximum length is 255 characters.
}

// Transitions represents the Transitions schema from the OpenAPI specification
type Transitions struct {
	Expand string `json:"expand,omitempty"` // Expand options that include additional transitions details in the response.
	Transitions []IssueTransition `json:"transitions,omitempty"` // List of issue transitions.
}

// PageBeanGroupDetails represents the PageBeanGroupDetails schema from the OpenAPI specification
type PageBeanGroupDetails struct {
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []GroupDetails `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
}

// CreateWorkflowTransitionScreenDetails represents the CreateWorkflowTransitionScreenDetails schema from the OpenAPI specification
type CreateWorkflowTransitionScreenDetails struct {
	Id string `json:"id"` // The ID of the screen.
}

// IssueFilterForBulkPropertySet represents the IssueFilterForBulkPropertySet schema from the OpenAPI specification
type IssueFilterForBulkPropertySet struct {
	Currentvalue interface{} `json:"currentValue,omitempty"` // The value of properties to perform the bulk operation on.
	Entityids []int64 `json:"entityIds,omitempty"` // List of issues to perform the bulk operation on.
	Hasproperty bool `json:"hasProperty,omitempty"` // Whether the bulk operation occurs only when the property is present on or absent from an issue.
}

// FieldChangedClause represents the FieldChangedClause schema from the OpenAPI specification
type FieldChangedClause struct {
	Predicates []JqlQueryClauseTimePredicate `json:"predicates"` // The list of time predicates.
	Field JqlQueryField `json:"field"` // A field used in a JQL query. See [Advanced searching - fields reference](https://confluence.atlassian.com/x/dAiiLQ) for more information about fields in JQL queries.
	Operator string `json:"operator"` // The operator applied to the field.
}

// AttachmentArchiveMetadataReadable represents the AttachmentArchiveMetadataReadable schema from the OpenAPI specification
type AttachmentArchiveMetadataReadable struct {
	Name string `json:"name,omitempty"` // The name of the archive file.
	Totalentrycount int64 `json:"totalEntryCount,omitempty"` // The number of items included in the archive.
	Entries []AttachmentArchiveItemReadable `json:"entries,omitempty"` // The list of the items included in the archive.
	Id int64 `json:"id,omitempty"` // The ID of the attachment.
	Mediatype string `json:"mediaType,omitempty"` // The MIME type of the attachment.
}

// IssueTypeIssueCreateMetadata represents the IssueTypeIssueCreateMetadata schema from the OpenAPI specification
type IssueTypeIssueCreateMetadata struct {
	Name string `json:"name,omitempty"` // The name of the issue type.
	Self string `json:"self,omitempty"` // The URL of these issue type details.
	Subtask bool `json:"subtask,omitempty"` // Whether this issue type is used to create subtasks.
	Avatarid int64 `json:"avatarId,omitempty"` // The ID of the issue type's avatar.
	Expand string `json:"expand,omitempty"` // Expand options that include additional issue type metadata details in the response.
	Hierarchylevel int `json:"hierarchyLevel,omitempty"` // Hierarchy level of the issue type.
	Iconurl string `json:"iconUrl,omitempty"` // The URL of the issue type's avatar.
	Scope interface{} `json:"scope,omitempty"` // Details of the next-gen projects the issue type is available in.
	Fields map[string]interface{} `json:"fields,omitempty"` // List of the fields available when creating an issue for the issue type.
	Id string `json:"id,omitempty"` // The ID of the issue type.
	Description string `json:"description,omitempty"` // The description of the issue type.
	Entityid string `json:"entityId,omitempty"` // Unique ID for next-gen projects.
}

// ProjectRoleActorsUpdateBean represents the ProjectRoleActorsUpdateBean schema from the OpenAPI specification
type ProjectRoleActorsUpdateBean struct {
	Id int64 `json:"id,omitempty"` // The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.
	Categorisedactors map[string]interface{} `json:"categorisedActors,omitempty"` // The actors to add to the project role. Add groups using: * `atlassian-group-role-actor` and a list of group names. * `atlassian-group-role-actor-id` and a list of group IDs. As a group's name can change, use of `atlassian-group-role-actor-id` is recommended. For example, `"atlassian-group-role-actor-id":["eef79f81-0b89-4fca-a736-4be531a10869","77f6ab39-e755-4570-a6ae-2d7a8df0bcb8"]`. Add users using `atlassian-user-role-actor` and a list of account IDs. For example, `"atlassian-user-role-actor":["12345678-9abc-def1-2345-6789abcdef12", "abcdef12-3456-789a-bcde-f123456789ab"]`.
}

// AnnouncementBannerConfigurationUpdate represents the AnnouncementBannerConfigurationUpdate schema from the OpenAPI specification
type AnnouncementBannerConfigurationUpdate struct {
	Isdismissible bool `json:"isDismissible,omitempty"` // Flag indicating if the announcement banner can be dismissed by the user.
	Isenabled bool `json:"isEnabled,omitempty"` // Flag indicating if the announcement banner is enabled or not.
	Message string `json:"message,omitempty"` // The text on the announcement banner.
	Visibility string `json:"visibility,omitempty"` // Visibility of the announcement banner. Can be public or private.
}

// AssociateFieldConfigurationsWithIssueTypesRequest represents the AssociateFieldConfigurationsWithIssueTypesRequest schema from the OpenAPI specification
type AssociateFieldConfigurationsWithIssueTypesRequest struct {
	Mappings []FieldConfigurationToIssueTypeMapping `json:"mappings"` // Field configuration to issue type mappings.
}

// ContainerForWebhookIDs represents the ContainerForWebhookIDs schema from the OpenAPI specification
type ContainerForWebhookIDs struct {
	Webhookids []int64 `json:"webhookIds"` // A list of webhook IDs.
}

// TimeTrackingProvider represents the TimeTrackingProvider schema from the OpenAPI specification
type TimeTrackingProvider struct {
	Key string `json:"key"` // The key for the time tracking provider. For example, *JIRA*.
	Name string `json:"name,omitempty"` // The name of the time tracking provider. For example, *JIRA provided time tracking*.
	Url string `json:"url,omitempty"` // The URL of the configuration page for the time tracking provider app. For example, */example/config/url*. This property is only returned if the `adminPageKey` property is set in the module descriptor of the time tracking provider app.
}

// EntityProperty represents the EntityProperty schema from the OpenAPI specification
type EntityProperty struct {
	Key string `json:"key,omitempty"` // The key of the property. Required on create and update.
	Value interface{} `json:"value,omitempty"` // The value of the property. Required on create and update.
}

// CustomFieldContextDefaultValueSingleGroupPicker represents the CustomFieldContextDefaultValueSingleGroupPicker schema from the OpenAPI specification
type CustomFieldContextDefaultValueSingleGroupPicker struct {
	TypeField string `json:"type"`
	Contextid string `json:"contextId"` // The ID of the context.
	Groupid string `json:"groupId"` // The ID of the the default group.
}

// ProjectIssueTypeMappings represents the ProjectIssueTypeMappings schema from the OpenAPI specification
type ProjectIssueTypeMappings struct {
	Mappings []ProjectIssueTypeMapping `json:"mappings"` // The project and issue type mappings.
}

// WorkManagementNavigationInfo represents the WorkManagementNavigationInfo schema from the OpenAPI specification
type WorkManagementNavigationInfo struct {
	Boardname string `json:"boardName,omitempty"`
}

// PageBeanNotificationSchemeAndProjectMappingJsonBean represents the PageBeanNotificationSchemeAndProjectMappingJsonBean schema from the OpenAPI specification
type PageBeanNotificationSchemeAndProjectMappingJsonBean struct {
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []NotificationSchemeAndProjectMappingJsonBean `json:"values,omitempty"` // The list of items.
}

// SimpleListWrapperGroupName represents the SimpleListWrapperGroupName schema from the OpenAPI specification
type SimpleListWrapperGroupName struct {
	Size int `json:"size,omitempty"`
	Callback ListWrapperCallbackGroupName `json:"callback,omitempty"`
	Items []GroupName `json:"items,omitempty"`
	Max_results int `json:"max-results,omitempty"`
	Pagingcallback ListWrapperCallbackGroupName `json:"pagingCallback,omitempty"`
}

// JiraExpressionEvalContextBean represents the JiraExpressionEvalContextBean schema from the OpenAPI specification
type JiraExpressionEvalContextBean struct {
	Issue interface{} `json:"issue,omitempty"` // The issue that is available under the `issue` variable when evaluating the expression.
	Issues interface{} `json:"issues,omitempty"` // The collection of issues that is available under the `issues` variable when evaluating the expression.
	Project interface{} `json:"project,omitempty"` // The project that is available under the `project` variable when evaluating the expression.
	Servicedesk int64 `json:"serviceDesk,omitempty"` // The ID of the service desk that is available under the `serviceDesk` variable when evaluating the expression.
	Sprint int64 `json:"sprint,omitempty"` // The ID of the sprint that is available under the `sprint` variable when evaluating the expression.
	Board int64 `json:"board,omitempty"` // The ID of the board that is available under the `board` variable when evaluating the expression.
	Custom []CustomContextVariable `json:"custom,omitempty"` // Custom context variables and their types. These variable types are available for use in a custom context: * `user`: A [user](https://developer.atlassian.com/cloud/jira/platform/jira-expressions-type-reference#user) specified as an Atlassian account ID. * `issue`: An [issue](https://developer.atlassian.com/cloud/jira/platform/jira-expressions-type-reference#issue) specified by ID or key. All the fields of the issue object are available in the Jira expression. * `json`: A JSON object containing custom content. * `list`: A JSON list of `user`, `issue`, or `json` variable types.
	Customerrequest int64 `json:"customerRequest,omitempty"` // The ID of the customer request that is available under the `customerRequest` variable when evaluating the expression. This is the same as the ID of the underlying Jira issue, but the customer request context variable will have a different type.
}

// RemoteIssueLinkIdentifies represents the RemoteIssueLinkIdentifies schema from the OpenAPI specification
type RemoteIssueLinkIdentifies struct {
	Self string `json:"self,omitempty"` // The URL of the remote issue link.
	Id int64 `json:"id,omitempty"` // The ID of the remote issue link, such as the ID of the item on the remote system.
}

// JqlFunctionPrecomputationBean represents the JqlFunctionPrecomputationBean schema from the OpenAPI specification
type JqlFunctionPrecomputationBean struct {
	Operator string `json:"operator,omitempty"`
	Field string `json:"field,omitempty"`
	Updated string `json:"updated,omitempty"`
	Used string `json:"used,omitempty"`
	Arguments []string `json:"arguments,omitempty"`
	Id string `json:"id,omitempty"`
	Created string `json:"created,omitempty"`
	Functionname string `json:"functionName,omitempty"`
	Value string `json:"value,omitempty"`
	Functionkey string `json:"functionKey,omitempty"`
}

// PageBeanNotificationScheme represents the PageBeanNotificationScheme schema from the OpenAPI specification
type PageBeanNotificationScheme struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []NotificationScheme `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// CustomFieldDefinitionJsonBean represents the CustomFieldDefinitionJsonBean schema from the OpenAPI specification
type CustomFieldDefinitionJsonBean struct {
	Searcherkey string `json:"searcherKey,omitempty"` // The searcher defines the way the field is searched in Jira. For example, *com.atlassian.jira.plugin.system.customfieldtypes:grouppickersearcher*. The search UI (basic search and JQL search) will display different operations and values for the field, based on the field searcher. You must specify a searcher that is valid for the field type, as listed below (abbreviated values shown): * `cascadingselect`: `cascadingselectsearcher` * `datepicker`: `daterange` * `datetime`: `datetimerange` * `float`: `exactnumber` or `numberrange` * `grouppicker`: `grouppickersearcher` * `importid`: `exactnumber` or `numberrange` * `labels`: `labelsearcher` * `multicheckboxes`: `multiselectsearcher` * `multigrouppicker`: `multiselectsearcher` * `multiselect`: `multiselectsearcher` * `multiuserpicker`: `userpickergroupsearcher` * `multiversion`: `versionsearcher` * `project`: `projectsearcher` * `radiobuttons`: `multiselectsearcher` * `readonlyfield`: `textsearcher` * `select`: `multiselectsearcher` * `textarea`: `textsearcher` * `textfield`: `textsearcher` * `url`: `exacttextsearcher` * `userpicker`: `userpickergroupsearcher` * `version`: `versionsearcher` If no searcher is provided, the field isn't searchable. However, [Forge custom fields](https://developer.atlassian.com/platform/forge/manifest-reference/modules/#jira-custom-field-type--beta-) have a searcher set automatically, so are always searchable.
	TypeField string `json:"type"` // The type of the custom field. These built-in custom field types are available: * `cascadingselect`: Enables values to be selected from two levels of select lists (value: `com.atlassian.jira.plugin.system.customfieldtypes:cascadingselect`) * `datepicker`: Stores a date using a picker control (value: `com.atlassian.jira.plugin.system.customfieldtypes:datepicker`) * `datetime`: Stores a date with a time component (value: `com.atlassian.jira.plugin.system.customfieldtypes:datetime`) * `float`: Stores and validates a numeric (floating point) input (value: `com.atlassian.jira.plugin.system.customfieldtypes:float`) * `grouppicker`: Stores a user group using a picker control (value: `com.atlassian.jira.plugin.system.customfieldtypes:grouppicker`) * `importid`: A read-only field that stores the ID the issue had in the system it was imported from (value: `com.atlassian.jira.plugin.system.customfieldtypes:importid`) * `labels`: Stores labels (value: `com.atlassian.jira.plugin.system.customfieldtypes:labels`) * `multicheckboxes`: Stores multiple values using checkboxes (value: ``) * `multigrouppicker`: Stores multiple user groups using a picker control (value: ``) * `multiselect`: Stores multiple values using a select list (value: `com.atlassian.jira.plugin.system.customfieldtypes:multicheckboxes`) * `multiuserpicker`: Stores multiple users using a picker control (value: `com.atlassian.jira.plugin.system.customfieldtypes:multigrouppicker`) * `multiversion`: Stores multiple versions from the versions available in a project using a picker control (value: `com.atlassian.jira.plugin.system.customfieldtypes:multiversion`) * `project`: Stores a project from a list of projects that the user is permitted to view (value: `com.atlassian.jira.plugin.system.customfieldtypes:project`) * `radiobuttons`: Stores a value using radio buttons (value: `com.atlassian.jira.plugin.system.customfieldtypes:radiobuttons`) * `readonlyfield`: Stores a read-only text value, which can only be populated via the API (value: `com.atlassian.jira.plugin.system.customfieldtypes:readonlyfield`) * `select`: Stores a value from a configurable list of options (value: `com.atlassian.jira.plugin.system.customfieldtypes:select`) * `textarea`: Stores a long text string using a multiline text area (value: `com.atlassian.jira.plugin.system.customfieldtypes:textarea`) * `textfield`: Stores a text string using a single-line text box (value: `com.atlassian.jira.plugin.system.customfieldtypes:textfield`) * `url`: Stores a URL (value: `com.atlassian.jira.plugin.system.customfieldtypes:url`) * `userpicker`: Stores a user using a picker control (value: `com.atlassian.jira.plugin.system.customfieldtypes:userpicker`) * `version`: Stores a version using a picker control (value: `com.atlassian.jira.plugin.system.customfieldtypes:version`) To create a field based on a [Forge custom field type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/#jira-custom-field-type--beta-), use the ID of the Forge custom field type as the value. For example, `ari:cloud:ecosystem::extension/e62f20a2-4b61-4dbe-bfb9-9a88b5e3ac84/548c5df1-24aa-4f7c-bbbb-3038d947cb05/static/my-cf-type-key`.
	Description string `json:"description,omitempty"` // The description of the custom field, which is displayed in Jira.
	Name string `json:"name"` // The name of the custom field, which is displayed in Jira. This is not the unique identifier.
}

// IssueSecurityLevelMember represents the IssueSecurityLevelMember schema from the OpenAPI specification
type IssueSecurityLevelMember struct {
	Issuesecuritylevelid int64 `json:"issueSecurityLevelId"` // The ID of the issue security level.
	Holder interface{} `json:"holder"` // The user or group being granted the permission. It consists of a `type` and a type-dependent `parameter`. See [Holder object](../api-group-permission-schemes/#holder-object) in *Get all permission schemes* for more information.
	Id int64 `json:"id"` // The ID of the issue security level member.
}

// IncludedFields represents the IncludedFields schema from the OpenAPI specification
type IncludedFields struct {
	Included []string `json:"included,omitempty"`
	Actuallyincluded []string `json:"actuallyIncluded,omitempty"`
	Excluded []string `json:"excluded,omitempty"`
}

// PageBeanIssueTypeScreenSchemesProjects represents the PageBeanIssueTypeScreenSchemesProjects schema from the OpenAPI specification
type PageBeanIssueTypeScreenSchemesProjects struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []IssueTypeScreenSchemesProjects `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// PriorityId represents the PriorityId schema from the OpenAPI specification
type PriorityId struct {
	Id string `json:"id"` // The ID of the issue priority.
}

// PageBeanDashboard represents the PageBeanDashboard schema from the OpenAPI specification
type PageBeanDashboard struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []Dashboard `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// CreateWorkflowDetails represents the CreateWorkflowDetails schema from the OpenAPI specification
type CreateWorkflowDetails struct {
	Statuses []CreateWorkflowStatusDetails `json:"statuses"` // The statuses of the workflow. Any status that does not include a transition is added to the workflow without a transition.
	Transitions []CreateWorkflowTransitionDetails `json:"transitions"` // The transitions of the workflow. For the request to be valid, these transitions must: * include one *initial* transition. * not use the same name for a *global* and *directed* transition. * have a unique name for each *global* transition. * have a unique 'to' status for each *global* transition. * have unique names for each transition from a status. * not have a 'from' status on *initial* and *global* transitions. * have a 'from' status on *directed* transitions. All the transition statuses must be included in `statuses`.
	Description string `json:"description,omitempty"` // The description of the workflow. The maximum length is 1000 characters.
	Name string `json:"name"` // The name of the workflow. The name must be unique. The maximum length is 255 characters. Characters can be separated by a whitespace but the name cannot start or end with a whitespace.
}

// UserDetails represents the UserDetails schema from the OpenAPI specification
type UserDetails struct {
	Accountid string `json:"accountId,omitempty"` // The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*.
	Key string `json:"key,omitempty"` // This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Name string `json:"name,omitempty"` // This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Timezone string `json:"timeZone,omitempty"` // The time zone specified in the user's profile. Depending on the user’s privacy settings, this may be returned as null.
	Active bool `json:"active,omitempty"` // Whether the user is active.
	Accounttype string `json:"accountType,omitempty"` // The type of account represented by this user. This will be one of 'atlassian' (normal users), 'app' (application user) or 'customer' (Jira Service Desk customer user)
	Avatarurls interface{} `json:"avatarUrls,omitempty"` // The avatars of the user.
	Displayname string `json:"displayName,omitempty"` // The display name of the user. Depending on the user’s privacy settings, this may return an alternative value.
	Emailaddress string `json:"emailAddress,omitempty"` // The email address of the user. Depending on the user’s privacy settings, this may be returned as null.
	Self string `json:"self,omitempty"` // The URL of the user.
}

// JqlFunctionPrecomputationUpdateBean represents the JqlFunctionPrecomputationUpdateBean schema from the OpenAPI specification
type JqlFunctionPrecomputationUpdateBean struct {
	Id int64 `json:"id"`
	Value string `json:"value"`
}

// PageBeanScreenWithTab represents the PageBeanScreenWithTab schema from the OpenAPI specification
type PageBeanScreenWithTab struct {
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []ScreenWithTab `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
}

// PageBeanComponentWithIssueCount represents the PageBeanComponentWithIssueCount schema from the OpenAPI specification
type PageBeanComponentWithIssueCount struct {
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []ComponentWithIssueCount `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
}

// FailedWebhooks represents the FailedWebhooks schema from the OpenAPI specification
type FailedWebhooks struct {
	Next string `json:"next,omitempty"` // The URL to the next page of results. Present only if the request returned at least one result.The next page may be empty at the time of receiving the response, but new failed webhooks may appear in time. You can save the URL to the next page and query for new results periodically (for example, every hour).
	Values []FailedWebhook `json:"values"` // The list of webhooks.
	Maxresults int `json:"maxResults"` // The maximum number of items on the page. If the list of values is shorter than this number, then there are no more pages.
}

// ContextForProjectAndIssueType represents the ContextForProjectAndIssueType schema from the OpenAPI specification
type ContextForProjectAndIssueType struct {
	Contextid string `json:"contextId"` // The ID of the custom field context.
	Issuetypeid string `json:"issueTypeId"` // The ID of the issue type.
	Projectid string `json:"projectId"` // The ID of the project.
}

// CustomFieldContextDefaultValueUpdate represents the CustomFieldContextDefaultValueUpdate schema from the OpenAPI specification
type CustomFieldContextDefaultValueUpdate struct {
	Defaultvalues []CustomFieldContextDefaultValue `json:"defaultValues,omitempty"`
}

// Screen represents the Screen schema from the OpenAPI specification
type Screen struct {
	Scope interface{} `json:"scope,omitempty"` // The scope of the screen.
	Description string `json:"description,omitempty"` // The description of the screen.
	Id int64 `json:"id,omitempty"` // The ID of the screen.
	Name string `json:"name,omitempty"` // The name of the screen.
}

// DashboardGadgetPosition represents the DashboardGadgetPosition schema from the OpenAPI specification
type DashboardGadgetPosition struct {
	The_column_position_of_the_gadget int `json:"The column position of the gadget."`
	The_row_position_of_the_gadget int `json:"The row position of the gadget."`
}

// PermissionGrants represents the PermissionGrants schema from the OpenAPI specification
type PermissionGrants struct {
	Expand string `json:"expand,omitempty"` // Expand options that include additional permission grant details in the response.
	Permissions []PermissionGrant `json:"permissions,omitempty"` // Permission grants list.
}

// IssueLink represents the IssueLink schema from the OpenAPI specification
type IssueLink struct {
	Id string `json:"id,omitempty"` // The ID of the issue link.
	Inwardissue interface{} `json:"inwardIssue"` // Provides details about the linked issue. If presenting this link in a user interface, use the `inward` field of the issue link type to label the link.
	Outwardissue interface{} `json:"outwardIssue"` // Provides details about the linked issue. If presenting this link in a user interface, use the `outward` field of the issue link type to label the link.
	Self string `json:"self,omitempty"` // The URL of the issue link.
	TypeField interface{} `json:"type"` // The type of link between the issues.
}

// JiraStatus represents the JiraStatus schema from the OpenAPI specification
type JiraStatus struct {
	Id string `json:"id,omitempty"` // The ID of the status.
	Name string `json:"name,omitempty"` // The name of the status.
	Scope StatusScope `json:"scope,omitempty"` // The scope of the status.
	Statuscategory string `json:"statusCategory,omitempty"` // The category of the status.
	Usages []ProjectIssueTypes `json:"usages,omitempty"` // Projects and issue types where the status is used. Only available if the `usages` expand is requested.
	Description string `json:"description,omitempty"` // The description of the status.
}

// IssueTypeWithStatus represents the IssueTypeWithStatus schema from the OpenAPI specification
type IssueTypeWithStatus struct {
	Id string `json:"id"` // The ID of the issue type.
	Name string `json:"name"` // The name of the issue type.
	Self string `json:"self"` // The URL of the issue type's status details.
	Statuses []StatusDetails `json:"statuses"` // List of status details for the issue type.
	Subtask bool `json:"subtask"` // Whether this issue type represents subtasks.
}

// ValueOperand represents the ValueOperand schema from the OpenAPI specification
type ValueOperand struct {
	Encodedvalue string `json:"encodedValue,omitempty"` // Encoded value, which can be used directly in a JQL query.
	Value string `json:"value"` // The operand value.
}

// ProjectScopeBean represents the ProjectScopeBean schema from the OpenAPI specification
type ProjectScopeBean struct {
	Attributes []string `json:"attributes,omitempty"` // Defines the behavior of the option in the project.If notSelectable is set, the option cannot be set as the field's value. This is useful for archiving an option that has previously been selected but shouldn't be used anymore.If defaultValue is set, the option is selected by default.
	Id int64 `json:"id,omitempty"` // The ID of the project that the option's behavior applies to.
}

// CreateUiModificationDetails represents the CreateUiModificationDetails schema from the OpenAPI specification
type CreateUiModificationDetails struct {
	Contexts []UiModificationContextDetails `json:"contexts,omitempty"` // List of contexts of the UI modification. The maximum number of contexts is 1000.
	Data string `json:"data,omitempty"` // The data of the UI modification. The maximum size of the data is 50000 characters.
	Description string `json:"description,omitempty"` // The description of the UI modification. The maximum length is 255 characters.
	Name string `json:"name"` // The name of the UI modification. The maximum length is 255 characters.
}

// ScreenTypes represents the ScreenTypes schema from the OpenAPI specification
type ScreenTypes struct {
	View int64 `json:"view,omitempty"` // The ID of the view screen.
	Create int64 `json:"create,omitempty"` // The ID of the create screen.
	DefaultField int64 `json:"default,omitempty"` // The ID of the default screen. Required when creating a screen scheme.
	Edit int64 `json:"edit,omitempty"` // The ID of the edit screen.
}

// WorkflowTransitionRulesDetails represents the WorkflowTransitionRulesDetails schema from the OpenAPI specification
type WorkflowTransitionRulesDetails struct {
	Workflowruleids []string `json:"workflowRuleIds"` // The list of connect workflow rule IDs.
	Workflowid WorkflowId `json:"workflowId"` // Properties that identify a workflow.
}

// AttachmentArchiveEntry represents the AttachmentArchiveEntry schema from the OpenAPI specification
type AttachmentArchiveEntry struct {
	Size int64 `json:"size,omitempty"`
	Abbreviatedname string `json:"abbreviatedName,omitempty"`
	Entryindex int64 `json:"entryIndex,omitempty"`
	Mediatype string `json:"mediaType,omitempty"`
	Name string `json:"name,omitempty"`
}

// JiraExpressionsComplexityValueBean represents the JiraExpressionsComplexityValueBean schema from the OpenAPI specification
type JiraExpressionsComplexityValueBean struct {
	Value int `json:"value"` // The complexity value of the current expression.
	Limit int `json:"limit"` // The maximum allowed complexity. The evaluation will fail if this value is exceeded.
}

// IssueTypeCreateBean represents the IssueTypeCreateBean schema from the OpenAPI specification
type IssueTypeCreateBean struct {
	Description string `json:"description,omitempty"` // The description of the issue type.
	Hierarchylevel int `json:"hierarchyLevel,omitempty"` // The hierarchy level of the issue type. Use: * `-1` for Subtask. * `0` for Base. Defaults to `0`.
	Name string `json:"name"` // The unique name for the issue type. The maximum length is 60 characters.
	TypeField string `json:"type,omitempty"` // Deprecated. Use `hierarchyLevel` instead. See the [deprecation notice](https://community.developer.atlassian.com/t/deprecation-of-the-epic-link-parent-link-and-other-related-fields-in-rest-apis-and-webhooks/54048) for details. Whether the issue type is `subtype` or `standard`. Defaults to `standard`.
}

// CreateWorkflowCondition represents the CreateWorkflowCondition schema from the OpenAPI specification
type CreateWorkflowCondition struct {
	TypeField string `json:"type,omitempty"` // The type of the transition rule.
	Conditions []CreateWorkflowCondition `json:"conditions,omitempty"` // The list of workflow conditions.
	Configuration map[string]interface{} `json:"configuration,omitempty"` // EXPERIMENTAL. The configuration of the transition rule.
	Operator string `json:"operator,omitempty"` // The compound condition operator.
}

// JiraExpressionAnalysis represents the JiraExpressionAnalysis schema from the OpenAPI specification
type JiraExpressionAnalysis struct {
	Errors []JiraExpressionValidationError `json:"errors,omitempty"` // A list of validation errors. Not included if the expression is valid.
	Expression string `json:"expression"` // The analysed expression.
	TypeField string `json:"type,omitempty"` // EXPERIMENTAL. The inferred type of the expression.
	Valid bool `json:"valid"` // Whether the expression is valid and the interpreter will evaluate it. Note that the expression may fail at runtime (for example, if it executes too many expensive operations).
	Complexity JiraExpressionComplexity `json:"complexity,omitempty"` // Details about the complexity of the analysed Jira expression.
}

// RestrictedPermission represents the RestrictedPermission schema from the OpenAPI specification
type RestrictedPermission struct {
	Id string `json:"id,omitempty"` // The ID of the permission. Either `id` or `key` must be specified. Use [Get all permissions](#api-rest-api-3-permissions-get) to get the list of permissions.
	Key string `json:"key,omitempty"` // The key of the permission. Either `id` or `key` must be specified. Use [Get all permissions](#api-rest-api-3-permissions-get) to get the list of permissions.
}

// CustomFieldOptionUpdate represents the CustomFieldOptionUpdate schema from the OpenAPI specification
type CustomFieldOptionUpdate struct {
	Disabled bool `json:"disabled,omitempty"` // Whether the option is disabled.
	Id string `json:"id"` // The ID of the custom field option.
	Value string `json:"value,omitempty"` // The value of the custom field option.
}

// JQLQueryWithUnknownUsers represents the JQLQueryWithUnknownUsers schema from the OpenAPI specification
type JQLQueryWithUnknownUsers struct {
	Originalquery string `json:"originalQuery,omitempty"` // The original query, for reference
	Convertedquery string `json:"convertedQuery,omitempty"` // The converted query, with accountIDs instead of user identifiers, or 'unknown' for users that could not be found
}

// NotificationSchemeEventTypeId represents the NotificationSchemeEventTypeId schema from the OpenAPI specification
type NotificationSchemeEventTypeId struct {
	Id string `json:"id"` // The ID of the notification scheme event.
}

// JsonNode represents the JsonNode schema from the OpenAPI specification
type JsonNode struct {
	Binary bool `json:"binary,omitempty"`
	Textual bool `json:"textual,omitempty"`
	Missingnode bool `json:"missingNode,omitempty"`
	Null bool `json:"null,omitempty"`
	Elements map[string]interface{} `json:"elements,omitempty"`
	Biginteger bool `json:"bigInteger,omitempty"`
	Numbertype string `json:"numberType,omitempty"`
	Floatingpointnumber bool `json:"floatingPointNumber,omitempty"`
	Valueasdouble float64 `json:"valueAsDouble,omitempty"`
	Booleanvalue bool `json:"booleanValue,omitempty"`
	Pojo bool `json:"pojo,omitempty"`
	Bigdecimal bool `json:"bigDecimal,omitempty"`
	Containernode bool `json:"containerNode,omitempty"`
	Intvalue int `json:"intValue,omitempty"`
	Longvalue int64 `json:"longValue,omitempty"`
	Valuenode bool `json:"valueNode,omitempty"`
	Decimalvalue float64 `json:"decimalValue,omitempty"`
	Fieldnames map[string]interface{} `json:"fieldNames,omitempty"`
	Valueasboolean bool `json:"valueAsBoolean,omitempty"`
	Double bool `json:"double,omitempty"`
	Boolean bool `json:"boolean,omitempty"`
	Valueasint int `json:"valueAsInt,omitempty"`
	Array bool `json:"array,omitempty"`
	Long bool `json:"long,omitempty"`
	Bigintegervalue int `json:"bigIntegerValue,omitempty"`
	IntField bool `json:"int,omitempty"`
	Object bool `json:"object,omitempty"`
	Textvalue string `json:"textValue,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
	Integralnumber bool `json:"integralNumber,omitempty"`
	Valueastext string `json:"valueAsText,omitempty"`
	Binaryvalue []string `json:"binaryValue,omitempty"`
	Numbervalue float64 `json:"numberValue,omitempty"`
	Number bool `json:"number,omitempty"`
	Valueaslong int64 `json:"valueAsLong,omitempty"`
	Doublevalue float64 `json:"doubleValue,omitempty"`
}

// SecuritySchemeLevelMemberBean represents the SecuritySchemeLevelMemberBean schema from the OpenAPI specification
type SecuritySchemeLevelMemberBean struct {
	Parameter string `json:"parameter,omitempty"` // The value corresponding to the specified member type.
	TypeField string `json:"type"` // The issue security level member type, e.g `reporter`, `group`, `user`.
}

// AppWorkflowTransitionRule represents the AppWorkflowTransitionRule schema from the OpenAPI specification
type AppWorkflowTransitionRule struct {
	Key string `json:"key"` // The key of the rule, as defined in the Connect or the Forge app descriptor.
	Transition interface{} `json:"transition,omitempty"`
	Configuration RuleConfiguration `json:"configuration"` // A rule configuration.
	Id string `json:"id"` // The ID of the transition rule.
}

// ProjectIdentifierBean represents the ProjectIdentifierBean schema from the OpenAPI specification
type ProjectIdentifierBean struct {
	Key string `json:"key,omitempty"` // The key of the project.
	Id int64 `json:"id,omitempty"` // The ID of the project.
}

// CustomFieldContextDefaultValueURL represents the CustomFieldContextDefaultValueURL schema from the OpenAPI specification
type CustomFieldContextDefaultValueURL struct {
	Url string `json:"url"` // The default URL.
	Contextid string `json:"contextId"` // The ID of the context.
	TypeField string `json:"type"`
}

// Icon represents the Icon schema from the OpenAPI specification
type Icon struct {
	Title string `json:"title,omitempty"` // The title of the icon. This is used as follows: * For a status icon it is used as a tooltip on the icon. If not set, the status icon doesn't display a tooltip in Jira. * For the remote object icon it is used in conjunction with the application name to display a tooltip for the link's icon. The tooltip takes the format "\[application name\] icon title". Blank itemsare excluded from the tooltip title. If both items are blank, the icon tooltop displays as "Web Link".
	Url16x16 string `json:"url16x16,omitempty"` // The URL of an icon that displays at 16x16 pixel in Jira.
	Link string `json:"link,omitempty"` // The URL of the tooltip, used only for a status icon. If not set, the status icon in Jira is not clickable.
}

// UserMigrationBean represents the UserMigrationBean schema from the OpenAPI specification
type UserMigrationBean struct {
	Accountid string `json:"accountId,omitempty"`
	Key string `json:"key,omitempty"`
	Username string `json:"username,omitempty"`
}

// ChangeDetails represents the ChangeDetails schema from the OpenAPI specification
type ChangeDetails struct {
	Field string `json:"field,omitempty"` // The name of the field changed.
	Fieldid string `json:"fieldId,omitempty"` // The ID of the field changed.
	Fieldtype string `json:"fieldtype,omitempty"` // The type of the field changed.
	From string `json:"from,omitempty"` // The details of the original value.
	Fromstring string `json:"fromString,omitempty"` // The details of the original value as a string.
	To string `json:"to,omitempty"` // The details of the new value.
}

// SearchResults represents the SearchResults schema from the OpenAPI specification
type SearchResults struct {
	Total int `json:"total,omitempty"` // The number of results on the page.
	Warningmessages []string `json:"warningMessages,omitempty"` // Any warnings related to the JQL query.
	Expand string `json:"expand,omitempty"` // Expand options that include additional search result details in the response.
	Issues []IssueBean `json:"issues,omitempty"` // The list of issues found by the search.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of results that could be on the page.
	Names map[string]interface{} `json:"names,omitempty"` // The ID and name of each field in the search results.
	Schema map[string]interface{} `json:"schema,omitempty"` // The schema describing the field types in the search results.
	Startat int `json:"startAt,omitempty"` // The index of the first item returned on the page.
}

// ProjectRoleGroup represents the ProjectRoleGroup schema from the OpenAPI specification
type ProjectRoleGroup struct {
	Name string `json:"name,omitempty"` // The name of the group. As a group's name can change, use of `groupId` is recommended to identify the group.
	Displayname string `json:"displayName,omitempty"` // The display name of the group.
	Groupid string `json:"groupId,omitempty"` // The ID of the group.
}

// PageOfStatuses represents the PageOfStatuses schema from the OpenAPI specification
type PageOfStatuses struct {
	Self string `json:"self,omitempty"` // The URL of this page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned on the page.
	Total int64 `json:"total,omitempty"` // Number of items that satisfy the search.
	Values []JiraStatus `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // The URL of the next page of results, if any.
}

// AttachmentArchiveImpl represents the AttachmentArchiveImpl schema from the OpenAPI specification
type AttachmentArchiveImpl struct {
	Totalentrycount int `json:"totalEntryCount,omitempty"` // The number of items in the archive.
	Entries []AttachmentArchiveEntry `json:"entries,omitempty"` // The list of the items included in the archive.
}

// JiraExpressionValidationError represents the JiraExpressionValidationError schema from the OpenAPI specification
type JiraExpressionValidationError struct {
	Line int `json:"line,omitempty"` // The text line in which the error occurred.
	Message string `json:"message"` // Details about the error.
	TypeField string `json:"type"` // The error type.
	Column int `json:"column,omitempty"` // The text column in which the error occurred.
	Expression string `json:"expression,omitempty"` // The part of the expression in which the error occurred.
}

// FieldConfigurationSchemeProjectAssociation represents the FieldConfigurationSchemeProjectAssociation schema from the OpenAPI specification
type FieldConfigurationSchemeProjectAssociation struct {
	Fieldconfigurationschemeid string `json:"fieldConfigurationSchemeId,omitempty"` // The ID of the field configuration scheme. If the field configuration scheme ID is `null`, the operation assigns the default field configuration scheme.
	Projectid string `json:"projectId"` // The ID of the project.
}

// PageBeanScreen represents the PageBeanScreen schema from the OpenAPI specification
type PageBeanScreen struct {
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []Screen `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
}

// IconBean represents the IconBean schema from the OpenAPI specification
type IconBean struct {
	Title string `json:"title,omitempty"` // The title of the icon, for use as a tooltip on the icon.
	Url16x16 string `json:"url16x16,omitempty"` // The URL of a 16x16 pixel icon.
	Link string `json:"link,omitempty"` // The URL of the tooltip, used only for a status icon.
}

// ListOperand represents the ListOperand schema from the OpenAPI specification
type ListOperand struct {
	Values []JqlQueryUnitaryOperand `json:"values"` // The list of operand values.
	Encodedoperand string `json:"encodedOperand,omitempty"` // Encoded operand, which can be used directly in a JQL query.
}

// Visibility represents the Visibility schema from the OpenAPI specification
type Visibility struct {
	TypeField string `json:"type,omitempty"` // Whether visibility of this item is restricted to a group or role.
	Value string `json:"value,omitempty"` // The name of the group or role that visibility of this item is restricted to. Please note that the name of a group is mutable, to reliably identify a group use `identifier`.
	Identifier string `json:"identifier,omitempty"` // The ID of the group or the name of the role that visibility of this item is restricted to.
}

// ProjectRole represents the ProjectRole schema from the OpenAPI specification
type ProjectRole struct {
	Currentuserrole bool `json:"currentUserRole,omitempty"` // Whether the calling user is part of this role.
	Description string `json:"description,omitempty"` // The description of the project role.
	Translatedname string `json:"translatedName,omitempty"` // The translated name of the project role.
	DefaultField bool `json:"default,omitempty"` // Whether this role is the default role for the project
	Name string `json:"name,omitempty"` // The name of the project role.
	Roleconfigurable bool `json:"roleConfigurable,omitempty"` // Whether the roles are configurable for this project.
	Scope interface{} `json:"scope,omitempty"` // The scope of the role. Indicated for roles associated with [next-gen projects](https://confluence.atlassian.com/x/loMyO).
	Self string `json:"self,omitempty"` // The URL the project role details.
	Actors []RoleActor `json:"actors,omitempty"` // The list of users who act in this role.
	Admin bool `json:"admin,omitempty"` // Whether this role is the admin role for the project.
	Id int64 `json:"id,omitempty"` // The ID of the project role.
}

// NewUserDetails represents the NewUserDetails schema from the OpenAPI specification
type NewUserDetails struct {
	Password string `json:"password,omitempty"` // This property is no longer available. If the user has an Atlassian account, their password is not changed. If the user does not have an Atlassian account, they are sent an email asking them set up an account.
	Products []string `json:"products,omitempty"` // Products the new user has access to. Valid products are: jira-core, jira-servicedesk, jira-product-discovery, jira-software. If left empty, the user will get default product access. To create a user without product access, set this field to be an empty array.
	Self string `json:"self,omitempty"` // The URL of the user.
	Applicationkeys []string `json:"applicationKeys,omitempty"` // Deprecated, do not use.
	Displayname string `json:"displayName,omitempty"` // This property is no longer available. If the user has an Atlassian account, their display name is not changed. If the user does not have an Atlassian account, they are sent an email asking them set up an account.
	Emailaddress string `json:"emailAddress"` // The email address for the user.
	Key string `json:"key,omitempty"` // This property is no longer available. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Name string `json:"name,omitempty"` // This property is no longer available. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
}

// DashboardGadgetUpdateRequest represents the DashboardGadgetUpdateRequest schema from the OpenAPI specification
type DashboardGadgetUpdateRequest struct {
	Position interface{} `json:"position,omitempty"` // The position of the gadget.
	Title string `json:"title,omitempty"` // The title of the gadget.
	Color string `json:"color,omitempty"` // The color of the gadget. Should be one of `blue`, `red`, `yellow`, `green`, `cyan`, `purple`, `gray`, or `white`.
}

// PageBeanFieldConfigurationIssueTypeItem represents the PageBeanFieldConfigurationIssueTypeItem schema from the OpenAPI specification
type PageBeanFieldConfigurationIssueTypeItem struct {
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []FieldConfigurationIssueTypeItem `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
}

// ProjectCategory represents the ProjectCategory schema from the OpenAPI specification
type ProjectCategory struct {
	Description string `json:"description,omitempty"` // The description of the project category.
	Id string `json:"id,omitempty"` // The ID of the project category.
	Name string `json:"name,omitempty"` // The name of the project category. Required on create, optional on update.
	Self string `json:"self,omitempty"` // The URL of the project category.
}

// MultipleCustomFieldValuesUpdate represents the MultipleCustomFieldValuesUpdate schema from the OpenAPI specification
type MultipleCustomFieldValuesUpdate struct {
	Customfield string `json:"customField"` // The ID or key of the custom field. For example, `customfield_10010`.
	Issueids []int64 `json:"issueIds"` // The list of issue IDs.
	Value interface{} `json:"value"` // The value for the custom field. The value must be compatible with the [custom field type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/#data-types) as follows: * `string` the value must be a string. * `number` the value must be a number. * `datetime` the value must be a string that represents a date in the ISO format or the simplified extended ISO format. For example, `"2023-01-18T12:00:00-03:00"` or `"2023-01-18T12:00:00.000Z"`. However, the milliseconds part is ignored. * `user` the value must be an object that contains the `accountId` field. * `group` the value must be an object that contains the group `name` or `groupId` field. Because group names can change, we recommend using `groupId`. A list of appropriate values must be provided if the field is of the `list` [collection type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/#collection-types).
}

// FilterSubscription represents the FilterSubscription schema from the OpenAPI specification
type FilterSubscription struct {
	Group interface{} `json:"group,omitempty"` // The group subscribing to filter.
	Id int64 `json:"id,omitempty"` // The ID of the filter subscription.
	User interface{} `json:"user,omitempty"` // The user subscribing to filter.
}

// JiraExpressionsComplexityBean represents the JiraExpressionsComplexityBean schema from the OpenAPI specification
type JiraExpressionsComplexityBean struct {
	Steps interface{} `json:"steps"` // The number of steps it took to evaluate the expression, where a step is a high-level operation performed by the expression. A step is an operation such as arithmetic, accessing a property, accessing a context variable, or calling a function.
	Beans interface{} `json:"beans"` // The number of Jira REST API beans returned in the response.
	Expensiveoperations interface{} `json:"expensiveOperations"` // The number of expensive operations executed while evaluating the expression. Expensive operations are those that load additional data, such as entity properties, comments, or custom fields.
	Primitivevalues interface{} `json:"primitiveValues"` // The number of primitive values returned in the response.
}

// UpdateNotificationSchemeDetails represents the UpdateNotificationSchemeDetails schema from the OpenAPI specification
type UpdateNotificationSchemeDetails struct {
	Description string `json:"description,omitempty"` // The description of the notification scheme.
	Name string `json:"name,omitempty"` // The name of the notification scheme. Must be unique.
}

// PageBeanCustomFieldContextOption represents the PageBeanCustomFieldContextOption schema from the OpenAPI specification
type PageBeanCustomFieldContextOption struct {
	Values []CustomFieldContextOption `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// IssueChangelogIds represents the IssueChangelogIds schema from the OpenAPI specification
type IssueChangelogIds struct {
	Changelogids []int64 `json:"changelogIds"` // The list of changelog IDs.
}

// PageBeanPriority represents the PageBeanPriority schema from the OpenAPI specification
type PageBeanPriority struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []Priority `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// ContainerForRegisteredWebhooks represents the ContainerForRegisteredWebhooks schema from the OpenAPI specification
type ContainerForRegisteredWebhooks struct {
	Webhookregistrationresult []RegisteredWebhook `json:"webhookRegistrationResult,omitempty"` // A list of registered webhooks.
}

// ChangeFilterOwner represents the ChangeFilterOwner schema from the OpenAPI specification
type ChangeFilterOwner struct {
	Accountid string `json:"accountId"` // The account ID of the new owner.
}

// Configuration represents the Configuration schema from the OpenAPI specification
type Configuration struct {
	Attachmentsenabled bool `json:"attachmentsEnabled,omitempty"` // Whether the ability to add attachments to issues is enabled.
	Issuelinkingenabled bool `json:"issueLinkingEnabled,omitempty"` // Whether the ability to link issues is enabled.
	Subtasksenabled bool `json:"subTasksEnabled,omitempty"` // Whether the ability to create subtasks for issues is enabled.
	Timetrackingconfiguration interface{} `json:"timeTrackingConfiguration,omitempty"` // The configuration of time tracking.
	Timetrackingenabled bool `json:"timeTrackingEnabled,omitempty"` // Whether the ability to track time is enabled. This property is deprecated.
	Unassignedissuesallowed bool `json:"unassignedIssuesAllowed,omitempty"` // Whether the ability to create unassigned issues is enabled. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details.
	Votingenabled bool `json:"votingEnabled,omitempty"` // Whether the ability for users to vote on issues is enabled. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details.
	Watchingenabled bool `json:"watchingEnabled,omitempty"` // Whether the ability for users to watch issues is enabled. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details.
}

// CustomFieldContextUpdateDetails represents the CustomFieldContextUpdateDetails schema from the OpenAPI specification
type CustomFieldContextUpdateDetails struct {
	Description string `json:"description,omitempty"` // The description of the custom field context. The maximum length is 255 characters.
	Name string `json:"name,omitempty"` // The name of the custom field context. The name must be unique. The maximum length is 255 characters.
}

// CustomFieldUpdatedContextOptionsList represents the CustomFieldUpdatedContextOptionsList schema from the OpenAPI specification
type CustomFieldUpdatedContextOptionsList struct {
	Options []CustomFieldOptionUpdate `json:"options,omitempty"` // The updated custom field options.
}

// NotificationRecipients represents the NotificationRecipients schema from the OpenAPI specification
type NotificationRecipients struct {
	Watchers bool `json:"watchers,omitempty"` // Whether the notification should be sent to the issue's watchers.
	Assignee bool `json:"assignee,omitempty"` // Whether the notification should be sent to the issue's assignees.
	Groupids []string `json:"groupIds,omitempty"` // List of groupIds to receive the notification.
	Groups []GroupName `json:"groups,omitempty"` // List of groups to receive the notification.
	Reporter bool `json:"reporter,omitempty"` // Whether the notification should be sent to the issue's reporter.
	Users []UserDetails `json:"users,omitempty"` // List of users to receive the notification.
	Voters bool `json:"voters,omitempty"` // Whether the notification should be sent to the issue's voters.
}

// AttachmentArchive represents the AttachmentArchive schema from the OpenAPI specification
type AttachmentArchive struct {
	Totalnumberofentriesavailable int `json:"totalNumberOfEntriesAvailable,omitempty"`
	Entries []AttachmentArchiveEntry `json:"entries,omitempty"`
	Moreavailable bool `json:"moreAvailable,omitempty"`
	Totalentrycount int `json:"totalEntryCount,omitempty"`
}

// SetDefaultPriorityRequest represents the SetDefaultPriorityRequest schema from the OpenAPI specification
type SetDefaultPriorityRequest struct {
	Id string `json:"id"` // The ID of the new default issue priority. Must be an existing ID or null. Setting this to null erases the default priority setting.
}

// IssuePickerSuggestionsIssueType represents the IssuePickerSuggestionsIssueType schema from the OpenAPI specification
type IssuePickerSuggestionsIssueType struct {
	Id string `json:"id,omitempty"` // The ID of the type of issues suggested for use in auto-completion.
	Issues []SuggestedIssue `json:"issues,omitempty"` // A list of issues suggested for use in auto-completion.
	Label string `json:"label,omitempty"` // The label of the type of issues suggested for use in auto-completion.
	Msg string `json:"msg,omitempty"` // If no issue suggestions are found, returns a message indicating no suggestions were found,
	Sub string `json:"sub,omitempty"` // If issue suggestions are found, returns a message indicating the number of issues suggestions found and returned.
}

// RuleConfiguration represents the RuleConfiguration schema from the OpenAPI specification
type RuleConfiguration struct {
	Tag string `json:"tag,omitempty"` // EXPERIMENTAL: A tag used to filter rules in [Get workflow transition rule configurations](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-transition-rules/#api-rest-api-3-workflow-rule-config-get).
	Value string `json:"value"` // Configuration of the rule, as it is stored by the Connect or the Forge app on the rule configuration page.
	Disabled bool `json:"disabled,omitempty"` // EXPERIMENTAL: Whether the rule is disabled.
}

// NotificationScheme represents the NotificationScheme schema from the OpenAPI specification
type NotificationScheme struct {
	Scope interface{} `json:"scope,omitempty"` // The scope of the notification scheme.
	Self string `json:"self,omitempty"`
	Description string `json:"description,omitempty"` // The description of the notification scheme.
	Expand string `json:"expand,omitempty"` // Expand options that include additional notification scheme details in the response.
	Id int64 `json:"id,omitempty"` // The ID of the notification scheme.
	Name string `json:"name,omitempty"` // The name of the notification scheme.
	Notificationschemeevents []NotificationSchemeEvent `json:"notificationSchemeEvents,omitempty"` // The notification events and associated recipients.
	Projects []int64 `json:"projects,omitempty"` // The list of project IDs associated with the notification scheme.
}

// ActorInputBean represents the ActorInputBean schema from the OpenAPI specification
type ActorInputBean struct {
	Group []string `json:"group,omitempty"` // The name of the group to add as a default actor. This parameter cannot be used with the `groupId` parameter. As a group's name can change,use of `groupId` is recommended. This parameter accepts a comma-separated list. For example, `"group":["project-admin", "jira-developers"]`.
	Groupid []string `json:"groupId,omitempty"` // The ID of the group to add as a default actor. This parameter cannot be used with the `group` parameter This parameter accepts a comma-separated list. For example, `"groupId":["77f6ab39-e755-4570-a6ae-2d7a8df0bcb8", "0c011f85-69ed-49c4-a801-3b18d0f771bc"]`.
	User []string `json:"user,omitempty"` // The account IDs of the users to add as default actors. This parameter accepts a comma-separated list. For example, `"user":["5b10a2844c20165700ede21g", "5b109f2e9729b51b54dc274d"]`.
}

// FieldConfiguration represents the FieldConfiguration schema from the OpenAPI specification
type FieldConfiguration struct {
	Description string `json:"description"` // The description of the field configuration.
	Id int64 `json:"id"` // The ID of the field configuration.
	Isdefault bool `json:"isDefault,omitempty"` // Whether the field configuration is the default.
	Name string `json:"name"` // The name of the field configuration.
}

// HistoryMetadata represents the HistoryMetadata schema from the OpenAPI specification
type HistoryMetadata struct {
	Emaildescriptionkey string `json:"emailDescriptionKey,omitempty"` // The description key of the email address associated the history record.
	Actor interface{} `json:"actor,omitempty"` // Details of the user whose action created the history record.
	Descriptionkey string `json:"descriptionKey,omitempty"` // The description key of the history record.
	Extradata map[string]interface{} `json:"extraData,omitempty"` // Additional arbitrary information about the history record.
	Cause interface{} `json:"cause,omitempty"` // Details of the cause that triggered the creation the history record.
	Emaildescription string `json:"emailDescription,omitempty"` // The description of the email address associated the history record.
	Generator interface{} `json:"generator,omitempty"` // Details of the system that generated the history record.
	Activitydescription string `json:"activityDescription,omitempty"` // The activity described in the history record.
	Activitydescriptionkey string `json:"activityDescriptionKey,omitempty"` // The key of the activity described in the history record.
	TypeField string `json:"type,omitempty"` // The type of the history record.
	Description string `json:"description,omitempty"` // The description of the history record.
}

// UpdateIssueSecurityLevelDetails represents the UpdateIssueSecurityLevelDetails schema from the OpenAPI specification
type UpdateIssueSecurityLevelDetails struct {
	Description string `json:"description,omitempty"` // The description of the issue security scheme level.
	Name string `json:"name,omitempty"` // The name of the issue security scheme level. Must be unique.
}

// ConnectCustomFieldValue represents the ConnectCustomFieldValue schema from the OpenAPI specification
type ConnectCustomFieldValue struct {
	Fieldid int `json:"fieldID"` // The custom field ID.
	Issueid int `json:"issueID"` // The issue ID.
	Number float64 `json:"number,omitempty"` // The value of number type custom field when `_type` is `NumberIssueField`.
	Optionid string `json:"optionID,omitempty"` // The value of single select and multiselect custom field type when `_type` is `SingleSelectIssueField` or `MultiSelectIssueField`.
	Richtext string `json:"richText,omitempty"` // The value of richText type custom field when `_type` is `RichTextIssueField`.
	StringField string `json:"string,omitempty"` // The value of string type custom field when `_type` is `StringIssueField`.
	Text string `json:"text,omitempty"` // The value of of text custom field type when `_type` is `TextIssueField`.
	TypeField string `json:"_type"` // The type of custom field.
}

// ContainerOfWorkflowSchemeAssociations represents the ContainerOfWorkflowSchemeAssociations schema from the OpenAPI specification
type ContainerOfWorkflowSchemeAssociations struct {
	Values []WorkflowSchemeAssociations `json:"values"` // A list of workflow schemes together with projects they are associated with.
}

// LicenseMetric represents the LicenseMetric schema from the OpenAPI specification
type LicenseMetric struct {
	Key string `json:"key,omitempty"` // The key of the license metric.
	Value string `json:"value,omitempty"` // The value for the license metric.
}

// MultiIssueEntityProperties represents the MultiIssueEntityProperties schema from the OpenAPI specification
type MultiIssueEntityProperties struct {
	Issues []IssueEntityPropertiesForMultiUpdate `json:"issues,omitempty"` // A list of issue IDs and their respective properties.
}

// LinkGroup represents the LinkGroup schema from the OpenAPI specification
type LinkGroup struct {
	Header SimpleLink `json:"header,omitempty"` // Details about the operations available in this version.
	Id string `json:"id,omitempty"`
	Links []SimpleLink `json:"links,omitempty"`
	Styleclass string `json:"styleClass,omitempty"`
	Weight int `json:"weight,omitempty"`
	Groups []LinkGroup `json:"groups,omitempty"`
}

// CustomFieldContextOption represents the CustomFieldContextOption schema from the OpenAPI specification
type CustomFieldContextOption struct {
	Id string `json:"id"` // The ID of the custom field option.
	Optionid string `json:"optionId,omitempty"` // For cascading options, the ID of the custom field option containing the cascading option.
	Value string `json:"value"` // The value of the custom field option.
	Disabled bool `json:"disabled"` // Whether the option is disabled.
}

// CustomFieldContextDefaultValueForgeMultiGroupField represents the CustomFieldContextDefaultValueForgeMultiGroupField schema from the OpenAPI specification
type CustomFieldContextDefaultValueForgeMultiGroupField struct {
	Contextid string `json:"contextId"` // The ID of the context.
	Groupids []string `json:"groupIds"` // The IDs of the default groups.
	TypeField string `json:"type"`
}

// PageBeanFieldConfigurationSchemeProjects represents the PageBeanFieldConfigurationSchemeProjects schema from the OpenAPI specification
type PageBeanFieldConfigurationSchemeProjects struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []FieldConfigurationSchemeProjects `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// License represents the License schema from the OpenAPI specification
type License struct {
	Applications []LicensedApplication `json:"applications"` // The applications under this license.
}

// StatusUpdateRequest represents the StatusUpdateRequest schema from the OpenAPI specification
type StatusUpdateRequest struct {
	Statuses []StatusUpdate `json:"statuses,omitempty"` // The list of statuses that will be updated.
}

// UpdatePriorityDetails represents the UpdatePriorityDetails schema from the OpenAPI specification
type UpdatePriorityDetails struct {
	Description string `json:"description,omitempty"` // The description of the priority.
	Iconurl string `json:"iconUrl,omitempty"` // The URL of an icon for the priority. Accepted protocols are HTTP and HTTPS. Built in icons can also be used.
	Name string `json:"name,omitempty"` // The name of the priority. Must be unique.
	Statuscolor string `json:"statusColor,omitempty"` // The status color of the priority in 3-digit or 6-digit hexadecimal format.
}

// JqlQueriesToParse represents the JqlQueriesToParse schema from the OpenAPI specification
type JqlQueriesToParse struct {
	Queries []string `json:"queries"` // A list of queries to parse.
}

// ChangedWorklog represents the ChangedWorklog schema from the OpenAPI specification
type ChangedWorklog struct {
	Properties []EntityProperty `json:"properties,omitempty"` // Details of properties associated with the change.
	Updatedtime int64 `json:"updatedTime,omitempty"` // The datetime of the change.
	Worklogid int64 `json:"worklogId,omitempty"` // The ID of the worklog.
}

// ReorderIssuePriorities represents the ReorderIssuePriorities schema from the OpenAPI specification
type ReorderIssuePriorities struct {
	After string `json:"after,omitempty"` // The ID of the priority. Required if `position` isn't provided.
	Ids []string `json:"ids"` // The list of issue IDs to be reordered. Cannot contain duplicates nor after ID.
	Position string `json:"position,omitempty"` // The position for issue priorities to be moved to. Required if `after` isn't provided.
}

// SetDefaultResolutionRequest represents the SetDefaultResolutionRequest schema from the OpenAPI specification
type SetDefaultResolutionRequest struct {
	Id string `json:"id"` // The ID of the new default issue resolution. Must be an existing ID or null. Setting this to null erases the default resolution setting.
}

// PageBeanSecurityLevelMember represents the PageBeanSecurityLevelMember schema from the OpenAPI specification
type PageBeanSecurityLevelMember struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []SecurityLevelMember `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// SimpleApplicationPropertyBean represents the SimpleApplicationPropertyBean schema from the OpenAPI specification
type SimpleApplicationPropertyBean struct {
	Value string `json:"value,omitempty"` // The new value.
	Id string `json:"id,omitempty"` // The ID of the application property.
}

// PageBeanUserKey represents the PageBeanUserKey schema from the OpenAPI specification
type PageBeanUserKey struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []UserKey `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// SecurityLevel represents the SecurityLevel schema from the OpenAPI specification
type SecurityLevel struct {
	Issuesecurityschemeid string `json:"issueSecuritySchemeId,omitempty"` // The ID of the issue level security scheme.
	Name string `json:"name,omitempty"` // The name of the issue level security item.
	Self string `json:"self,omitempty"` // The URL of the issue level security item.
	Description string `json:"description,omitempty"` // The description of the issue level security item.
	Id string `json:"id,omitempty"` // The ID of the issue level security item.
	Isdefault bool `json:"isDefault,omitempty"` // Whether the issue level security item is the default.
}

// BulkCustomFieldOptionUpdateRequest represents the BulkCustomFieldOptionUpdateRequest schema from the OpenAPI specification
type BulkCustomFieldOptionUpdateRequest struct {
	Options []CustomFieldOptionUpdate `json:"options,omitempty"` // Details of the options to update.
}

// NestedResponse represents the NestedResponse schema from the OpenAPI specification
type NestedResponse struct {
	Status int `json:"status,omitempty"`
	Warningcollection WarningCollection `json:"warningCollection,omitempty"`
	Errorcollection ErrorCollection `json:"errorCollection,omitempty"` // Error messages from an operation.
}

// FieldMetadata represents the FieldMetadata schema from the OpenAPI specification
type FieldMetadata struct {
	Allowedvalues []interface{} `json:"allowedValues,omitempty"` // The list of values allowed in the field.
	Autocompleteurl string `json:"autoCompleteUrl,omitempty"` // The URL that can be used to automatically complete the field.
	Key string `json:"key"` // The key of the field.
	Configuration map[string]interface{} `json:"configuration,omitempty"` // The configuration properties.
	Name string `json:"name"` // The name of the field.
	Operations []string `json:"operations"` // The list of operations that can be performed on the field.
	Required bool `json:"required"` // Whether the field is required.
	Defaultvalue interface{} `json:"defaultValue,omitempty"` // The default value of the field.
	Hasdefaultvalue bool `json:"hasDefaultValue,omitempty"` // Whether the field has a default value.
	Schema interface{} `json:"schema"` // The data type of the field.
}

// JiraExpressionResult represents the JiraExpressionResult schema from the OpenAPI specification
type JiraExpressionResult struct {
	Meta interface{} `json:"meta,omitempty"` // Contains various characteristics of the performed expression evaluation.
	Value interface{} `json:"value"` // The value of the evaluated expression. It may be a primitive JSON value or a Jira REST API object. (Some expressions do not produce any meaningful results—for example, an expression that returns a lambda function—if that's the case a simple string representation is returned. These string representations should not be relied upon and may change without notice.)
}

// ListWrapperCallbackGroupName represents the ListWrapperCallbackGroupName schema from the OpenAPI specification
type ListWrapperCallbackGroupName struct {
}

// WorkflowSchemeProjectAssociation represents the WorkflowSchemeProjectAssociation schema from the OpenAPI specification
type WorkflowSchemeProjectAssociation struct {
	Workflowschemeid string `json:"workflowSchemeId,omitempty"` // The ID of the workflow scheme. If the workflow scheme ID is `null`, the operation assigns the default workflow scheme.
	Projectid string `json:"projectId"` // The ID of the project.
}

// AuditRecords represents the AuditRecords schema from the OpenAPI specification
type AuditRecords struct {
	Total int64 `json:"total,omitempty"` // The total number of audit items returned.
	Limit int `json:"limit,omitempty"` // The requested or default limit on the number of audit items to be returned.
	Offset int `json:"offset,omitempty"` // The number of audit items skipped before the first item in this list.
	Records []AuditRecordBean `json:"records,omitempty"` // The list of audit items.
}

// CreateResolutionDetails represents the CreateResolutionDetails schema from the OpenAPI specification
type CreateResolutionDetails struct {
	Name string `json:"name"` // The name of the resolution. Must be unique (case-insensitive).
	Description string `json:"description,omitempty"` // The description of the resolution.
}

// CustomFieldContextDefaultValueSingleVersionPicker represents the CustomFieldContextDefaultValueSingleVersionPicker schema from the OpenAPI specification
type CustomFieldContextDefaultValueSingleVersionPicker struct {
	Versionorder string `json:"versionOrder,omitempty"` // The order the pickable versions are displayed in. If not provided, the released-first order is used. Available version orders are `"releasedFirst"` and `"unreleasedFirst"`.
	TypeField string `json:"type"`
	Versionid string `json:"versionId"` // The ID of the default version.
}

// JqlQueryFieldEntityProperty represents the JqlQueryFieldEntityProperty schema from the OpenAPI specification
type JqlQueryFieldEntityProperty struct {
	Key string `json:"key"` // The key of the property.
	Path string `json:"path"` // The path in the property value to query.
	TypeField string `json:"type,omitempty"` // The type of the property value extraction. Not available if the extraction for the property is not registered on the instance with the [Entity property](https://developer.atlassian.com/cloud/jira/platform/modules/entity-property/) module.
	Entity string `json:"entity"` // The object on which the property is set.
}

// SuggestedIssue represents the SuggestedIssue schema from the OpenAPI specification
type SuggestedIssue struct {
	Img string `json:"img,omitempty"` // The URL of the issue type's avatar.
	Key string `json:"key,omitempty"` // The key of the issue.
	Keyhtml string `json:"keyHtml,omitempty"` // The key of the issue in HTML format.
	Summary string `json:"summary,omitempty"` // The phrase containing the query string in HTML format, with the string highlighted with HTML bold tags.
	Summarytext string `json:"summaryText,omitempty"` // The phrase containing the query string, as plain text.
	Id int64 `json:"id,omitempty"` // The ID of the issue.
}

// CustomFieldContextDefaultValueTextArea represents the CustomFieldContextDefaultValueTextArea schema from the OpenAPI specification
type CustomFieldContextDefaultValueTextArea struct {
	Text string `json:"text,omitempty"` // The default text. The maximum length is 32767 characters.
	TypeField string `json:"type"`
}

// SimplifiedHierarchyLevel represents the SimplifiedHierarchyLevel schema from the OpenAPI specification
type SimplifiedHierarchyLevel struct {
	Externaluuid string `json:"externalUuid,omitempty"` // The external UUID of the hierarchy level. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	Belowlevelid int64 `json:"belowLevelId,omitempty"` // The ID of the level below this one in the hierarchy. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	Id int64 `json:"id,omitempty"` // The ID of the hierarchy level. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	Issuetypeids []int64 `json:"issueTypeIds,omitempty"` // The issue types available in this hierarchy level.
	Projectconfigurationid int64 `json:"projectConfigurationId,omitempty"` // The ID of the project configuration. This property is deprecated, see [Change oticen: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	Abovelevelid int64 `json:"aboveLevelId,omitempty"` // The ID of the level above this one in the hierarchy. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	Hierarchylevelnumber int `json:"hierarchyLevelNumber,omitempty"`
	Level int `json:"level,omitempty"` // The level of this item in the hierarchy.
	Name string `json:"name,omitempty"` // The name of this hierarchy level.
}

// WorkflowId represents the WorkflowId schema from the OpenAPI specification
type WorkflowId struct {
	Draft bool `json:"draft"` // Whether the workflow is in the draft state.
	Name string `json:"name"` // The name of the workflow.
}

// CreateNotificationSchemeDetails represents the CreateNotificationSchemeDetails schema from the OpenAPI specification
type CreateNotificationSchemeDetails struct {
	Description string `json:"description,omitempty"` // The description of the notification scheme.
	Name string `json:"name"` // The name of the notification scheme. Must be unique (case-insensitive).
	Notificationschemeevents []NotificationSchemeEventDetails `json:"notificationSchemeEvents,omitempty"` // The list of notifications which should be added to the notification scheme.
}

// ProjectRoleUser represents the ProjectRoleUser schema from the OpenAPI specification
type ProjectRoleUser struct {
	Accountid string `json:"accountId,omitempty"` // The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. Returns *unknown* if the record is deleted and corrupted, for example, as the result of a server import.
}

// CustomFieldContextDefaultValueForgeGroupField represents the CustomFieldContextDefaultValueForgeGroupField schema from the OpenAPI specification
type CustomFieldContextDefaultValueForgeGroupField struct {
	Contextid string `json:"contextId"` // The ID of the context.
	Groupid string `json:"groupId"` // The ID of the the default group.
	TypeField string `json:"type"`
}

// ProjectRoleDetails represents the ProjectRoleDetails schema from the OpenAPI specification
type ProjectRoleDetails struct {
	DefaultField bool `json:"default,omitempty"` // Whether this role is the default role for the project.
	Name string `json:"name,omitempty"` // The name of the project role.
	Roleconfigurable bool `json:"roleConfigurable,omitempty"` // Whether the roles are configurable for this project.
	Scope interface{} `json:"scope,omitempty"` // The scope of the role. Indicated for roles associated with [next-gen projects](https://confluence.atlassian.com/x/loMyO).
	Translatedname string `json:"translatedName,omitempty"` // The translated name of the project role.
	Description string `json:"description,omitempty"` // The description of the project role.
	Admin bool `json:"admin,omitempty"` // Whether this role is the admin role for the project.
	Id int64 `json:"id,omitempty"` // The ID of the project role.
	Self string `json:"self,omitempty"` // The URL the project role details.
}

// FieldConfigurationItemsDetails represents the FieldConfigurationItemsDetails schema from the OpenAPI specification
type FieldConfigurationItemsDetails struct {
	Fieldconfigurationitems []FieldConfigurationItem `json:"fieldConfigurationItems"` // Details of fields in a field configuration.
}

// SecuritySchemeLevelBean represents the SecuritySchemeLevelBean schema from the OpenAPI specification
type SecuritySchemeLevelBean struct {
	Description string `json:"description,omitempty"` // The description of the issue security scheme level.
	Isdefault bool `json:"isDefault,omitempty"` // Specifies whether the level is the default level. False by default.
	Members []SecuritySchemeLevelMemberBean `json:"members,omitempty"` // The list of level members which should be added to the issue security scheme level.
	Name string `json:"name"` // The name of the issue security scheme level. Must be unique.
}

// AddFieldBean represents the AddFieldBean schema from the OpenAPI specification
type AddFieldBean struct {
	Fieldid string `json:"fieldId"` // The ID of the field to add.
}

// TimeTrackingConfiguration represents the TimeTrackingConfiguration schema from the OpenAPI specification
type TimeTrackingConfiguration struct {
	Defaultunit string `json:"defaultUnit"` // The default unit of time applied to logged time.
	Timeformat string `json:"timeFormat"` // The format that will appear on an issue's *Time Spent* field.
	Workingdaysperweek float64 `json:"workingDaysPerWeek"` // The number of days in a working week.
	Workinghoursperday float64 `json:"workingHoursPerDay"` // The number of hours in a working day.
}

// WarningCollection represents the WarningCollection schema from the OpenAPI specification
type WarningCollection struct {
	Warnings []string `json:"warnings,omitempty"`
}

// ParsedJqlQueries represents the ParsedJqlQueries schema from the OpenAPI specification
type ParsedJqlQueries struct {
	Queries []ParsedJqlQuery `json:"queries"` // A list of parsed JQL queries.
}

// PageBeanIssueSecuritySchemeToProjectMapping represents the PageBeanIssueSecuritySchemeToProjectMapping schema from the OpenAPI specification
type PageBeanIssueSecuritySchemeToProjectMapping struct {
	Values []IssueSecuritySchemeToProjectMapping `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// ProjectIssueTypeHierarchy represents the ProjectIssueTypeHierarchy schema from the OpenAPI specification
type ProjectIssueTypeHierarchy struct {
	Hierarchy []ProjectIssueTypesHierarchyLevel `json:"hierarchy,omitempty"` // Details of an issue type hierarchy level.
	Projectid int64 `json:"projectId,omitempty"` // The ID of the project.
}

// Webhook represents the Webhook schema from the OpenAPI specification
type Webhook struct {
	Jqlfilter string `json:"jqlFilter"` // The JQL filter that specifies which issues the webhook is sent for.
	Events []string `json:"events"` // The Jira events that trigger the webhook.
	Expirationdate int64 `json:"expirationDate,omitempty"` // The date after which the webhook is no longer sent. Use [Extend webhook life](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-webhooks/#api-rest-api-3-webhook-refresh-put) to extend the date.
	Fieldidsfilter []string `json:"fieldIdsFilter,omitempty"` // A list of field IDs. When the issue changelog contains any of the fields, the webhook `jira:issue_updated` is sent. If this parameter is not present, the app is notified about all field updates.
	Id int64 `json:"id"` // The ID of the webhook.
	Issuepropertykeysfilter []string `json:"issuePropertyKeysFilter,omitempty"` // A list of issue property keys. A change of those issue properties triggers the `issue_property_set` or `issue_property_deleted` webhooks. If this parameter is not present, the app is notified about all issue property updates.
}

// SecurityScheme represents the SecurityScheme schema from the OpenAPI specification
type SecurityScheme struct {
	Description string `json:"description,omitempty"` // The description of the issue security scheme.
	Id int64 `json:"id,omitempty"` // The ID of the issue security scheme.
	Levels []SecurityLevel `json:"levels,omitempty"`
	Name string `json:"name,omitempty"` // The name of the issue security scheme.
	Self string `json:"self,omitempty"` // The URL of the issue security scheme.
	Defaultsecuritylevelid int64 `json:"defaultSecurityLevelId,omitempty"` // The ID of the default security level.
}

// StatusUpdate represents the StatusUpdate schema from the OpenAPI specification
type StatusUpdate struct {
	Description string `json:"description,omitempty"` // The description of the status.
	Id string `json:"id"` // The ID of the status.
	Name string `json:"name"` // The name of the status.
	Statuscategory string `json:"statusCategory"` // The category of the status.
}

// CustomFieldOption represents the CustomFieldOption schema from the OpenAPI specification
type CustomFieldOption struct {
	Self string `json:"self,omitempty"` // The URL of these custom field option details.
	Value string `json:"value,omitempty"` // The value of the custom field option.
}

// CreateIssueSecuritySchemeDetails represents the CreateIssueSecuritySchemeDetails schema from the OpenAPI specification
type CreateIssueSecuritySchemeDetails struct {
	Description string `json:"description,omitempty"` // The description of the issue security scheme.
	Levels []SecuritySchemeLevelBean `json:"levels,omitempty"` // The list of scheme levels which should be added to the security scheme.
	Name string `json:"name"` // The name of the issue security scheme. Must be unique (case-insensitive).
}

// HistoryMetadataParticipant represents the HistoryMetadataParticipant schema from the OpenAPI specification
type HistoryMetadataParticipant struct {
	Id string `json:"id,omitempty"` // The ID of the user or system associated with a history record.
	TypeField string `json:"type,omitempty"` // The type of the user or system associated with a history record.
	Url string `json:"url,omitempty"` // The URL of the user or system associated with a history record.
	Avatarurl string `json:"avatarUrl,omitempty"` // The URL to an avatar for the user or system associated with a history record.
	Displayname string `json:"displayName,omitempty"` // The display name of the user or system associated with a history record.
	Displaynamekey string `json:"displayNameKey,omitempty"` // The key of the display name of the user or system associated with a history record.
}

// Application represents the Application schema from the OpenAPI specification
type Application struct {
	Name string `json:"name,omitempty"` // The name of the application. Used in conjunction with the (remote) object icon title to display a tooltip for the link's icon. The tooltip takes the format "\[application name\] icon title". Blank items are excluded from the tooltip title. If both items are blank, the icon tooltop displays as "Web Link". Grouping and sorting of links may place links without an application name last.
	TypeField string `json:"type,omitempty"` // The name-spaced type of the application, used by registered rendering apps.
}

// AddNotificationsDetails represents the AddNotificationsDetails schema from the OpenAPI specification
type AddNotificationsDetails struct {
	Notificationschemeevents []NotificationSchemeEventDetails `json:"notificationSchemeEvents"` // The list of notifications which should be added to the notification scheme.
}

// IssueFieldOptionConfiguration represents the IssueFieldOptionConfiguration schema from the OpenAPI specification
type IssueFieldOptionConfiguration struct {
	Attributes []string `json:"attributes,omitempty"` // DEPRECATED
	Scope interface{} `json:"scope,omitempty"` // Defines the projects that the option is available in. If the scope is not defined, then the option is available in all projects.
}

// IssueTypeScheme represents the IssueTypeScheme schema from the OpenAPI specification
type IssueTypeScheme struct {
	Defaultissuetypeid string `json:"defaultIssueTypeId,omitempty"` // The ID of the default issue type of the issue type scheme.
	Description string `json:"description,omitempty"` // The description of the issue type scheme.
	Id string `json:"id"` // The ID of the issue type scheme.
	Isdefault bool `json:"isDefault,omitempty"` // Whether the issue type scheme is the default.
	Name string `json:"name"` // The name of the issue type scheme.
}

// ApplicationRole represents the ApplicationRole schema from the OpenAPI specification
type ApplicationRole struct {
	Numberofseats int `json:"numberOfSeats,omitempty"` // The maximum count of users on your license.
	Platform bool `json:"platform,omitempty"` // Indicates if the application role belongs to Jira platform (`jira-core`).
	Remainingseats int `json:"remainingSeats,omitempty"` // The count of users remaining on your license.
	Groupdetails []GroupName `json:"groupDetails,omitempty"` // The groups associated with the application role.
	Selectedbydefault bool `json:"selectedByDefault,omitempty"` // Determines whether this application role should be selected by default on user creation.
	Usercount int `json:"userCount,omitempty"` // The number of users counting against your license.
	Usercountdescription string `json:"userCountDescription,omitempty"` // The [type of users](https://confluence.atlassian.com/x/lRW3Ng) being counted against your license.
	Hasunlimitedseats bool `json:"hasUnlimitedSeats,omitempty"`
	Defined bool `json:"defined,omitempty"` // Deprecated.
	Name string `json:"name,omitempty"` // The display name of the application role.
	Defaultgroups []string `json:"defaultGroups,omitempty"` // The groups that are granted default access for this application role. As a group's name can change, use of `defaultGroupsDetails` is recommended to identify a groups.
	Defaultgroupsdetails []GroupName `json:"defaultGroupsDetails,omitempty"` // The groups that are granted default access for this application role.
	Groups []string `json:"groups,omitempty"` // The groups associated with the application role. As a group's name can change, use of `groupDetails` is recommended to identify a groups.
	Key string `json:"key,omitempty"` // The key of the application role.
}

// Group represents the Group schema from the OpenAPI specification
type Group struct {
	Groupid string `json:"groupId,omitempty"` // The ID of the group, which uniquely identifies the group across all Atlassian products. For example, *952d12c3-5b5b-4d04-bb32-44d383afc4b2*.
	Name string `json:"name,omitempty"` // The name of group.
	Self string `json:"self,omitempty"` // The URL for these group details.
	Users interface{} `json:"users,omitempty"` // A paginated list of the users that are members of the group. A maximum of 50 users is returned in the list, to access additional users append `[start-index:end-index]` to the expand request. For example, to access the next 50 users, use`?expand=users[51:100]`.
	Expand string `json:"expand,omitempty"` // Expand options that include additional group details in the response.
}

// CreateCustomFieldContext represents the CreateCustomFieldContext schema from the OpenAPI specification
type CreateCustomFieldContext struct {
	Id string `json:"id,omitempty"` // The ID of the context.
	Issuetypeids []string `json:"issueTypeIds,omitempty"` // The list of issue types IDs for the context. If the list is empty, the context refers to all issue types.
	Name string `json:"name"` // The name of the context.
	Projectids []string `json:"projectIds,omitempty"` // The list of project IDs associated with the context. If the list is empty, the context is global.
	Description string `json:"description,omitempty"` // The description of the context.
}

// GroupName represents the GroupName schema from the OpenAPI specification
type GroupName struct {
	Self string `json:"self,omitempty"` // The URL for these group details.
	Groupid string `json:"groupId,omitempty"` // The ID of the group, which uniquely identifies the group across all Atlassian products. For example, *952d12c3-5b5b-4d04-bb32-44d383afc4b2*.
	Name string `json:"name,omitempty"` // The name of group.
}

// RichText represents the RichText schema from the OpenAPI specification
type RichText struct {
	Empty bool `json:"empty,omitempty"`
	Emptyadf bool `json:"emptyAdf,omitempty"`
	Finalised bool `json:"finalised,omitempty"`
	Valueset bool `json:"valueSet,omitempty"`
}

// CustomFieldContextDefaultValueForgeDateTimeField represents the CustomFieldContextDefaultValueForgeDateTimeField schema from the OpenAPI specification
type CustomFieldContextDefaultValueForgeDateTimeField struct {
	TypeField string `json:"type"`
	Usecurrent bool `json:"useCurrent,omitempty"` // Whether to use the current date.
	Contextid string `json:"contextId"` // The ID of the context.
	Datetime string `json:"dateTime,omitempty"` // The default date-time in ISO format. Ignored if `useCurrent` is true.
}

// PageBeanIssueTypeToContextMapping represents the PageBeanIssueTypeToContextMapping schema from the OpenAPI specification
type PageBeanIssueTypeToContextMapping struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []IssueTypeToContextMapping `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// UpdateDefaultScreenScheme represents the UpdateDefaultScreenScheme schema from the OpenAPI specification
type UpdateDefaultScreenScheme struct {
	Screenschemeid string `json:"screenSchemeId"` // The ID of the screen scheme.
}

// PermissionHolder represents the PermissionHolder schema from the OpenAPI specification
type PermissionHolder struct {
	Value string `json:"value,omitempty"` // The identifier associated with the `type` value that defines the holder of the permission.
	Expand string `json:"expand,omitempty"` // Expand options that include additional permission holder details in the response.
	Parameter string `json:"parameter,omitempty"` // As a group's name can change, use of `value` is recommended. The identifier associated withthe `type` value that defines the holder of the permission.
	TypeField string `json:"type"` // The type of permission holder.
}

// RemoveOptionFromIssuesResult represents the RemoveOptionFromIssuesResult schema from the OpenAPI specification
type RemoveOptionFromIssuesResult struct {
	Modifiedissues []int64 `json:"modifiedIssues,omitempty"` // The IDs of the modified issues.
	Unmodifiedissues []int64 `json:"unmodifiedIssues,omitempty"` // The IDs of the unchanged issues, those issues where errors prevent modification.
	Errors interface{} `json:"errors,omitempty"` // A collection of errors related to unchanged issues. The collection size is limited, which means not all errors may be returned.
}

// GroupLabel represents the GroupLabel schema from the OpenAPI specification
type GroupLabel struct {
	Text string `json:"text,omitempty"` // The group label name.
	Title string `json:"title,omitempty"` // The title of the group label.
	TypeField string `json:"type,omitempty"` // The type of the group label.
}

// CustomFieldContextDefaultValueFloat represents the CustomFieldContextDefaultValueFloat schema from the OpenAPI specification
type CustomFieldContextDefaultValueFloat struct {
	Number float64 `json:"number"` // The default floating-point number.
	TypeField string `json:"type"`
}

// RoleActor represents the RoleActor schema from the OpenAPI specification
type RoleActor struct {
	Avatarurl string `json:"avatarUrl,omitempty"` // The avatar of the role actor.
	Displayname string `json:"displayName,omitempty"` // The display name of the role actor. For users, depending on the user’s privacy setting, this may return an alternative value for the user's name.
	Id int64 `json:"id,omitempty"` // The ID of the role actor.
	Name string `json:"name,omitempty"` // This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	TypeField string `json:"type,omitempty"` // The type of role actor.
	Actorgroup interface{} `json:"actorGroup,omitempty"`
	Actoruser interface{} `json:"actorUser,omitempty"`
}

// IssueUpdateMetadata represents the IssueUpdateMetadata schema from the OpenAPI specification
type IssueUpdateMetadata struct {
	Fields map[string]interface{} `json:"fields,omitempty"`
}

// RemoteIssueLinkRequest represents the RemoteIssueLinkRequest schema from the OpenAPI specification
type RemoteIssueLinkRequest struct {
	Application interface{} `json:"application,omitempty"` // Details of the remote application the linked item is in. For example, trello.
	Globalid string `json:"globalId,omitempty"` // An identifier for the remote item in the remote system. For example, the global ID for a remote item in Confluence would consist of the app ID and page ID, like this: `appId=456&pageId=123`. Setting this field enables the remote issue link details to be updated or deleted using remote system and item details as the record identifier, rather than using the record's Jira ID. The maximum length is 255 characters.
	Object interface{} `json:"object"` // Details of the item linked to.
	Relationship string `json:"relationship,omitempty"` // Description of the relationship between the issue and the linked item. If not set, the relationship description "links to" is used in Jira.
}

// IssuesMetaBean represents the IssuesMetaBean schema from the OpenAPI specification
type IssuesMetaBean struct {
	Jql IssuesJqlMetaDataBean `json:"jql,omitempty"` // The description of the page of issues loaded by the provided JQL query.
}

// FieldConfigurationToIssueTypeMapping represents the FieldConfigurationToIssueTypeMapping schema from the OpenAPI specification
type FieldConfigurationToIssueTypeMapping struct {
	Fieldconfigurationid string `json:"fieldConfigurationId"` // The ID of the field configuration.
	Issuetypeid string `json:"issueTypeId"` // The ID of the issue type or *default*. When set to *default* this field configuration issue type item applies to all issue types without a field configuration. An issue type can be included only once in a request.
}

// WorkflowTransitionRulesUpdateErrorDetails represents the WorkflowTransitionRulesUpdateErrorDetails schema from the OpenAPI specification
type WorkflowTransitionRulesUpdateErrorDetails struct {
	Ruleupdateerrors map[string]interface{} `json:"ruleUpdateErrors"` // A list of transition rule update errors, indexed by the transition rule ID. Any transition rule that appears here wasn't updated.
	Updateerrors []string `json:"updateErrors"` // The list of errors that specify why the workflow update failed. The workflow was not updated if the list contains any entries.
	Workflowid WorkflowId `json:"workflowId"` // Properties that identify a workflow.
}

// JiraExpressionForAnalysis represents the JiraExpressionForAnalysis schema from the OpenAPI specification
type JiraExpressionForAnalysis struct {
	Contextvariables map[string]interface{} `json:"contextVariables,omitempty"` // Context variables and their types. The type checker assumes that [common context variables](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#context-variables), such as `issue` or `project`, are available in context and sets their type. Use this property to override the default types or provide details of new variables.
	Expressions []string `json:"expressions"` // The list of Jira expressions to analyse.
}

// AttachmentSettings represents the AttachmentSettings schema from the OpenAPI specification
type AttachmentSettings struct {
	Enabled bool `json:"enabled,omitempty"` // Whether the ability to add attachments is enabled.
	Uploadlimit int64 `json:"uploadLimit,omitempty"` // The maximum size of attachments permitted, in bytes.
}

// PageBeanIssueTypeSchemeProjects represents the PageBeanIssueTypeSchemeProjects schema from the OpenAPI specification
type PageBeanIssueTypeSchemeProjects struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []IssueTypeSchemeProjects `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// DashboardDetails represents the DashboardDetails schema from the OpenAPI specification
type DashboardDetails struct {
	Name string `json:"name"` // The name of the dashboard.
	Sharepermissions []SharePermission `json:"sharePermissions"` // The share permissions for the dashboard.
	Description string `json:"description,omitempty"` // The description of the dashboard.
	Editpermissions []SharePermission `json:"editPermissions"` // The edit permissions for the dashboard.
}

// CustomFieldContextDefaultValueSingleOption represents the CustomFieldContextDefaultValueSingleOption schema from the OpenAPI specification
type CustomFieldContextDefaultValueSingleOption struct {
	Contextid string `json:"contextId"` // The ID of the context.
	Optionid string `json:"optionId"` // The ID of the default option.
	TypeField string `json:"type"`
}

// IssuesAndJQLQueries represents the IssuesAndJQLQueries schema from the OpenAPI specification
type IssuesAndJQLQueries struct {
	Jqls []string `json:"jqls"` // A list of JQL queries.
	Issueids []int64 `json:"issueIds"` // A list of issue IDs.
}

// IssueTypeScreenSchemesProjects represents the IssueTypeScreenSchemesProjects schema from the OpenAPI specification
type IssueTypeScreenSchemesProjects struct {
	Issuetypescreenscheme interface{} `json:"issueTypeScreenScheme"` // Details of an issue type screen scheme.
	Projectids []string `json:"projectIds"` // The IDs of the projects using the issue type screen scheme.
}

// IssueTypesWorkflowMapping represents the IssueTypesWorkflowMapping schema from the OpenAPI specification
type IssueTypesWorkflowMapping struct {
	Issuetypes []string `json:"issueTypes,omitempty"` // The list of issue type IDs.
	Updatedraftifneeded bool `json:"updateDraftIfNeeded,omitempty"` // Whether a draft workflow scheme is created or updated when updating an active workflow scheme. The draft is updated with the new workflow-issue types mapping. Defaults to `false`.
	Workflow string `json:"workflow,omitempty"` // The name of the workflow. Optional if updating the workflow-issue types mapping.
	Defaultmapping bool `json:"defaultMapping,omitempty"` // Whether the workflow is the default workflow for the workflow scheme.
}

// IssueEvent represents the IssueEvent schema from the OpenAPI specification
type IssueEvent struct {
	Id int64 `json:"id,omitempty"` // The ID of the event.
	Name string `json:"name,omitempty"` // The name of the event.
}

// PagedListUserDetailsApplicationUser represents the PagedListUserDetailsApplicationUser schema from the OpenAPI specification
type PagedListUserDetailsApplicationUser struct {
	Items []UserDetails `json:"items,omitempty"` // The list of items.
	Max_results int `json:"max-results,omitempty"` // The maximum number of results that could be on the page.
	Size int `json:"size,omitempty"` // The number of items on the page.
	Start_index int `json:"start-index,omitempty"` // The index of the first item returned on the page.
	End_index int `json:"end-index,omitempty"` // The index of the last item returned on the page.
}

// DeleteAndReplaceVersionBean represents the DeleteAndReplaceVersionBean schema from the OpenAPI specification
type DeleteAndReplaceVersionBean struct {
	Movefixissuesto int64 `json:"moveFixIssuesTo,omitempty"` // The ID of the version to update `fixVersion` to when the field contains the deleted version.
	Customfieldreplacementlist []CustomFieldReplacement `json:"customFieldReplacementList,omitempty"` // An array of custom field IDs (`customFieldId`) and version IDs (`moveTo`) to update when the fields contain the deleted version.
	Moveaffectedissuesto int64 `json:"moveAffectedIssuesTo,omitempty"` // The ID of the version to update `affectedVersion` to when the field contains the deleted version.
}

// UpdateScreenSchemeDetails represents the UpdateScreenSchemeDetails schema from the OpenAPI specification
type UpdateScreenSchemeDetails struct {
	Screens interface{} `json:"screens,omitempty"` // The IDs of the screens for the screen types of the screen scheme. Only screens used in classic projects are accepted.
	Description string `json:"description,omitempty"` // The description of the screen scheme. The maximum length is 255 characters.
	Name string `json:"name,omitempty"` // The name of the screen scheme. The name must be unique. The maximum length is 255 characters.
}

// NotificationRecipientsRestrictions represents the NotificationRecipientsRestrictions schema from the OpenAPI specification
type NotificationRecipientsRestrictions struct {
	Groups []GroupName `json:"groups,omitempty"` // List of group memberships required to receive the notification.
	Permissions []RestrictedPermission `json:"permissions,omitempty"` // List of permissions required to receive the notification.
	Groupids []string `json:"groupIds,omitempty"` // List of groupId memberships required to receive the notification.
}

// IssuePickerSuggestions represents the IssuePickerSuggestions schema from the OpenAPI specification
type IssuePickerSuggestions struct {
	Sections []IssuePickerSuggestionsIssueType `json:"sections,omitempty"` // A list of issues for an issue type suggested for use in auto-completion.
}

// JqlQueriesToSanitize represents the JqlQueriesToSanitize schema from the OpenAPI specification
type JqlQueriesToSanitize struct {
	Queries []JqlQueryToSanitize `json:"queries"` // The list of JQL queries to sanitize. Must contain unique values. Maximum of 20 queries.
}

// WorkflowRules represents the WorkflowRules schema from the OpenAPI specification
type WorkflowRules struct {
	Postfunctions []WorkflowTransitionRule `json:"postFunctions,omitempty"` // The workflow post functions.
	Validators []WorkflowTransitionRule `json:"validators,omitempty"` // The workflow validators.
	Conditionstree WorkflowCondition `json:"conditionsTree,omitempty"` // The workflow transition rule conditions tree.
}

// Transition represents the Transition schema from the OpenAPI specification
type Transition struct {
	Properties map[string]interface{} `json:"properties,omitempty"` // The properties of the transition.
	To string `json:"to"` // The status the transition goes to.
	TypeField string `json:"type"` // The type of the transition.
	Description string `json:"description"` // The description of the transition.
	From []string `json:"from"` // The statuses the transition can start from.
	Id string `json:"id"` // The ID of the transition.
	Name string `json:"name"` // The name of the transition.
	Screen TransitionScreenDetails `json:"screen,omitempty"` // The details of a transition screen.
	Rules WorkflowRules `json:"rules,omitempty"` // A collection of transition rules.
}

// IssueLinkTypes represents the IssueLinkTypes schema from the OpenAPI specification
type IssueLinkTypes struct {
	Issuelinktypes []IssueLinkType `json:"issueLinkTypes,omitempty"` // The issue link type bean.
}

// PageBeanFieldConfigurationScheme represents the PageBeanFieldConfigurationScheme schema from the OpenAPI specification
type PageBeanFieldConfigurationScheme struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []FieldConfigurationScheme `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// TaskProgressBeanRemoveOptionFromIssuesResult represents the TaskProgressBeanRemoveOptionFromIssuesResult schema from the OpenAPI specification
type TaskProgressBeanRemoveOptionFromIssuesResult struct {
	Finished int64 `json:"finished,omitempty"` // A timestamp recording when the task was finished.
	Lastupdate int64 `json:"lastUpdate"` // A timestamp recording when the task progress was last updated.
	Message string `json:"message,omitempty"` // Information about the progress of the task.
	Progress int64 `json:"progress"` // The progress of the task, as a percentage complete.
	Result interface{} `json:"result,omitempty"` // The result of the task execution.
	Status string `json:"status"` // The status of the task.
	Description string `json:"description,omitempty"` // The description of the task.
	Self string `json:"self"` // The URL of the task.
	Started int64 `json:"started,omitempty"` // A timestamp recording when the task was started.
	Id string `json:"id"` // The ID of the task.
	Submitted int64 `json:"submitted"` // A timestamp recording when the task was submitted.
	Submittedby int64 `json:"submittedBy"` // The ID of the user who submitted the task.
	Elapsedruntime int64 `json:"elapsedRuntime"` // The execution time of the task, in milliseconds.
}

// CustomFieldCreatedContextOptionsList represents the CustomFieldCreatedContextOptionsList schema from the OpenAPI specification
type CustomFieldCreatedContextOptionsList struct {
	Options []CustomFieldContextOption `json:"options,omitempty"` // The created custom field options.
}

// JiraExpressionsAnalysis represents the JiraExpressionsAnalysis schema from the OpenAPI specification
type JiraExpressionsAnalysis struct {
	Results []JiraExpressionAnalysis `json:"results"` // The results of Jira expressions analysis.
}

// IssueCreateMetadata represents the IssueCreateMetadata schema from the OpenAPI specification
type IssueCreateMetadata struct {
	Expand string `json:"expand,omitempty"` // Expand options that include additional project details in the response.
	Projects []ProjectIssueCreateMetadata `json:"projects,omitempty"` // List of projects and their issue creation metadata.
}

// RemoteObject represents the RemoteObject schema from the OpenAPI specification
type RemoteObject struct {
	Icon interface{} `json:"icon,omitempty"` // Details of the icon for the item. If no icon is defined, the default link icon is used in Jira.
	Status interface{} `json:"status,omitempty"` // The status of the item.
	Summary string `json:"summary,omitempty"` // The summary details of the item.
	Title string `json:"title"` // The title of the item.
	Url string `json:"url"` // The URL of the item.
}

// FieldUpdateOperation represents the FieldUpdateOperation schema from the OpenAPI specification
type FieldUpdateOperation struct {
	Edit interface{} `json:"edit,omitempty"` // The value to edit in the field.
	Remove interface{} `json:"remove,omitempty"` // The value to removed from the field.
	Set interface{} `json:"set,omitempty"` // The value to set in the field.
	Add interface{} `json:"add,omitempty"` // The value to add to the field.
	CopyField interface{} `json:"copy,omitempty"` // The field value to copy from another issue.
}

// SecurityLevelMember represents the SecurityLevelMember schema from the OpenAPI specification
type SecurityLevelMember struct {
	Issuesecurityschemeid string `json:"issueSecuritySchemeId"` // The ID of the issue security scheme.
	Holder interface{} `json:"holder"` // The user or group being granted the permission. It consists of a `type` and a type-dependent `parameter`. See [Holder object](../api-group-permission-schemes/#holder-object) in *Get all permission schemes* for more information.
	Id string `json:"id"` // The ID of the issue security level member.
	Issuesecuritylevelid string `json:"issueSecurityLevelId"` // The ID of the issue security level.
}

// ScreenableTab represents the ScreenableTab schema from the OpenAPI specification
type ScreenableTab struct {
	Id int64 `json:"id,omitempty"` // The ID of the screen tab.
	Name string `json:"name"` // The name of the screen tab. The maximum length is 255 characters.
}

// JqlQueryOrderByClauseElement represents the JqlQueryOrderByClauseElement schema from the OpenAPI specification
type JqlQueryOrderByClauseElement struct {
	Direction string `json:"direction,omitempty"` // The direction in which to order the results.
	Field JqlQueryField `json:"field"` // A field used in a JQL query. See [Advanced searching - fields reference](https://confluence.atlassian.com/x/dAiiLQ) for more information about fields in JQL queries.
}

// AuditRecordBean represents the AuditRecordBean schema from the OpenAPI specification
type AuditRecordBean struct {
	Eventsource string `json:"eventSource,omitempty"` // The event the audit record originated from.
	Changedvalues []ChangedValueBean `json:"changedValues,omitempty"` // The list of values changed in the record event.
	Id int64 `json:"id,omitempty"` // The ID of the audit record.
	Remoteaddress string `json:"remoteAddress,omitempty"` // The URL of the computer where the creation of the audit record was initiated.
	Associateditems []AssociatedItemBean `json:"associatedItems,omitempty"` // The list of items associated with the changed record.
	Created string `json:"created,omitempty"` // The date and time on which the audit record was created.
	Authorkey string `json:"authorKey,omitempty"` // Deprecated, use `authorAccountId` instead. The key of the user who created the audit record.
	Category string `json:"category,omitempty"` // The category of the audit record. For a list of these categories, see the help article [Auditing in Jira applications](https://confluence.atlassian.com/x/noXKM).
	Objectitem AssociatedItemBean `json:"objectItem,omitempty"` // Details of an item associated with the changed record.
	Summary string `json:"summary,omitempty"` // The summary of the audit record.
	Description string `json:"description,omitempty"` // The description of the audit record.
}

// PageBeanCustomFieldContext represents the PageBeanCustomFieldContext schema from the OpenAPI specification
type PageBeanCustomFieldContext struct {
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []CustomFieldContext `json:"values,omitempty"` // The list of items.
}

// FieldConfigurationDetails represents the FieldConfigurationDetails schema from the OpenAPI specification
type FieldConfigurationDetails struct {
	Description string `json:"description,omitempty"` // The description of the field configuration.
	Name string `json:"name"` // The name of the field configuration. Must be unique.
}

// ScreenSchemeId represents the ScreenSchemeId schema from the OpenAPI specification
type ScreenSchemeId struct {
	Id int64 `json:"id"` // The ID of the screen scheme.
}

// CustomFieldContextDefaultValueMultiUserPicker represents the CustomFieldContextDefaultValueMultiUserPicker schema from the OpenAPI specification
type CustomFieldContextDefaultValueMultiUserPicker struct {
	Contextid string `json:"contextId"` // The ID of the context.
	TypeField string `json:"type"`
	Accountids []string `json:"accountIds"` // The IDs of the default users.
}

// WorkflowSchemeAssociations represents the WorkflowSchemeAssociations schema from the OpenAPI specification
type WorkflowSchemeAssociations struct {
	Projectids []string `json:"projectIds"` // The list of projects that use the workflow scheme.
	Workflowscheme interface{} `json:"workflowScheme"` // The workflow scheme.
}

// ProjectIssueSecurityLevels represents the ProjectIssueSecurityLevels schema from the OpenAPI specification
type ProjectIssueSecurityLevels struct {
	Levels []SecurityLevel `json:"levels"` // Issue level security items list.
}

// SanitizedJqlQueries represents the SanitizedJqlQueries schema from the OpenAPI specification
type SanitizedJqlQueries struct {
	Queries []SanitizedJqlQuery `json:"queries,omitempty"` // The list of sanitized JQL queries.
}

// WorkflowRulesSearchDetails represents the WorkflowRulesSearchDetails schema from the OpenAPI specification
type WorkflowRulesSearchDetails struct {
	Invalidrules []string `json:"invalidRules,omitempty"` // List of workflow rule IDs that do not belong to the workflow or can not be found.
	Validrules []WorkflowTransitionRules `json:"validRules,omitempty"` // List of valid workflow transition rules.
	Workflowentityid string `json:"workflowEntityId,omitempty"` // The workflow ID.
}

// UpdateResolutionDetails represents the UpdateResolutionDetails schema from the OpenAPI specification
type UpdateResolutionDetails struct {
	Description string `json:"description,omitempty"` // The description of the resolution.
	Name string `json:"name"` // The name of the resolution. Must be unique.
}

// CreatedIssues represents the CreatedIssues schema from the OpenAPI specification
type CreatedIssues struct {
	Errors []BulkOperationErrorResult `json:"errors,omitempty"` // Error details for failed issue creation requests.
	Issues []CreatedIssue `json:"issues,omitempty"` // Details of the issues created.
}

// PageBeanFieldConfigurationDetails represents the PageBeanFieldConfigurationDetails schema from the OpenAPI specification
type PageBeanFieldConfigurationDetails struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []FieldConfigurationDetails `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// Version represents the Version schema from the OpenAPI specification
type Version struct {
	Issuesstatusforfixversion interface{} `json:"issuesStatusForFixVersion,omitempty"` // If the expand option `issuesstatus` is used, returns the count of issues in this version for each of the status categories *to do*, *in progress*, *done*, and *unmapped*. The *unmapped* property contains a count of issues with a status other than *to do*, *in progress*, and *done*.
	Project string `json:"project,omitempty"` // Deprecated. Use `projectId`.
	Expand string `json:"expand,omitempty"` // Use [expand](em>#expansion) to include additional information about version in the response. This parameter accepts a comma-separated list. Expand options include: * `operations` Returns the list of operations available for this version. * `issuesstatus` Returns the count of issues in this version for each of the status categories *to do*, *in progress*, *done*, and *unmapped*. The *unmapped* property contains a count of issues with a status other than *to do*, *in progress*, and *done*. Optional for create and update.
	Userreleasedate string `json:"userReleaseDate,omitempty"` // The date on which work on this version is expected to finish, expressed in the instance's *Day/Month/Year Format* date format.
	Archived bool `json:"archived,omitempty"` // Indicates that the version is archived. Optional when creating or updating a version.
	Description string `json:"description,omitempty"` // The description of the version. Optional when creating or updating a version.
	Projectid int64 `json:"projectId,omitempty"` // The ID of the project to which this version is attached. Required when creating a version. Not applicable when updating a version.
	Releasedate string `json:"releaseDate,omitempty"` // The release date of the version. Expressed in ISO 8601 format (yyyy-mm-dd). Optional when creating or updating a version.
	Released bool `json:"released,omitempty"` // Indicates that the version is released. If the version is released a request to release again is ignored. Not applicable when creating a version. Optional when updating a version.
	Userstartdate string `json:"userStartDate,omitempty"` // The date on which work on this version is expected to start, expressed in the instance's *Day/Month/Year Format* date format.
	Operations []SimpleLink `json:"operations,omitempty"` // If the expand option `operations` is used, returns the list of operations available for this version.
	Name string `json:"name,omitempty"` // The unique name of the version. Required when creating a version. Optional when updating a version. The maximum length is 255 characters.
	Overdue bool `json:"overdue,omitempty"` // Indicates that the version is overdue.
	Self string `json:"self,omitempty"` // The URL of the version.
	Startdate string `json:"startDate,omitempty"` // The start date of the version. Expressed in ISO 8601 format (yyyy-mm-dd). Optional when creating or updating a version.
	Id string `json:"id,omitempty"` // The ID of the version.
	Moveunfixedissuesto string `json:"moveUnfixedIssuesTo,omitempty"` // The URL of the self link to the version to which all unfixed issues are moved when a version is released. Not applicable when creating a version. Optional when updating a version.
}

// CreateWorkflowTransitionRulesDetails represents the CreateWorkflowTransitionRulesDetails schema from the OpenAPI specification
type CreateWorkflowTransitionRulesDetails struct {
	Conditions interface{} `json:"conditions,omitempty"` // The workflow conditions.
	Postfunctions []CreateWorkflowTransitionRule `json:"postFunctions,omitempty"` // The workflow post functions. **Note:** The default post functions are always added to the *initial* transition, as in: "postFunctions": [ { "type": "IssueCreateFunction" }, { "type": "IssueReindexFunction" }, { "type": "FireIssueEventFunction", "configuration": { "event": { "id": "1", "name": "issue_created" } } } ] **Note:** The default post functions are always added to the *global* and *directed* transitions, as in: "postFunctions": [ { "type": "UpdateIssueStatusFunction" }, { "type": "CreateCommentFunction" }, { "type": "GenerateChangeHistoryFunction" }, { "type": "IssueReindexFunction" }, { "type": "FireIssueEventFunction", "configuration": { "event": { "id": "13", "name": "issue_generic" } } } ]
	Validators []CreateWorkflowTransitionRule `json:"validators,omitempty"` // The workflow validators. **Note:** The default permission validator is always added to the *initial* transition, as in: "validators": [ { "type": "PermissionValidator", "configuration": { "permissionKey": "CREATE_ISSUES" } } ]
}

// CustomFieldContextDefaultValueForgeMultiUserField represents the CustomFieldContextDefaultValueForgeMultiUserField schema from the OpenAPI specification
type CustomFieldContextDefaultValueForgeMultiUserField struct {
	Accountids []string `json:"accountIds"` // The IDs of the default users.
	Contextid string `json:"contextId"` // The ID of the context.
	TypeField string `json:"type"`
}

// BulkIssuePropertyUpdateRequest represents the BulkIssuePropertyUpdateRequest schema from the OpenAPI specification
type BulkIssuePropertyUpdateRequest struct {
	Expression string `json:"expression,omitempty"` // EXPERIMENTAL. The Jira expression to calculate the value of the property. The value of the expression must be an object that can be converted to JSON, such as a number, boolean, string, list, or map. The context variables available to the expression are `issue` and `user`. Issues for which the expression returns a value whose JSON representation is longer than 32768 characters are ignored.
	Filter interface{} `json:"filter,omitempty"` // The bulk operation filter.
	Value interface{} `json:"value,omitempty"` // The value of the property. The value must be a [valid](https://tools.ietf.org/html/rfc4627), non-empty JSON blob. The maximum length is 32768 characters.
}

// WorkflowTransitionRulesUpdateErrors represents the WorkflowTransitionRulesUpdateErrors schema from the OpenAPI specification
type WorkflowTransitionRulesUpdateErrors struct {
	Updateresults []WorkflowTransitionRulesUpdateErrorDetails `json:"updateResults"` // A list of workflows.
}

// CreateWorkflowTransitionRule represents the CreateWorkflowTransitionRule schema from the OpenAPI specification
type CreateWorkflowTransitionRule struct {
	Configuration map[string]interface{} `json:"configuration,omitempty"` // EXPERIMENTAL. The configuration of the transition rule.
	TypeField string `json:"type"` // The type of the transition rule.
}

// NotificationSchemeAndProjectMappingJsonBean represents the NotificationSchemeAndProjectMappingJsonBean schema from the OpenAPI specification
type NotificationSchemeAndProjectMappingJsonBean struct {
	Notificationschemeid string `json:"notificationSchemeId,omitempty"`
	Projectid string `json:"projectId,omitempty"`
}

// PageBeanProject represents the PageBeanProject schema from the OpenAPI specification
type PageBeanProject struct {
	Values []Project `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// WorkflowCondition represents the WorkflowCondition schema from the OpenAPI specification
type WorkflowCondition struct {
}

// FoundGroup represents the FoundGroup schema from the OpenAPI specification
type FoundGroup struct {
	Labels []GroupLabel `json:"labels,omitempty"`
	Name string `json:"name,omitempty"` // The name of the group. The name of a group is mutable, to reliably identify a group use ``groupId`.`
	Groupid string `json:"groupId,omitempty"` // The ID of the group, which uniquely identifies the group across all Atlassian products. For example, *952d12c3-5b5b-4d04-bb32-44d383afc4b2*.
	Html string `json:"html,omitempty"` // The group name with the matched query string highlighted with the HTML bold tag.
}

// AutoCompleteSuggestions represents the AutoCompleteSuggestions schema from the OpenAPI specification
type AutoCompleteSuggestions struct {
	Results []AutoCompleteSuggestion `json:"results,omitempty"` // The list of suggested item.
}

// DefaultWorkflow represents the DefaultWorkflow schema from the OpenAPI specification
type DefaultWorkflow struct {
	Updatedraftifneeded bool `json:"updateDraftIfNeeded,omitempty"` // Whether a draft workflow scheme is created or updated when updating an active workflow scheme. The draft is updated with the new default workflow. Defaults to `false`.
	Workflow string `json:"workflow"` // The name of the workflow to set as the default workflow.
}

// PageBeanWorkflow represents the PageBeanWorkflow schema from the OpenAPI specification
type PageBeanWorkflow struct {
	Values []Workflow `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// AddGroupBean represents the AddGroupBean schema from the OpenAPI specification
type AddGroupBean struct {
	Name string `json:"name"` // The name of the group.
}

// IssueTypeScreenSchemeUpdateDetails represents the IssueTypeScreenSchemeUpdateDetails schema from the OpenAPI specification
type IssueTypeScreenSchemeUpdateDetails struct {
	Name string `json:"name,omitempty"` // The name of the issue type screen scheme. The name must be unique. The maximum length is 255 characters.
	Description string `json:"description,omitempty"` // The description of the issue type screen scheme. The maximum length is 255 characters.
}

// ScreenSchemeDetails represents the ScreenSchemeDetails schema from the OpenAPI specification
type ScreenSchemeDetails struct {
	Name string `json:"name"` // The name of the screen scheme. The name must be unique. The maximum length is 255 characters.
	Screens interface{} `json:"screens"` // The IDs of the screens for the screen types of the screen scheme. Only screens used in classic projects are accepted.
	Description string `json:"description,omitempty"` // The description of the screen scheme. The maximum length is 255 characters.
}

// PermittedProjects represents the PermittedProjects schema from the OpenAPI specification
type PermittedProjects struct {
	Projects []ProjectIdentifierBean `json:"projects,omitempty"` // A list of projects.
}

// TimeTrackingDetails represents the TimeTrackingDetails schema from the OpenAPI specification
type TimeTrackingDetails struct {
	Remainingestimateseconds int64 `json:"remainingEstimateSeconds,omitempty"` // The remaining estimate of time needed for this issue in seconds.
	Timespent string `json:"timeSpent,omitempty"` // Time worked on this issue in readable format.
	Timespentseconds int64 `json:"timeSpentSeconds,omitempty"` // Time worked on this issue in seconds.
	Originalestimate string `json:"originalEstimate,omitempty"` // The original estimate of time needed for this issue in readable format.
	Originalestimateseconds int64 `json:"originalEstimateSeconds,omitempty"` // The original estimate of time needed for this issue in seconds.
	Remainingestimate string `json:"remainingEstimate,omitempty"` // The remaining estimate of time needed for this issue in readable format.
}

// PageBeanSecurityLevel represents the PageBeanSecurityLevel schema from the OpenAPI specification
type PageBeanSecurityLevel struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []SecurityLevel `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// SearchRequestBean represents the SearchRequestBean schema from the OpenAPI specification
type SearchRequestBean struct {
	Startat int `json:"startAt,omitempty"` // The index of the first item to return in the page of results (page offset). The base index is `0`.
	Validatequery string `json:"validateQuery,omitempty"` // Determines how to validate the JQL query and treat the validation results. Supported values: * `strict` Returns a 400 response code if any errors are found, along with a list of all errors (and warnings). * `warn` Returns all errors as warnings. * `none` No validation is performed. * `true` *Deprecated* A legacy synonym for `strict`. * `false` *Deprecated* A legacy synonym for `warn`. The default is `strict`. Note: If the JQL is not correctly formed a 400 response code is returned, regardless of the `validateQuery` value.
	Expand []string `json:"expand,omitempty"` // Use [expand](em>#expansion) to include additional information about issues in the response. Note that, unlike the majority of instances where `expand` is specified, `expand` is defined as a list of values. The expand options are: * `renderedFields` Returns field values rendered in HTML format. * `names` Returns the display name of each field. * `schema` Returns the schema describing a field type. * `transitions` Returns all possible transitions for the issue. * `operations` Returns all possible operations for the issue. * `editmeta` Returns information about how each field can be edited. * `changelog` Returns a list of recent updates to an issue, sorted by date, starting from the most recent. * `versionedRepresentations` Instead of `fields`, returns `versionedRepresentations` a JSON array containing each version of a field's value, with the highest numbered item representing the most recent version.
	Fields []string `json:"fields,omitempty"` // A list of fields to return for each issue, use it to retrieve a subset of fields. This parameter accepts a comma-separated list. Expand options include: * `*all` Returns all fields. * `*navigable` Returns navigable fields. * Any issue field, prefixed with a minus to exclude. The default is `*navigable`. Examples: * `summary,comment` Returns the summary and comments fields only. * `-description` Returns all navigable (default) fields except description. * `*all,-comment` Returns all fields except comments. Multiple `fields` parameters can be included in a request. Note: All navigable fields are returned by default. This differs from [GET issue](#api-rest-api-3-issue-issueIdOrKey-get) where the default is all fields.
	Fieldsbykeys bool `json:"fieldsByKeys,omitempty"` // Reference fields by their key (rather than ID). The default is `false`.
	Jql string `json:"jql,omitempty"` // A [JQL](https://confluence.atlassian.com/x/egORLQ) expression.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items to return per page.
	Properties []string `json:"properties,omitempty"` // A list of up to 5 issue properties to include in the results. This parameter accepts a comma-separated list.
}

// IssueTypeInfo represents the IssueTypeInfo schema from the OpenAPI specification
type IssueTypeInfo struct {
	Avatarid int64 `json:"avatarId,omitempty"` // The avatar of the issue type.
	Id int64 `json:"id,omitempty"` // The ID of the issue type.
	Name string `json:"name,omitempty"` // The name of the issue type.
}

// IssueTypeSchemeProjects represents the IssueTypeSchemeProjects schema from the OpenAPI specification
type IssueTypeSchemeProjects struct {
	Issuetypescheme interface{} `json:"issueTypeScheme"` // Details of an issue type scheme.
	Projectids []string `json:"projectIds"` // The IDs of the projects using the issue type scheme.
}

// PageOfDashboards represents the PageOfDashboards schema from the OpenAPI specification
type PageOfDashboards struct {
	Next string `json:"next,omitempty"` // The URL of the next page of results, if any.
	Prev string `json:"prev,omitempty"` // The URL of the previous page of results, if any.
	Startat int `json:"startAt,omitempty"` // The index of the first item returned on the page.
	Total int `json:"total,omitempty"` // The number of results on the page.
	Dashboards []Dashboard `json:"dashboards,omitempty"` // List of dashboards.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of results that could be on the page.
}

// DashboardGadget represents the DashboardGadget schema from the OpenAPI specification
type DashboardGadget struct {
	Id int64 `json:"id"` // The ID of the gadget instance.
	Modulekey string `json:"moduleKey,omitempty"` // The module key of the gadget type.
	Position interface{} `json:"position"` // The position of the gadget.
	Title string `json:"title"` // The title of the gadget.
	Uri string `json:"uri,omitempty"` // The URI of the gadget type.
	Color string `json:"color"` // The color of the gadget. Should be one of `blue`, `red`, `yellow`, `green`, `cyan`, `purple`, `gray`, or `white`.
}

// ComponentIssuesCount represents the ComponentIssuesCount schema from the OpenAPI specification
type ComponentIssuesCount struct {
	Issuecount int64 `json:"issueCount,omitempty"` // The count of issues assigned to a component.
	Self string `json:"self,omitempty"` // The URL for this count of issues for a component.
}

// CustomFieldContextDefaultValue represents the CustomFieldContextDefaultValue schema from the OpenAPI specification
type CustomFieldContextDefaultValue struct {
}

// CustomFieldContextDefaultValueForgeNumberField represents the CustomFieldContextDefaultValueForgeNumberField schema from the OpenAPI specification
type CustomFieldContextDefaultValueForgeNumberField struct {
	Contextid string `json:"contextId"` // The ID of the context.
	Number float64 `json:"number"` // The default floating-point number.
	TypeField string `json:"type"`
}

// IssueTypeToContextMapping represents the IssueTypeToContextMapping schema from the OpenAPI specification
type IssueTypeToContextMapping struct {
	Contextid string `json:"contextId"` // The ID of the context.
	Isanyissuetype bool `json:"isAnyIssueType,omitempty"` // Whether the context is mapped to any issue type.
	Issuetypeid string `json:"issueTypeId,omitempty"` // The ID of the issue type.
}

// IssueTypeIdsToRemove represents the IssueTypeIdsToRemove schema from the OpenAPI specification
type IssueTypeIdsToRemove struct {
	Issuetypeids []string `json:"issueTypeIds"` // The list of issue type IDs. Must contain unique values not longer than 255 characters and not be empty. Maximum of 100 IDs.
}

// SecuritySchemeMembersRequest represents the SecuritySchemeMembersRequest schema from the OpenAPI specification
type SecuritySchemeMembersRequest struct {
	Members []SecuritySchemeLevelMemberBean `json:"members,omitempty"` // The list of level members which should be added to the issue security scheme level.
}

// BulkIssueIsWatching represents the BulkIssueIsWatching schema from the OpenAPI specification
type BulkIssueIsWatching struct {
	Issuesiswatching map[string]interface{} `json:"issuesIsWatching,omitempty"` // The map of issue ID to boolean watch status.
}

// ApplicationProperty represents the ApplicationProperty schema from the OpenAPI specification
type ApplicationProperty struct {
	Defaultvalue string `json:"defaultValue,omitempty"` // The default value of the application property.
	Example string `json:"example,omitempty"`
	TypeField string `json:"type,omitempty"` // The data type of the application property.
	Id string `json:"id,omitempty"` // The ID of the application property. The ID and key are the same.
	Allowedvalues []string `json:"allowedValues,omitempty"` // The allowed values, if applicable.
	Key string `json:"key,omitempty"` // The key of the application property. The ID and key are the same.
	Desc string `json:"desc,omitempty"` // The description of the application property.
	Name string `json:"name,omitempty"` // The name of the application property.
	Value string `json:"value,omitempty"` // The new value.
}

// IssueTypeScreenSchemeMapping represents the IssueTypeScreenSchemeMapping schema from the OpenAPI specification
type IssueTypeScreenSchemeMapping struct {
	Screenschemeid string `json:"screenSchemeId"` // The ID of the screen scheme. Only screen schemes used in classic projects are accepted.
	Issuetypeid string `json:"issueTypeId"` // The ID of the issue type or *default*. Only issue types used in classic projects are accepted. An entry for *default* must be provided and defines the mapping for all issue types without a screen scheme.
}

// PermissionScheme represents the PermissionScheme schema from the OpenAPI specification
type PermissionScheme struct {
	Self string `json:"self,omitempty"` // The URL of the permission scheme.
	Description string `json:"description,omitempty"` // A description for the permission scheme.
	Expand string `json:"expand,omitempty"` // The expand options available for the permission scheme.
	Id int64 `json:"id,omitempty"` // The ID of the permission scheme.
	Name string `json:"name"` // The name of the permission scheme. Must be unique.
	Permissions []PermissionGrant `json:"permissions,omitempty"` // The permission scheme to create or update. See [About permission schemes and grants](../api-group-permission-schemes/#about-permission-schemes-and-grants) for more information.
	Scope interface{} `json:"scope,omitempty"` // The scope of the permission scheme.
}

// ResolutionId represents the ResolutionId schema from the OpenAPI specification
type ResolutionId struct {
	Id string `json:"id"` // The ID of the issue resolution.
}

// IssueList represents the IssueList schema from the OpenAPI specification
type IssueList struct {
	Issueids []string `json:"issueIds"` // The list of issue IDs.
}

// AddSecuritySchemeLevelsRequestBean represents the AddSecuritySchemeLevelsRequestBean schema from the OpenAPI specification
type AddSecuritySchemeLevelsRequestBean struct {
	Levels []SecuritySchemeLevelBean `json:"levels,omitempty"` // The list of scheme levels which should be added to the security scheme.
}

// IssueContextVariable represents the IssueContextVariable schema from the OpenAPI specification
type IssueContextVariable struct {
	Id int64 `json:"id,omitempty"` // The issue ID.
	Key string `json:"key,omitempty"` // The issue key.
	TypeField string `json:"type"` // Type of custom context variable.
}

// PageOfChangelogs represents the PageOfChangelogs schema from the OpenAPI specification
type PageOfChangelogs struct {
	Total int `json:"total,omitempty"` // The number of results on the page.
	Histories []Changelog `json:"histories,omitempty"` // The list of changelogs.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of results that could be on the page.
	Startat int `json:"startAt,omitempty"` // The index of the first item returned on the page.
}

// IssueMatches represents the IssueMatches schema from the OpenAPI specification
type IssueMatches struct {
	Matches []IssueMatchesForJQL `json:"matches"`
}

// StatusMapping represents the StatusMapping schema from the OpenAPI specification
type StatusMapping struct {
	Issuetypeid string `json:"issueTypeId"` // The ID of the issue type.
	Newstatusid string `json:"newStatusId"` // The ID of the new status.
	Statusid string `json:"statusId"` // The ID of the status.
}

// ActorsMap represents the ActorsMap schema from the OpenAPI specification
type ActorsMap struct {
	User []string `json:"user,omitempty"` // The user account ID of the user to add.
	Group []string `json:"group,omitempty"` // The name of the group to add. This parameter cannot be used with the `groupId` parameter. As a group's name can change, use of `groupId` is recommended.
	Groupid []string `json:"groupId,omitempty"` // The ID of the group to add. This parameter cannot be used with the `group` parameter.
}

// PageBeanString represents the PageBeanString schema from the OpenAPI specification
type PageBeanString struct {
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []string `json:"values,omitempty"` // The list of items.
}

// PublishDraftWorkflowScheme represents the PublishDraftWorkflowScheme schema from the OpenAPI specification
type PublishDraftWorkflowScheme struct {
	Statusmappings []StatusMapping `json:"statusMappings,omitempty"` // Mappings of statuses to new statuses for issue types.
}

// WorkflowSchemeIdName represents the WorkflowSchemeIdName schema from the OpenAPI specification
type WorkflowSchemeIdName struct {
	Id string `json:"id"` // The ID of the workflow scheme.
	Name string `json:"name"` // The name of the workflow scheme.
}

// ProjectAvatars represents the ProjectAvatars schema from the OpenAPI specification
type ProjectAvatars struct {
	Custom []Avatar `json:"custom,omitempty"` // List of avatars added to Jira. These avatars may be deleted.
	System []Avatar `json:"system,omitempty"` // List of avatars included with Jira. These avatars cannot be deleted.
}

// PageBeanContextualConfiguration represents the PageBeanContextualConfiguration schema from the OpenAPI specification
type PageBeanContextualConfiguration struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []ContextualConfiguration `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// FailedWebhook represents the FailedWebhook schema from the OpenAPI specification
type FailedWebhook struct {
	Id string `json:"id"` // The webhook ID, as sent in the `X-Atlassian-Webhook-Identifier` header with the webhook.
	Url string `json:"url"` // The original webhook destination.
	Body string `json:"body,omitempty"` // The webhook body.
	Failuretime int64 `json:"failureTime"` // The time the webhook was added to the list of failed webhooks (that is, the time of the last failed retry).
}

// IssueUpdateDetails represents the IssueUpdateDetails schema from the OpenAPI specification
type IssueUpdateDetails struct {
	Historymetadata interface{} `json:"historyMetadata,omitempty"` // Additional issue history details.
	Properties []EntityProperty `json:"properties,omitempty"` // Details of issue properties to be add or update.
	Transition interface{} `json:"transition,omitempty"` // Details of a transition. Required when performing a transition, optional when creating or editing an issue.
	Update map[string]interface{} `json:"update,omitempty"` // A Map containing the field field name and a list of operations to perform on the issue screen field. Note that fields included in here cannot be included in `fields`.
	Fields map[string]interface{} `json:"fields,omitempty"` // List of issue screen fields to update, specifying the sub-field to update and its value for each field. This field provides a straightforward option when setting a sub-field. When multiple sub-fields or other operations are required, use `update`. Fields included in here cannot be included in `update`.
}

// AssociatedItemBean represents the AssociatedItemBean schema from the OpenAPI specification
type AssociatedItemBean struct {
	Typename string `json:"typeName,omitempty"` // The type of the associated record.
	Id string `json:"id,omitempty"` // The ID of the associated record.
	Name string `json:"name,omitempty"` // The name of the associated record.
	Parentid string `json:"parentId,omitempty"` // The ID of the associated parent record.
	Parentname string `json:"parentName,omitempty"` // The name of the associated parent record.
}

// WorkflowStatus represents the WorkflowStatus schema from the OpenAPI specification
type WorkflowStatus struct {
	Id string `json:"id"` // The ID of the issue status.
	Name string `json:"name"` // The name of the status in the workflow.
	Properties map[string]interface{} `json:"properties,omitempty"` // Additional properties that modify the behavior of issues in this status. Supports the properties `jira.issue.editable` and `issueEditable` (deprecated) that indicate whether issues are editable.
}

// ErrorCollection represents the ErrorCollection schema from the OpenAPI specification
type ErrorCollection struct {
	Errors map[string]interface{} `json:"errors,omitempty"` // The list of errors by parameter returned by the operation. For example,"projectKey": "Project keys must start with an uppercase letter, followed by one or more uppercase alphanumeric characters."
	Status int `json:"status,omitempty"`
	Errormessages []string `json:"errorMessages,omitempty"` // The list of error messages produced by this operation. For example, "input parameter 'key' must be provided"
}

// Avatars represents the Avatars schema from the OpenAPI specification
type Avatars struct {
	Custom []Avatar `json:"custom,omitempty"` // Custom avatars list.
	System []Avatar `json:"system,omitempty"` // System avatars list.
}

// WebhookDetails represents the WebhookDetails schema from the OpenAPI specification
type WebhookDetails struct {
	Jqlfilter string `json:"jqlFilter"` // The JQL filter that specifies which issues the webhook is sent for. Only a subset of JQL can be used. The supported elements are: * Fields: `issueKey`, `project`, `issuetype`, `status`, `assignee`, `reporter`, `issue.property`, and `cf[id]`. For custom fields (`cf[id]`), only the epic label custom field is supported.". * Operators: `=`, `!=`, `IN`, and `NOT IN`.
	Events []string `json:"events"` // The Jira events that trigger the webhook.
	Fieldidsfilter []string `json:"fieldIdsFilter,omitempty"` // A list of field IDs. When the issue changelog contains any of the fields, the webhook `jira:issue_updated` is sent. If this parameter is not present, the app is notified about all field updates.
	Issuepropertykeysfilter []string `json:"issuePropertyKeysFilter,omitempty"` // A list of issue property keys. A change of those issue properties triggers the `issue_property_set` or `issue_property_deleted` webhooks. If this parameter is not present, the app is notified about all issue property updates.
}

// UiModificationIdentifiers represents the UiModificationIdentifiers schema from the OpenAPI specification
type UiModificationIdentifiers struct {
	Self string `json:"self"` // The URL of the UI modification.
	Id string `json:"id"` // The ID of the UI modification.
}

// DefaultShareScope represents the DefaultShareScope schema from the OpenAPI specification
type DefaultShareScope struct {
	Scope string `json:"scope"` // The scope of the default sharing for new filters and dashboards: * `AUTHENTICATED` Shared with all logged-in users. * `GLOBAL` Shared with all logged-in users. This shows as `AUTHENTICATED` in the response. * `PRIVATE` Not shared with any users.
}

// PageBeanIssueTypeScheme represents the PageBeanIssueTypeScheme schema from the OpenAPI specification
type PageBeanIssueTypeScheme struct {
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []IssueTypeScheme `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
}

// AvailableDashboardGadgetsResponse represents the AvailableDashboardGadgetsResponse schema from the OpenAPI specification
type AvailableDashboardGadgetsResponse struct {
	Gadgets []AvailableDashboardGadget `json:"gadgets"` // The list of available gadgets.
}

// FieldConfigurationIssueTypeItem represents the FieldConfigurationIssueTypeItem schema from the OpenAPI specification
type FieldConfigurationIssueTypeItem struct {
	Fieldconfigurationid string `json:"fieldConfigurationId"` // The ID of the field configuration.
	Fieldconfigurationschemeid string `json:"fieldConfigurationSchemeId"` // The ID of the field configuration scheme.
	Issuetypeid string `json:"issueTypeId"` // The ID of the issue type or *default*. When set to *default* this field configuration issue type item applies to all issue types without a field configuration.
}

// IssueFilterForBulkPropertyDelete represents the IssueFilterForBulkPropertyDelete schema from the OpenAPI specification
type IssueFilterForBulkPropertyDelete struct {
	Currentvalue interface{} `json:"currentValue,omitempty"` // The value of properties to perform the bulk operation on.
	Entityids []int64 `json:"entityIds,omitempty"` // List of issues to perform the bulk delete operation on.
}

// DashboardGadgetResponse represents the DashboardGadgetResponse schema from the OpenAPI specification
type DashboardGadgetResponse struct {
	Gadgets []DashboardGadget `json:"gadgets"` // The list of gadgets.
}

// ServerInformation represents the ServerInformation schema from the OpenAPI specification
type ServerInformation struct {
	Scminfo string `json:"scmInfo,omitempty"` // The unique identifier of the Jira version.
	Buildnumber int `json:"buildNumber,omitempty"` // The build number of the Jira version.
	Healthchecks []HealthCheckResult `json:"healthChecks,omitempty"` // Jira instance health check results. Deprecated and no longer returned.
	Versionnumbers []int `json:"versionNumbers,omitempty"` // The major, minor, and revision version numbers of the Jira version.
	Baseurl string `json:"baseUrl,omitempty"` // The base URL of the Jira instance.
	Servertime string `json:"serverTime,omitempty"` // The time in Jira when this request was responded to.
	Servertitle string `json:"serverTitle,omitempty"` // The name of the Jira instance.
	Version string `json:"version,omitempty"` // The version of Jira.
	Builddate string `json:"buildDate,omitempty"` // The timestamp when the Jira version was built.
	Deploymenttype string `json:"deploymentType,omitempty"` // The type of server deployment. This is always returned as *Cloud*.
}

// IssueMatchesForJQL represents the IssueMatchesForJQL schema from the OpenAPI specification
type IssueMatchesForJQL struct {
	Matchedissues []int64 `json:"matchedIssues"` // A list of issue IDs.
	Errors []string `json:"errors"` // A list of errors.
}

// IdBean represents the IdBean schema from the OpenAPI specification
type IdBean struct {
	Id int64 `json:"id"` // The ID of the permission scheme to associate with the project. Use the [Get all permission schemes](#api-rest-api-3-permissionscheme-get) resource to get a list of permission scheme IDs.
}

// PageBeanField represents the PageBeanField schema from the OpenAPI specification
type PageBeanField struct {
	Values []Field `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// LinkIssueRequestJsonBean represents the LinkIssueRequestJsonBean schema from the OpenAPI specification
type LinkIssueRequestJsonBean struct {
	TypeField IssueLinkType `json:"type"` // This object is used as follows: * In the [ issueLink](#api-rest-api-3-issueLink-post) resource it defines and reports on the type of link between the issues. Find a list of issue link types with [Get issue link types](#api-rest-api-3-issueLinkType-get). * In the [ issueLinkType](#api-rest-api-3-issueLinkType-post) resource it defines and reports on issue link types.
	Comment Comment `json:"comment,omitempty"` // A comment.
	Inwardissue LinkedIssue `json:"inwardIssue"` // The ID or key of a linked issue.
	Outwardissue LinkedIssue `json:"outwardIssue"` // The ID or key of a linked issue.
}

// BulkCustomFieldOptionCreateRequest represents the BulkCustomFieldOptionCreateRequest schema from the OpenAPI specification
type BulkCustomFieldOptionCreateRequest struct {
	Options []CustomFieldOptionCreate `json:"options,omitempty"` // Details of options to create.
}

// IssueTypeWorkflowMapping represents the IssueTypeWorkflowMapping schema from the OpenAPI specification
type IssueTypeWorkflowMapping struct {
	Updatedraftifneeded bool `json:"updateDraftIfNeeded,omitempty"` // Set to true to create or update the draft of a workflow scheme and update the mapping in the draft, when the workflow scheme cannot be edited. Defaults to `false`. Only applicable when updating the workflow-issue types mapping.
	Workflow string `json:"workflow,omitempty"` // The name of the workflow.
	Issuetype string `json:"issueType,omitempty"` // The ID of the issue type. Not required if updating the issue type-workflow mapping.
}

// UiModificationDetails represents the UiModificationDetails schema from the OpenAPI specification
type UiModificationDetails struct {
	Self string `json:"self"` // The URL of the UI modification.
	Contexts []UiModificationContextDetails `json:"contexts,omitempty"` // List of contexts of the UI modification. The maximum number of contexts is 1000.
	Data string `json:"data,omitempty"` // The data of the UI modification. The maximum size of the data is 50000 characters.
	Description string `json:"description,omitempty"` // The description of the UI modification. The maximum length is 255 characters.
	Id string `json:"id"` // The ID of the UI modification.
	Name string `json:"name"` // The name of the UI modification. The maximum length is 255 characters.
}

// CustomFieldContextDefaultValueForgeStringField represents the CustomFieldContextDefaultValueForgeStringField schema from the OpenAPI specification
type CustomFieldContextDefaultValueForgeStringField struct {
	Contextid string `json:"contextId"` // The ID of the context.
	Text string `json:"text,omitempty"` // The default text. The maximum length is 254 characters.
	TypeField string `json:"type"`
}

// UpdateUiModificationDetails represents the UpdateUiModificationDetails schema from the OpenAPI specification
type UpdateUiModificationDetails struct {
	Name string `json:"name,omitempty"` // The name of the UI modification. The maximum length is 255 characters.
	Contexts []UiModificationContextDetails `json:"contexts,omitempty"` // List of contexts of the UI modification. The maximum number of contexts is 1000. If provided, replaces all existing contexts.
	Data string `json:"data,omitempty"` // The data of the UI modification. The maximum size of the data is 50000 characters.
	Description string `json:"description,omitempty"` // The description of the UI modification. The maximum length is 255 characters.
}

// Votes represents the Votes schema from the OpenAPI specification
type Votes struct {
	Votes int64 `json:"votes,omitempty"` // The number of votes on the issue.
	Hasvoted bool `json:"hasVoted,omitempty"` // Whether the user making this request has voted on the issue.
	Self string `json:"self,omitempty"` // The URL of these issue vote details.
	Voters []User `json:"voters,omitempty"` // List of the users who have voted on this issue. An empty list is returned when the calling user doesn't have the *View voters and watchers* project permission.
}

// NotificationSchemeEventDetails represents the NotificationSchemeEventDetails schema from the OpenAPI specification
type NotificationSchemeEventDetails struct {
	Event interface{} `json:"event"` // The ID of the event.
	Notifications []NotificationSchemeNotificationDetails `json:"notifications"` // The list of notifications mapped to a specified event.
}

// IssueEntityProperties represents the IssueEntityProperties schema from the OpenAPI specification
type IssueEntityProperties struct {
	Entitiesids []int64 `json:"entitiesIds,omitempty"` // A list of entity property IDs.
	Properties map[string]interface{} `json:"properties,omitempty"` // A list of entity property keys and values.
}

// IssueTransition represents the IssueTransition schema from the OpenAPI specification
type IssueTransition struct {
	Fields map[string]interface{} `json:"fields,omitempty"` // Details of the fields associated with the issue transition screen. Use this information to populate `fields` and `update` in a transition request.
	Isavailable bool `json:"isAvailable,omitempty"` // Whether the transition is available to be performed.
	Isglobal bool `json:"isGlobal,omitempty"` // Whether the issue transition is global, that is, the transition is applied to issues regardless of their status.
	Name string `json:"name,omitempty"` // The name of the issue transition.
	Id string `json:"id,omitempty"` // The ID of the issue transition. Required when specifying a transition to undertake.
	Isconditional bool `json:"isConditional,omitempty"` // Whether the issue has to meet criteria before the issue transition is applied.
	Looped bool `json:"looped,omitempty"`
	Expand string `json:"expand,omitempty"` // Expand options that include additional transition details in the response.
	Hasscreen bool `json:"hasScreen,omitempty"` // Whether there is a screen associated with the issue transition.
	Isinitial bool `json:"isInitial,omitempty"` // Whether this is the initial issue transition for the workflow.
	To interface{} `json:"to,omitempty"` // Details of the issue status after the transition.
}

// SearchAutoCompleteFilter represents the SearchAutoCompleteFilter schema from the OpenAPI specification
type SearchAutoCompleteFilter struct {
	Includecollapsedfields bool `json:"includeCollapsedFields,omitempty"` // Include collapsed fields for fields that have non-unique names.
	Projectids []int64 `json:"projectIds,omitempty"` // List of project IDs used to filter the visible field details returned.
}

// Watchers represents the Watchers schema from the OpenAPI specification
type Watchers struct {
	Watchcount int `json:"watchCount,omitempty"` // The number of users watching this issue.
	Watchers []UserDetails `json:"watchers,omitempty"` // Details of the users watching this issue.
	Iswatching bool `json:"isWatching,omitempty"` // Whether the calling user is watching this issue.
	Self string `json:"self,omitempty"` // The URL of these issue watcher details.
}

// PageBeanScreenScheme represents the PageBeanScreenScheme schema from the OpenAPI specification
type PageBeanScreenScheme struct {
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []ScreenScheme `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
}

// CustomFieldReplacement represents the CustomFieldReplacement schema from the OpenAPI specification
type CustomFieldReplacement struct {
	Moveto int64 `json:"moveTo,omitempty"` // The version number to use as a replacement for the deleted version.
	Customfieldid int64 `json:"customFieldId,omitempty"` // The ID of the custom field in which to replace the version number.
}

// SecuritySchemeWithProjects represents the SecuritySchemeWithProjects schema from the OpenAPI specification
type SecuritySchemeWithProjects struct {
	Name string `json:"name"` // The name of the issue security scheme.
	Projectids []int64 `json:"projectIds,omitempty"` // The list of project IDs associated with the issue security scheme.
	Self string `json:"self"` // The URL of the issue security scheme.
	Defaultlevel int64 `json:"defaultLevel,omitempty"` // The default level ID of the issue security scheme.
	Description string `json:"description,omitempty"` // The description of the issue security scheme.
	Id int64 `json:"id"` // The ID of the issue security scheme.
}

// ConnectWorkflowTransitionRule represents the ConnectWorkflowTransitionRule schema from the OpenAPI specification
type ConnectWorkflowTransitionRule struct {
	Key string `json:"key"` // The key of the rule, as defined in the Connect app descriptor.
	Transition WorkflowTransition `json:"transition,omitempty"` // A workflow transition.
	Configuration RuleConfiguration `json:"configuration"` // A rule configuration.
	Id string `json:"id"` // The ID of the transition rule.
}

// CustomFieldContextSingleUserPickerDefaults represents the CustomFieldContextSingleUserPickerDefaults schema from the OpenAPI specification
type CustomFieldContextSingleUserPickerDefaults struct {
	Accountid string `json:"accountId"` // The ID of the default user.
	Contextid string `json:"contextId"` // The ID of the context.
	TypeField string `json:"type"`
	Userfilter UserFilter `json:"userFilter"` // Filter for a User Picker (single) custom field.
}

// CustomFieldContextDefaultValueReadOnly represents the CustomFieldContextDefaultValueReadOnly schema from the OpenAPI specification
type CustomFieldContextDefaultValueReadOnly struct {
	TypeField string `json:"type"`
	Text string `json:"text,omitempty"` // The default text. The maximum length is 255 characters.
}

// UserList represents the UserList schema from the OpenAPI specification
type UserList struct {
	End_index int `json:"end-index,omitempty"` // The index of the last item returned on the page.
	Items []User `json:"items,omitempty"` // The list of items.
	Max_results int `json:"max-results,omitempty"` // The maximum number of results that could be on the page.
	Size int `json:"size,omitempty"` // The number of items on the page.
	Start_index int `json:"start-index,omitempty"` // The index of the first item returned on the page.
}

// BulkOperationErrorResult represents the BulkOperationErrorResult schema from the OpenAPI specification
type BulkOperationErrorResult struct {
	Elementerrors ErrorCollection `json:"elementErrors,omitempty"` // Error messages from an operation.
	Failedelementnumber int `json:"failedElementNumber,omitempty"`
	Status int `json:"status,omitempty"`
}

// PageBeanComment represents the PageBeanComment schema from the OpenAPI specification
type PageBeanComment struct {
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []Comment `json:"values,omitempty"` // The list of items.
}

// NotificationSchemeNotificationDetails represents the NotificationSchemeNotificationDetails schema from the OpenAPI specification
type NotificationSchemeNotificationDetails struct {
	Notificationtype string `json:"notificationType"` // The notification type, e.g `CurrentAssignee`, `Group`, `EmailAddress`.
	Parameter string `json:"parameter,omitempty"` // The value corresponding to the specified notification type.
}

// TaskProgressBeanObject represents the TaskProgressBeanObject schema from the OpenAPI specification
type TaskProgressBeanObject struct {
	Lastupdate int64 `json:"lastUpdate"` // A timestamp recording when the task progress was last updated.
	Self string `json:"self"` // The URL of the task.
	Status string `json:"status"` // The status of the task.
	Finished int64 `json:"finished,omitempty"` // A timestamp recording when the task was finished.
	Message string `json:"message,omitempty"` // Information about the progress of the task.
	Progress int64 `json:"progress"` // The progress of the task, as a percentage complete.
	Started int64 `json:"started,omitempty"` // A timestamp recording when the task was started.
	Elapsedruntime int64 `json:"elapsedRuntime"` // The execution time of the task, in milliseconds.
	Result interface{} `json:"result,omitempty"` // The result of the task execution.
	Submitted int64 `json:"submitted"` // A timestamp recording when the task was submitted.
	Submittedby int64 `json:"submittedBy"` // The ID of the user who submitted the task.
	Description string `json:"description,omitempty"` // The description of the task.
	Id string `json:"id"` // The ID of the task.
}

// Notification represents the Notification schema from the OpenAPI specification
type Notification struct {
	Restrict interface{} `json:"restrict,omitempty"` // Restricts the notifications to users with the specified permissions.
	Subject string `json:"subject,omitempty"` // The subject of the email notification for the issue. If this is not specified, then the subject is set to the issue key and summary.
	Textbody string `json:"textBody,omitempty"` // The plain text body of the email notification for the issue.
	To interface{} `json:"to,omitempty"` // The recipients of the email notification for the issue.
	Htmlbody string `json:"htmlBody,omitempty"` // The HTML body of the email notification for the issue.
}

// TransitionScreenDetails represents the TransitionScreenDetails schema from the OpenAPI specification
type TransitionScreenDetails struct {
	Id string `json:"id"` // The ID of the screen.
	Name string `json:"name,omitempty"` // The name of the screen.
}

// SetDefaultLevelsRequest represents the SetDefaultLevelsRequest schema from the OpenAPI specification
type SetDefaultLevelsRequest struct {
	Defaultvalues []DefaultLevelValue `json:"defaultValues"` // List of objects with issue security scheme ID and new default level ID.
}

// ProjectType represents the ProjectType schema from the OpenAPI specification
type ProjectType struct {
	Descriptioni18nkey string `json:"descriptionI18nKey,omitempty"` // The key of the project type's description.
	Formattedkey string `json:"formattedKey,omitempty"` // The formatted key of the project type.
	Icon string `json:"icon,omitempty"` // The icon of the project type.
	Key string `json:"key,omitempty"` // The key of the project type.
	Color string `json:"color,omitempty"` // The color of the project type.
}

// BulkPermissionGrants represents the BulkPermissionGrants schema from the OpenAPI specification
type BulkPermissionGrants struct {
	Projectpermissions []BulkProjectPermissionGrants `json:"projectPermissions"` // List of project permissions and the projects and issues those permissions provide access to.
	Globalpermissions []string `json:"globalPermissions"` // List of permissions granted to the user.
}

// GroupDetails represents the GroupDetails schema from the OpenAPI specification
type GroupDetails struct {
	Name string `json:"name,omitempty"` // The name of the group.
	Groupid string `json:"groupId,omitempty"` // The ID of the group, which uniquely identifies the group across all Atlassian products. For example, *952d12c3-5b5b-4d04-bb32-44d383afc4b2*.
}

// ConnectModule represents the ConnectModule schema from the OpenAPI specification
type ConnectModule struct {
}

// PropertyKeys represents the PropertyKeys schema from the OpenAPI specification
type PropertyKeys struct {
	Keys []PropertyKey `json:"keys,omitempty"` // Property key details.
}

// JqlQueryToSanitize represents the JqlQueryToSanitize schema from the OpenAPI specification
type JqlQueryToSanitize struct {
	Accountid string `json:"accountId,omitempty"` // The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*.
	Query string `json:"query"` // The query to sanitize.
}

// DeprecatedWorkflow represents the DeprecatedWorkflow schema from the OpenAPI specification
type DeprecatedWorkflow struct {
	Description string `json:"description,omitempty"` // The description of the workflow.
	Lastmodifieddate string `json:"lastModifiedDate,omitempty"` // The datetime the workflow was last modified.
	Lastmodifieduser string `json:"lastModifiedUser,omitempty"` // This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Lastmodifieduseraccountid string `json:"lastModifiedUserAccountId,omitempty"` // The account ID of the user that last modified the workflow.
	Name string `json:"name,omitempty"` // The name of the workflow.
	Scope interface{} `json:"scope,omitempty"` // The scope where this workflow applies
	Steps int `json:"steps,omitempty"` // The number of steps included in the workflow.
	DefaultField bool `json:"default,omitempty"`
}

// ContainerForProjectFeatures represents the ContainerForProjectFeatures schema from the OpenAPI specification
type ContainerForProjectFeatures struct {
	Features []ProjectFeature `json:"features,omitempty"` // The project features.
}

// PermissionsKeysBean represents the PermissionsKeysBean schema from the OpenAPI specification
type PermissionsKeysBean struct {
	Permissions []string `json:"permissions"` // A list of permission keys.
}

// AutoCompleteSuggestion represents the AutoCompleteSuggestion schema from the OpenAPI specification
type AutoCompleteSuggestion struct {
	Displayname string `json:"displayName,omitempty"` // The display name of a suggested item. If `fieldValue` or `predicateValue` are provided, the matching text is highlighted with the HTML bold tag.
	Value string `json:"value,omitempty"` // The value of a suggested item.
}

// IssuesUpdateBean represents the IssuesUpdateBean schema from the OpenAPI specification
type IssuesUpdateBean struct {
	Issueupdates []IssueUpdateDetails `json:"issueUpdates,omitempty"`
}

// GlobalScopeBean represents the GlobalScopeBean schema from the OpenAPI specification
type GlobalScopeBean struct {
	Attributes []string `json:"attributes,omitempty"` // Defines the behavior of the option in the global context.If notSelectable is set, the option cannot be set as the field's value. This is useful for archiving an option that has previously been selected but shouldn't be used anymore.If defaultValue is set, the option is selected by default.
}

// JqlQueryClauseTimePredicate represents the JqlQueryClauseTimePredicate schema from the OpenAPI specification
type JqlQueryClauseTimePredicate struct {
	Operand JqlQueryClauseOperand `json:"operand"` // Details of an operand in a JQL clause.
	Operator string `json:"operator"` // The operator between the field and the operand.
}

// EntityPropertyDetails represents the EntityPropertyDetails schema from the OpenAPI specification
type EntityPropertyDetails struct {
	Entityid float64 `json:"entityId"` // The entity property ID.
	Key string `json:"key"` // The entity property key.
	Value string `json:"value"` // The new value of the entity property.
}

// MoveFieldBean represents the MoveFieldBean schema from the OpenAPI specification
type MoveFieldBean struct {
	After string `json:"after,omitempty"` // The ID of the screen tab field after which to place the moved screen tab field. Required if `position` isn't provided.
	Position string `json:"position,omitempty"` // The named position to which the screen tab field should be moved. Required if `after` isn't provided.
}

// FoundGroups represents the FoundGroups schema from the OpenAPI specification
type FoundGroups struct {
	Groups []FoundGroup `json:"groups,omitempty"`
	Header string `json:"header,omitempty"` // Header text indicating the number of groups in the response and the total number of groups found in the search.
	Total int `json:"total,omitempty"` // The total number of groups found in the search.
}

// ScreenWithTab represents the ScreenWithTab schema from the OpenAPI specification
type ScreenWithTab struct {
	Scope interface{} `json:"scope,omitempty"` // The scope of the screen.
	Tab interface{} `json:"tab,omitempty"` // The tab for the screen.
	Description string `json:"description,omitempty"` // The description of the screen.
	Id int64 `json:"id,omitempty"` // The ID of the screen.
	Name string `json:"name,omitempty"` // The name of the screen.
}

// StatusDetails represents the StatusDetails schema from the OpenAPI specification
type StatusDetails struct {
	Self string `json:"self,omitempty"` // The URL of the status.
	Statuscategory interface{} `json:"statusCategory,omitempty"` // The category assigned to the status.
	Description string `json:"description,omitempty"` // The description of the status.
	Iconurl string `json:"iconUrl,omitempty"` // The URL of the icon used to represent the status.
	Id string `json:"id,omitempty"` // The ID of the status.
	Name string `json:"name,omitempty"` // The name of the status.
}

// FieldWasClause represents the FieldWasClause schema from the OpenAPI specification
type FieldWasClause struct {
	Operand JqlQueryClauseOperand `json:"operand"` // Details of an operand in a JQL clause.
	Operator string `json:"operator"` // The operator between the field and operand.
	Predicates []JqlQueryClauseTimePredicate `json:"predicates"` // The list of time predicates.
	Field JqlQueryField `json:"field"` // A field used in a JQL query. See [Advanced searching - fields reference](https://confluence.atlassian.com/x/dAiiLQ) for more information about fields in JQL queries.
}

// Changelog represents the Changelog schema from the OpenAPI specification
type Changelog struct {
	Historymetadata interface{} `json:"historyMetadata,omitempty"` // The history metadata associated with the changed.
	Id string `json:"id,omitempty"` // The ID of the changelog.
	Items []ChangeDetails `json:"items,omitempty"` // The list of items changed.
	Author interface{} `json:"author,omitempty"` // The user who made the change.
	Created string `json:"created,omitempty"` // The date on which the change took place.
}

// FieldConfigurationItem represents the FieldConfigurationItem schema from the OpenAPI specification
type FieldConfigurationItem struct {
	Description string `json:"description,omitempty"` // The description of the field within the field configuration.
	Id string `json:"id"` // The ID of the field within the field configuration.
	Ishidden bool `json:"isHidden,omitempty"` // Whether the field is hidden in the field configuration.
	Isrequired bool `json:"isRequired,omitempty"` // Whether the field is required in the field configuration.
	Renderer string `json:"renderer,omitempty"` // The renderer type for the field within the field configuration.
}

// ProjectIssueCreateMetadata represents the ProjectIssueCreateMetadata schema from the OpenAPI specification
type ProjectIssueCreateMetadata struct {
	Self string `json:"self,omitempty"` // The URL of the project.
	Avatarurls interface{} `json:"avatarUrls,omitempty"` // List of the project's avatars, returning the avatar size and associated URL.
	Expand string `json:"expand,omitempty"` // Expand options that include additional project issue create metadata details in the response.
	Id string `json:"id,omitempty"` // The ID of the project.
	Issuetypes []IssueTypeIssueCreateMetadata `json:"issuetypes,omitempty"` // List of the issue types supported by the project.
	Key string `json:"key,omitempty"` // The key of the project.
	Name string `json:"name,omitempty"` // The name of the project.
}

// JsonContextVariable represents the JsonContextVariable schema from the OpenAPI specification
type JsonContextVariable struct {
	TypeField string `json:"type"` // Type of custom context variable.
	Value map[string]interface{} `json:"value,omitempty"` // A JSON object containing custom content.
}

// NotificationSchemeEvent represents the NotificationSchemeEvent schema from the OpenAPI specification
type NotificationSchemeEvent struct {
	Event NotificationEvent `json:"event,omitempty"` // Details about a notification event.
	Notifications []EventNotification `json:"notifications,omitempty"`
}

// DashboardGadgetSettings represents the DashboardGadgetSettings schema from the OpenAPI specification
type DashboardGadgetSettings struct {
	Modulekey string `json:"moduleKey,omitempty"` // The module key of the gadget type. Can't be provided with `uri`.
	Position interface{} `json:"position,omitempty"` // The position of the gadget. When the gadget is placed into the position, other gadgets in the same column are moved down to accommodate it.
	Title string `json:"title,omitempty"` // The title of the gadget.
	Uri string `json:"uri,omitempty"` // The URI of the gadget type. Can't be provided with `moduleKey`.
	Color string `json:"color,omitempty"` // The color of the gadget. Should be one of `blue`, `red`, `yellow`, `green`, `cyan`, `purple`, `gray`, or `white`.
	Ignoreuriandmodulekeyvalidation bool `json:"ignoreUriAndModuleKeyValidation,omitempty"` // Whether to ignore the validation of module key and URI. For example, when a gadget is created that is a part of an application that isn't installed.
}

// LinkedIssue represents the LinkedIssue schema from the OpenAPI specification
type LinkedIssue struct {
	Fields interface{} `json:"fields,omitempty"` // The fields associated with the issue.
	Id string `json:"id,omitempty"` // The ID of an issue. Required if `key` isn't provided.
	Key string `json:"key,omitempty"` // The key of an issue. Required if `id` isn't provided.
	Self string `json:"self,omitempty"` // The URL of the issue.
}

// HealthCheckResult represents the HealthCheckResult schema from the OpenAPI specification
type HealthCheckResult struct {
	Description string `json:"description,omitempty"` // The description of the Jira health check item.
	Name string `json:"name,omitempty"` // The name of the Jira health check item.
	Passed bool `json:"passed,omitempty"` // Whether the Jira health check item passed or failed.
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Locale string `json:"locale,omitempty"` // The locale of the user. Depending on the user’s privacy setting, this may be returned as null.
	Self string `json:"self,omitempty"` // The URL of the user.
	Emailaddress string `json:"emailAddress,omitempty"` // The email address of the user. Depending on the user’s privacy setting, this may be returned as null.
	Key string `json:"key,omitempty"` // This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Accounttype string `json:"accountType,omitempty"` // The user account type. Can take the following values: * `atlassian` regular Atlassian user account * `app` system account used for Connect applications and OAuth to represent external systems * `customer` Jira Service Desk account representing an external service desk
	Applicationroles interface{} `json:"applicationRoles,omitempty"` // The application roles the user is assigned to.
	Avatarurls interface{} `json:"avatarUrls,omitempty"` // The avatars of the user.
	Timezone string `json:"timeZone,omitempty"` // The time zone specified in the user's profile. Depending on the user’s privacy setting, this may be returned as null.
	Active bool `json:"active,omitempty"` // Whether the user is active.
	Expand string `json:"expand,omitempty"` // Expand options that include additional user details in the response.
	Groups interface{} `json:"groups,omitempty"` // The groups that the user belongs to.
	Name string `json:"name,omitempty"` // This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Accountid string `json:"accountId,omitempty"` // The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. Required in requests.
	Displayname string `json:"displayName,omitempty"` // The display name of the user. Depending on the user’s privacy setting, this may return an alternative value.
}

// PermissionSchemes represents the PermissionSchemes schema from the OpenAPI specification
type PermissionSchemes struct {
	Permissionschemes []PermissionScheme `json:"permissionSchemes,omitempty"` // Permission schemes list.
}

// UpdateScreenDetails represents the UpdateScreenDetails schema from the OpenAPI specification
type UpdateScreenDetails struct {
	Description string `json:"description,omitempty"` // The description of the screen. The maximum length is 255 characters.
	Name string `json:"name,omitempty"` // The name of the screen. The name must be unique. The maximum length is 255 characters.
}

// ProjectFeature represents the ProjectFeature schema from the OpenAPI specification
type ProjectFeature struct {
	Togglelocked bool `json:"toggleLocked,omitempty"` // Whether the state of the feature can be updated.
	Feature string `json:"feature,omitempty"` // The key of the feature.
	Imageuri string `json:"imageUri,omitempty"` // URI for the image representing the feature.
	Localiseddescription string `json:"localisedDescription,omitempty"` // Localized display description for the feature.
	Localisedname string `json:"localisedName,omitempty"` // Localized display name for the feature.
	Prerequisites []string `json:"prerequisites,omitempty"` // List of keys of the features required to enable the feature.
	Projectid int64 `json:"projectId,omitempty"` // The ID of the project.
	State string `json:"state,omitempty"` // The state of the feature. When updating the state of a feature, only ENABLED and DISABLED are supported. Responses can contain all values
}

// SanitizedJqlQuery represents the SanitizedJqlQuery schema from the OpenAPI specification
type SanitizedJqlQuery struct {
	Errors interface{} `json:"errors,omitempty"` // The list of errors.
	Initialquery string `json:"initialQuery,omitempty"` // The initial query.
	Sanitizedquery string `json:"sanitizedQuery,omitempty"` // The sanitized query, if there were no errors.
	Accountid string `json:"accountId,omitempty"` // The account ID of the user for whom sanitization was performed.
}

// CustomFieldContextDefaultValueMultipleVersionPicker represents the CustomFieldContextDefaultValueMultipleVersionPicker schema from the OpenAPI specification
type CustomFieldContextDefaultValueMultipleVersionPicker struct {
	Versionids []string `json:"versionIds"` // The IDs of the default versions.
	Versionorder string `json:"versionOrder,omitempty"` // The order the pickable versions are displayed in. If not provided, the released-first order is used. Available version orders are `"releasedFirst"` and `"unreleasedFirst"`.
	TypeField string `json:"type"`
}

// CustomFieldContextDefaultValueMultipleGroupPicker represents the CustomFieldContextDefaultValueMultipleGroupPicker schema from the OpenAPI specification
type CustomFieldContextDefaultValueMultipleGroupPicker struct {
	Contextid string `json:"contextId"` // The ID of the context.
	Groupids []string `json:"groupIds"` // The IDs of the default groups.
	TypeField string `json:"type"`
}

// PageBeanProjectDetails represents the PageBeanProjectDetails schema from the OpenAPI specification
type PageBeanProjectDetails struct {
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []ProjectDetails `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
}

// CustomFieldContextDefaultValueForgeMultiStringField represents the CustomFieldContextDefaultValueForgeMultiStringField schema from the OpenAPI specification
type CustomFieldContextDefaultValueForgeMultiStringField struct {
	TypeField string `json:"type"`
	Values []string `json:"values,omitempty"` // List of string values. The maximum length for a value is 254 characters.
}

// PageOfWorklogs represents the PageOfWorklogs schema from the OpenAPI specification
type PageOfWorklogs struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of results that could be on the page.
	Startat int `json:"startAt,omitempty"` // The index of the first item returned on the page.
	Total int `json:"total,omitempty"` // The number of results on the page.
	Worklogs []Worklog `json:"worklogs,omitempty"` // List of worklogs.
}

// BulkPermissionsRequestBean represents the BulkPermissionsRequestBean schema from the OpenAPI specification
type BulkPermissionsRequestBean struct {
	Accountid string `json:"accountId,omitempty"` // The account ID of a user.
	Globalpermissions []string `json:"globalPermissions,omitempty"` // Global permissions to look up.
	Projectpermissions []BulkProjectPermissions `json:"projectPermissions,omitempty"` // Project permissions with associated projects and issues to look up.
}

// Dashboard represents the Dashboard schema from the OpenAPI specification
type Dashboard struct {
	Iswritable bool `json:"isWritable,omitempty"` // Whether the current user has permission to edit the dashboard.
	Popularity int64 `json:"popularity,omitempty"` // The number of users who have this dashboard as a favorite.
	Name string `json:"name,omitempty"` // The name of the dashboard.
	Rank int `json:"rank,omitempty"` // The rank of this dashboard.
	Systemdashboard bool `json:"systemDashboard,omitempty"` // Whether the current dashboard is system dashboard.
	Owner interface{} `json:"owner,omitempty"` // The owner of the dashboard.
	Self string `json:"self,omitempty"` // The URL of these dashboard details.
	Sharepermissions []SharePermission `json:"sharePermissions,omitempty"` // The details of any view share permissions for the dashboard.
	Description string `json:"description,omitempty"`
	Editpermissions []SharePermission `json:"editPermissions,omitempty"` // The details of any edit share permissions for the dashboard.
	View string `json:"view,omitempty"` // The URL of the dashboard.
	Automaticrefreshms int `json:"automaticRefreshMs,omitempty"` // The automatic refresh interval for the dashboard in milliseconds.
	Id string `json:"id,omitempty"` // The ID of the dashboard.
	Isfavourite bool `json:"isFavourite,omitempty"` // Whether the dashboard is selected as a favorite by the user.
}

// VersionIssuesStatus represents the VersionIssuesStatus schema from the OpenAPI specification
type VersionIssuesStatus struct {
	Done int64 `json:"done,omitempty"` // Count of issues with status *done*.
	Inprogress int64 `json:"inProgress,omitempty"` // Count of issues with status *in progress*.
	Todo int64 `json:"toDo,omitempty"` // Count of issues with status *to do*.
	Unmapped int64 `json:"unmapped,omitempty"` // Count of issues with a status other than *to do*, *in progress*, and *done*.
}

// IssueTypeScreenSchemeProjectAssociation represents the IssueTypeScreenSchemeProjectAssociation schema from the OpenAPI specification
type IssueTypeScreenSchemeProjectAssociation struct {
	Issuetypescreenschemeid string `json:"issueTypeScreenSchemeId,omitempty"` // The ID of the issue type screen scheme.
	Projectid string `json:"projectId,omitempty"` // The ID of the project.
}

// JexpIssues represents the JexpIssues schema from the OpenAPI specification
type JexpIssues struct {
	Jql interface{} `json:"jql,omitempty"` // The JQL query that specifies the set of issues available in the Jira expression.
}

// CreateWorkflowTransitionDetails represents the CreateWorkflowTransitionDetails schema from the OpenAPI specification
type CreateWorkflowTransitionDetails struct {
	Screen interface{} `json:"screen,omitempty"` // The screen of the transition.
	To string `json:"to"` // The status the transition goes to.
	TypeField string `json:"type"` // The type of the transition.
	Description string `json:"description,omitempty"` // The description of the transition. The maximum length is 1000 characters.
	From []string `json:"from,omitempty"` // The statuses the transition can start from.
	Name string `json:"name"` // The name of the transition. The maximum length is 60 characters.
	Properties map[string]interface{} `json:"properties,omitempty"` // The properties of the transition.
	Rules interface{} `json:"rules,omitempty"` // The rules of the transition.
}

// EventNotification represents the EventNotification schema from the OpenAPI specification
type EventNotification struct {
	Recipient string `json:"recipient,omitempty"` // The identifier associated with the `notificationType` value that defines the receiver of the notification, where the receiver isn't implied by the `notificationType` value. So, when `notificationType` is: * `User`, `recipient` is the user account ID. * `Group`, `recipient` is the group ID. * `ProjectRole`, `recipient` is the project role ID. * `UserCustomField`, `recipient` is the ID of the custom field. * `GroupCustomField`, `recipient` is the ID of the custom field.
	Emailaddress string `json:"emailAddress,omitempty"` // The email address.
	Projectrole interface{} `json:"projectRole,omitempty"` // The specified project role.
	Expand string `json:"expand,omitempty"` // Expand options that include additional event notification details in the response.
	Group interface{} `json:"group,omitempty"` // The specified group.
	Id int64 `json:"id,omitempty"` // The ID of the notification.
	Parameter string `json:"parameter,omitempty"` // As a group's name can change, use of `recipient` is recommended. The identifier associated with the `notificationType` value that defines the receiver of the notification, where the receiver isn't implied by `notificationType` value. So, when `notificationType` is: * `User` The `parameter` is the user account ID. * `Group` The `parameter` is the group name. * `ProjectRole` The `parameter` is the project role ID. * `UserCustomField` The `parameter` is the ID of the custom field. * `GroupCustomField` The `parameter` is the ID of the custom field.
	User interface{} `json:"user,omitempty"` // The specified user.
	Field interface{} `json:"field,omitempty"` // The custom user or group field.
	Notificationtype string `json:"notificationType,omitempty"` // Identifies the recipients of the notification.
}

// ProjectLandingPageInfo represents the ProjectLandingPageInfo schema from the OpenAPI specification
type ProjectLandingPageInfo struct {
	Boardname string `json:"boardName,omitempty"`
	Projectkey string `json:"projectKey,omitempty"`
	Queueid int64 `json:"queueId,omitempty"`
	Queuename string `json:"queueName,omitempty"`
	Simplified bool `json:"simplified,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Projecttype string `json:"projectType,omitempty"`
	Simpleboard bool `json:"simpleBoard,omitempty"`
	Boardid int64 `json:"boardId,omitempty"`
	Queuecategory string `json:"queueCategory,omitempty"`
	Url string `json:"url,omitempty"`
}

// UserPickerUser represents the UserPickerUser schema from the OpenAPI specification
type UserPickerUser struct {
	Name string `json:"name,omitempty"` // This property is no longer available . See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Accountid string `json:"accountId,omitempty"` // The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*.
	Avatarurl string `json:"avatarUrl,omitempty"` // The avatar URL of the user.
	Displayname string `json:"displayName,omitempty"` // The display name of the user. Depending on the user’s privacy setting, this may be returned as null.
	Html string `json:"html,omitempty"` // The display name, email address, and key of the user with the matched query string highlighted with the HTML bold tag.
	Key string `json:"key,omitempty"` // This property is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
}

// ProjectIdentifiers represents the ProjectIdentifiers schema from the OpenAPI specification
type ProjectIdentifiers struct {
	Id int64 `json:"id"` // The ID of the created project.
	Key string `json:"key"` // The key of the created project.
	Self string `json:"self"` // The URL of the created project.
}

// Hierarchy represents the Hierarchy schema from the OpenAPI specification
type Hierarchy struct {
	Baselevelid int64 `json:"baseLevelId,omitempty"` // The ID of the base level. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	Levels []SimplifiedHierarchyLevel `json:"levels,omitempty"` // Details about the hierarchy level.
}

// JexpJqlIssues represents the JexpJqlIssues schema from the OpenAPI specification
type JexpJqlIssues struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of issues to return from the JQL query. Inspect `meta.issues.jql.maxResults` in the response to ensure the maximum value has not been exceeded.
	Query string `json:"query,omitempty"` // The JQL query.
	Startat int64 `json:"startAt,omitempty"` // The index of the first issue to return from the JQL query.
	Validation string `json:"validation,omitempty"` // Determines how to validate the JQL query and treat the validation results.
}

// SystemAvatars represents the SystemAvatars schema from the OpenAPI specification
type SystemAvatars struct {
	System []Avatar `json:"system,omitempty"` // A list of avatar details.
}

// PaginatedResponseComment represents the PaginatedResponseComment schema from the OpenAPI specification
type PaginatedResponseComment struct {
	Startat int64 `json:"startAt,omitempty"`
	Total int64 `json:"total,omitempty"`
	Maxresults int `json:"maxResults,omitempty"`
	Results []Comment `json:"results,omitempty"`
}

// CustomFieldContextDefaultValueMultipleOption represents the CustomFieldContextDefaultValueMultipleOption schema from the OpenAPI specification
type CustomFieldContextDefaultValueMultipleOption struct {
	TypeField string `json:"type"`
	Contextid string `json:"contextId"` // The ID of the context.
	Optionids []string `json:"optionIds"` // The list of IDs of the default options.
}

// StatusCreate represents the StatusCreate schema from the OpenAPI specification
type StatusCreate struct {
	Description string `json:"description,omitempty"` // The description of the status.
	Name string `json:"name"` // The name of the status.
	Statuscategory string `json:"statusCategory"` // The category of the status.
}

// IssuesJqlMetaDataBean represents the IssuesJqlMetaDataBean schema from the OpenAPI specification
type IssuesJqlMetaDataBean struct {
	Count int `json:"count"` // The number of issues that were loaded in this evaluation.
	Maxresults int `json:"maxResults"` // The maximum number of issues that could be loaded in this evaluation.
	Startat int64 `json:"startAt"` // The index of the first issue.
	Totalcount int64 `json:"totalCount"` // The total number of issues the JQL returned.
	Validationwarnings []string `json:"validationWarnings,omitempty"` // Any warnings related to the JQL query. Present only if the validation mode was set to `warn`.
}

// IssueFieldOptionCreateBean represents the IssueFieldOptionCreateBean schema from the OpenAPI specification
type IssueFieldOptionCreateBean struct {
	Value string `json:"value"` // The option's name, which is displayed in Jira.
	Config IssueFieldOptionConfiguration `json:"config,omitempty"` // Details of the projects the option is available in.
	Properties map[string]interface{} `json:"properties,omitempty"` // The properties of the option as arbitrary key-value pairs. These properties can be searched using JQL, if the extractions (see https://developer.atlassian.com/cloud/jira/platform/modules/issue-field-option-property-index/) are defined in the descriptor for the issue field module.
}

// UserContextVariable represents the UserContextVariable schema from the OpenAPI specification
type UserContextVariable struct {
	TypeField string `json:"type"` // Type of custom context variable.
	Accountid string `json:"accountId"` // The account ID of the user.
}

// Avatar represents the Avatar schema from the OpenAPI specification
type Avatar struct {
	Filename string `json:"fileName,omitempty"` // The file name of the avatar icon. Returned for system avatars.
	Id string `json:"id"` // The ID of the avatar.
	Isdeletable bool `json:"isDeletable,omitempty"` // Whether the avatar can be deleted.
	Isselected bool `json:"isSelected,omitempty"` // Whether the avatar is used in Jira. For example, shown as a project's avatar.
	Issystemavatar bool `json:"isSystemAvatar,omitempty"` // Whether the avatar is a system avatar.
	Owner string `json:"owner,omitempty"` // The owner of the avatar. For a system avatar the owner is null (and nothing is returned). For non-system avatars this is the appropriate identifier, such as the ID for a project or the account ID for a user.
	Urls map[string]interface{} `json:"urls,omitempty"` // The list of avatar icon URLs.
}

// IssueTypeScreenSchemeId represents the IssueTypeScreenSchemeId schema from the OpenAPI specification
type IssueTypeScreenSchemeId struct {
	Id string `json:"id"` // The ID of the issue type screen scheme.
}

// UpdatedProjectCategory represents the UpdatedProjectCategory schema from the OpenAPI specification
type UpdatedProjectCategory struct {
	Description string `json:"description,omitempty"` // The name of the project category.
	Id string `json:"id,omitempty"` // The ID of the project category.
	Name string `json:"name,omitempty"` // The description of the project category.
	Self string `json:"self,omitempty"` // The URL of the project category.
}

// Comment represents the Comment schema from the OpenAPI specification
type Comment struct {
	Body interface{} `json:"body,omitempty"` // The comment text in [Atlassian Document Format](https://developer.atlassian.com/cloud/jira/platform/apis/document/structure/).
	Jsdpublic bool `json:"jsdPublic,omitempty"` // Whether the comment is visible in Jira Service Desk. Defaults to true when comments are created in the Jira Cloud Platform. This includes when the site doesn't use Jira Service Desk or the project isn't a Jira Service Desk project and, therefore, there is no Jira Service Desk for the issue to be visible on. To create a comment with its visibility in Jira Service Desk set to false, use the Jira Service Desk REST API [Create request comment](https://developer.atlassian.com/cloud/jira/service-desk/rest/#api-rest-servicedeskapi-request-issueIdOrKey-comment-post) operation.
	Updated string `json:"updated,omitempty"` // The date and time at which the comment was updated last.
	Properties []EntityProperty `json:"properties,omitempty"` // A list of comment properties. Optional on create and update.
	Self string `json:"self,omitempty"` // The URL of the comment.
	Created string `json:"created,omitempty"` // The date and time at which the comment was created.
	Id string `json:"id,omitempty"` // The ID of the comment.
	Updateauthor interface{} `json:"updateAuthor,omitempty"` // The ID of the user who updated the comment last.
	Author interface{} `json:"author,omitempty"` // The ID of the user who created the comment.
	Jsdauthorcanseerequest bool `json:"jsdAuthorCanSeeRequest,omitempty"` // Whether the comment was added from an email sent by a person who is not part of the issue. See [Allow external emails to be added as comments on issues](https://support.atlassian.com/jira-service-management-cloud/docs/allow-external-emails-to-be-added-as-comments-on-issues/)for information on setting up this feature.
	Renderedbody string `json:"renderedBody,omitempty"` // The rendered version of the comment.
	Visibility interface{} `json:"visibility,omitempty"` // The group or role to which this comment is visible. Optional on create and update.
}

// CustomFieldContextDefaultValueLabels represents the CustomFieldContextDefaultValueLabels schema from the OpenAPI specification
type CustomFieldContextDefaultValueLabels struct {
	Labels []string `json:"labels"` // The default labels value.
	TypeField string `json:"type"`
}

// PageBeanCustomFieldContextDefaultValue represents the PageBeanCustomFieldContextDefaultValue schema from the OpenAPI specification
type PageBeanCustomFieldContextDefaultValue struct {
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []CustomFieldContextDefaultValue `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
}

// IssueTypeScreenSchemeDetails represents the IssueTypeScreenSchemeDetails schema from the OpenAPI specification
type IssueTypeScreenSchemeDetails struct {
	Description string `json:"description,omitempty"` // The description of the issue type screen scheme. The maximum length is 255 characters.
	Issuetypemappings []IssueTypeScreenSchemeMapping `json:"issueTypeMappings"` // The IDs of the screen schemes for the issue type IDs and *default*. A *default* entry is required to create an issue type screen scheme, it defines the mapping for all issue types without a screen scheme.
	Name string `json:"name"` // The name of the issue type screen scheme. The name must be unique. The maximum length is 255 characters.
}

// PageBeanUser represents the PageBeanUser schema from the OpenAPI specification
type PageBeanUser struct {
	Values []User `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// PageBeanIssueFieldOption represents the PageBeanIssueFieldOption schema from the OpenAPI specification
type PageBeanIssueFieldOption struct {
	Values []IssueFieldOption `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// Fields represents the Fields schema from the OpenAPI specification
type Fields struct {
	Issuetype interface{} `json:"issueType,omitempty"` // The type of the linked issue.
	Issuetype IssueTypeDetails `json:"issuetype,omitempty"` // Details about an issue type.
	Priority interface{} `json:"priority,omitempty"` // The priority of the linked issue.
	Status interface{} `json:"status,omitempty"` // The status of the linked issue.
	Summary string `json:"summary,omitempty"` // The summary description of the linked issue.
	Timetracking interface{} `json:"timetracking,omitempty"` // The time tracking of the linked issue.
	Assignee interface{} `json:"assignee,omitempty"` // The assignee of the linked issue.
}

// FilterDetails represents the FilterDetails schema from the OpenAPI specification
type FilterDetails struct {
	Sharepermissions []SharePermission `json:"sharePermissions,omitempty"` // The groups and projects that the filter is shared with. This can be specified when updating a filter, but not when creating a filter.
	Subscriptions []FilterSubscription `json:"subscriptions,omitempty"` // The users that are subscribed to the filter.
	Description string `json:"description,omitempty"` // The description of the filter.
	Favourite bool `json:"favourite,omitempty"` // Whether the filter is selected as a favorite by any users, not including the filter owner.
	Favouritedcount int64 `json:"favouritedCount,omitempty"` // The count of how many users have selected this filter as a favorite, including the filter owner.
	Self string `json:"self,omitempty"` // The URL of the filter.
	Searchurl string `json:"searchUrl,omitempty"` // A URL to view the filter results in Jira, using the [Search for issues using JQL](#api-rest-api-3-filter-search-get) operation with the filter's JQL string to return the filter results. For example, *https://your-domain.atlassian.net/rest/api/3/search?jql=project+%3D+SSP+AND+issuetype+%3D+Bug*.
	Viewurl string `json:"viewUrl,omitempty"` // A URL to view the filter results in Jira, using the ID of the filter. For example, *https://your-domain.atlassian.net/issues/?filter=10100*.
	Expand string `json:"expand,omitempty"` // Expand options that include additional filter details in the response.
	Jql string `json:"jql,omitempty"` // The JQL query for the filter. For example, *project = SSP AND issuetype = Bug*.
	Name string `json:"name"` // The name of the filter.
	Owner interface{} `json:"owner,omitempty"` // The user who owns the filter. Defaults to the creator of the filter, however, Jira administrators can change the owner of a shared filter in the admin settings.
	Editpermissions []SharePermission `json:"editPermissions,omitempty"` // The groups and projects that can edit the filter. This can be specified when updating a filter, but not when creating a filter.
	Id string `json:"id,omitempty"` // The unique identifier for the filter.
}

// CreateUpdateRoleRequestBean represents the CreateUpdateRoleRequestBean schema from the OpenAPI specification
type CreateUpdateRoleRequestBean struct {
	Name string `json:"name,omitempty"` // The name of the project role. Must be unique. Cannot begin or end with whitespace. The maximum length is 255 characters. Required when creating a project role. Optional when partially updating a project role.
	Description string `json:"description,omitempty"` // A description of the project role. Required when fully updating a project role. Optional when creating or partially updating a project role.
}

// CustomFieldContextDefaultValueForgeObjectField represents the CustomFieldContextDefaultValueForgeObjectField schema from the OpenAPI specification
type CustomFieldContextDefaultValueForgeObjectField struct {
	TypeField string `json:"type"`
	Object map[string]interface{} `json:"object,omitempty"` // The default JSON object.
}

// CustomFieldContext represents the CustomFieldContext schema from the OpenAPI specification
type CustomFieldContext struct {
	Id string `json:"id"` // The ID of the context.
	Isanyissuetype bool `json:"isAnyIssueType"` // Whether the context apply to all issue types.
	Isglobalcontext bool `json:"isGlobalContext"` // Whether the context is global.
	Name string `json:"name"` // The name of the context.
	Description string `json:"description"` // The description of the context.
}

// IdOrKeyBean represents the IdOrKeyBean schema from the OpenAPI specification
type IdOrKeyBean struct {
	Key string `json:"key,omitempty"` // The key of the referenced item.
	Id int64 `json:"id,omitempty"` // The ID of the referenced item.
}

// SimpleErrorCollection represents the SimpleErrorCollection schema from the OpenAPI specification
type SimpleErrorCollection struct {
	Httpstatuscode int `json:"httpStatusCode,omitempty"`
	Errormessages []string `json:"errorMessages,omitempty"` // The list of error messages produced by this operation. For example, "input parameter 'key' must be provided"
	Errors map[string]interface{} `json:"errors,omitempty"` // The list of errors by parameter returned by the operation. For example,"projectKey": "Project keys must start with an uppercase letter, followed by one or more uppercase alphanumeric characters."
}

// PageBeanIssueTypeScreenSchemeItem represents the PageBeanIssueTypeScreenSchemeItem schema from the OpenAPI specification
type PageBeanIssueTypeScreenSchemeItem struct {
	Values []IssueTypeScreenSchemeItem `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// PageOfComments represents the PageOfComments schema from the OpenAPI specification
type PageOfComments struct {
	Comments []Comment `json:"comments,omitempty"` // The list of comments.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// JsonTypeBean represents the JsonTypeBean schema from the OpenAPI specification
type JsonTypeBean struct {
	Items string `json:"items,omitempty"` // When the data type is an array, the name of the field items within the array.
	System string `json:"system,omitempty"` // If the field is a system field, the name of the field.
	TypeField string `json:"type"` // The data type of the field.
	Configuration map[string]interface{} `json:"configuration,omitempty"` // If the field is a custom field, the configuration of the field.
	Custom string `json:"custom,omitempty"` // If the field is a custom field, the URI of the field.
	Customid int64 `json:"customId,omitempty"` // If the field is a custom field, the custom ID of the field.
}

// ContextualConfiguration represents the ContextualConfiguration schema from the OpenAPI specification
type ContextualConfiguration struct {
	Configuration interface{} `json:"configuration,omitempty"` // The field configuration.
	Fieldcontextid string `json:"fieldContextId"` // The ID of the field context the configuration is associated with.
	Id string `json:"id"` // The ID of the configuration.
	Schema interface{} `json:"schema,omitempty"` // The field value schema.
}

// CustomFieldContextDefaultValueDateTime represents the CustomFieldContextDefaultValueDateTime schema from the OpenAPI specification
type CustomFieldContextDefaultValueDateTime struct {
	Datetime string `json:"dateTime,omitempty"` // The default date-time in ISO format. Ignored if `useCurrent` is true.
	TypeField string `json:"type"`
	Usecurrent bool `json:"useCurrent,omitempty"` // Whether to use the current date.
}

// PropertyKey represents the PropertyKey schema from the OpenAPI specification
type PropertyKey struct {
	Key string `json:"key,omitempty"` // The key of the property.
	Self string `json:"self,omitempty"` // The URL of the property.
}

// StatusScope represents the StatusScope schema from the OpenAPI specification
type StatusScope struct {
	TypeField string `json:"type"` // The scope of the status. `GLOBAL` for company-managed projects and `PROJECT` for team-managed projects.
	Project ProjectId `json:"project,omitempty"` // Project ID details.
}

// WorkflowScheme represents the WorkflowScheme schema from the OpenAPI specification
type WorkflowScheme struct {
	Draft bool `json:"draft,omitempty"` // Whether the workflow scheme is a draft or not.
	Lastmodifieduser interface{} `json:"lastModifiedUser,omitempty"` // The user that last modified the draft workflow scheme. A modification is a change to the issue type-project mappings only. This property does not apply to non-draft workflows.
	Updatedraftifneeded bool `json:"updateDraftIfNeeded,omitempty"` // Whether to create or update a draft workflow scheme when updating an active workflow scheme. An active workflow scheme is a workflow scheme that is used by at least one project. The following examples show how this property works: * Update an active workflow scheme with `updateDraftIfNeeded` set to `true`: If a draft workflow scheme exists, it is updated. Otherwise, a draft workflow scheme is created. * Update an active workflow scheme with `updateDraftIfNeeded` set to `false`: An error is returned, as active workflow schemes cannot be updated. * Update an inactive workflow scheme with `updateDraftIfNeeded` set to `true`: The workflow scheme is updated, as inactive workflow schemes do not require drafts to update. Defaults to `false`.
	Issuetypemappings map[string]interface{} `json:"issueTypeMappings,omitempty"` // The issue type to workflow mappings, where each mapping is an issue type ID and workflow name pair. Note that an issue type can only be mapped to one workflow in a workflow scheme.
	Issuetypes map[string]interface{} `json:"issueTypes,omitempty"` // The issue types available in Jira.
	Name string `json:"name,omitempty"` // The name of the workflow scheme. The name must be unique. The maximum length is 255 characters. Required when creating a workflow scheme.
	Originaldefaultworkflow string `json:"originalDefaultWorkflow,omitempty"` // For draft workflow schemes, this property is the name of the default workflow for the original workflow scheme. The default workflow has *All Unassigned Issue Types* assigned to it in Jira.
	Self string `json:"self,omitempty"`
	Defaultworkflow string `json:"defaultWorkflow,omitempty"` // The name of the default workflow for the workflow scheme. The default workflow has *All Unassigned Issue Types* assigned to it in Jira. If `defaultWorkflow` is not specified when creating a workflow scheme, it is set to *Jira Workflow (jira)*.
	Description string `json:"description,omitempty"` // The description of the workflow scheme.
	Originalissuetypemappings map[string]interface{} `json:"originalIssueTypeMappings,omitempty"` // For draft workflow schemes, this property is the issue type to workflow mappings for the original workflow scheme, where each mapping is an issue type ID and workflow name pair. Note that an issue type can only be mapped to one workflow in a workflow scheme.
	Id int64 `json:"id,omitempty"` // The ID of the workflow scheme.
	Lastmodified string `json:"lastModified,omitempty"` // The date-time that the draft workflow scheme was last modified. A modification is a change to the issue type-project mappings only. This property does not apply to non-draft workflows.
}

// WorkflowCompoundCondition represents the WorkflowCompoundCondition schema from the OpenAPI specification
type WorkflowCompoundCondition struct {
	Operator string `json:"operator"` // The compound condition operator.
	Conditions []WorkflowCondition `json:"conditions"` // The list of workflow conditions.
	Nodetype string `json:"nodeType"`
}

// StatusCreateRequest represents the StatusCreateRequest schema from the OpenAPI specification
type StatusCreateRequest struct {
	Scope StatusScope `json:"scope"` // The scope of the status.
	Statuses []StatusCreate `json:"statuses"` // Details of the statuses being created.
}

// WorklogIdsRequestBean represents the WorklogIdsRequestBean schema from the OpenAPI specification
type WorklogIdsRequestBean struct {
	Ids []int64 `json:"ids"` // A list of worklog IDs.
}

// PageBeanCustomFieldContextProjectMapping represents the PageBeanCustomFieldContextProjectMapping schema from the OpenAPI specification
type PageBeanCustomFieldContextProjectMapping struct {
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []CustomFieldContextProjectMapping `json:"values,omitempty"` // The list of items.
}

// Resolution represents the Resolution schema from the OpenAPI specification
type Resolution struct {
	Id string `json:"id,omitempty"` // The ID of the issue resolution.
	Name string `json:"name,omitempty"` // The name of the issue resolution.
	Self string `json:"self,omitempty"` // The URL of the issue resolution.
	Description string `json:"description,omitempty"` // The description of the issue resolution.
}

// FoundUsersAndGroups represents the FoundUsersAndGroups schema from the OpenAPI specification
type FoundUsersAndGroups struct {
	Users FoundUsers `json:"users,omitempty"` // The list of users found in a search, including header text (Showing X of Y matching users) and total of matched users.
	Groups FoundGroups `json:"groups,omitempty"` // The list of groups found in a search, including header text (Showing X of Y matching groups) and total of matched groups.
}

// UpdateUserToGroupBean represents the UpdateUserToGroupBean schema from the OpenAPI specification
type UpdateUserToGroupBean struct {
	Accountid string `json:"accountId,omitempty"` // The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*.
	Name string `json:"name,omitempty"` // This property is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
}

// Filter represents the Filter schema from the OpenAPI specification
type Filter struct {
	Subscriptions interface{} `json:"subscriptions,omitempty"` // A paginated list of the users that are subscribed to the filter.
	Description string `json:"description,omitempty"` // A description of the filter.
	Favourite bool `json:"favourite,omitempty"` // Whether the filter is selected as a favorite.
	Self string `json:"self,omitempty"` // The URL of the filter.
	Searchurl string `json:"searchUrl,omitempty"` // A URL to view the filter results in Jira, using the [Search for issues using JQL](#api-rest-api-3-filter-search-get) operation with the filter's JQL string to return the filter results. For example, *https://your-domain.atlassian.net/rest/api/3/search?jql=project+%3D+SSP+AND+issuetype+%3D+Bug*.
	Sharepermissions []SharePermission `json:"sharePermissions,omitempty"` // The groups and projects that the filter is shared with.
	Viewurl string `json:"viewUrl,omitempty"` // A URL to view the filter results in Jira, using the ID of the filter. For example, *https://your-domain.atlassian.net/issues/?filter=10100*.
	Editpermissions []SharePermission `json:"editPermissions,omitempty"` // The groups and projects that can edit the filter.
	Id string `json:"id,omitempty"` // The unique identifier for the filter.
	Jql string `json:"jql,omitempty"` // The JQL query for the filter. For example, *project = SSP AND issuetype = Bug*.
	Owner interface{} `json:"owner,omitempty"` // The user who owns the filter. This is defaulted to the creator of the filter, however Jira administrators can change the owner of a shared filter in the admin settings.
	Sharedusers interface{} `json:"sharedUsers,omitempty"` // A paginated list of the users that the filter is shared with. This includes users that are members of the groups or can browse the projects that the filter is shared with.
	Favouritedcount int64 `json:"favouritedCount,omitempty"` // The count of how many users have selected this filter as a favorite, including the filter owner.
	Name string `json:"name"` // The name of the filter. Must be unique.
}

// Operations represents the Operations schema from the OpenAPI specification
type Operations struct {
	Linkgroups []LinkGroup `json:"linkGroups,omitempty"` // Details of the link groups defining issue operations.
}

// CustomFieldContextProjectMapping represents the CustomFieldContextProjectMapping schema from the OpenAPI specification
type CustomFieldContextProjectMapping struct {
	Contextid string `json:"contextId"` // The ID of the context.
	Isglobalcontext bool `json:"isGlobalContext,omitempty"` // Whether context is global.
	Projectid string `json:"projectId,omitempty"` // The ID of the project.
}

// AttachmentMetadata represents the AttachmentMetadata schema from the OpenAPI specification
type AttachmentMetadata struct {
	Created string `json:"created,omitempty"` // The datetime the attachment was created.
	Size int64 `json:"size,omitempty"` // The size of the attachment.
	Author interface{} `json:"author,omitempty"` // Details of the user who attached the file.
	Id int64 `json:"id,omitempty"` // The ID of the attachment.
	Properties map[string]interface{} `json:"properties,omitempty"` // Additional properties of the attachment.
	Filename string `json:"filename,omitempty"` // The name of the attachment file.
	Mimetype string `json:"mimeType,omitempty"` // The MIME type of the attachment.
	Thumbnail string `json:"thumbnail,omitempty"` // The URL of a thumbnail representing the attachment.
	Content string `json:"content,omitempty"` // The URL of the attachment.
	Self string `json:"self,omitempty"` // The URL of the attachment metadata details.
}

// JqlQueryField represents the JqlQueryField schema from the OpenAPI specification
type JqlQueryField struct {
	Encodedname string `json:"encodedName,omitempty"` // The encoded name of the field, which can be used directly in a JQL query.
	Name string `json:"name"` // The name of the field.
	Property []JqlQueryFieldEntityProperty `json:"property,omitempty"` // When the field refers to a value in an entity property, details of the entity property value.
}

// PageBeanIssueTypeSchemeMapping represents the PageBeanIssueTypeSchemeMapping schema from the OpenAPI specification
type PageBeanIssueTypeSchemeMapping struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []IssueTypeSchemeMapping `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// IssueTypeScreenScheme represents the IssueTypeScreenScheme schema from the OpenAPI specification
type IssueTypeScreenScheme struct {
	Description string `json:"description,omitempty"` // The description of the issue type screen scheme.
	Id string `json:"id"` // The ID of the issue type screen scheme.
	Name string `json:"name"` // The name of the issue type screen scheme.
}

// Priority represents the Priority schema from the OpenAPI specification
type Priority struct {
	Name string `json:"name,omitempty"` // The name of the issue priority.
	Self string `json:"self,omitempty"` // The URL of the issue priority.
	Statuscolor string `json:"statusColor,omitempty"` // The color used to indicate the issue priority.
	Description string `json:"description,omitempty"` // The description of the issue priority.
	Iconurl string `json:"iconUrl,omitempty"` // The URL of the icon for the issue priority.
	Id string `json:"id,omitempty"` // The ID of the issue priority.
	Isdefault bool `json:"isDefault,omitempty"` // Whether this priority is the default.
}

// JqlQuery represents the JqlQuery schema from the OpenAPI specification
type JqlQuery struct {
	Where JqlQueryClause `json:"where,omitempty"` // A JQL query clause.
	Orderby JqlQueryOrderByClause `json:"orderBy,omitempty"` // Details of the order-by JQL clause.
}

// ConvertedJQLQueries represents the ConvertedJQLQueries schema from the OpenAPI specification
type ConvertedJQLQueries struct {
	Querieswithunknownusers []JQLQueryWithUnknownUsers `json:"queriesWithUnknownUsers,omitempty"` // List of queries containing user information that could not be mapped to an existing user
	Querystrings []string `json:"queryStrings,omitempty"` // The list of converted query strings with account IDs in place of user identifiers.
}

// CustomFieldValueUpdateDetails represents the CustomFieldValueUpdateDetails schema from the OpenAPI specification
type CustomFieldValueUpdateDetails struct {
	Updates []CustomFieldValueUpdate `json:"updates,omitempty"` // The list of custom field update details.
}

// WorkflowTransitionRules represents the WorkflowTransitionRules schema from the OpenAPI specification
type WorkflowTransitionRules struct {
	Conditions []AppWorkflowTransitionRule `json:"conditions,omitempty"` // The list of conditions within the workflow.
	Postfunctions []AppWorkflowTransitionRule `json:"postFunctions,omitempty"` // The list of post functions within the workflow.
	Validators []AppWorkflowTransitionRule `json:"validators,omitempty"` // The list of validators within the workflow.
	Workflowid WorkflowId `json:"workflowId"` // Properties that identify a workflow.
}

// VersionUnresolvedIssuesCount represents the VersionUnresolvedIssuesCount schema from the OpenAPI specification
type VersionUnresolvedIssuesCount struct {
	Self string `json:"self,omitempty"` // The URL of these count details.
	Issuescount int64 `json:"issuesCount,omitempty"` // Count of issues.
	Issuesunresolvedcount int64 `json:"issuesUnresolvedCount,omitempty"` // Count of unresolved issues.
}

// VersionUsageInCustomField represents the VersionUsageInCustomField schema from the OpenAPI specification
type VersionUsageInCustomField struct {
	Customfieldid int64 `json:"customFieldId,omitempty"` // The ID of the custom field.
	Fieldname string `json:"fieldName,omitempty"` // The name of the custom field.
	Issuecountwithversionincustomfield int64 `json:"issueCountWithVersionInCustomField,omitempty"` // Count of the issues where the custom field contains the version.
}

// PageBeanWorkflowTransitionRules represents the PageBeanWorkflowTransitionRules schema from the OpenAPI specification
type PageBeanWorkflowTransitionRules struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []WorkflowTransitionRules `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// ConnectCustomFieldValues represents the ConnectCustomFieldValues schema from the OpenAPI specification
type ConnectCustomFieldValues struct {
	Updatevaluelist []ConnectCustomFieldValue `json:"updateValueList,omitempty"` // The list of custom field update details.
}

// PageBeanIssueTypeScreenScheme represents the PageBeanIssueTypeScreenScheme schema from the OpenAPI specification
type PageBeanIssueTypeScreenScheme struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []IssueTypeScreenScheme `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// WorkflowTransition represents the WorkflowTransition schema from the OpenAPI specification
type WorkflowTransition struct {
	Name string `json:"name"` // The transition name.
	Id int `json:"id"` // The transition ID.
}

// PageBeanResolutionJsonBean represents the PageBeanResolutionJsonBean schema from the OpenAPI specification
type PageBeanResolutionJsonBean struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []ResolutionJsonBean `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// ChangedValueBean represents the ChangedValueBean schema from the OpenAPI specification
type ChangedValueBean struct {
	Changedfrom string `json:"changedFrom,omitempty"` // The value of the field before the change.
	Changedto string `json:"changedTo,omitempty"` // The value of the field after the change.
	Fieldname string `json:"fieldName,omitempty"` // The name of the field changed.
}

// ProjectComponent represents the ProjectComponent schema from the OpenAPI specification
type ProjectComponent struct {
	Assignee interface{} `json:"assignee,omitempty"` // The details of the user associated with `assigneeType`, if any. See `realAssignee` for details of the user assigned to issues created with this component.
	Leadusername string `json:"leadUserName,omitempty"` // This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Self string `json:"self,omitempty"` // The URL of the component.
	Id string `json:"id,omitempty"` // The unique identifier for the component.
	Realassignee interface{} `json:"realAssignee,omitempty"` // The user assigned to issues created with this component, when `assigneeType` does not identify a valid assignee.
	Description string `json:"description,omitempty"` // The description for the component. Optional when creating or updating a component.
	Isassigneetypevalid bool `json:"isAssigneeTypeValid,omitempty"` // Whether a user is associated with `assigneeType`. For example, if the `assigneeType` is set to `COMPONENT_LEAD` but the component lead is not set, then `false` is returned.
	Name string `json:"name,omitempty"` // The unique name for the component in the project. Required when creating a component. Optional when updating a component. The maximum length is 255 characters.
	Project string `json:"project,omitempty"` // The key of the project the component is assigned to. Required when creating a component. Can't be updated.
	Realassigneetype string `json:"realAssigneeType,omitempty"` // The type of the assignee that is assigned to issues created with this component, when an assignee cannot be set from the `assigneeType`. For example, `assigneeType` is set to `COMPONENT_LEAD` but no component lead is set. This property is set to one of the following values: * `PROJECT_LEAD` when `assigneeType` is `PROJECT_LEAD` and the project lead has permission to be assigned issues in the project that the component is in. * `COMPONENT_LEAD` when `assignee`Type is `COMPONENT_LEAD` and the component lead has permission to be assigned issues in the project that the component is in. * `UNASSIGNED` when `assigneeType` is `UNASSIGNED` and Jira is configured to allow unassigned issues. * `PROJECT_DEFAULT` when none of the preceding cases are true.
	Assigneetype string `json:"assigneeType,omitempty"` // The nominal user type used to determine the assignee for issues created with this component. See `realAssigneeType` for details on how the type of the user, and hence the user, assigned to issues is determined. Can take the following values: * `PROJECT_LEAD` the assignee to any issues created with this component is nominally the lead for the project the component is in. * `COMPONENT_LEAD` the assignee to any issues created with this component is nominally the lead for the component. * `UNASSIGNED` an assignee is not set for issues created with this component. * `PROJECT_DEFAULT` the assignee to any issues created with this component is nominally the default assignee for the project that the component is in. Default value: `PROJECT_DEFAULT`. Optional when creating or updating a component.
	Lead interface{} `json:"lead,omitempty"` // The user details for the component's lead user.
	Leadaccountid string `json:"leadAccountId,omitempty"` // The accountId of the component's lead user. The accountId uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*.
	Projectid int64 `json:"projectId,omitempty"` // The ID of the project the component is assigned to.
}

// AttachmentArchiveItemReadable represents the AttachmentArchiveItemReadable schema from the OpenAPI specification
type AttachmentArchiveItemReadable struct {
	Size string `json:"size,omitempty"` // The size of the archive item.
	Index int64 `json:"index,omitempty"` // The position of the item within the archive.
	Label string `json:"label,omitempty"` // The label for the archive item.
	Mediatype string `json:"mediaType,omitempty"` // The MIME type of the archive item.
	Path string `json:"path,omitempty"` // The path of the archive item.
}

// FieldValueClause represents the FieldValueClause schema from the OpenAPI specification
type FieldValueClause struct {
	Field JqlQueryField `json:"field"` // A field used in a JQL query. See [Advanced searching - fields reference](https://confluence.atlassian.com/x/dAiiLQ) for more information about fields in JQL queries.
	Operand JqlQueryClauseOperand `json:"operand"` // Details of an operand in a JQL clause.
	Operator string `json:"operator"` // The operator between the field and operand.
}

// ProjectPermissions represents the ProjectPermissions schema from the OpenAPI specification
type ProjectPermissions struct {
	Canedit bool `json:"canEdit,omitempty"` // Whether the logged user can edit the project.
}

// RemoteIssueLink represents the RemoteIssueLink schema from the OpenAPI specification
type RemoteIssueLink struct {
	Relationship string `json:"relationship,omitempty"` // Description of the relationship between the issue and the linked item.
	Self string `json:"self,omitempty"` // The URL of the link.
	Application interface{} `json:"application,omitempty"` // Details of the remote application the linked item is in.
	Globalid string `json:"globalId,omitempty"` // The global ID of the link, such as the ID of the item on the remote system.
	Id int64 `json:"id,omitempty"` // The ID of the link.
	Object interface{} `json:"object,omitempty"` // Details of the item linked to.
}

// DefaultLevelValue represents the DefaultLevelValue schema from the OpenAPI specification
type DefaultLevelValue struct {
	Defaultlevelid string `json:"defaultLevelId"` // The ID of the issue security level to set as default for the specified scheme. Providing null will reset the default level.
	Issuesecurityschemeid string `json:"issueSecuritySchemeId"` // The ID of the issue security scheme to set default level for.
}

// CustomFieldContextDefaultValueForgeUserField represents the CustomFieldContextDefaultValueForgeUserField schema from the OpenAPI specification
type CustomFieldContextDefaultValueForgeUserField struct {
	TypeField string `json:"type"`
	Userfilter UserFilter `json:"userFilter"` // Filter for a User Picker (single) custom field.
	Accountid string `json:"accountId"` // The ID of the default user.
	Contextid string `json:"contextId"` // The ID of the context.
}

// SecuritySchemes represents the SecuritySchemes schema from the OpenAPI specification
type SecuritySchemes struct {
	Issuesecurityschemes []SecurityScheme `json:"issueSecuritySchemes,omitempty"` // List of security schemes.
}

// PageBeanChangelog represents the PageBeanChangelog schema from the OpenAPI specification
type PageBeanChangelog struct {
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []Changelog `json:"values,omitempty"` // The list of items.
}

// CustomFieldContextDefaultValueTextField represents the CustomFieldContextDefaultValueTextField schema from the OpenAPI specification
type CustomFieldContextDefaultValueTextField struct {
	Text string `json:"text,omitempty"` // The default text. The maximum length is 254 characters.
	TypeField string `json:"type"`
}

// BulkProjectPermissions represents the BulkProjectPermissions schema from the OpenAPI specification
type BulkProjectPermissions struct {
	Issues []int64 `json:"issues,omitempty"` // List of issue IDs.
	Permissions []string `json:"permissions"` // List of project permissions.
	Projects []int64 `json:"projects,omitempty"` // List of project IDs.
}

// OperationMessage represents the OperationMessage schema from the OpenAPI specification
type OperationMessage struct {
	Message string `json:"message"` // The human-readable message that describes the result.
	Statuscode int `json:"statusCode"` // The status code of the response.
}

// ComponentWithIssueCount represents the ComponentWithIssueCount schema from the OpenAPI specification
type ComponentWithIssueCount struct {
	Id string `json:"id,omitempty"` // The unique identifier for the component.
	Issuecount int64 `json:"issueCount,omitempty"` // Count of issues for the component.
	Project string `json:"project,omitempty"` // The key of the project to which the component is assigned.
	Realassigneetype string `json:"realAssigneeType,omitempty"` // The type of the assignee that is assigned to issues created with this component, when an assignee cannot be set from the `assigneeType`. For example, `assigneeType` is set to `COMPONENT_LEAD` but no component lead is set. This property is set to one of the following values: * `PROJECT_LEAD` when `assigneeType` is `PROJECT_LEAD` and the project lead has permission to be assigned issues in the project that the component is in. * `COMPONENT_LEAD` when `assignee`Type is `COMPONENT_LEAD` and the component lead has permission to be assigned issues in the project that the component is in. * `UNASSIGNED` when `assigneeType` is `UNASSIGNED` and Jira is configured to allow unassigned issues. * `PROJECT_DEFAULT` when none of the preceding cases are true.
	Self string `json:"self,omitempty"` // The URL for this count of the issues contained in the component.
	Assignee interface{} `json:"assignee,omitempty"` // The details of the user associated with `assigneeType`, if any. See `realAssignee` for details of the user assigned to issues created with this component.
	Assigneetype string `json:"assigneeType,omitempty"` // The nominal user type used to determine the assignee for issues created with this component. See `realAssigneeType` for details on how the type of the user, and hence the user, assigned to issues is determined. Takes the following values: * `PROJECT_LEAD` the assignee to any issues created with this component is nominally the lead for the project the component is in. * `COMPONENT_LEAD` the assignee to any issues created with this component is nominally the lead for the component. * `UNASSIGNED` an assignee is not set for issues created with this component. * `PROJECT_DEFAULT` the assignee to any issues created with this component is nominally the default assignee for the project that the component is in.
	Description string `json:"description,omitempty"` // The description for the component.
	Projectid int64 `json:"projectId,omitempty"` // Not used.
	Isassigneetypevalid bool `json:"isAssigneeTypeValid,omitempty"` // Whether a user is associated with `assigneeType`. For example, if the `assigneeType` is set to `COMPONENT_LEAD` but the component lead is not set, then `false` is returned.
	Lead interface{} `json:"lead,omitempty"` // The user details for the component's lead user.
	Name string `json:"name,omitempty"` // The name for the component.
	Realassignee interface{} `json:"realAssignee,omitempty"` // The user assigned to issues created with this component, when `assigneeType` does not identify a valid assignee.
}

// StringList represents the StringList schema from the OpenAPI specification
type StringList struct {
}

// Scope represents the Scope schema from the OpenAPI specification
type Scope struct {
	Project interface{} `json:"project,omitempty"` // The project the item has scope in.
	TypeField string `json:"type,omitempty"` // The type of scope.
}

// ProjectDetails represents the ProjectDetails schema from the OpenAPI specification
type ProjectDetails struct {
	Key string `json:"key,omitempty"` // The key of the project.
	Name string `json:"name,omitempty"` // The name of the project.
	Projectcategory interface{} `json:"projectCategory,omitempty"` // The category the project belongs to.
	Projecttypekey string `json:"projectTypeKey,omitempty"` // The [project type](https://confluence.atlassian.com/x/GwiiLQ#Jiraapplicationsoverview-Productfeaturesandprojecttypes) of the project.
	Self string `json:"self,omitempty"` // The URL of the project details.
	Simplified bool `json:"simplified,omitempty"` // Whether or not the project is simplified.
	Avatarurls interface{} `json:"avatarUrls,omitempty"` // The URLs of the project's avatars.
	Id string `json:"id,omitempty"` // The ID of the project.
}

// IssueSecuritySchemeToProjectMapping represents the IssueSecuritySchemeToProjectMapping schema from the OpenAPI specification
type IssueSecuritySchemeToProjectMapping struct {
	Issuesecurityschemeid string `json:"issueSecuritySchemeId,omitempty"`
	Projectid string `json:"projectId,omitempty"`
}

// ServiceManagementNavigationInfo represents the ServiceManagementNavigationInfo schema from the OpenAPI specification
type ServiceManagementNavigationInfo struct {
	Queuecategory string `json:"queueCategory,omitempty"`
	Queueid int64 `json:"queueId,omitempty"`
	Queuename string `json:"queueName,omitempty"`
}

// Permissions represents the Permissions schema from the OpenAPI specification
type Permissions struct {
	Permissions map[string]interface{} `json:"permissions,omitempty"` // List of permissions.
}

// UserFilter represents the UserFilter schema from the OpenAPI specification
type UserFilter struct {
	Groups []string `json:"groups,omitempty"` // User groups autocomplete suggestion users must belong to. If not provided, the default values are used. A maximum of 10 groups can be provided.
	Roleids []int64 `json:"roleIds,omitempty"` // Roles that autocomplete suggestion users must belong to. If not provided, the default values are used. A maximum of 10 roles can be provided.
	Enabled bool `json:"enabled"` // Whether the filter is enabled.
}

// ListWrapperCallbackApplicationRole represents the ListWrapperCallbackApplicationRole schema from the OpenAPI specification
type ListWrapperCallbackApplicationRole struct {
}

// WorkflowOperations represents the WorkflowOperations schema from the OpenAPI specification
type WorkflowOperations struct {
	Candelete bool `json:"canDelete"` // Whether the workflow can be deleted.
	Canedit bool `json:"canEdit"` // Whether the workflow can be updated.
}

// FieldLastUsed represents the FieldLastUsed schema from the OpenAPI specification
type FieldLastUsed struct {
	TypeField string `json:"type,omitempty"` // Last used value type: * *TRACKED*: field is tracked and a last used date is available. * *NOT\_TRACKED*: field is not tracked, last used date is not available. * *NO\_INFORMATION*: field is tracked, but no last used date is available.
	Value string `json:"value,omitempty"` // The date when the value of the field last changed.
}

// Workflow represents the Workflow schema from the OpenAPI specification
type Workflow struct {
	Description string `json:"description"` // The description of the workflow.
	Schemes []WorkflowSchemeIdName `json:"schemes,omitempty"` // The workflow schemes the workflow is assigned to.
	Id PublishedWorkflowId `json:"id"` // Properties that identify a published workflow.
	Operations WorkflowOperations `json:"operations,omitempty"` // Operations allowed on a workflow
	Projects []ProjectDetails `json:"projects,omitempty"` // The projects the workflow is assigned to, through workflow schemes.
	Updated string `json:"updated,omitempty"` // The last edited date of the workflow.
	Created string `json:"created,omitempty"` // The creation date of the workflow.
	Hasdraftworkflow bool `json:"hasDraftWorkflow,omitempty"` // Whether the workflow has a draft version.
	Statuses []WorkflowStatus `json:"statuses,omitempty"` // The statuses of the workflow.
	Transitions []Transition `json:"transitions,omitempty"` // The transitions of the workflow.
	Isdefault bool `json:"isDefault,omitempty"` // Whether this is the default workflow.
}

// JQLReferenceData represents the JQLReferenceData schema from the OpenAPI specification
type JQLReferenceData struct {
	Jqlreservedwords []string `json:"jqlReservedWords,omitempty"` // List of JQL query reserved words.
	Visiblefieldnames []FieldReferenceData `json:"visibleFieldNames,omitempty"` // List of fields usable in JQL queries.
	Visiblefunctionnames []FunctionReferenceData `json:"visibleFunctionNames,omitempty"` // List of functions usable in JQL queries.
}

// CustomFieldConfigurations represents the CustomFieldConfigurations schema from the OpenAPI specification
type CustomFieldConfigurations struct {
	Configurations []ContextualConfiguration `json:"configurations"` // The list of custom field configuration details.
}

// IssueTypeIds represents the IssueTypeIds schema from the OpenAPI specification
type IssueTypeIds struct {
	Issuetypeids []string `json:"issueTypeIds"` // The list of issue type IDs.
}

// PageBeanFilterDetails represents the PageBeanFilterDetails schema from the OpenAPI specification
type PageBeanFilterDetails struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []FilterDetails `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// SecuritySchemeId represents the SecuritySchemeId schema from the OpenAPI specification
type SecuritySchemeId struct {
	Id string `json:"id"` // The ID of the issue security scheme.
}

// NotificationEvent represents the NotificationEvent schema from the OpenAPI specification
type NotificationEvent struct {
	Description string `json:"description,omitempty"` // The description of the event.
	Id int64 `json:"id,omitempty"` // The ID of the event. The event can be a [Jira system event](https://confluence.atlassian.com/x/8YdKLg#Creatinganotificationscheme-eventsEvents) or a [custom event](https://confluence.atlassian.com/x/AIlKLg).
	Name string `json:"name,omitempty"` // The name of the event.
	Templateevent interface{} `json:"templateEvent,omitempty"` // The template of the event. Only custom events configured by Jira administrators have template.
}

// PageBeanUserDetails represents the PageBeanUserDetails schema from the OpenAPI specification
type PageBeanUserDetails struct {
	Values []UserDetails `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// ScreenDetails represents the ScreenDetails schema from the OpenAPI specification
type ScreenDetails struct {
	Name string `json:"name"` // The name of the screen. The name must be unique. The maximum length is 255 characters.
	Description string `json:"description,omitempty"` // The description of the screen. The maximum length is 255 characters.
}

// IssueEntityPropertiesForMultiUpdate represents the IssueEntityPropertiesForMultiUpdate schema from the OpenAPI specification
type IssueEntityPropertiesForMultiUpdate struct {
	Issueid int64 `json:"issueID,omitempty"` // The ID of the issue.
	Properties map[string]interface{} `json:"properties,omitempty"` // Entity properties to set on the issue. The maximum length of an issue property value is 32768 characters.
}

// PageBeanSecuritySchemeWithProjects represents the PageBeanSecuritySchemeWithProjects schema from the OpenAPI specification
type PageBeanSecuritySchemeWithProjects struct {
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []SecuritySchemeWithProjects `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
}

// IssueTypeDetails represents the IssueTypeDetails schema from the OpenAPI specification
type IssueTypeDetails struct {
	Name string `json:"name,omitempty"` // The name of the issue type.
	Self string `json:"self,omitempty"` // The URL of these issue type details.
	Subtask bool `json:"subtask,omitempty"` // Whether this issue type is used to create subtasks.
	Hierarchylevel int `json:"hierarchyLevel,omitempty"` // Hierarchy level of the issue type.
	Description string `json:"description,omitempty"` // The description of the issue type.
	Id string `json:"id,omitempty"` // The ID of the issue type.
	Avatarid int64 `json:"avatarId,omitempty"` // The ID of the issue type's avatar.
	Entityid string `json:"entityId,omitempty"` // Unique ID for next-gen projects.
	Iconurl string `json:"iconUrl,omitempty"` // The URL of the issue type's avatar.
	Scope interface{} `json:"scope,omitempty"` // Details of the next-gen projects the issue type is available in.
}

// IssueTypeUpdateBean represents the IssueTypeUpdateBean schema from the OpenAPI specification
type IssueTypeUpdateBean struct {
	Avatarid int64 `json:"avatarId,omitempty"` // The ID of an issue type avatar.
	Description string `json:"description,omitempty"` // The description of the issue type.
	Name string `json:"name,omitempty"` // The unique name for the issue type. The maximum length is 60 characters.
}

// AvailableDashboardGadget represents the AvailableDashboardGadget schema from the OpenAPI specification
type AvailableDashboardGadget struct {
	Title string `json:"title"` // The title of the gadget.
	Uri string `json:"uri,omitempty"` // The URI of the gadget type.
	Modulekey string `json:"moduleKey,omitempty"` // The module key of the gadget type.
}

// WorkflowIDs represents the WorkflowIDs schema from the OpenAPI specification
type WorkflowIDs struct {
	Name string `json:"name"` // The name of the workflow.
	Entityid string `json:"entityId,omitempty"` // The entity ID of the workflow.
}

// JiraExpressionComplexity represents the JiraExpressionComplexity schema from the OpenAPI specification
type JiraExpressionComplexity struct {
	Expensiveoperations string `json:"expensiveOperations"` // Information that can be used to determine how many [expensive operations](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#expensive-operations) the evaluation of the expression will perform. This information may be a formula or number. For example: * `issues.map(i => i.comments)` performs as many expensive operations as there are issues on the issues list. So this parameter returns `N`, where `N` is the size of issue list. * `new Issue(10010).comments` gets comments for one issue, so its complexity is `2` (`1` to retrieve issue 10010 from the database plus `1` to get its comments).
	Variables map[string]interface{} `json:"variables,omitempty"` // Variables used in the formula, mapped to the parts of the expression they refer to.
}

// PageBeanContextForProjectAndIssueType represents the PageBeanContextForProjectAndIssueType schema from the OpenAPI specification
type PageBeanContextForProjectAndIssueType struct {
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []ContextForProjectAndIssueType `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
}

// CompoundClause represents the CompoundClause schema from the OpenAPI specification
type CompoundClause struct {
	Clauses []JqlQueryClause `json:"clauses"` // The list of nested clauses.
	Operator string `json:"operator"` // The operator between the clauses.
}

// UpdateFieldConfigurationSchemeDetails represents the UpdateFieldConfigurationSchemeDetails schema from the OpenAPI specification
type UpdateFieldConfigurationSchemeDetails struct {
	Name string `json:"name"` // The name of the field configuration scheme. The name must be unique.
	Description string `json:"description,omitempty"` // The description of the field configuration scheme.
}

// StatusCategory represents the StatusCategory schema from the OpenAPI specification
type StatusCategory struct {
	Key string `json:"key,omitempty"` // The key of the status category.
	Name string `json:"name,omitempty"` // The name of the status category.
	Self string `json:"self,omitempty"` // The URL of the status category.
	Colorname string `json:"colorName,omitempty"` // The name of the color used to represent the status category.
	Id int64 `json:"id,omitempty"` // The ID of the status category.
}

// JiraExpressionEvaluationMetaDataBean represents the JiraExpressionEvaluationMetaDataBean schema from the OpenAPI specification
type JiraExpressionEvaluationMetaDataBean struct {
	Complexity interface{} `json:"complexity,omitempty"` // Contains information about the expression complexity. For example, the number of steps it took to evaluate the expression.
	Issues interface{} `json:"issues,omitempty"` // Contains information about the `issues` variable in the context. For example, is the issues were loaded with JQL, information about the page will be included here.
}

// FunctionReferenceData represents the FunctionReferenceData schema from the OpenAPI specification
type FunctionReferenceData struct {
	Displayname string `json:"displayName,omitempty"` // The display name of the function.
	Islist string `json:"isList,omitempty"` // Whether the function can take a list of arguments.
	Types []string `json:"types,omitempty"` // The data types returned by the function.
	Value string `json:"value,omitempty"` // The function identifier.
}

// OrderOfIssueTypes represents the OrderOfIssueTypes schema from the OpenAPI specification
type OrderOfIssueTypes struct {
	After string `json:"after,omitempty"` // The ID of the issue type to place the moved issue types after. Required if `position` isn't provided.
	Issuetypeids []string `json:"issueTypeIds"` // A list of the issue type IDs to move. The order of the issue type IDs in the list is the order they are given after the move.
	Position string `json:"position,omitempty"` // The position the issue types should be moved to. Required if `after` isn't provided.
}

// CreatePriorityDetails represents the CreatePriorityDetails schema from the OpenAPI specification
type CreatePriorityDetails struct {
	Name string `json:"name"` // The name of the priority. Must be unique.
	Statuscolor string `json:"statusColor"` // The status color of the priority in 3-digit or 6-digit hexadecimal format.
	Description string `json:"description,omitempty"` // The description of the priority.
	Iconurl string `json:"iconUrl,omitempty"` // The URL of an icon for the priority. Accepted protocols are HTTP and HTTPS. Built in icons can also be used.
}

// BulkProjectPermissionGrants represents the BulkProjectPermissionGrants schema from the OpenAPI specification
type BulkProjectPermissionGrants struct {
	Issues []int64 `json:"issues"` // IDs of the issues the user has the permission for.
	Permission string `json:"permission"` // A project permission,
	Projects []int64 `json:"projects"` // IDs of the projects the user has the permission for.
}

// PageBeanWorkflowScheme represents the PageBeanWorkflowScheme schema from the OpenAPI specification
type PageBeanWorkflowScheme struct {
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []WorkflowScheme `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
}

// SharePermissionInputBean represents the SharePermissionInputBean schema from the OpenAPI specification
type SharePermissionInputBean struct {
	Rights int `json:"rights,omitempty"` // The rights for the share permission.
	TypeField string `json:"type"` // The type of the share permission.Specify the type as follows: * `user` Share with a user. * `group` Share with a group. Specify `groupname` as well. * `project` Share with a project. Specify `projectId` as well. * `projectRole` Share with a project role in a project. Specify `projectId` and `projectRoleId` as well. * `global` Share globally, including anonymous users. If set, this type overrides all existing share permissions and must be deleted before any non-global share permissions is set. * `authenticated` Share with all logged-in users. This shows as `loggedin` in the response. If set, this type overrides all existing share permissions and must be deleted before any non-global share permissions is set.
	Accountid string `json:"accountId,omitempty"` // The user account ID that the filter is shared with. For a request, specify the `accountId` property for the user.
	Groupid string `json:"groupId,omitempty"` // The ID of the group, which uniquely identifies the group across all Atlassian products.For example, *952d12c3-5b5b-4d04-bb32-44d383afc4b2*. Cannot be provided with `groupname`.
	Groupname string `json:"groupname,omitempty"` // The name of the group to share the filter with. Set `type` to `group`. Please note that the name of a group is mutable, to reliably identify a group use `groupId`.
	Projectid string `json:"projectId,omitempty"` // The ID of the project to share the filter with. Set `type` to `project`.
	Projectroleid string `json:"projectRoleId,omitempty"` // The ID of the project role to share the filter with. Set `type` to `projectRole` and the `projectId` for the project that the role is in.
}

// UpdateIssueSecuritySchemeRequestBean represents the UpdateIssueSecuritySchemeRequestBean schema from the OpenAPI specification
type UpdateIssueSecuritySchemeRequestBean struct {
	Description string `json:"description,omitempty"` // The description of the security scheme scheme.
	Name string `json:"name,omitempty"` // The name of the security scheme scheme. Must be unique.
}

// CreatedIssue represents the CreatedIssue schema from the OpenAPI specification
type CreatedIssue struct {
	Transition interface{} `json:"transition,omitempty"` // The response code and messages related to any requested transition.
	Watchers interface{} `json:"watchers,omitempty"` // The response code and messages related to any requested watchers.
	Id string `json:"id,omitempty"` // The ID of the created issue or subtask.
	Key string `json:"key,omitempty"` // The key of the created issue or subtask.
	Self string `json:"self,omitempty"` // The URL of the created issue or subtask.
}

// UiModificationContextDetails represents the UiModificationContextDetails schema from the OpenAPI specification
type UiModificationContextDetails struct {
	Id string `json:"id,omitempty"` // The ID of the UI modification context.
	Isavailable bool `json:"isAvailable,omitempty"` // Whether a context is available. For example, when a project is deleted the context becomes unavailable.
	Issuetypeid string `json:"issueTypeId"` // The issue type ID of the context.
	Projectid string `json:"projectId"` // The project ID of the context.
	Viewtype string `json:"viewType"` // The view type of the context. Only `GIC` (Global Issue Create) is supported.
}

// WorkflowSimpleCondition represents the WorkflowSimpleCondition schema from the OpenAPI specification
type WorkflowSimpleCondition struct {
	Nodetype string `json:"nodeType"`
	TypeField string `json:"type"` // The type of the transition rule.
	Configuration map[string]interface{} `json:"configuration,omitempty"` // EXPERIMENTAL. The configuration of the transition rule.
}

// ProjectIssueTypesHierarchyLevel represents the ProjectIssueTypesHierarchyLevel schema from the OpenAPI specification
type ProjectIssueTypesHierarchyLevel struct {
	Entityid string `json:"entityId,omitempty"` // The ID of the issue type hierarchy level. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	Issuetypes []IssueTypeInfo `json:"issueTypes,omitempty"` // The list of issue types in the hierarchy level.
	Level int `json:"level,omitempty"` // The level of the issue type hierarchy level.
	Name string `json:"name,omitempty"` // The name of the issue type hierarchy level.
}

// PageBeanJqlFunctionPrecomputationBean represents the PageBeanJqlFunctionPrecomputationBean schema from the OpenAPI specification
type PageBeanJqlFunctionPrecomputationBean struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []JqlFunctionPrecomputationBean `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// CustomFieldOptionCreate represents the CustomFieldOptionCreate schema from the OpenAPI specification
type CustomFieldOptionCreate struct {
	Disabled bool `json:"disabled,omitempty"` // Whether the option is disabled.
	Optionid string `json:"optionId,omitempty"` // For cascading options, the ID of the custom field object containing the cascading option.
	Value string `json:"value"` // The value of the custom field option.
}

// FieldConfigurationScheme represents the FieldConfigurationScheme schema from the OpenAPI specification
type FieldConfigurationScheme struct {
	Description string `json:"description,omitempty"` // The description of the field configuration scheme.
	Id string `json:"id"` // The ID of the field configuration scheme.
	Name string `json:"name"` // The name of the field configuration scheme.
}

// JqlQueryOrderByClause represents the JqlQueryOrderByClause schema from the OpenAPI specification
type JqlQueryOrderByClause struct {
	Fields []JqlQueryOrderByClauseElement `json:"fields"` // The list of order-by clause fields and their ordering directives.
}

// JqlQueryClauseOperand represents the JqlQueryClauseOperand schema from the OpenAPI specification
type JqlQueryClauseOperand struct {
}

// WorkflowsWithTransitionRulesDetails represents the WorkflowsWithTransitionRulesDetails schema from the OpenAPI specification
type WorkflowsWithTransitionRulesDetails struct {
	Workflows []WorkflowTransitionRulesDetails `json:"workflows"` // The list of workflows with transition rules to delete.
}

// CustomFieldContextDefaultValueDate represents the CustomFieldContextDefaultValueDate schema from the OpenAPI specification
type CustomFieldContextDefaultValueDate struct {
	TypeField string `json:"type"`
	Usecurrent bool `json:"useCurrent,omitempty"` // Whether to use the current date.
	Date string `json:"date,omitempty"` // The default date in ISO format. Ignored if `useCurrent` is true.
}

// ConnectModules represents the ConnectModules schema from the OpenAPI specification
type ConnectModules struct {
	Modules []ConnectModule `json:"modules"` // A list of app modules in the same format as the `modules` property in the [app descriptor](https://developer.atlassian.com/cloud/jira/platform/app-descriptor/).
}

// SoftwareNavigationInfo represents the SoftwareNavigationInfo schema from the OpenAPI specification
type SoftwareNavigationInfo struct {
	Boardname string `json:"boardName,omitempty"`
	Simpleboard bool `json:"simpleBoard,omitempty"`
	Totalboardsinproject int64 `json:"totalBoardsInProject,omitempty"`
	Boardid int64 `json:"boardId,omitempty"`
}

// Worklog represents the Worklog schema from the OpenAPI specification
type Worklog struct {
	Timespent string `json:"timeSpent,omitempty"` // The time spent working on the issue as days (\#d), hours (\#h), or minutes (\#m or \#). Required when creating a worklog if `timeSpentSeconds` isn't provided. Optional when updating a worklog. Cannot be provided if `timeSpentSecond` is provided.
	Visibility interface{} `json:"visibility,omitempty"` // Details about any restrictions in the visibility of the worklog. Optional when creating or updating a worklog.
	Comment interface{} `json:"comment,omitempty"` // A comment about the worklog in [Atlassian Document Format](https://developer.atlassian.com/cloud/jira/platform/apis/document/structure/). Optional when creating or updating a worklog.
	Created string `json:"created,omitempty"` // The datetime on which the worklog was created.
	Self string `json:"self,omitempty"` // The URL of the worklog item.
	Started string `json:"started,omitempty"` // The datetime on which the worklog effort was started. Required when creating a worklog. Optional when updating a worklog.
	Updateauthor interface{} `json:"updateAuthor,omitempty"` // Details of the user who last updated the worklog.
	Issueid string `json:"issueId,omitempty"` // The ID of the issue this worklog is for.
	Timespentseconds int64 `json:"timeSpentSeconds,omitempty"` // The time in seconds spent working on the issue. Required when creating a worklog if `timeSpent` isn't provided. Optional when updating a worklog. Cannot be provided if `timeSpent` is provided.
	Updated string `json:"updated,omitempty"` // The datetime on which the worklog was last updated.
	Author interface{} `json:"author,omitempty"` // Details of the user who created the worklog.
	Id string `json:"id,omitempty"` // The ID of the worklog record.
	Properties []EntityProperty `json:"properties,omitempty"` // Details of properties for the worklog. Optional when creating or updating a worklog.
}

// CustomFieldContextDefaultValueProject represents the CustomFieldContextDefaultValueProject schema from the OpenAPI specification
type CustomFieldContextDefaultValueProject struct {
	Contextid string `json:"contextId"` // The ID of the context.
	Projectid string `json:"projectId"` // The ID of the default project.
	TypeField string `json:"type"`
}

// CreateProjectDetails represents the CreateProjectDetails schema from the OpenAPI specification
type CreateProjectDetails struct {
	Avatarid int64 `json:"avatarId,omitempty"` // An integer value for the project's avatar.
	Categoryid int64 `json:"categoryId,omitempty"` // The ID of the project's category. A complete list of category IDs is found using the [Get all project categories](#api-rest-api-3-projectCategory-get) operation.
	Projecttemplatekey string `json:"projectTemplateKey,omitempty"` // A predefined configuration for a project. The type of the `projectTemplateKey` must match with the type of the `projectTypeKey`.
	Url string `json:"url,omitempty"` // A link to information about this project, such as project documentation
	Issuetypescheme int64 `json:"issueTypeScheme,omitempty"` // The ID of the issue type scheme for the project. Use the [Get all issue type schemes](#api-rest-api-3-issuetypescheme-get) operation to get a list of issue type scheme IDs. If you specify the issue type scheme you cannot specify the project template key.
	Lead string `json:"lead,omitempty"` // This parameter is deprecated because of privacy changes. Use `leadAccountId` instead. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. The user name of the project lead. Either `lead` or `leadAccountId` must be set when creating a project. Cannot be provided with `leadAccountId`.
	Issuetypescreenscheme int64 `json:"issueTypeScreenScheme,omitempty"` // The ID of the issue type screen scheme for the project. Use the [Get all issue type screen schemes](#api-rest-api-3-issuetypescreenscheme-get) operation to get a list of issue type screen scheme IDs. If you specify the issue type screen scheme you cannot specify the project template key.
	Workflowscheme int64 `json:"workflowScheme,omitempty"` // The ID of the workflow scheme for the project. Use the [Get all workflow schemes](#api-rest-api-3-workflowscheme-get) operation to get a list of workflow scheme IDs. If you specify the workflow scheme you cannot specify the project template key.
	Fieldconfigurationscheme int64 `json:"fieldConfigurationScheme,omitempty"` // The ID of the field configuration scheme for the project. Use the [Get all field configuration schemes](#api-rest-api-3-fieldconfigurationscheme-get) operation to get a list of field configuration scheme IDs. If you specify the field configuration scheme you cannot specify the project template key.
	Notificationscheme int64 `json:"notificationScheme,omitempty"` // The ID of the notification scheme for the project. Use the [Get notification schemes](#api-rest-api-3-notificationscheme-get) resource to get a list of notification scheme IDs.
	Permissionscheme int64 `json:"permissionScheme,omitempty"` // The ID of the permission scheme for the project. Use the [Get all permission schemes](#api-rest-api-3-permissionscheme-get) resource to see a list of all permission scheme IDs.
	Key string `json:"key"` // Project keys must be unique and start with an uppercase letter followed by one or more uppercase alphanumeric characters. The maximum length is 10 characters.
	Leadaccountid string `json:"leadAccountId,omitempty"` // The account ID of the project lead. Either `lead` or `leadAccountId` must be set when creating a project. Cannot be provided with `lead`.
	Description string `json:"description,omitempty"` // A brief description of the project.
	Assigneetype string `json:"assigneeType,omitempty"` // The default assignee when creating issues for this project.
	Issuesecurityscheme int64 `json:"issueSecurityScheme,omitempty"` // The ID of the issue security scheme for the project, which enables you to control who can and cannot view issues. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) resource to get all issue security scheme IDs.
	Name string `json:"name"` // The name of the project.
	Projecttypekey string `json:"projectTypeKey,omitempty"` // The [project type](https://confluence.atlassian.com/x/GwiiLQ#Jiraapplicationsoverview-Productfeaturesandprojecttypes), which defines the application-specific feature set. If you don't specify the project template you have to specify the project type.
}

// IssueTypeSchemeMapping represents the IssueTypeSchemeMapping schema from the OpenAPI specification
type IssueTypeSchemeMapping struct {
	Issuetypeid string `json:"issueTypeId"` // The ID of the issue type.
	Issuetypeschemeid string `json:"issueTypeSchemeId"` // The ID of the issue type scheme.
}

// ScreenScheme represents the ScreenScheme schema from the OpenAPI specification
type ScreenScheme struct {
	Description string `json:"description,omitempty"` // The description of the screen scheme.
	Id int64 `json:"id,omitempty"` // The ID of the screen scheme.
	Issuetypescreenschemes interface{} `json:"issueTypeScreenSchemes,omitempty"` // Details of the issue type screen schemes associated with the screen scheme.
	Name string `json:"name,omitempty"` // The name of the screen scheme.
	Screens interface{} `json:"screens,omitempty"` // The IDs of the screens for the screen types of the screen scheme.
}

// CustomFieldContextDefaultValueCascadingOption represents the CustomFieldContextDefaultValueCascadingOption schema from the OpenAPI specification
type CustomFieldContextDefaultValueCascadingOption struct {
	Optionid string `json:"optionId"` // The ID of the default option.
	TypeField string `json:"type"`
	Cascadingoptionid string `json:"cascadingOptionId,omitempty"` // The ID of the default cascading option.
	Contextid string `json:"contextId"` // The ID of the context.
}

// WorkflowTransitionProperty represents the WorkflowTransitionProperty schema from the OpenAPI specification
type WorkflowTransitionProperty struct {
	Id string `json:"id,omitempty"` // The ID of the transition property.
	Key string `json:"key,omitempty"` // The key of the transition property. Also known as the name of the transition property.
	Value string `json:"value"` // The value of the transition property.
}

// WorkflowTransitionRule represents the WorkflowTransitionRule schema from the OpenAPI specification
type WorkflowTransitionRule struct {
	Configuration interface{} `json:"configuration,omitempty"` // EXPERIMENTAL. The configuration of the transition rule.
	TypeField string `json:"type"` // The type of the transition rule.
}

// WebhookRegistrationDetails represents the WebhookRegistrationDetails schema from the OpenAPI specification
type WebhookRegistrationDetails struct {
	Url string `json:"url"` // The URL that specifies where to send the webhooks. This URL must use the same base URL as the Connect app. Only a single URL per app is allowed to be registered.
	Webhooks []WebhookDetails `json:"webhooks"` // A list of webhooks.
}

// JQLPersonalDataMigrationRequest represents the JQLPersonalDataMigrationRequest schema from the OpenAPI specification
type JQLPersonalDataMigrationRequest struct {
	Querystrings []string `json:"queryStrings,omitempty"` // A list of queries with user identifiers. Maximum of 100 queries.
}

// PermissionGrant represents the PermissionGrant schema from the OpenAPI specification
type PermissionGrant struct {
	Holder interface{} `json:"holder,omitempty"` // The user or group being granted the permission. It consists of a `type`, a type-dependent `parameter` and a type-dependent `value`. See [Holder object](../api-group-permission-schemes/#holder-object) in *Get all permission schemes* for more information.
	Id int64 `json:"id,omitempty"` // The ID of the permission granted details.
	Permission string `json:"permission,omitempty"` // The permission to grant. This permission can be one of the built-in permissions or a custom permission added by an app. See [Built-in permissions](../api-group-permission-schemes/#built-in-permissions) in *Get all permission schemes* for more information about the built-in permissions. See the [project permission](https://developer.atlassian.com/cloud/jira/platform/modules/project-permission/) and [global permission](https://developer.atlassian.com/cloud/jira/platform/modules/global-permission/) module documentation for more information about custom permissions.
	Self string `json:"self,omitempty"` // The URL of the permission granted details.
}

// ParsedJqlQuery represents the ParsedJqlQuery schema from the OpenAPI specification
type ParsedJqlQuery struct {
	Errors []string `json:"errors,omitempty"` // The list of syntax or validation errors.
	Query string `json:"query"` // The JQL query that was parsed and validated.
	Structure interface{} `json:"structure,omitempty"` // The syntax tree of the query. Empty if the query was invalid.
}

// CustomContextVariable represents the CustomContextVariable schema from the OpenAPI specification
type CustomContextVariable struct {
	TypeField string `json:"type"` // Type of custom context variable.
}

// ChangedWorklogs represents the ChangedWorklogs schema from the OpenAPI specification
type ChangedWorklogs struct {
	Lastpage bool `json:"lastPage,omitempty"`
	Nextpage string `json:"nextPage,omitempty"` // The URL of the next list of changed worklogs.
	Self string `json:"self,omitempty"` // The URL of this changed worklogs list.
	Since int64 `json:"since,omitempty"` // The datetime of the first worklog item in the list.
	Until int64 `json:"until,omitempty"` // The datetime of the last worklog item in the list.
	Values []ChangedWorklog `json:"values,omitempty"` // Changed worklog list.
}

// IssueFieldOptionScopeBean represents the IssueFieldOptionScopeBean schema from the OpenAPI specification
type IssueFieldOptionScopeBean struct {
	Global interface{} `json:"global,omitempty"` // Defines the behavior of the option within the global context. If this property is set, even if set to an empty object, then the option is available in all projects.
	Projects []int64 `json:"projects,omitempty"` // DEPRECATED
	Projects2 []ProjectScopeBean `json:"projects2,omitempty"` // Defines the projects in which the option is available and the behavior of the option within each project. Specify one object per project. The behavior of the option in a project context overrides the behavior in the global context.
}

// UserPermission represents the UserPermission schema from the OpenAPI specification
type UserPermission struct {
	Deprecatedkey bool `json:"deprecatedKey,omitempty"` // Indicate whether the permission key is deprecated. Note that deprecated keys cannot be used in the `permissions parameter of Get my permissions. Deprecated keys are not returned by Get all permissions.`
	Description string `json:"description,omitempty"` // The description of the permission.
	Havepermission bool `json:"havePermission,omitempty"` // Whether the permission is available to the user in the queried context.
	Id string `json:"id,omitempty"` // The ID of the permission. Either `id` or `key` must be specified. Use [Get all permissions](#api-rest-api-3-permissions-get) to get the list of permissions.
	Key string `json:"key,omitempty"` // The key of the permission. Either `id` or `key` must be specified. Use [Get all permissions](#api-rest-api-3-permissions-get) to get the list of permissions.
	Name string `json:"name,omitempty"` // The name of the permission.
	TypeField string `json:"type,omitempty"` // The type of the permission.
}

// UserBean represents the UserBean schema from the OpenAPI specification
type UserBean struct {
	Avatarurls interface{} `json:"avatarUrls,omitempty"` // The avatars of the user.
	Displayname string `json:"displayName,omitempty"` // The display name of the user. Depending on the user’s privacy setting, this may return an alternative value.
	Key string `json:"key,omitempty"` // This property is deprecated in favor of `accountId` because of privacy changes. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. The key of the user.
	Name string `json:"name,omitempty"` // This property is deprecated in favor of `accountId` because of privacy changes. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. The username of the user.
	Self string `json:"self,omitempty"` // The URL of the user.
	Accountid string `json:"accountId,omitempty"` // The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*.
	Active bool `json:"active,omitempty"` // Whether the user is active.
}

// ProjectId represents the ProjectId schema from the OpenAPI specification
type ProjectId struct {
	Id string `json:"id"` // The ID of the project.
}

// RegisteredWebhook represents the RegisteredWebhook schema from the OpenAPI specification
type RegisteredWebhook struct {
	Createdwebhookid int64 `json:"createdWebhookId,omitempty"` // The ID of the webhook. Returned if the webhook is created.
	Errors []string `json:"errors,omitempty"` // Error messages specifying why the webhook creation failed.
}

// ProjectInsight represents the ProjectInsight schema from the OpenAPI specification
type ProjectInsight struct {
	Lastissueupdatetime string `json:"lastIssueUpdateTime,omitempty"` // The last issue update time.
	Totalissuecount int64 `json:"totalIssueCount,omitempty"` // Total issue count.
}

// UpdateCustomFieldDetails represents the UpdateCustomFieldDetails schema from the OpenAPI specification
type UpdateCustomFieldDetails struct {
	Description string `json:"description,omitempty"` // The description of the custom field. The maximum length is 40000 characters.
	Name string `json:"name,omitempty"` // The name of the custom field. It doesn't have to be unique. The maximum length is 255 characters.
	Searcherkey string `json:"searcherKey,omitempty"` // The searcher that defines the way the field is searched in Jira. It can be set to `null`, otherwise you must specify the valid searcher for the field type, as listed below (abbreviated values shown): * `cascadingselect`: `cascadingselectsearcher` * `datepicker`: `daterange` * `datetime`: `datetimerange` * `float`: `exactnumber` or `numberrange` * `grouppicker`: `grouppickersearcher` * `importid`: `exactnumber` or `numberrange` * `labels`: `labelsearcher` * `multicheckboxes`: `multiselectsearcher` * `multigrouppicker`: `multiselectsearcher` * `multiselect`: `multiselectsearcher` * `multiuserpicker`: `userpickergroupsearcher` * `multiversion`: `versionsearcher` * `project`: `projectsearcher` * `radiobuttons`: `multiselectsearcher` * `readonlyfield`: `textsearcher` * `select`: `multiselectsearcher` * `textarea`: `textsearcher` * `textfield`: `textsearcher` * `url`: `exacttextsearcher` * `userpicker`: `userpickergroupsearcher` * `version`: `versionsearcher`
}

// CustomFieldValueUpdate represents the CustomFieldValueUpdate schema from the OpenAPI specification
type CustomFieldValueUpdate struct {
	Issueids []int64 `json:"issueIds"` // The list of issue IDs.
	Value interface{} `json:"value"` // The value for the custom field. The value must be compatible with the [custom field type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/#data-types) as follows: * `string` the value must be a string. * `number` the value must be a number. * `datetime` the value must be a string that represents a date in the ISO format or the simplified extended ISO format. For example, `"2023-01-18T12:00:00-03:00"` or `"2023-01-18T12:00:00.000Z"`. However, the milliseconds part is ignored. * `user` the value must be an object that contains the `accountId` field. * `group` the value must be an object that contains the group `name` or `groupId` field. Because group names can change, we recommend using `groupId`. A list of appropriate values must be provided if the field is of the `list` [collection type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/#collection-types).
}

// WebhooksExpirationDate represents the WebhooksExpirationDate schema from the OpenAPI specification
type WebhooksExpirationDate struct {
	Expirationdate int64 `json:"expirationDate"` // The expiration date of all the refreshed webhooks.
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Projectcategory interface{} `json:"projectCategory,omitempty"` // The category the project belongs to.
	Issuetypes []IssueTypeDetails `json:"issueTypes,omitempty"` // List of the issue types available in the project.
	Key string `json:"key,omitempty"` // The key of the project.
	Uuid string `json:"uuid,omitempty"` // Unique ID for next-gen projects.
	Style string `json:"style,omitempty"` // The type of the project.
	Expand string `json:"expand,omitempty"` // Expand options that include additional project details in the response.
	Archiveddate string `json:"archivedDate,omitempty"` // The date when the project was archived.
	Url string `json:"url,omitempty"` // A link to information about this project, such as project documentation.
	Roles map[string]interface{} `json:"roles,omitempty"` // The name and self URL for each role defined in the project. For more information, see [Create project role](#api-rest-api-3-role-post).
	Lead interface{} `json:"lead,omitempty"` // The username of the project lead.
	Avatarurls interface{} `json:"avatarUrls,omitempty"` // The URLs of the project's avatars.
	Simplified bool `json:"simplified,omitempty"` // Whether the project is simplified.
	Assigneetype string `json:"assigneeType,omitempty"` // The default assignee when creating issues for this project.
	Self string `json:"self,omitempty"` // The URL of the project details.
	Email string `json:"email,omitempty"` // An email address associated with the project.
	Favourite bool `json:"favourite,omitempty"` // Whether the project is selected as a favorite.
	Id string `json:"id,omitempty"` // The ID of the project.
	Deleted bool `json:"deleted,omitempty"` // Whether the project is marked as deleted.
	Deletedby interface{} `json:"deletedBy,omitempty"` // The user who marked the project as deleted.
	Versions []Version `json:"versions,omitempty"` // The versions defined in the project. For more information, see [Create version](#api-rest-api-3-version-post).
	Components []ProjectComponent `json:"components,omitempty"` // List of the components contained in the project.
	Name string `json:"name,omitempty"` // The name of the project.
	Projecttypekey string `json:"projectTypeKey,omitempty"` // The [project type](https://confluence.atlassian.com/x/GwiiLQ#Jiraapplicationsoverview-Productfeaturesandprojecttypes) of the project.
	Properties map[string]interface{} `json:"properties,omitempty"` // Map of project properties
	Landingpageinfo interface{} `json:"landingPageInfo,omitempty"` // The project landing page info.
	Insight interface{} `json:"insight,omitempty"` // Insights about the project.
	Deleteddate string `json:"deletedDate,omitempty"` // The date when the project was marked as deleted.
	Description string `json:"description,omitempty"` // A brief description of the project.
	Permissions interface{} `json:"permissions,omitempty"` // User permissions on the project
	Archived bool `json:"archived,omitempty"` // Whether the project is archived.
	Issuetypehierarchy interface{} `json:"issueTypeHierarchy,omitempty"` // The issue type hierarchy for the project.
	Retentiontilldate string `json:"retentionTillDate,omitempty"` // The date when the project is deleted permanently.
	Archivedby interface{} `json:"archivedBy,omitempty"` // The user who archived the project.
	Isprivate bool `json:"isPrivate,omitempty"` // Whether the project is private.
}

// ResolutionJsonBean represents the ResolutionJsonBean schema from the OpenAPI specification
type ResolutionJsonBean struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Self string `json:"self,omitempty"`
	DefaultField bool `json:"default,omitempty"`
	Description string `json:"description,omitempty"`
	Iconurl string `json:"iconUrl,omitempty"`
}

// IssueLinkType represents the IssueLinkType schema from the OpenAPI specification
type IssueLinkType struct {
	Name string `json:"name,omitempty"` // The name of the issue link type and is used as follows: * In the [ issueLink](#api-rest-api-3-issueLink-post) resource it is the type of issue link. Required on create when `id` isn't provided. Otherwise, read only. * In the [ issueLinkType](#api-rest-api-3-issueLinkType-post) resource it is required on create and optional on update. Otherwise, read only.
	Outward string `json:"outward,omitempty"` // The description of the issue link type outward link and is used as follows: * In the [ issueLink](#api-rest-api-3-issueLink-post) resource it is read only. * In the [ issueLinkType](#api-rest-api-3-issueLinkType-post) resource it is required on create and optional on update. Otherwise, read only.
	Self string `json:"self,omitempty"` // The URL of the issue link type. Read only.
	Id string `json:"id,omitempty"` // The ID of the issue link type and is used as follows: * In the [ issueLink](#api-rest-api-3-issueLink-post) resource it is the type of issue link. Required on create when `name` isn't provided. Otherwise, read only. * In the [ issueLinkType](#api-rest-api-3-issueLinkType-post) resource it is read only.
	Inward string `json:"inward,omitempty"` // The description of the issue link type inward link and is used as follows: * In the [ issueLink](#api-rest-api-3-issueLink-post) resource it is read only. * In the [ issueLinkType](#api-rest-api-3-issueLinkType-post) resource it is required on create and optional on update. Otherwise, read only.
}

// ProjectIssueTypes represents the ProjectIssueTypes schema from the OpenAPI specification
type ProjectIssueTypes struct {
	Project ProjectId `json:"project,omitempty"` // Project ID details.
	Issuetypes []string `json:"issueTypes,omitempty"` // IDs of the issue types
}

// PageBeanContext represents the PageBeanContext schema from the OpenAPI specification
type PageBeanContext struct {
	Values []Context `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
}

// WorkflowRulesSearch represents the WorkflowRulesSearch schema from the OpenAPI specification
type WorkflowRulesSearch struct {
	Expand string `json:"expand,omitempty"` // Use expand to include additional information in the response. This parameter accepts `transition` which, for each rule, returns information about the transition the rule is assigned to.
	Ruleids []string `json:"ruleIds"` // The list of workflow rule IDs.
	Workflowentityid string `json:"workflowEntityId"` // The workflow ID.
}

// UpdateProjectDetails represents the UpdateProjectDetails schema from the OpenAPI specification
type UpdateProjectDetails struct {
	Assigneetype string `json:"assigneeType,omitempty"` // The default assignee when creating issues for this project.
	Avatarid int64 `json:"avatarId,omitempty"` // An integer value for the project's avatar.
	Issuesecurityscheme int64 `json:"issueSecurityScheme,omitempty"` // The ID of the issue security scheme for the project, which enables you to control who can and cannot view issues. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) resource to get all issue security scheme IDs.
	Leadaccountid string `json:"leadAccountId,omitempty"` // The account ID of the project lead. Cannot be provided with `lead`.
	Permissionscheme int64 `json:"permissionScheme,omitempty"` // The ID of the permission scheme for the project. Use the [Get all permission schemes](#api-rest-api-3-permissionscheme-get) resource to see a list of all permission scheme IDs.
	Key string `json:"key,omitempty"` // Project keys must be unique and start with an uppercase letter followed by one or more uppercase alphanumeric characters. The maximum length is 10 characters.
	Lead string `json:"lead,omitempty"` // This parameter is deprecated because of privacy changes. Use `leadAccountId` instead. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. The user name of the project lead. Cannot be provided with `leadAccountId`.
	Name string `json:"name,omitempty"` // The name of the project.
	Notificationscheme int64 `json:"notificationScheme,omitempty"` // The ID of the notification scheme for the project. Use the [Get notification schemes](#api-rest-api-3-notificationscheme-get) resource to get a list of notification scheme IDs.
	Url string `json:"url,omitempty"` // A link to information about this project, such as project documentation
	Description string `json:"description,omitempty"` // A brief description of the project.
	Categoryid int64 `json:"categoryId,omitempty"` // The ID of the project's category. A complete list of category IDs is found using the [Get all project categories](#api-rest-api-3-projectCategory-get) operation. To remove the project category from the project, set the value to `-1.`
}

// CreateWorkflowStatusDetails represents the CreateWorkflowStatusDetails schema from the OpenAPI specification
type CreateWorkflowStatusDetails struct {
	Id string `json:"id"` // The ID of the status.
	Properties map[string]interface{} `json:"properties,omitempty"` // The properties of the status.
}

// VersionIssueCounts represents the VersionIssueCounts schema from the OpenAPI specification
type VersionIssueCounts struct {
	Issuesaffectedcount int64 `json:"issuesAffectedCount,omitempty"` // Count of issues where the `affectedVersion` is set to the version.
	Issuesfixedcount int64 `json:"issuesFixedCount,omitempty"` // Count of issues where the `fixVersion` is set to the version.
	Self string `json:"self,omitempty"` // The URL of these count details.
	Customfieldusage []VersionUsageInCustomField `json:"customFieldUsage,omitempty"` // List of custom fields using the version.
	Issuecountwithcustomfieldsshowingversion int64 `json:"issueCountWithCustomFieldsShowingVersion,omitempty"` // Count of issues where a version custom field is set to the version.
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Resolved bool `json:"resolved,omitempty"` // Whether the item is resolved. If set to "true", the link to the issue is displayed in a strikethrough font, otherwise the link displays in normal font.
	Icon interface{} `json:"icon,omitempty"` // Details of the icon representing the status. If not provided, no status icon displays in Jira.
}

// IssueTypeScreenSchemeItem represents the IssueTypeScreenSchemeItem schema from the OpenAPI specification
type IssueTypeScreenSchemeItem struct {
	Issuetypeid string `json:"issueTypeId"` // The ID of the issue type or *default*. Only issue types used in classic projects are accepted. When creating an issue screen scheme, an entry for *default* must be provided and defines the mapping for all issue types without a screen scheme. Otherwise, a *default* entry can't be provided.
	Issuetypescreenschemeid string `json:"issueTypeScreenSchemeId"` // The ID of the issue type screen scheme.
	Screenschemeid string `json:"screenSchemeId"` // The ID of the screen scheme.
}

// ScreenableField represents the ScreenableField schema from the OpenAPI specification
type ScreenableField struct {
	Id string `json:"id,omitempty"` // The ID of the screen tab field.
	Name string `json:"name,omitempty"` // The name of the screen tab field. Required on create and update. The maximum length is 255 characters.
}

// SimpleListWrapperApplicationRole represents the SimpleListWrapperApplicationRole schema from the OpenAPI specification
type SimpleListWrapperApplicationRole struct {
	Max_results int `json:"max-results,omitempty"`
	Pagingcallback ListWrapperCallbackApplicationRole `json:"pagingCallback,omitempty"`
	Size int `json:"size,omitempty"`
	Callback ListWrapperCallbackApplicationRole `json:"callback,omitempty"`
	Items []ApplicationRole `json:"items,omitempty"`
}

// Locale represents the Locale schema from the OpenAPI specification
type Locale struct {
	Locale string `json:"locale,omitempty"` // The locale code. The Java the locale format is used: a two character language code (ISO 639), an underscore, and two letter country code (ISO 3166). For example, en\_US represents a locale of English (United States). Required on create.
}

// SimpleLink represents the SimpleLink schema from the OpenAPI specification
type SimpleLink struct {
	Href string `json:"href,omitempty"`
	Iconclass string `json:"iconClass,omitempty"`
	Id string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Styleclass string `json:"styleClass,omitempty"`
	Title string `json:"title,omitempty"`
	Weight int `json:"weight,omitempty"`
}

// AvatarUrlsBean represents the AvatarUrlsBean schema from the OpenAPI specification
type AvatarUrlsBean struct {
	Field24x24 string `json:"24x24,omitempty"` // The URL of the item's 24x24 pixel avatar.
	Field32x32 string `json:"32x32,omitempty"` // The URL of the item's 32x32 pixel avatar.
	Field48x48 string `json:"48x48,omitempty"` // The URL of the item's 48x48 pixel avatar.
	Field16x16 string `json:"16x16,omitempty"` // The URL of the item's 16x16 pixel avatar.
}

// ProjectFeatureState represents the ProjectFeatureState schema from the OpenAPI specification
type ProjectFeatureState struct {
	State string `json:"state,omitempty"` // The feature state.
}

// FunctionOperand represents the FunctionOperand schema from the OpenAPI specification
type FunctionOperand struct {
	Function string `json:"function"` // The name of the function.
	Arguments []string `json:"arguments"` // The list of function arguments.
	Encodedoperand string `json:"encodedOperand,omitempty"` // Encoded operand, which can be used directly in a JQL query.
}

// AnnouncementBannerConfiguration represents the AnnouncementBannerConfiguration schema from the OpenAPI specification
type AnnouncementBannerConfiguration struct {
	Visibility string `json:"visibility,omitempty"` // Visibility of the announcement banner.
	Hashid string `json:"hashId,omitempty"` // Hash of the banner data. The client detects updates by comparing hash IDs.
	Isdismissible bool `json:"isDismissible,omitempty"` // Flag indicating if the announcement banner can be dismissed by the user.
	Isenabled bool `json:"isEnabled,omitempty"` // Flag indicating if the announcement banner is enabled or not.
	Message string `json:"message,omitempty"` // The text on the announcement banner.
}

// FieldDetails represents the FieldDetails schema from the OpenAPI specification
type FieldDetails struct {
	Custom bool `json:"custom,omitempty"` // Whether the field is a custom field.
	Orderable bool `json:"orderable,omitempty"` // Whether the content of the field can be used to order lists.
	Searchable bool `json:"searchable,omitempty"` // Whether the content of the field can be searched.
	Name string `json:"name,omitempty"` // The name of the field.
	Clausenames []string `json:"clauseNames,omitempty"` // The names that can be used to reference the field in an advanced search. For more information, see [Advanced searching - fields reference](https://confluence.atlassian.com/x/gwORLQ).
	Key string `json:"key,omitempty"` // The key of the field.
	Scope interface{} `json:"scope,omitempty"` // The scope of the field.
	Id string `json:"id,omitempty"` // The ID of the field.
	Navigable bool `json:"navigable,omitempty"` // Whether the field can be used as a column on the issue navigator.
	Schema interface{} `json:"schema,omitempty"` // The data schema for the field.
}

// NotificationSchemeId represents the NotificationSchemeId schema from the OpenAPI specification
type NotificationSchemeId struct {
	Id string `json:"id"` // The ID of a notification scheme.
}

// PageBeanVersion represents the PageBeanVersion schema from the OpenAPI specification
type PageBeanVersion struct {
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []Version `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
}

// FoundUsers represents the FoundUsers schema from the OpenAPI specification
type FoundUsers struct {
	Header string `json:"header,omitempty"` // Header text indicating the number of users in the response and the total number of users found in the search.
	Total int `json:"total,omitempty"` // The total number of users found in the search.
	Users []UserPickerUser `json:"users,omitempty"`
}

// FilterSubscriptionsList represents the FilterSubscriptionsList schema from the OpenAPI specification
type FilterSubscriptionsList struct {
	End_index int `json:"end-index,omitempty"` // The index of the last item returned on the page.
	Items []FilterSubscription `json:"items,omitempty"` // The list of items.
	Max_results int `json:"max-results,omitempty"` // The maximum number of results that could be on the page.
	Size int `json:"size,omitempty"` // The number of items on the page.
	Start_index int `json:"start-index,omitempty"` // The index of the first item returned on the page.
}

// PublishedWorkflowId represents the PublishedWorkflowId schema from the OpenAPI specification
type PublishedWorkflowId struct {
	Entityid string `json:"entityId,omitempty"` // The entity ID of the workflow.
	Name string `json:"name"` // The name of the workflow.
}

// PageBeanWebhook represents the PageBeanWebhook schema from the OpenAPI specification
type PageBeanWebhook struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []Webhook `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
}

// JqlFunctionPrecomputationUpdateRequestBean represents the JqlFunctionPrecomputationUpdateRequestBean schema from the OpenAPI specification
type JqlFunctionPrecomputationUpdateRequestBean struct {
	Values []JqlFunctionPrecomputationUpdateBean `json:"values,omitempty"`
}

// KeywordOperand represents the KeywordOperand schema from the OpenAPI specification
type KeywordOperand struct {
	Keyword string `json:"keyword"` // The keyword that is the operand value.
}

// IssueTypeSchemeDetails represents the IssueTypeSchemeDetails schema from the OpenAPI specification
type IssueTypeSchemeDetails struct {
	Defaultissuetypeid string `json:"defaultIssueTypeId,omitempty"` // The ID of the default issue type of the issue type scheme. This ID must be included in `issueTypeIds`.
	Description string `json:"description,omitempty"` // The description of the issue type scheme. The maximum length is 4000 characters.
	Issuetypeids []string `json:"issueTypeIds"` // The list of issue types IDs of the issue type scheme. At least one standard issue type ID is required.
	Name string `json:"name"` // The name of the issue type scheme. The name must be unique. The maximum length is 255 characters.
}

// UserKey represents the UserKey schema from the OpenAPI specification
type UserKey struct {
	Key string `json:"key,omitempty"` // This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Accountid string `json:"accountId,omitempty"` // The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. Returns *unknown* if the record is deleted and corrupted, for example, as the result of a server import.
}

// JqlQueryUnitaryOperand represents the JqlQueryUnitaryOperand schema from the OpenAPI specification
type JqlQueryUnitaryOperand struct {
}

// FieldReferenceData represents the FieldReferenceData schema from the OpenAPI specification
type FieldReferenceData struct {
	Types []string `json:"types,omitempty"` // The data types of items in the field.
	Auto string `json:"auto,omitempty"` // Whether the field provide auto-complete suggestions.
	Orderable string `json:"orderable,omitempty"` // Whether the field can be used in a query's `ORDER BY` clause.
	Value string `json:"value,omitempty"` // The field identifier.
	Deprecated string `json:"deprecated,omitempty"` // Whether this field has been deprecated.
	Displayname string `json:"displayName,omitempty"` // The display name contains the following: * for system fields, the field name. For example, `Summary`. * for collapsed custom fields, the field name followed by a hyphen and then the field name and field type. For example, `Component - Component[Dropdown]`. * for other custom fields, the field name followed by a hyphen and then the custom field ID. For example, `Component - cf[10061]`.
	Cfid string `json:"cfid,omitempty"` // If the item is a custom field, the ID of the custom field.
	Deprecatedsearcherkey string `json:"deprecatedSearcherKey,omitempty"` // The searcher key of the field, only passed when the field is deprecated.
	Operators []string `json:"operators,omitempty"` // The valid search operators for the field.
	Searchable string `json:"searchable,omitempty"` // Whether the content of this field can be searched.
}

// Context represents the Context schema from the OpenAPI specification
type Context struct {
	Id int64 `json:"id,omitempty"` // The ID of the context.
	Name string `json:"name,omitempty"` // The name of the context.
	Scope interface{} `json:"scope,omitempty"` // The scope of the context.
}

// OrderOfCustomFieldOptions represents the OrderOfCustomFieldOptions schema from the OpenAPI specification
type OrderOfCustomFieldOptions struct {
	After string `json:"after,omitempty"` // The ID of the custom field option or cascading option to place the moved options after. Required if `position` isn't provided.
	Customfieldoptionids []string `json:"customFieldOptionIds"` // A list of IDs of custom field options to move. The order of the custom field option IDs in the list is the order they are given after the move. The list must contain custom field options or cascading options, but not both.
	Position string `json:"position,omitempty"` // The position the custom field options should be moved to. Required if `after` isn't provided.
}

// UpdateScreenTypes represents the UpdateScreenTypes schema from the OpenAPI specification
type UpdateScreenTypes struct {
	DefaultField string `json:"default,omitempty"` // The ID of the default screen. When specified, must include a screen ID as a default screen is required.
	Edit string `json:"edit,omitempty"` // The ID of the edit screen. To remove the screen association, pass a null.
	View string `json:"view,omitempty"` // The ID of the view screen. To remove the screen association, pass a null.
	Create string `json:"create,omitempty"` // The ID of the create screen. To remove the screen association, pass a null.
}

// ProjectEmailAddress represents the ProjectEmailAddress schema from the OpenAPI specification
type ProjectEmailAddress struct {
	Emailaddressstatus []string `json:"emailAddressStatus,omitempty"` // When using a custom domain, the status of the email address.
	Emailaddress string `json:"emailAddress,omitempty"` // The email address.
}

// LicensedApplication represents the LicensedApplication schema from the OpenAPI specification
type LicensedApplication struct {
	Id string `json:"id"` // The ID of the application.
	Plan string `json:"plan"` // The licensing plan.
}

// ProjectIssueTypeMapping represents the ProjectIssueTypeMapping schema from the OpenAPI specification
type ProjectIssueTypeMapping struct {
	Issuetypeid string `json:"issueTypeId"` // The ID of the issue type.
	Projectid string `json:"projectId"` // The ID of the project.
}

// ColumnItem represents the ColumnItem schema from the OpenAPI specification
type ColumnItem struct {
	Value string `json:"value,omitempty"` // The issue navigator column value.
	Label string `json:"label,omitempty"` // The issue navigator column label.
}

// IssueTypeScreenSchemeMappingDetails represents the IssueTypeScreenSchemeMappingDetails schema from the OpenAPI specification
type IssueTypeScreenSchemeMappingDetails struct {
	Issuetypemappings []IssueTypeScreenSchemeMapping `json:"issueTypeMappings"` // The list of issue type to screen scheme mappings. A *default* entry cannot be specified because a default entry is added when an issue type screen scheme is created.
}

// UnrestrictedUserEmail represents the UnrestrictedUserEmail schema from the OpenAPI specification
type UnrestrictedUserEmail struct {
	Accountid string `json:"accountId,omitempty"` // The accountId of the user
	Email string `json:"email,omitempty"` // The email of the user
}

// WorkflowTransitionRulesUpdate represents the WorkflowTransitionRulesUpdate schema from the OpenAPI specification
type WorkflowTransitionRulesUpdate struct {
	Workflows []WorkflowTransitionRules `json:"workflows"` // The list of workflows with transition rules to update.
}

// IssueBean represents the IssueBean schema from the OpenAPI specification
type IssueBean struct {
	Self string `json:"self,omitempty"` // The URL of the issue details.
	Versionedrepresentations map[string]interface{} `json:"versionedRepresentations,omitempty"` // The versions of each field on the issue.
	Properties map[string]interface{} `json:"properties,omitempty"` // Details of the issue properties identified in the request.
	Names map[string]interface{} `json:"names,omitempty"` // The ID and name of each field present on the issue.
	Key string `json:"key,omitempty"` // The key of the issue.
	Changelog interface{} `json:"changelog,omitempty"` // Details of changelogs associated with the issue.
	Editmeta interface{} `json:"editmeta,omitempty"` // The metadata for the fields on the issue that can be amended.
	Renderedfields map[string]interface{} `json:"renderedFields,omitempty"` // The rendered value of each field present on the issue.
	Id string `json:"id,omitempty"` // The ID of the issue.
	Fields map[string]interface{} `json:"fields,omitempty"`
	Fieldstoinclude IncludedFields `json:"fieldsToInclude,omitempty"`
	Schema map[string]interface{} `json:"schema,omitempty"` // The schema describing each field present on the issue.
	Transitions []IssueTransition `json:"transitions,omitempty"` // The transitions that can be performed on the issue.
	Operations interface{} `json:"operations,omitempty"` // The operations that can be performed on the issue.
	Expand string `json:"expand,omitempty"` // Expand options that include additional issue details in the response.
}

// JqlQueryClause represents the JqlQueryClause schema from the OpenAPI specification
type JqlQueryClause struct {
}

// MultipleCustomFieldValuesUpdateDetails represents the MultipleCustomFieldValuesUpdateDetails schema from the OpenAPI specification
type MultipleCustomFieldValuesUpdateDetails struct {
	Updates []MultipleCustomFieldValuesUpdate `json:"updates,omitempty"`
}

// ProjectIds represents the ProjectIds schema from the OpenAPI specification
type ProjectIds struct {
	Projectids []string `json:"projectIds"` // The IDs of projects.
}

// UserBeanAvatarUrls represents the UserBeanAvatarUrls schema from the OpenAPI specification
type UserBeanAvatarUrls struct {
	Field32x32 string `json:"32x32,omitempty"` // The URL of the user's 32x32 pixel avatar.
	Field48x48 string `json:"48x48,omitempty"` // The URL of the user's 48x48 pixel avatar.
	Field16x16 string `json:"16x16,omitempty"` // The URL of the user's 16x16 pixel avatar.
	Field24x24 string `json:"24x24,omitempty"` // The URL of the user's 24x24 pixel avatar.
}

// SharePermission represents the SharePermission schema from the OpenAPI specification
type SharePermission struct {
	Role interface{} `json:"role,omitempty"` // The project role that the filter is shared with. For a request, specify the `id` for the role. You must also specify the `project` object and `id` for the project that the role is in.
	TypeField string `json:"type"` // The type of share permission: * `user` Shared with a user. * `group` Shared with a group. If set in a request, then specify `sharePermission.group` as well. * `project` Shared with a project. If set in a request, then specify `sharePermission.project` as well. * `projectRole` Share with a project role in a project. This value is not returned in responses. It is used in requests, where it needs to be specify with `projectId` and `projectRoleId`. * `global` Shared globally. If set in a request, no other `sharePermission` properties need to be specified. * `loggedin` Shared with all logged-in users. Note: This value is set in a request by specifying `authenticated` as the `type`. * `project-unknown` Shared with a project that the user does not have access to. Cannot be set in a request.
	User interface{} `json:"user,omitempty"` // The user account ID that the filter is shared with. For a request, specify the `accountId` property for the user.
	Group interface{} `json:"group,omitempty"` // The group that the filter is shared with. For a request, specify the `groupId` or `name` property for the group. As a group's name can change, use of `groupId` is recommended.
	Id int64 `json:"id,omitempty"` // The unique identifier of the share permission.
	Project interface{} `json:"project,omitempty"` // The project that the filter is shared with. This is similar to the project object returned by [Get project](#api-rest-api-3-project-projectIdOrKey-get) but it contains a subset of the properties, which are: `self`, `id`, `key`, `assigneeType`, `name`, `roles`, `avatarUrls`, `projectType`, `simplified`. For a request, specify the `id` for the project.
}

// ErrorMessage represents the ErrorMessage schema from the OpenAPI specification
type ErrorMessage struct {
	Message string `json:"message"` // The error message.
}

// Field represents the Field schema from the OpenAPI specification
type Field struct {
	Key string `json:"key,omitempty"` // The key of the field.
	Name string `json:"name"` // The name of the field.
	Lastused FieldLastUsed `json:"lastUsed,omitempty"` // Information about the most recent use of a field.
	Searcherkey string `json:"searcherKey,omitempty"` // The searcher key of the field. Returned for custom fields.
	Contextscount int64 `json:"contextsCount,omitempty"` // Number of contexts where the field is used.
	Description string `json:"description,omitempty"` // The description of the field.
	Isunscreenable bool `json:"isUnscreenable,omitempty"` // Whether the field is shown on screen or not.
	Projectscount int64 `json:"projectsCount,omitempty"` // Number of projects where the field is used.
	Schema JsonTypeBean `json:"schema"` // The schema of a field.
	Id string `json:"id"` // The ID of the field.
	Screenscount int64 `json:"screensCount,omitempty"` // Number of screens where the field is used.
	Islocked bool `json:"isLocked,omitempty"` // Whether the field is locked.
}

// IssueFieldOption represents the IssueFieldOption schema from the OpenAPI specification
type IssueFieldOption struct {
	Config IssueFieldOptionConfiguration `json:"config,omitempty"` // Details of the projects the option is available in.
	Id int64 `json:"id"` // The unique identifier for the option. This is only unique within the select field's set of options.
	Properties map[string]interface{} `json:"properties,omitempty"` // The properties of the object, as arbitrary key-value pairs. These properties can be searched using JQL, if the extractions (see [Issue Field Option Property Index](https://developer.atlassian.com/cloud/jira/platform/modules/issue-field-option-property-index/)) are defined in the descriptor for the issue field module.
	Value string `json:"value"` // The option's name, which is displayed in Jira.
}

// PageBeanFieldConfigurationItem represents the PageBeanFieldConfigurationItem schema from the OpenAPI specification
type PageBeanFieldConfigurationItem struct {
	Nextpage string `json:"nextPage,omitempty"` // If there is another page of results, the URL of the next page.
	Self string `json:"self,omitempty"` // The URL of the page.
	Startat int64 `json:"startAt,omitempty"` // The index of the first item returned.
	Total int64 `json:"total,omitempty"` // The number of items returned.
	Values []FieldConfigurationItem `json:"values,omitempty"` // The list of items.
	Islast bool `json:"isLast,omitempty"` // Whether this is the last page.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that could be returned.
}

// Attachment represents the Attachment schema from the OpenAPI specification
type Attachment struct {
	Thumbnail string `json:"thumbnail,omitempty"` // The URL of a thumbnail representing the attachment.
	Mimetype string `json:"mimeType,omitempty"` // The MIME type of the attachment.
	Size int64 `json:"size,omitempty"` // The size of the attachment.
	Author interface{} `json:"author,omitempty"` // Details of the user who added the attachment.
	Content string `json:"content,omitempty"` // The content of the attachment.
	Filename string `json:"filename,omitempty"` // The file name of the attachment.
	Id string `json:"id,omitempty"` // The ID of the attachment.
	Created string `json:"created,omitempty"` // The datetime the attachment was created.
	Self string `json:"self,omitempty"` // The URL of the attachment details response.
}
