package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/the-jira-cloud-platform-rest-api/mcp-server/config"
	"github.com/the-jira-cloud-platform-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func StoreavatarHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["x"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("x=%v", val))
		}
		if val, ok := args["y"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("y=%v", val))
		}
		if val, ok := args["size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("size=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody interface{}
		
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
		url := fmt.Sprintf("%s/rest/api/3/universal_avatar/type/%s/owner/%s%s", cfg.BaseURL, type, entityId, queryString)
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
		var result models.Avatar
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

func CreateStoreavatarTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_rest_api_3_universal_avatar_type_type_owner_entityId",
		mcp.WithDescription("Load avatar"),
		mcp.WithString("type", mcp.Required(), mcp.Description("The avatar type.")),
		mcp.WithString("entityId", mcp.Required(), mcp.Description("The ID of the item the avatar is associated with.")),
		mcp.WithNumber("x", mcp.Description("The X coordinate of the top-left corner of the crop region.")),
		mcp.WithNumber("y", mcp.Description("The Y coordinate of the top-left corner of the crop region.")),
		mcp.WithNumber("size", mcp.Required(), mcp.Description("The length of each side of the crop region.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    StoreavatarHandler(cfg),
	}
}
