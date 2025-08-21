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

func BulkgetgroupsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["startAt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startAt=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["groupId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("groupId=%v", val))
		}
		if val, ok := args["groupName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("groupName=%v", val))
		}
		if val, ok := args["accessType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("accessType=%v", val))
		}
		if val, ok := args["applicationKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("applicationKey=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/group/bulk%s", cfg.BaseURL, queryString)
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
		var result models.PageBeanGroupDetails
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

func CreateBulkgetgroupsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_group_bulk",
		mcp.WithDescription("Bulk get groups"),
		mcp.WithNumber("startAt", mcp.Description("The index of the first item to return in a page of results (page offset).")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of items to return per page.")),
		mcp.WithArray("groupId", mcp.Description("The ID of a group. To specify multiple IDs, pass multiple `groupId` parameters. For example, `groupId=5b10a2844c20165700ede21g&groupId=5b10ac8d82e05b22cc7d4ef5`.")),
		mcp.WithArray("groupName", mcp.Description("The name of a group. To specify multiple names, pass multiple `groupName` parameters. For example, `groupName=administrators&groupName=jira-software-users`.")),
		mcp.WithString("accessType", mcp.Description("The access level of a group. Valid values: 'site-admin', 'admin', 'user'.")),
		mcp.WithString("applicationKey", mcp.Description("The application key of the product user groups to search for. Valid values: 'jira-servicedesk', 'jira-software', 'jira-product-discovery', 'jira-core'.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    BulkgetgroupsHandler(cfg),
	}
}
