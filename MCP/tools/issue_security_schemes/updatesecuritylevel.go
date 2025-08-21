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

func UpdatesecuritylevelHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		schemeIdVal, ok := args["schemeId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: schemeId"), nil
		}
		schemeId, ok := schemeIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: schemeId"), nil
		}
		levelIdVal, ok := args["levelId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: levelId"), nil
		}
		levelId, ok := levelIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: levelId"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.UpdateIssueSecurityLevelDetails
		
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
		url := fmt.Sprintf("%s/rest/api/3/issuesecurityschemes/%s/level/%s", cfg.BaseURL, schemeId, levelId)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
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

func CreateUpdatesecuritylevelTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_rest_api_3_issuesecurityschemes_schemeId_level_levelId",
		mcp.WithDescription("Update issue security level"),
		mcp.WithString("schemeId", mcp.Required(), mcp.Description("The ID of the issue security scheme level belongs to.")),
		mcp.WithString("levelId", mcp.Required(), mcp.Description("The ID of the issue security level to update.")),
		mcp.WithString("description", mcp.Description("Input parameter: The description of the issue security scheme level.")),
		mcp.WithString("name", mcp.Description("Input parameter: The name of the issue security scheme level. Must be unique.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatesecuritylevelHandler(cfg),
	}
}
