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

func UpdateversionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
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
		var requestBody models.Version
		
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
		url := fmt.Sprintf("%s/rest/api/3/version/%s", cfg.BaseURL, id)
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
		var result models.Version
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

func CreateUpdateversionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_rest_api_3_version_id",
		mcp.WithDescription("Update version"),
		mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the version.")),
		mcp.WithString("releaseDate", mcp.Description("Input parameter: The release date of the version. Expressed in ISO 8601 format (yyyy-mm-dd). Optional when creating or updating a version.")),
		mcp.WithBoolean("released", mcp.Description("Input parameter: Indicates that the version is released. If the version is released a request to release again is ignored. Not applicable when creating a version. Optional when updating a version.")),
		mcp.WithString("userStartDate", mcp.Description("Input parameter: The date on which work on this version is expected to start, expressed in the instance's *Day/Month/Year Format* date format.")),
		mcp.WithString("name", mcp.Description("Input parameter: The unique name of the version. Required when creating a version. Optional when updating a version. The maximum length is 255 characters.")),
		mcp.WithBoolean("overdue", mcp.Description("Input parameter: Indicates that the version is overdue.")),
		mcp.WithString("self", mcp.Description("Input parameter: The URL of the version.")),
		mcp.WithString("startDate", mcp.Description("Input parameter: The start date of the version. Expressed in ISO 8601 format (yyyy-mm-dd). Optional when creating or updating a version.")),
		mcp.WithString("id", mcp.Description("Input parameter: The ID of the version.")),
		mcp.WithString("moveUnfixedIssuesTo", mcp.Description("Input parameter: The URL of the self link to the version to which all unfixed issues are moved when a version is released. Not applicable when creating a version. Optional when updating a version.")),
		mcp.WithString("project", mcp.Description("Input parameter: Deprecated. Use `projectId`.")),
		mcp.WithString("userReleaseDate", mcp.Description("Input parameter: The date on which work on this version is expected to finish, expressed in the instance's *Day/Month/Year Format* date format.")),
		mcp.WithBoolean("archived", mcp.Description("Input parameter: Indicates that the version is archived. Optional when creating or updating a version.")),
		mcp.WithString("description", mcp.Description("Input parameter: The description of the version. Optional when creating or updating a version.")),
		mcp.WithNumber("projectId", mcp.Description("Input parameter: The ID of the project to which this version is attached. Required when creating a version. Not applicable when updating a version.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateversionHandler(cfg),
	}
}
