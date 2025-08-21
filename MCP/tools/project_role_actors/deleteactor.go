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

func DeleteactorHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["user"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user=%v", val))
		}
		if val, ok := args["group"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("group=%v", val))
		}
		if val, ok := args["groupId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("groupId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/project/%s/role/%s%s", cfg.BaseURL, projectIdOrKey, id, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreateDeleteactorTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_rest_api_3_project_projectIdOrKey_role_id",
		mcp.WithDescription("Delete actors from project role"),
		mcp.WithString("projectIdOrKey", mcp.Required(), mcp.Description("The project ID or project key (case sensitive).")),
		mcp.WithNumber("id", mcp.Required(), mcp.Description("The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.")),
		mcp.WithString("user", mcp.Description("The user account ID of the user to remove from the project role.")),
		mcp.WithString("group", mcp.Description("The name of the group to remove from the project role. This parameter cannot be used with the `groupId` parameter. As a group's name can change, use of `groupId` is recommended.")),
		mcp.WithString("groupId", mcp.Description("The ID of the group to remove from the project role. This parameter cannot be used with the `group` parameter.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteactorHandler(cfg),
	}
}
