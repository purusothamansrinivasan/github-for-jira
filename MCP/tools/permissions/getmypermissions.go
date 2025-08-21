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

func GetmypermissionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["projectKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projectKey=%v", val))
		}
		if val, ok := args["projectId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projectId=%v", val))
		}
		if val, ok := args["issueKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("issueKey=%v", val))
		}
		if val, ok := args["issueId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("issueId=%v", val))
		}
		if val, ok := args["permissions"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("permissions=%v", val))
		}
		if val, ok := args["projectUuid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projectUuid=%v", val))
		}
		if val, ok := args["projectConfigurationUuid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projectConfigurationUuid=%v", val))
		}
		if val, ok := args["commentId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("commentId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/mypermissions%s", cfg.BaseURL, queryString)
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
		var result models.Permissions
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

func CreateGetmypermissionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_mypermissions",
		mcp.WithDescription("Get my permissions"),
		mcp.WithString("projectKey", mcp.Description("The key of project. Ignored if `projectId` is provided.")),
		mcp.WithString("projectId", mcp.Description("The ID of project.")),
		mcp.WithString("issueKey", mcp.Description("The key of the issue. Ignored if `issueId` is provided.")),
		mcp.WithString("issueId", mcp.Description("The ID of the issue.")),
		mcp.WithString("permissions", mcp.Description("A list of permission keys. (Required) This parameter accepts a comma-separated list. To get the list of available permissions, use [Get all permissions](#api-rest-api-3-permissions-get).")),
		mcp.WithString("projectUuid", mcp.Description("")),
		mcp.WithString("projectConfigurationUuid", mcp.Description("")),
		mcp.WithString("commentId", mcp.Description("The ID of the comment.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetmypermissionsHandler(cfg),
	}
}
