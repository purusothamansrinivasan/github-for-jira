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

func CreateworkflowtransitionpropertyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		transitionIdVal, ok := args["transitionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: transitionId"), nil
		}
		transitionId, ok := transitionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: transitionId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("key=%v", val))
		}
		if val, ok := args["workflowName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("workflowName=%v", val))
		}
		if val, ok := args["workflowMode"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("workflowMode=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.WorkflowTransitionProperty
		
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
		url := fmt.Sprintf("%s/rest/api/3/workflow/transitions/%s/properties%s", cfg.BaseURL, transitionId, queryString)
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
		var result models.WorkflowTransitionProperty
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

func CreateCreateworkflowtransitionpropertyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_rest_api_3_workflow_transitions_transitionId_properties",
		mcp.WithDescription("Create workflow transition property"),
		mcp.WithNumber("transitionId", mcp.Required(), mcp.Description("The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next to the transition.")),
		mcp.WithString("key", mcp.Required(), mcp.Description("The key of the property being added, also known as the name of the property. Set this to the same value as the `key` defined in the request body.")),
		mcp.WithString("workflowName", mcp.Required(), mcp.Description("The name of the workflow that the transition belongs to.")),
		mcp.WithString("workflowMode", mcp.Description("The workflow status. Set to *live* for inactive workflows or *draft* for draft workflows. Active workflows cannot be edited.")),
		mcp.WithString("id", mcp.Description("Input parameter: The ID of the transition property.")),
		mcp.WithString("key", mcp.Description("Input parameter: The key of the transition property. Also known as the name of the transition property.")),
		mcp.WithString("value", mcp.Required(), mcp.Description("Input parameter: The value of the transition property.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateworkflowtransitionpropertyHandler(cfg),
	}
}
