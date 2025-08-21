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

func GetfilterspaginatedHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["filterName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filterName=%v", val))
		}
		if val, ok := args["accountId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("accountId=%v", val))
		}
		if val, ok := args["owner"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owner=%v", val))
		}
		if val, ok := args["groupname"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("groupname=%v", val))
		}
		if val, ok := args["groupId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("groupId=%v", val))
		}
		if val, ok := args["projectId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projectId=%v", val))
		}
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["orderBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderBy=%v", val))
		}
		if val, ok := args["startAt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startAt=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		if val, ok := args["overrideSharePermissions"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("overrideSharePermissions=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/filter/search%s", cfg.BaseURL, queryString)
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
		var result models.PageBeanFilterDetails
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

func CreateGetfilterspaginatedTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_filter_search",
		mcp.WithDescription("Search for filters"),
		mcp.WithString("filterName", mcp.Description("String used to perform a case-insensitive partial match with `name`.")),
		mcp.WithString("accountId", mcp.Description("User account ID used to return filters with the matching `owner.accountId`. This parameter cannot be used with `owner`.")),
		mcp.WithString("owner", mcp.Description("This parameter is deprecated because of privacy changes. Use `accountId` instead. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. User name used to return filters with the matching `owner.name`. This parameter cannot be used with `accountId`.")),
		mcp.WithString("groupname", mcp.Description("As a group's name can change, use of `groupId` is recommended to identify a group. Group name used to returns filters that are shared with a group that matches `sharePermissions.group.groupname`. This parameter cannot be used with the `groupId` parameter.")),
		mcp.WithString("groupId", mcp.Description("Group ID used to returns filters that are shared with a group that matches `sharePermissions.group.groupId`. This parameter cannot be used with the `groupname` parameter.")),
		mcp.WithNumber("projectId", mcp.Description("Project ID used to returns filters that are shared with a project that matches `sharePermissions.project.id`.")),
		mcp.WithArray("id", mcp.Description("The list of filter IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=10000&id=10001`. Do not exceed 200 filter IDs.")),
		mcp.WithString("orderBy", mcp.Description("[Order](#ordering) the results by a field:\n\n *  `description` Sorts by filter description. Note that this sorting works independently of whether the expand to display the description field is in use.\n *  `favourite_count` Sorts by the count of how many users have this filter as a favorite.\n *  `is_favourite` Sorts by whether the filter is marked as a favorite.\n *  `id` Sorts by filter ID.\n *  `name` Sorts by filter name.\n *  `owner` Sorts by the ID of the filter owner.\n *  `is_shared` Sorts by whether the filter is shared.")),
		mcp.WithNumber("startAt", mcp.Description("The index of the first item to return in a page of results (page offset).")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of items to return per page.")),
		mcp.WithString("expand", mcp.Description("Use [expand](#expansion) to include additional information about filter in the response. This parameter accepts a comma-separated list. Expand options include:\n\n *  `description` Returns the description of the filter.\n *  `favourite` Returns an indicator of whether the user has set the filter as a favorite.\n *  `favouritedCount` Returns a count of how many users have set this filter as a favorite.\n *  `jql` Returns the JQL query that the filter uses.\n *  `owner` Returns the owner of the filter.\n *  `searchUrl` Returns a URL to perform the filter's JQL query.\n *  `sharePermissions` Returns the share permissions defined for the filter.\n *  `editPermissions` Returns the edit permissions defined for the filter.\n *  `isWritable` Returns whether the current user has permission to edit the filter.\n *  `subscriptions` Returns the users that are subscribed to the filter.\n *  `viewUrl` Returns a URL to view the filter.")),
		mcp.WithBoolean("overrideSharePermissions", mcp.Description("EXPERIMENTAL: Whether share permissions are overridden to enable filters with any share permissions to be returned. Available to users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetfilterspaginatedHandler(cfg),
	}
}
