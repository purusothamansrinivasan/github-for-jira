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

func FinduserswithallpermissionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["username"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("username=%v", val))
		}
		if val, ok := args["accountId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("accountId=%v", val))
		}
		if val, ok := args["permissions"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("permissions=%v", val))
		}
		if val, ok := args["issueKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("issueKey=%v", val))
		}
		if val, ok := args["projectKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projectKey=%v", val))
		}
		if val, ok := args["startAt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startAt=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/user/permission/search%s", cfg.BaseURL, queryString)
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
		var result []User
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

func CreateFinduserswithallpermissionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_user_permission_search",
		mcp.WithDescription("Find users with permissions"),
		mcp.WithString("query", mcp.Description("A query string that is matched against user attributes, such as `displayName` and `emailAddress`, to find relevant users. The string can match the prefix of the attribute's value. For example, *query=john* matches a user with a `displayName` of *John Smith* and a user with an `emailAddress` of *johnson@example.com*. Required, unless `accountId` is specified.")),
		mcp.WithString("username", mcp.Description("This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.")),
		mcp.WithString("accountId", mcp.Description("A query string that is matched exactly against user `accountId`. Required, unless `query` is specified.")),
		mcp.WithString("permissions", mcp.Required(), mcp.Description("A comma separated list of permissions. Permissions can be specified as any:\n\n *  permission returned by [Get all permissions](#api-rest-api-3-permissions-get).\n *  custom project permission added by Connect apps.\n *  (deprecated) one of the following:\n    \n     *  ASSIGNABLE\\_USER\n     *  ASSIGN\\_ISSUE\n     *  ATTACHMENT\\_DELETE\\_ALL\n     *  ATTACHMENT\\_DELETE\\_OWN\n     *  BROWSE\n     *  CLOSE\\_ISSUE\n     *  COMMENT\\_DELETE\\_ALL\n     *  COMMENT\\_DELETE\\_OWN\n     *  COMMENT\\_EDIT\\_ALL\n     *  COMMENT\\_EDIT\\_OWN\n     *  COMMENT\\_ISSUE\n     *  CREATE\\_ATTACHMENT\n     *  CREATE\\_ISSUE\n     *  DELETE\\_ISSUE\n     *  EDIT\\_ISSUE\n     *  LINK\\_ISSUE\n     *  MANAGE\\_WATCHER\\_LIST\n     *  MODIFY\\_REPORTER\n     *  MOVE\\_ISSUE\n     *  PROJECT\\_ADMIN\n     *  RESOLVE\\_ISSUE\n     *  SCHEDULE\\_ISSUE\n     *  SET\\_ISSUE\\_SECURITY\n     *  TRANSITION\\_ISSUE\n     *  VIEW\\_VERSION\\_CONTROL\n     *  VIEW\\_VOTERS\\_AND\\_WATCHERS\n     *  VIEW\\_WORKFLOW\\_READONLY\n     *  WORKLOG\\_DELETE\\_ALL\n     *  WORKLOG\\_DELETE\\_OWN\n     *  WORKLOG\\_EDIT\\_ALL\n     *  WORKLOG\\_EDIT\\_OWN\n     *  WORK\\_ISSUE")),
		mcp.WithString("issueKey", mcp.Description("The issue key for the issue.")),
		mcp.WithString("projectKey", mcp.Description("The project key for the project (case sensitive).")),
		mcp.WithNumber("startAt", mcp.Description("The index of the first item to return in a page of results (page offset).")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of items to return per page.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    FinduserswithallpermissionsHandler(cfg),
	}
}
