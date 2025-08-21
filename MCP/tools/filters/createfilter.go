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

func CreatefilterHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		if val, ok := args["overrideSharePermissions"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("overrideSharePermissions=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Filter
		
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
		url := fmt.Sprintf("%s/rest/api/3/filter%s", cfg.BaseURL, queryString)
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
		var result models.Filter
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

func CreateCreatefilterTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_rest_api_3_filter",
		mcp.WithDescription("Create filter"),
		mcp.WithString("expand", mcp.Description("Use [expand](#expansion) to include additional information about filter in the response. This parameter accepts a comma-separated list. Expand options include:\n\n *  `sharedUsers` Returns the users that the filter is shared with. This includes users that can browse projects that the filter is shared with. If you don't specify `sharedUsers`, then the `sharedUsers` object is returned but it doesn't list any users. The list of users returned is limited to 1000, to access additional users append `[start-index:end-index]` to the expand request. For example, to access the next 1000 users, use `?expand=sharedUsers[1001:2000]`.\n *  `subscriptions` Returns the users that are subscribed to the filter. If you don't specify `subscriptions`, the `subscriptions` object is returned but it doesn't list any subscriptions. The list of subscriptions returned is limited to 1000, to access additional subscriptions append `[start-index:end-index]` to the expand request. For example, to access the next 1000 subscriptions, use `?expand=subscriptions[1001:2000]`.")),
		mcp.WithBoolean("overrideSharePermissions", mcp.Description("EXPERIMENTAL: Whether share permissions are overridden to enable filters with any share permissions to be created. Available to users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).")),
		mcp.WithString("viewUrl", mcp.Description("Input parameter: A URL to view the filter results in Jira, using the ID of the filter. For example, *https://your-domain.atlassian.net/issues/?filter=10100*.")),
		mcp.WithArray("editPermissions", mcp.Description("Input parameter: The groups and projects that can edit the filter.")),
		mcp.WithString("id", mcp.Description("Input parameter: The unique identifier for the filter.")),
		mcp.WithString("jql", mcp.Description("Input parameter: The JQL query for the filter. For example, *project = SSP AND issuetype = Bug*.")),
		mcp.WithString("owner", mcp.Description("Input parameter: The user who owns the filter. This is defaulted to the creator of the filter, however Jira administrators can change the owner of a shared filter in the admin settings.")),
		mcp.WithString("sharedUsers", mcp.Description("Input parameter: A paginated list of the users that the filter is shared with. This includes users that are members of the groups or can browse the projects that the filter is shared with.")),
		mcp.WithNumber("favouritedCount", mcp.Description("Input parameter: The count of how many users have selected this filter as a favorite, including the filter owner.")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: The name of the filter. Must be unique.")),
		mcp.WithString("subscriptions", mcp.Description("Input parameter: A paginated list of the users that are subscribed to the filter.")),
		mcp.WithString("description", mcp.Description("Input parameter: A description of the filter.")),
		mcp.WithBoolean("favourite", mcp.Description("Input parameter: Whether the filter is selected as a favorite.")),
		mcp.WithString("self", mcp.Description("Input parameter: The URL of the filter.")),
		mcp.WithString("searchUrl", mcp.Description("Input parameter: A URL to view the filter results in Jira, using the [Search for issues using JQL](#api-rest-api-3-filter-search-get) operation with the filter's JQL string to return the filter results. For example, *https://your-domain.atlassian.net/rest/api/3/search?jql=project+%3D+SSP+AND+issuetype+%3D+Bug*.")),
		mcp.WithArray("sharePermissions", mcp.Description("Input parameter: The groups and projects that the filter is shared with.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatefilterHandler(cfg),
	}
}
