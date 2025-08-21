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

func GetsecuritylevelmembersHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["schemeId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("schemeId=%v", val))
		}
		if val, ok := args["levelId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("levelId=%v", val))
		}
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/issuesecurityschemes/level/member%s", cfg.BaseURL, queryString)
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
		var result models.PageBeanSecurityLevelMember
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

func CreateGetsecuritylevelmembersTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_issuesecurityschemes_level_member",
		mcp.WithDescription("Get issue security level members"),
		mcp.WithString("startAt", mcp.Description("The index of the first item to return in a page of results (page offset).")),
		mcp.WithString("maxResults", mcp.Description("The maximum number of items to return per page.")),
		mcp.WithArray("id", mcp.Description("The list of issue security level member IDs. To include multiple issue security level members separate IDs with an ampersand: `id=10000&id=10001`.")),
		mcp.WithArray("schemeId", mcp.Description("The list of issue security scheme IDs. To include multiple issue security schemes separate IDs with an ampersand: `schemeId=10000&schemeId=10001`.")),
		mcp.WithArray("levelId", mcp.Description("The list of issue security level IDs. To include multiple issue security levels separate IDs with an ampersand: `levelId=10000&levelId=10001`.")),
		mcp.WithString("expand", mcp.Description("Use expand to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:\n\n *  `all` Returns all expandable information\n *  `field` Returns information about the custom field granted the permission\n *  `group` Returns information about the group that is granted the permission\n *  `projectRole` Returns information about the project role granted the permission\n *  `user` Returns information about the user who is granted the permission")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsecuritylevelmembersHandler(cfg),
	}
}
