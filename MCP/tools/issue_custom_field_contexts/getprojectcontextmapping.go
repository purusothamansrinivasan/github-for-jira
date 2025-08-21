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

func GetprojectcontextmappingHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		fieldIdVal, ok := args["fieldId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: fieldId"), nil
		}
		fieldId, ok := fieldIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: fieldId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["contextId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("contextId=%v", val))
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
		url := fmt.Sprintf("%s/rest/api/3/field/%s/context/projectmapping%s", cfg.BaseURL, fieldId, queryString)
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
		var result models.PageBeanCustomFieldContextProjectMapping
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

func CreateGetprojectcontextmappingTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_field_fieldId_context_projectmapping",
		mcp.WithDescription("Get project mappings for custom field context"),
		mcp.WithString("fieldId", mcp.Required(), mcp.Description("The ID of the custom field, for example `customfield\\_10000`.")),
		mcp.WithArray("contextId", mcp.Description("The list of context IDs. To include multiple context, separate IDs with ampersand: `contextId=10000&contextId=10001`.")),
		mcp.WithNumber("startAt", mcp.Description("The index of the first item to return in a page of results (page offset).")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of items to return per page.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetprojectcontextmappingHandler(cfg),
	}
}
