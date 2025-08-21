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

func UpdatecommentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["notifyUsers"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("notifyUsers=%v", val))
		}
		if val, ok := args["overrideEditableFlag"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("overrideEditableFlag=%v", val))
		}
		if val, ok := args["expand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expand=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Comment
		
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
		url := fmt.Sprintf("%s/rest/api/3/issue/%s/comment/%s%s", cfg.BaseURL, issueIdOrKey, id, queryString)
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
		var result models.Comment
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

func CreateUpdatecommentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_rest_api_3_issue_issueIdOrKey_comment_id",
		mcp.WithDescription("Update comment"),
		mcp.WithString("issueIdOrKey", mcp.Required(), mcp.Description("The ID or key of the issue.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the comment.")),
		mcp.WithBoolean("notifyUsers", mcp.Description("Whether users are notified when a comment is updated.")),
		mcp.WithBoolean("overrideEditableFlag", mcp.Description("Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect app users with the *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).")),
		mcp.WithString("expand", mcp.Description("Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts `renderedBody`, which returns the comment body rendered in HTML.")),
		mcp.WithString("id", mcp.Description("Input parameter: The ID of the comment.")),
		mcp.WithString("updateAuthor", mcp.Description("Input parameter: The ID of the user who updated the comment last.")),
		mcp.WithString("author", mcp.Description("Input parameter: The ID of the user who created the comment.")),
		mcp.WithBoolean("jsdAuthorCanSeeRequest", mcp.Description("Input parameter: Whether the comment was added from an email sent by a person who is not part of the issue. See [Allow external emails to be added as comments on issues](https://support.atlassian.com/jira-service-management-cloud/docs/allow-external-emails-to-be-added-as-comments-on-issues/)for information on setting up this feature.")),
		mcp.WithString("renderedBody", mcp.Description("Input parameter: The rendered version of the comment.")),
		mcp.WithString("visibility", mcp.Description("Input parameter: The group or role to which this comment is visible. Optional on create and update.")),
		mcp.WithString("body", mcp.Description("Input parameter: The comment text in [Atlassian Document Format](https://developer.atlassian.com/cloud/jira/platform/apis/document/structure/).")),
		mcp.WithBoolean("jsdPublic", mcp.Description("Input parameter: Whether the comment is visible in Jira Service Desk. Defaults to true when comments are created in the Jira Cloud Platform. This includes when the site doesn't use Jira Service Desk or the project isn't a Jira Service Desk project and, therefore, there is no Jira Service Desk for the issue to be visible on. To create a comment with its visibility in Jira Service Desk set to false, use the Jira Service Desk REST API [Create request comment](https://developer.atlassian.com/cloud/jira/service-desk/rest/#api-rest-servicedeskapi-request-issueIdOrKey-comment-post) operation.")),
		mcp.WithString("updated", mcp.Description("Input parameter: The date and time at which the comment was updated last.")),
		mcp.WithArray("properties", mcp.Description("Input parameter: A list of comment properties. Optional on create and update.")),
		mcp.WithString("self", mcp.Description("Input parameter: The URL of the comment.")),
		mcp.WithString("created", mcp.Description("Input parameter: The date and time at which the comment was created.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatecommentHandler(cfg),
	}
}
