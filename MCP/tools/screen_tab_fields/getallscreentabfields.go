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

func GetallscreentabfieldsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		screenIdVal, ok := args["screenId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: screenId"), nil
		}
		screenId, ok := screenIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: screenId"), nil
		}
		tabIdVal, ok := args["tabId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: tabId"), nil
		}
		tabId, ok := tabIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: tabId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["projectKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projectKey=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/screens/%s/tabs/%s/fields%s", cfg.BaseURL, screenId, tabId, queryString)
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
		var result []ScreenableField
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

func CreateGetallscreentabfieldsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_screens_screenId_tabs_tabId_fields",
		mcp.WithDescription("Get all screen tab fields"),
		mcp.WithNumber("screenId", mcp.Required(), mcp.Description("The ID of the screen.")),
		mcp.WithNumber("tabId", mcp.Required(), mcp.Description("The ID of the screen tab.")),
		mcp.WithString("projectKey", mcp.Description("The key of the project.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetallscreentabfieldsHandler(cfg),
	}
}
