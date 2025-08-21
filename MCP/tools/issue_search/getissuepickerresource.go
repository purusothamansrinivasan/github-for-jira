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

func GetissuepickerresourceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["currentJQL"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("currentJQL=%v", val))
		}
		if val, ok := args["currentIssueKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("currentIssueKey=%v", val))
		}
		if val, ok := args["currentProjectId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("currentProjectId=%v", val))
		}
		if val, ok := args["showSubTasks"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("showSubTasks=%v", val))
		}
		if val, ok := args["showSubTaskParent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("showSubTaskParent=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/issue/picker%s", cfg.BaseURL, queryString)
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
		var result models.IssuePickerSuggestions
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

func CreateGetissuepickerresourceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_issue_picker",
		mcp.WithDescription("Get issue picker suggestions"),
		mcp.WithString("query", mcp.Description("A string to match against text fields in the issue such as title, description, or comments.")),
		mcp.WithString("currentJQL", mcp.Description("A JQL query defining a list of issues to search for the query term. Note that `username` and `userkey` cannot be used as search terms for this parameter, due to privacy reasons. Use `accountId` instead.")),
		mcp.WithString("currentIssueKey", mcp.Description("The key of an issue to exclude from search results. For example, the issue the user is viewing when they perform this query.")),
		mcp.WithString("currentProjectId", mcp.Description("The ID of a project that suggested issues must belong to.")),
		mcp.WithBoolean("showSubTasks", mcp.Description("Indicate whether to include subtasks in the suggestions list.")),
		mcp.WithBoolean("showSubTaskParent", mcp.Description("When `currentIssueKey` is a subtask, whether to include the parent issue in the suggestions if it matches the query.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetissuepickerresourceHandler(cfg),
	}
}
