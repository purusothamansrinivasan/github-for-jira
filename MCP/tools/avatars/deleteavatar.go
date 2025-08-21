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

func DeleteavatarHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		typeVal, ok := args["type"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: type"), nil
		}
		type, ok := typeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: type"), nil
		}
		owningObjectIdVal, ok := args["owningObjectId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: owningObjectId"), nil
		}
		owningObjectId, ok := owningObjectIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: owningObjectId"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		url := fmt.Sprintf("%s/rest/api/3/universal_avatar/type/%s/owner/%s/avatar/%s", cfg.BaseURL, type, owningObjectId, id)
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

func CreateDeleteavatarTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_rest_api_3_universal_avatar_type_type_owner_owningObjectId_avatar_id",
		mcp.WithDescription("Delete avatar"),
		mcp.WithString("type", mcp.Required(), mcp.Description("The avatar type.")),
		mcp.WithString("owningObjectId", mcp.Required(), mcp.Description("The ID of the item the avatar is associated with.")),
		mcp.WithNumber("id", mcp.Required(), mcp.Description("The ID of the avatar.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteavatarHandler(cfg),
	}
}
