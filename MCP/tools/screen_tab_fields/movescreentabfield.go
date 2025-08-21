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

func MovescreentabfieldHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.MoveFieldBean
		
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
		url := fmt.Sprintf("%s/rest/api/3/screens/%s/tabs/%s/fields/%s/move", cfg.BaseURL, screenId, tabId, id)
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

func CreateMovescreentabfieldTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_rest_api_3_screens_screenId_tabs_tabId_fields_id_move",
		mcp.WithDescription("Move screen tab field"),
		mcp.WithNumber("screenId", mcp.Required(), mcp.Description("The ID of the screen.")),
		mcp.WithNumber("tabId", mcp.Required(), mcp.Description("The ID of the screen tab.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the field.")),
		mcp.WithString("after", mcp.Description("Input parameter: The ID of the screen tab field after which to place the moved screen tab field. Required if `position` isn't provided.")),
		mcp.WithString("position", mcp.Description("Input parameter: The named position to which the screen tab field should be moved. Required if `after` isn't provided.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    MovescreentabfieldHandler(cfg),
	}
}
