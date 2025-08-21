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

func GettransitionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		issueIdOrKeyVal, ok := args["issueIdOrKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: issueIdOrKey"), nil
		}
		issueIdOrKey, ok := issueIdOrKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: issueIdOrKey"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		if val, ok := args["transitionId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("transitionId=%v", val))
		}
		if val, ok := args["skipRemoteOnlyCondition"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("skipRemoteOnlyCondition=%v", val))
		}
		if val, ok := args["includeUnavailableTransitions"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeUnavailableTransitions=%v", val))
		}
		if val, ok := args["sortByOpsBarAndStatus"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sortByOpsBarAndStatus=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/issue/%s/transitions%s", cfg.BaseURL, issueIdOrKey, queryString)
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
		var result models.Transitions
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

func CreateGettransitionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_issue_issueIdOrKey_transitions",
		mcp.WithDescription("Get transitions"),
		mcp.WithString("issueIdOrKey", mcp.Required(), mcp.Description("The ID or key of the issue.")),
		mcp.WithString("expand", mcp.Description("Use [expand](#expansion) to include additional information about transitions in the response. This parameter accepts `transitions.fields`, which returns information about the fields in the transition screen for each transition. Fields hidden from the screen are not returned. Use this information to populate the `fields` and `update` fields in [Transition issue](#api-rest-api-3-issue-issueIdOrKey-transitions-post).")),
		mcp.WithString("transitionId", mcp.Description("The ID of the transition.")),
		mcp.WithBoolean("skipRemoteOnlyCondition", mcp.Description("Whether transitions with the condition *Hide From User Condition* are included in the response.")),
		mcp.WithBoolean("includeUnavailableTransitions", mcp.Description("Whether details of transitions that fail a condition are included in the response")),
		mcp.WithBoolean("sortByOpsBarAndStatus", mcp.Description("Whether the transitions are sorted by ops-bar sequence value first then category order (Todo, In Progress, Done) or only by ops-bar sequence value.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GettransitionsHandler(cfg),
	}
}
