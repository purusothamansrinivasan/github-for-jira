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

func GetworkflowtransitionruleconfigurationsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["startAt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startAt=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["types"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("types=%v", val))
		}
		if val, ok := args["keys"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("keys=%v", val))
		}
		if val, ok := args["workflowNames"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("workflowNames=%v", val))
		}
		if val, ok := args["withTags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("withTags=%v", val))
		}
		if val, ok := args["draft"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("draft=%v", val))
		}
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/workflow/rule/config%s", cfg.BaseURL, queryString)
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
		var result models.PageBeanWorkflowTransitionRules
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

func CreateGetworkflowtransitionruleconfigurationsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_workflow_rule_config",
		mcp.WithDescription("Get workflow transition rule configurations"),
		mcp.WithNumber("startAt", mcp.Description("The index of the first item to return in a page of results (page offset).")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of items to return per page.")),
		mcp.WithArray("types", mcp.Required(), mcp.Description("The types of the transition rules to return.")),
		mcp.WithArray("keys", mcp.Description("The transition rule class keys, as defined in the Connect or the Forge app descriptor, of the transition rules to return.")),
		mcp.WithArray("workflowNames", mcp.Description("EXPERIMENTAL: The list of workflow names to filter by.")),
		mcp.WithArray("withTags", mcp.Description("EXPERIMENTAL: The list of `tags` to filter by.")),
		mcp.WithBoolean("draft", mcp.Description("EXPERIMENTAL: Whether draft or published workflows are returned. If not provided, both workflow types are returned.")),
		mcp.WithString("expand", mcp.Description("Use [expand](#expansion) to include additional information in the response. This parameter accepts `transition`, which, for each rule, returns information about the transition the rule is assigned to.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetworkflowtransitionruleconfigurationsHandler(cfg),
	}
}
