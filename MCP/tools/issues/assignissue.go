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

func AssignissueHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		issueIdOrKeyVal, ok := args["issueIdOrKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: issueIdOrKey"), nil
		}
		issueIdOrKey, ok := issueIdOrKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: issueIdOrKey"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.User
		
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
		url := fmt.Sprintf("%s/rest/api/3/issue/%s/assignee", cfg.BaseURL, issueIdOrKey)
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
		var result map[string]interface{}
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

func CreateAssignissueTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_rest_api_3_issue_issueIdOrKey_assignee",
		mcp.WithDescription("Assign issue"),
		mcp.WithString("issueIdOrKey", mcp.Required(), mcp.Description("The ID or key of the issue to be assigned.")),
		mcp.WithBoolean("active", mcp.Description("Input parameter: Whether the user is active.")),
		mcp.WithString("expand", mcp.Description("Input parameter: Expand options that include additional user details in the response.")),
		mcp.WithString("groups", mcp.Description("Input parameter: The groups that the user belongs to.")),
		mcp.WithString("name", mcp.Description("Input parameter: This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.")),
		mcp.WithString("accountId", mcp.Description("Input parameter: The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. Required in requests.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: The display name of the user. Depending on the user’s privacy setting, this may return an alternative value.")),
		mcp.WithString("locale", mcp.Description("Input parameter: The locale of the user. Depending on the user’s privacy setting, this may be returned as null.")),
		mcp.WithString("self", mcp.Description("Input parameter: The URL of the user.")),
		mcp.WithString("emailAddress", mcp.Description("Input parameter: The email address of the user. Depending on the user’s privacy setting, this may be returned as null.")),
		mcp.WithString("key", mcp.Description("Input parameter: This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.")),
		mcp.WithString("accountType", mcp.Description("Input parameter: The user account type. Can take the following values:\n\n *  `atlassian` regular Atlassian user account\n *  `app` system account used for Connect applications and OAuth to represent external systems\n *  `customer` Jira Service Desk account representing an external service desk")),
		mcp.WithString("applicationRoles", mcp.Description("Input parameter: The application roles the user is assigned to.")),
		mcp.WithString("avatarUrls", mcp.Description("Input parameter: The avatars of the user.")),
		mcp.WithString("timeZone", mcp.Description("Input parameter: The time zone specified in the user's profile. Depending on the user’s privacy setting, this may be returned as null.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AssignissueHandler(cfg),
	}
}
