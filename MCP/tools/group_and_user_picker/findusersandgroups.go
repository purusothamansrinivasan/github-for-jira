package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/the-jira-cloud-platform-rest-api/mcp-server/config"
	"github.com/the-jira-cloud-platform-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func FindusersandgroupsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["showAvatar"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("showAvatar=%v", val))
		}
		if val, ok := args["fieldId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fieldId=%v", val))
		}
		if val, ok := args["projectId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projectId=%v", val))
		}
		if val, ok := args["issueTypeId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("issueTypeId=%v", val))
		}
		if val, ok := args["avatarSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("avatarSize=%v", val))
		}
		if val, ok := args["caseInsensitive"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("caseInsensitive=%v", val))
		}
		if val, ok := args["excludeConnectAddons"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("excludeConnectAddons=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/groupuserpicker%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.FoundUsersAndGroups
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

func CreateFindusersandgroupsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_groupuserpicker",
		mcp.WithDescription("Find users and groups"),
		mcp.WithString("query", mcp.Required(), mcp.Description("The search string.")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of items to return in each list.")),
		mcp.WithBoolean("showAvatar", mcp.Description("Whether the user avatar should be returned. If an invalid value is provided, the default value is used.")),
		mcp.WithString("fieldId", mcp.Description("The custom field ID of the field this request is for.")),
		mcp.WithArray("projectId", mcp.Description("The ID of a project that returned users and groups must have permission to view. To include multiple projects, provide an ampersand-separated list. For example, `projectId=10000&projectId=10001`. This parameter is only used when `fieldId` is present.")),
		mcp.WithArray("issueTypeId", mcp.Description("The ID of an issue type that returned users and groups must have permission to view. To include multiple issue types, provide an ampersand-separated list. For example, `issueTypeId=10000&issueTypeId=10001`. Special values, such as `-1` (all standard issue types) and `-2` (all subtask issue types), are supported. This parameter is only used when `fieldId` is present.")),
		mcp.WithString("avatarSize", mcp.Description("The size of the avatar to return. If an invalid value is provided, the default value is used.")),
		mcp.WithBoolean("caseInsensitive", mcp.Description("Whether the search for groups should be case insensitive.")),
		mcp.WithBoolean("excludeConnectAddons", mcp.Description("Whether Connect app users and groups should be excluded from the search results. If an invalid value is provided, the default value is used.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    FindusersandgroupsHandler(cfg),
	}
}
