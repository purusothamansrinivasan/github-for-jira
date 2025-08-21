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

func GeteditissuemetaHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/rest/api/3/issue/%s/editmeta%s", cfg.BaseURL, issueIdOrKey, queryString)
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
		var result models.IssueUpdateMetadata
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

func CreateGeteditissuemetaTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_issue_issueIdOrKey_editmeta",
		mcp.WithDescription("Get edit issue metadata"),
		mcp.WithString("issueIdOrKey", mcp.Required(), mcp.Description("The ID or key of the issue.")),
		mcp.WithBoolean("overrideScreenSecurity", mcp.Description("Whether hidden fields are returned. Available to Connect app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).")),
		mcp.WithBoolean("overrideEditableFlag", mcp.Description("Whether non-editable fields are returned. Available to Connect app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GeteditissuemetaHandler(cfg),
	}
}
