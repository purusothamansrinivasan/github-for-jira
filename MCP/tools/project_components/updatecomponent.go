package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/the-jira-cloud-platform-rest-api/mcp-server/config"
	"github.com/the-jira-cloud-platform-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdatecomponentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ProjectComponent
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/rest/api/3/component/%s", cfg.BaseURL, id)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BasicAuth != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Basic %s", cfg.BasicAuth))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.ProjectComponent
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateUpdatecomponentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_rest_api_3_component_id",
		mcp.WithDescription("Update component"),
		mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the component.")),
		mcp.WithString("assignee", mcp.Description("Input parameter: The details of the user associated with `assigneeType`, if any. See `realAssignee` for details of the user assigned to issues created with this component.")),
		mcp.WithString("leadUserName", mcp.Description("Input parameter: This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.")),
		mcp.WithString("self", mcp.Description("Input parameter: The URL of the component.")),
		mcp.WithString("id", mcp.Description("Input parameter: The unique identifier for the component.")),
		mcp.WithString("realAssignee", mcp.Description("Input parameter: The user assigned to issues created with this component, when `assigneeType` does not identify a valid assignee.")),
		mcp.WithString("description", mcp.Description("Input parameter: The description for the component. Optional when creating or updating a component.")),
		mcp.WithBoolean("isAssigneeTypeValid", mcp.Description("Input parameter: Whether a user is associated with `assigneeType`. For example, if the `assigneeType` is set to `COMPONENT_LEAD` but the component lead is not set, then `false` is returned.")),
		mcp.WithString("name", mcp.Description("Input parameter: The unique name for the component in the project. Required when creating a component. Optional when updating a component. The maximum length is 255 characters.")),
		mcp.WithString("project", mcp.Description("Input parameter: The key of the project the component is assigned to. Required when creating a component. Can't be updated.")),
		mcp.WithString("realAssigneeType", mcp.Description("Input parameter: The type of the assignee that is assigned to issues created with this component, when an assignee cannot be set from the `assigneeType`. For example, `assigneeType` is set to `COMPONENT_LEAD` but no component lead is set. This property is set to one of the following values:\n\n *  `PROJECT_LEAD` when `assigneeType` is `PROJECT_LEAD` and the project lead has permission to be assigned issues in the project that the component is in.\n *  `COMPONENT_LEAD` when `assignee`Type is `COMPONENT_LEAD` and the component lead has permission to be assigned issues in the project that the component is in.\n *  `UNASSIGNED` when `assigneeType` is `UNASSIGNED` and Jira is configured to allow unassigned issues.\n *  `PROJECT_DEFAULT` when none of the preceding cases are true.")),
		mcp.WithString("assigneeType", mcp.Description("Input parameter: The nominal user type used to determine the assignee for issues created with this component. See `realAssigneeType` for details on how the type of the user, and hence the user, assigned to issues is determined. Can take the following values:\n\n *  `PROJECT_LEAD` the assignee to any issues created with this component is nominally the lead for the project the component is in.\n *  `COMPONENT_LEAD` the assignee to any issues created with this component is nominally the lead for the component.\n *  `UNASSIGNED` an assignee is not set for issues created with this component.\n *  `PROJECT_DEFAULT` the assignee to any issues created with this component is nominally the default assignee for the project that the component is in.\n\nDefault value: `PROJECT_DEFAULT`.  \nOptional when creating or updating a component.")),
		mcp.WithString("lead", mcp.Description("Input parameter: The user details for the component's lead user.")),
		mcp.WithString("leadAccountId", mcp.Description("Input parameter: The accountId of the component's lead user. The accountId uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*.")),
		mcp.WithNumber("projectId", mcp.Description("Input parameter: The ID of the project the component is assigned to.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatecomponentHandler(cfg),
	}
}
