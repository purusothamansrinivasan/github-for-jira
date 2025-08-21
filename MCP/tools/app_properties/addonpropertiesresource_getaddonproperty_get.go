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

func Addonpropertiesresource_getaddonproperty_getHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		addonKeyVal, ok := args["addonKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: addonKey"), nil
		}
		addonKey, ok := addonKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: addonKey"), nil
		}
		propertyKeyVal, ok := args["propertyKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: propertyKey"), nil
		}
		propertyKey, ok := propertyKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: propertyKey"), nil
		}
		url := fmt.Sprintf("%s/rest/atlassian-connect/1/addons/%s/properties/%s", cfg.BaseURL, addonKey, propertyKey)
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
		var result models.EntityProperty
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

func CreateAddonpropertiesresource_getaddonproperty_getTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_atlassian-connect_1_addons_addonKey_properties_propertyKey",
		mcp.WithDescription("Get app property"),
		mcp.WithString("addonKey", mcp.Required(), mcp.Description("The key of the app, as defined in its descriptor.")),
		mcp.WithString("propertyKey", mcp.Required(), mcp.Description("The key of the property.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Addonpropertiesresource_getaddonproperty_getHandler(cfg),
	}
}
