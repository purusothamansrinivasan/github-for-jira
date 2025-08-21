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

func GetprojectversionspaginatedHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["startAt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startAt=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["orderBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderBy=%v", val))
		}
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status=%v", val))
		}
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/project/%s/version%s", cfg.BaseURL, projectIdOrKey, queryString)
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
		var result models.PageBeanVersion
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

func CreateGetprojectversionspaginatedTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_project_projectIdOrKey_version",
		mcp.WithDescription("Get project versions paginated"),
		mcp.WithString("projectIdOrKey", mcp.Required(), mcp.Description("The project ID or project key (case sensitive).")),
		mcp.WithNumber("startAt", mcp.Description("The index of the first item to return in a page of results (page offset).")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of items to return per page.")),
		mcp.WithString("orderBy", mcp.Description("[Order](#ordering) the results by a field:\n\n *  `description` Sorts by version description.\n *  `name` Sorts by version name.\n *  `releaseDate` Sorts by release date, starting with the oldest date. Versions with no release date are listed last.\n *  `sequence` Sorts by the order of appearance in the user interface.\n *  `startDate` Sorts by start date, starting with the oldest date. Versions with no start date are listed last.")),
		mcp.WithString("query", mcp.Description("Filter the results using a literal string. Versions with matching `name` or `description` are returned (case insensitive).")),
		mcp.WithString("status", mcp.Description("A list of status values used to filter the results by version status. This parameter accepts a comma-separated list. The status values are `released`, `unreleased`, and `archived`.")),
		mcp.WithString("expand", mcp.Description("Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:\n\n *  `issuesstatus` Returns the number of issues in each status category for each version.\n *  `operations` Returns actions that can be performed on the specified version.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetprojectversionspaginatedHandler(cfg),
	}
}
