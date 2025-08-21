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

func GetworkflowtransitionpropertiesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["includeReservedKeys"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeReservedKeys=%v", val))
		}
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
		url := fmt.Sprintf("%s/rest/api/3/workflow/transitions/%s/properties%s", cfg.BaseURL, transitionId, queryString)
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

func CreateGetworkflowtransitionpropertiesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_workflow_transitions_transitionId_properties",
		mcp.WithDescription("Get workflow transition properties"),
		mcp.WithNumber("transitionId", mcp.Required(), mcp.Description("The ID of the transition. To get the ID, view the workflow in text mode in the Jira administration console. The ID is shown next to the transition.")),
		mcp.WithBoolean("includeReservedKeys", mcp.Description("Some properties with keys that have the *jira.* prefix are reserved, which means they are not editable. To include these properties in the results, set this parameter to *true*.")),
		mcp.WithString("key", mcp.Description("The key of the property being returned, also known as the name of the property. If this parameter is not specified, all properties on the transition are returned.")),
		mcp.WithString("workflowName", mcp.Required(), mcp.Description("The name of the workflow that the transition belongs to.")),
		mcp.WithString("workflowMode", mcp.Description("The workflow status. Set to *live* for active and inactive workflows, or *draft* for draft workflows.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetworkflowtransitionpropertiesHandler(cfg),
	}
}
