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

func DeleteissuetypepropertyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		issueTypeIdVal, ok := args["issueTypeId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: issueTypeId"), nil
		}
		issueTypeId, ok := issueTypeIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: issueTypeId"), nil
		}
		propertyKeyVal, ok := args["propertyKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: propertyKey"), nil
		}
		propertyKey, ok := propertyKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: propertyKey"), nil
		}
		url := fmt.Sprintf("%s/rest/api/3/issuetype/%s/properties/%s", cfg.BaseURL, issueTypeId, propertyKey)
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

func CreateDeleteissuetypepropertyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_rest_api_3_issuetype_issueTypeId_properties_propertyKey",
		mcp.WithDescription("Delete issue type property"),
		mcp.WithString("issueTypeId", mcp.Required(), mcp.Description("The ID of the issue type.")),
		mcp.WithString("propertyKey", mcp.Required(), mcp.Description("The key of the property. Use [Get issue type property keys](#api-rest-api-3-issuetype-issueTypeId-properties-get) to get a list of all issue type property keys.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteissuetypepropertyHandler(cfg),
	}
}
