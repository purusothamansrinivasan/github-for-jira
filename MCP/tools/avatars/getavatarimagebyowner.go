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

func GetavatarimagebyownerHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		entityIdVal, ok := args["entityId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: entityId"), nil
		}
		entityId, ok := entityIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: entityId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("size=%v", val))
		}
		if val, ok := args["format"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("format=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/universal_avatar/view/type/%s/owner/%s%s", cfg.BaseURL, type, entityId, queryString)
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

func CreateGetavatarimagebyownerTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_universal_avatar_view_type_type_owner_entityId",
		mcp.WithDescription("Get avatar image by owner"),
		mcp.WithString("type", mcp.Required(), mcp.Description("The icon type of the avatar.")),
		mcp.WithString("entityId", mcp.Required(), mcp.Description("The ID of the project or issue type the avatar belongs to.")),
		mcp.WithString("size", mcp.Description("The size of the avatar image. If not provided the default size is returned.")),
		mcp.WithString("format", mcp.Description("The format to return the avatar image in. If not provided the original content format is returned.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetavatarimagebyownerHandler(cfg),
	}
}
