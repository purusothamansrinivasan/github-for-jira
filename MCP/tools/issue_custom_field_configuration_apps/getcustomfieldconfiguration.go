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

func GetcustomfieldconfigurationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		fieldIdOrKeyVal, ok := args["fieldIdOrKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: fieldIdOrKey"), nil
		}
		fieldIdOrKey, ok := fieldIdOrKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: fieldIdOrKey"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["fieldContextId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fieldContextId=%v", val))
		}
		if val, ok := args["issueId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("issueId=%v", val))
		}
		if val, ok := args["projectKeyOrId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projectKeyOrId=%v", val))
		}
		if val, ok := args["issueTypeId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("issueTypeId=%v", val))
		}
		if val, ok := args["startAt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startAt=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/app/field/%s/context/configuration%s", cfg.BaseURL, fieldIdOrKey, queryString)
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
		var result models.PageBeanContextualConfiguration
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

func CreateGetcustomfieldconfigurationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_app_field_fieldIdOrKey_context_configuration",
		mcp.WithDescription("Get custom field configurations"),
		mcp.WithString("fieldIdOrKey", mcp.Required(), mcp.Description("The ID or key of the custom field, for example `customfield_10000`.")),
		mcp.WithArray("id", mcp.Description("The list of configuration IDs. To include multiple configurations, separate IDs with an ampersand: `id=10000&id=10001`. Can't be provided with `fieldContextId`, `issueId`, `projectKeyOrId`, or `issueTypeId`.")),
		mcp.WithArray("fieldContextId", mcp.Description("The list of field context IDs. To include multiple field contexts, separate IDs with an ampersand: `fieldContextId=10000&fieldContextId=10001`. Can't be provided with `id`, `issueId`, `projectKeyOrId`, or `issueTypeId`.")),
		mcp.WithNumber("issueId", mcp.Description("The ID of the issue to filter results by. If the issue doesn't exist, an empty list is returned. Can't be provided with `projectKeyOrId`, or `issueTypeId`.")),
		mcp.WithString("projectKeyOrId", mcp.Description("The ID or key of the project to filter results by. Must be provided with `issueTypeId`. Can't be provided with `issueId`.")),
		mcp.WithString("issueTypeId", mcp.Description("The ID of the issue type to filter results by. Must be provided with `projectKeyOrId`. Can't be provided with `issueId`.")),
		mcp.WithNumber("startAt", mcp.Description("The index of the first item to return in a page of results (page offset).")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of items to return per page.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetcustomfieldconfigurationHandler(cfg),
	}
}
