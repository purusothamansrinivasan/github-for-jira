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

func AddworklogHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["notifyUsers"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("notifyUsers=%v", val))
		}
		if val, ok := args["adjustEstimate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("adjustEstimate=%v", val))
		}
		if val, ok := args["newEstimate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("newEstimate=%v", val))
		}
		if val, ok := args["reduceBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("reduceBy=%v", val))
		}
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		if val, ok := args["overrideEditableFlag"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("overrideEditableFlag=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Worklog
		
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
		url := fmt.Sprintf("%s/rest/api/3/issue/%s/worklog%s", cfg.BaseURL, issueIdOrKey, queryString)
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
		var result models.Worklog
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

func CreateAddworklogTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_rest_api_3_issue_issueIdOrKey_worklog",
		mcp.WithDescription("Add worklog"),
		mcp.WithString("issueIdOrKey", mcp.Required(), mcp.Description("The ID or key the issue.")),
		mcp.WithBoolean("notifyUsers", mcp.Description("Whether users watching the issue are notified by email.")),
		mcp.WithString("adjustEstimate", mcp.Description("Defines how to update the issue's time estimate, the options are:\n\n *  `new` Sets the estimate to a specific value, defined in `newEstimate`.\n *  `leave` Leaves the estimate unchanged.\n *  `manual` Reduces the estimate by amount specified in `reduceBy`.\n *  `auto` Reduces the estimate by the value of `timeSpent` in the worklog.")),
		mcp.WithString("newEstimate", mcp.Description("The value to set as the issue's remaining time estimate, as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). For example, *2d*. Required when `adjustEstimate` is `new`.")),
		mcp.WithString("reduceBy", mcp.Description("The amount to reduce the issue's remaining estimate by, as days (\\#d), hours (\\#h), or minutes (\\#m). For example, *2d*. Required when `adjustEstimate` is `manual`.")),
		mcp.WithString("expand", mcp.Description("Use [expand](#expansion) to include additional information about work logs in the response. This parameter accepts `properties`, which returns worklog properties.")),
		mcp.WithBoolean("overrideEditableFlag", mcp.Description("Whether the worklog entry should be added to the issue even if the issue is not editable, because jira.issue.editable set to false or missing. For example, the issue is closed. Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) can use this flag.")),
		mcp.WithString("id", mcp.Description("Input parameter: The ID of the worklog record.")),
		mcp.WithArray("properties", mcp.Description("Input parameter: Details of properties for the worklog. Optional when creating or updating a worklog.")),
		mcp.WithString("timeSpent", mcp.Description("Input parameter: The time spent working on the issue as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). Required when creating a worklog if `timeSpentSeconds` isn't provided. Optional when updating a worklog. Cannot be provided if `timeSpentSecond` is provided.")),
		mcp.WithString("visibility", mcp.Description("Input parameter: Details about any restrictions in the visibility of the worklog. Optional when creating or updating a worklog.")),
		mcp.WithString("comment", mcp.Description("Input parameter: A comment about the worklog in [Atlassian Document Format](https://developer.atlassian.com/cloud/jira/platform/apis/document/structure/). Optional when creating or updating a worklog.")),
		mcp.WithString("created", mcp.Description("Input parameter: The datetime on which the worklog was created.")),
		mcp.WithString("self", mcp.Description("Input parameter: The URL of the worklog item.")),
		mcp.WithString("started", mcp.Description("Input parameter: The datetime on which the worklog effort was started. Required when creating a worklog. Optional when updating a worklog.")),
		mcp.WithString("updateAuthor", mcp.Description("Input parameter: Details of the user who last updated the worklog.")),
		mcp.WithString("issueId", mcp.Description("Input parameter: The ID of the issue this worklog is for.")),
		mcp.WithNumber("timeSpentSeconds", mcp.Description("Input parameter: The time in seconds spent working on the issue. Required when creating a worklog if `timeSpent` isn't provided. Optional when updating a worklog. Cannot be provided if `timeSpent` is provided.")),
		mcp.WithString("updated", mcp.Description("Input parameter: The datetime on which the worklog was last updated.")),
		mcp.WithString("author", mcp.Description("Input parameter: Details of the user who created the worklog.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AddworklogHandler(cfg),
	}
}
