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

func DeleteworkflowschemeissuetypeHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		issueTypeVal, ok := args["issueType"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: issueType"), nil
		}
		issueType, ok := issueTypeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: issueType"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["updateDraftIfNeeded"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updateDraftIfNeeded=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/workflowscheme/%s/issuetype/%s%s", cfg.BaseURL, id, issueType, queryString)
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
		var result models.WorkflowScheme
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

func CreateDeleteworkflowschemeissuetypeTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_rest_api_3_workflowscheme_id_issuetype_issueType",
		mcp.WithDescription("Delete workflow for issue type in workflow scheme"),
		mcp.WithNumber("id", mcp.Required(), mcp.Description("The ID of the workflow scheme.")),
		mcp.WithString("issueType", mcp.Required(), mcp.Description("The ID of the issue type.")),
		mcp.WithBoolean("updateDraftIfNeeded", mcp.Description("Set to true to create or update the draft of a workflow scheme and update the mapping in the draft, when the workflow scheme cannot be edited. Defaults to `false`.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteworkflowschemeissuetypeHandler(cfg),
	}
}
