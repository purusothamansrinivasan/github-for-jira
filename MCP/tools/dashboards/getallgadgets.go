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

func GetallgadgetsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		dashboardIdVal, ok := args["dashboardId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: dashboardId"), nil
		}
		dashboardId, ok := dashboardIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: dashboardId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["moduleKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("moduleKey=%v", val))
		}
		if val, ok := args["uri"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("uri=%v", val))
		}
		if val, ok := args["gadgetId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("gadgetId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/dashboard/%s/gadget%s", cfg.BaseURL, dashboardId, queryString)
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
		var result models.DashboardGadgetResponse
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

func CreateGetallgadgetsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_dashboard_dashboardId_gadget",
		mcp.WithDescription("Get gadgets"),
		mcp.WithNumber("dashboardId", mcp.Required(), mcp.Description("The ID of the dashboard.")),
		mcp.WithArray("moduleKey", mcp.Description("The list of gadgets module keys. To include multiple module keys, separate module keys with ampersand: `moduleKey=key:one&moduleKey=key:two`.")),
		mcp.WithArray("uri", mcp.Description("The list of gadgets URIs. To include multiple URIs, separate URIs with ampersand: `uri=/rest/example/uri/1&uri=/rest/example/uri/2`.")),
		mcp.WithArray("gadgetId", mcp.Description("The list of gadgets IDs. To include multiple IDs, separate IDs with ampersand: `gadgetId=10000&gadgetId=10001`.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetallgadgetsHandler(cfg),
	}
}
