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

func GetworkflowspaginatedHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["workflowName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("workflowName=%v", val))
		}
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		if val, ok := args["queryString"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("queryString=%v", val))
		}
		if val, ok := args["orderBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderBy=%v", val))
		}
		if val, ok := args["isActive"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("isActive=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/workflow/search%s", cfg.BaseURL, queryString)
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
		var result models.PageBeanWorkflow
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

func CreateGetworkflowspaginatedTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_workflow_search",
		mcp.WithDescription("Get workflows paginated"),
		mcp.WithNumber("startAt", mcp.Description("The index of the first item to return in a page of results (page offset).")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of items to return per page.")),
		mcp.WithArray("workflowName", mcp.Description("The name of a workflow to return. To include multiple workflows, provide an ampersand-separated list. For example, `workflowName=name1&workflowName=name2`.")),
		mcp.WithString("expand", mcp.Description("Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:\n\n *  `transitions` For each workflow, returns information about the transitions inside the workflow.\n *  `transitions.rules` For each workflow transition, returns information about its rules. Transitions are included automatically if this expand is requested.\n *  `transitions.properties` For each workflow transition, returns information about its properties. Transitions are included automatically if this expand is requested.\n *  `statuses` For each workflow, returns information about the statuses inside the workflow.\n *  `statuses.properties` For each workflow status, returns information about its properties. Statuses are included automatically if this expand is requested.\n *  `default` For each workflow, returns information about whether this is the default workflow.\n *  `schemes` For each workflow, returns information about the workflow schemes the workflow is assigned to.\n *  `projects` For each workflow, returns information about the projects the workflow is assigned to, through workflow schemes.\n *  `hasDraftWorkflow` For each workflow, returns information about whether the workflow has a draft version.\n *  `operations` For each workflow, returns information about the actions that can be undertaken on the workflow.")),
		mcp.WithString("queryString", mcp.Description("String used to perform a case-insensitive partial match with workflow name.")),
		mcp.WithString("orderBy", mcp.Description("[Order](#ordering) the results by a field:\n\n *  `name` Sorts by workflow name.\n *  `created` Sorts by create time.\n *  `updated` Sorts by update time.")),
		mcp.WithBoolean("isActive", mcp.Description("Filters active and inactive workflows.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetworkflowspaginatedHandler(cfg),
	}
}
