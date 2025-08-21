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

func SetworkflowschemedraftissuetypeHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.IssueTypeWorkflowMapping
		
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
		url := fmt.Sprintf("%s/rest/api/3/workflowscheme/%s/draft/issuetype/%s", cfg.BaseURL, id, issueType)
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

func CreateSetworkflowschemedraftissuetypeTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_rest_api_3_workflowscheme_id_draft_issuetype_issueType",
		mcp.WithDescription("Set workflow for issue type in draft workflow scheme"),
		mcp.WithNumber("id", mcp.Required(), mcp.Description("The ID of the workflow scheme that the draft belongs to.")),
		mcp.WithString("issueType", mcp.Required(), mcp.Description("The ID of the issue type.")),
		mcp.WithBoolean("updateDraftIfNeeded", mcp.Description("Input parameter: Set to true to create or update the draft of a workflow scheme and update the mapping in the draft, when the workflow scheme cannot be edited. Defaults to `false`. Only applicable when updating the workflow-issue types mapping.")),
		mcp.WithString("workflow", mcp.Description("Input parameter: The name of the workflow.")),
		mcp.WithString("issueType", mcp.Description("Input parameter: The ID of the issue type. Not required if updating the issue type-workflow mapping.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SetworkflowschemedraftissuetypeHandler(cfg),
	}
}
