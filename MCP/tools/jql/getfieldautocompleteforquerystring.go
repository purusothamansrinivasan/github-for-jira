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

func GetfieldautocompleteforquerystringHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["fieldName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fieldName=%v", val))
		}
		if val, ok := args["fieldValue"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fieldValue=%v", val))
		}
		if val, ok := args["predicateName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("predicateName=%v", val))
		}
		if val, ok := args["predicateValue"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("predicateValue=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rest/api/3/jql/autocompletedata/suggestions%s", cfg.BaseURL, queryString)
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
		var result models.AutoCompleteSuggestions
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

func CreateGetfieldautocompleteforquerystringTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rest_api_3_jql_autocompletedata_suggestions",
		mcp.WithDescription("Get field auto complete suggestions"),
		mcp.WithString("fieldName", mcp.Description("The name of the field.")),
		mcp.WithString("fieldValue", mcp.Description("The partial field item name entered by the user.")),
		mcp.WithString("predicateName", mcp.Description("The name of the [ CHANGED operator predicate](https://confluence.atlassian.com/x/hQORLQ#Advancedsearching-operatorsreference-CHANGEDCHANGED) for which the suggestions are generated. The valid predicate operators are *by*, *from*, and *to*.")),
		mcp.WithString("predicateValue", mcp.Description("The partial predicate item name entered by the user.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetfieldautocompleteforquerystringHandler(cfg),
	}
}
