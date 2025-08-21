package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/the-jira-cloud-platform-rest-api/mcp-server/config"
	"github.com/the-jira-cloud-platform-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func AddactorusersHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.ActorsMap
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/rest/api/3/project/%s/role/%s", cfg.BaseURL, projectIdOrKey, id)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.ProjectRole
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

func CreateAddactorusersTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_rest_api_3_project_projectIdOrKey_role_id",
		mcp.WithDescription("Add actors to project role"),
		mcp.WithString("projectIdOrKey", mcp.Required(), mcp.Description("The project ID or project key (case sensitive).")),
		mcp.WithNumber("id", mcp.Required(), mcp.Description("The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.")),
		mcp.WithArray("groupId", mcp.Description("Input parameter: The ID of the group to add. This parameter cannot be used with the `group` parameter.")),
		mcp.WithArray("user", mcp.Description("Input parameter: The user account ID of the user to add.")),
		mcp.WithArray("group", mcp.Description("Input parameter: The name of the group to add. This parameter cannot be used with the `groupId` parameter. As a group's name can change, use of `groupId` is recommended.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AddactorusersHandler(cfg),
	}
}
