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

func ReplaceissuefieldoptionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		fieldKeyVal, ok := args["fieldKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: fieldKey"), nil
		}
		fieldKey, ok := fieldKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: fieldKey"), nil
		}
		optionIdVal, ok := args["optionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: optionId"), nil
		}
		optionId, ok := optionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: optionId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["replaceWith"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("replaceWith=%v", val))
		}
		if val, ok := args["jql"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jql=%v", val))
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
		url := fmt.Sprintf("%s/rest/api/3/field/%s/option/%s/issue%s", cfg.BaseURL, fieldKey, optionId, queryString)
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

func CreateReplaceissuefieldoptionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_rest_api_3_field_fieldKey_option_optionId_issue",
		mcp.WithDescription("Replace issue field option"),
		mcp.WithNumber("replaceWith", mcp.Description("The ID of the option that will replace the currently selected option.")),
		mcp.WithString("jql", mcp.Description("A JQL query that specifies the issues to be updated. For example, *project=10000*.")),
		mcp.WithBoolean("overrideScreenSecurity", mcp.Description("Whether screen security is overridden to enable hidden fields to be edited. Available to Connect and Forge app users with admin permission.")),
		mcp.WithBoolean("overrideEditableFlag", mcp.Description("Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).")),
		mcp.WithString("fieldKey", mcp.Required(), mcp.Description("The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the `fieldKey` value, do one of the following:\n\n *  open the app's plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the `jiraIssueFields` module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.\n *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in `key`. For example, `\"key\": \"teams-add-on__team-issue-field\"`")),
		mcp.WithNumber("optionId", mcp.Required(), mcp.Description("The ID of the option to be deselected.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ReplaceissuefieldoptionHandler(cfg),
	}
}
