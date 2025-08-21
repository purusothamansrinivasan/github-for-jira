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

func EditissueHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["overrideScreenSecurity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("overrideScreenSecurity=%v", val))
		}
		if val, ok := args["overrideEditableFlag"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("overrideEditableFlag=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.IssueUpdateDetails
		
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
		url := fmt.Sprintf("%s/rest/api/3/issue/%s%s", cfg.BaseURL, issueIdOrKey, queryString)
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
		var result map[string]interface{}
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

func CreateEditissueTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_rest_api_3_issue_issueIdOrKey",
		mcp.WithDescription("Edit issue"),
		mcp.WithString("issueIdOrKey", mcp.Required(), mcp.Description("The ID or key of the issue.")),
		mcp.WithBoolean("notifyUsers", mcp.Description("Whether a notification email about the issue update is sent to all watchers. To disable the notification, administer Jira or administer project permissions are required. If the user doesn't have the necessary permission the request is ignored.")),
		mcp.WithBoolean("overrideScreenSecurity", mcp.Description("Whether screen security is overridden to enable hidden fields to be edited. Available to Connect app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).")),
		mcp.WithBoolean("overrideEditableFlag", mcp.Description("Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).")),
		mcp.WithString("historyMetadata", mcp.Description("Input parameter: Additional issue history details.")),
		mcp.WithArray("properties", mcp.Description("Input parameter: Details of issue properties to be add or update.")),
		mcp.WithString("transition", mcp.Description("Input parameter: Details of a transition. Required when performing a transition, optional when creating or editing an issue.")),
		mcp.WithObject("update", mcp.Description("Input parameter: A Map containing the field field name and a list of operations to perform on the issue screen field. Note that fields included in here cannot be included in `fields`.")),
		mcp.WithObject("fields", mcp.Description("Input parameter: List of issue screen fields to update, specifying the sub-field to update and its value for each field. This field provides a straightforward option when setting a sub-field. When multiple sub-fields or other operations are required, use `update`. Fields included in here cannot be included in `update`.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    EditissueHandler(cfg),
	}
}
