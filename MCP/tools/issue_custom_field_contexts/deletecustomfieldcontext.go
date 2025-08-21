package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/the-jira-cloud-platform-rest-api/mcp-server/config"
	"github.com/the-jira-cloud-platform-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DeletecustomfieldcontextHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		fieldIdVal, ok := args["fieldId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: fieldId"), nil
		}
		fieldId, ok := fieldIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: fieldId"), nil
		}
		contextIdVal, ok := args["contextId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: contextId"), nil
		}
		contextId, ok := contextIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: contextId"), nil
		}
		url := fmt.Sprintf("%s/rest/api/3/field/%s/context/%s", cfg.BaseURL, fieldId, contextId)
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

func CreateDeletecustomfieldcontextTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_rest_api_3_field_fieldId_context_contextId",
		mcp.WithDescription("Delete custom field context"),
		mcp.WithString("fieldId", mcp.Required(), mcp.Description("The ID of the custom field.")),
		mcp.WithNumber("contextId", mcp.Required(), mcp.Description("The ID of the context.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletecustomfieldcontextHandler(cfg),
	}
}
