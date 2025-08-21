package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/the-jira-cloud-platform-rest-api/mcp-server/config"
	"github.com/the-jira-cloud-platform-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdateprojectHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		projectIdOrKeyVal, ok := args["projectIdOrKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: projectIdOrKey"), nil
		}
		projectIdOrKey, ok := projectIdOrKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: projectIdOrKey"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.UpdateProjectDetails
		
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
		url := fmt.Sprintf("%s/rest/api/3/project/%s%s", cfg.BaseURL, projectIdOrKey, queryString)
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
		var result models.Project
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

func CreateUpdateprojectTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_rest_api_3_project_projectIdOrKey",
		mcp.WithDescription("Update project"),
		mcp.WithString("projectIdOrKey", mcp.Required(), mcp.Description("The project ID or project key (case sensitive).")),
		mcp.WithString("expand", mcp.Description("Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Note that the project description, issue types, and project lead are included in all responses by default. Expand options include:\n\n *  `description` The project description.\n *  `issueTypes` The issue types associated with the project.\n *  `lead` The project lead.\n *  `projectKeys` All project keys associated with the project.")),
		mcp.WithString("leadAccountId", mcp.Description("Input parameter: The account ID of the project lead. Cannot be provided with `lead`.")),
		mcp.WithNumber("permissionScheme", mcp.Description("Input parameter: The ID of the permission scheme for the project. Use the [Get all permission schemes](#api-rest-api-3-permissionscheme-get) resource to see a list of all permission scheme IDs.")),
		mcp.WithString("key", mcp.Description("Input parameter: Project keys must be unique and start with an uppercase letter followed by one or more uppercase alphanumeric characters. The maximum length is 10 characters.")),
		mcp.WithString("lead", mcp.Description("Input parameter: This parameter is deprecated because of privacy changes. Use `leadAccountId` instead. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. The user name of the project lead. Cannot be provided with `leadAccountId`.")),
		mcp.WithString("name", mcp.Description("Input parameter: The name of the project.")),
		mcp.WithNumber("notificationScheme", mcp.Description("Input parameter: The ID of the notification scheme for the project. Use the [Get notification schemes](#api-rest-api-3-notificationscheme-get) resource to get a list of notification scheme IDs.")),
		mcp.WithString("url", mcp.Description("Input parameter: A link to information about this project, such as project documentation")),
		mcp.WithString("description", mcp.Description("Input parameter: A brief description of the project.")),
		mcp.WithNumber("categoryId", mcp.Description("Input parameter: The ID of the project's category. A complete list of category IDs is found using the [Get all project categories](#api-rest-api-3-projectCategory-get) operation. To remove the project category from the project, set the value to `-1.`")),
		mcp.WithString("assigneeType", mcp.Description("Input parameter: The default assignee when creating issues for this project.")),
		mcp.WithNumber("avatarId", mcp.Description("Input parameter: An integer value for the project's avatar.")),
		mcp.WithNumber("issueSecurityScheme", mcp.Description("Input parameter: The ID of the issue security scheme for the project, which enables you to control who can and cannot view issues. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) resource to get all issue security scheme IDs.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateprojectHandler(cfg),
	}
}
