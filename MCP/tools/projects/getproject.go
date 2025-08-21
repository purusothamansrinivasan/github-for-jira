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

func GetprojectHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		if val, ok := args["properties"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("properties=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/project/%s%s", cfg.BaseURL, projectIdOrKey, queryString)
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
		var result models.Project
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

func CreateGetprojectTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_project_projectIdOrKey",
		mcp.WithDescription("Get project"),
		mcp.WithString("projectIdOrKey", mcp.Required(), mcp.Description("The project ID or project key (case sensitive).")),
		mcp.WithString("expand", mcp.Description("Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Note that the project description, issue types, and project lead are included in all responses by default. Expand options include:\n\n *  `description` The project description.\n *  `issueTypes` The issue types associated with the project.\n *  `lead` The project lead.\n *  `projectKeys` All project keys associated with the project.\n *  `issueTypeHierarchy` The project issue type hierarchy.")),
		mcp.WithArray("properties", mcp.Description("A list of project properties to return for the project. This parameter accepts a comma-separated list.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetprojectHandler(cfg),
	}
}
