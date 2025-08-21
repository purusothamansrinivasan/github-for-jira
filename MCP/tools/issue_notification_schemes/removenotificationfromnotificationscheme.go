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

func RemovenotificationfromnotificationschemeHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		notificationSchemeIdVal, ok := args["notificationSchemeId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: notificationSchemeId"), nil
		}
		notificationSchemeId, ok := notificationSchemeIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: notificationSchemeId"), nil
		}
		notificationIdVal, ok := args["notificationId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: notificationId"), nil
		}
		notificationId, ok := notificationIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: notificationId"), nil
		}
		url := fmt.Sprintf("%s/rest/api/3/notificationscheme/%s/notification/%s", cfg.BaseURL, notificationSchemeId, notificationId)
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

func CreateRemovenotificationfromnotificationschemeTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_rest_api_3_notificationscheme_notificationSchemeId_notification_notificationId",
		mcp.WithDescription("Remove notification from notification scheme"),
		mcp.WithString("notificationSchemeId", mcp.Required(), mcp.Description("The ID of the notification scheme.")),
		mcp.WithString("notificationId", mcp.Required(), mcp.Description("The ID of the notification.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    RemovenotificationfromnotificationschemeHandler(cfg),
	}
}
