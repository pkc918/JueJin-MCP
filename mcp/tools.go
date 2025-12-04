package mcp

import (
	"context"

	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/browser"
	"github.com/unomcp/JueJin-MCP/juejin"
)

// 发布掘金文章
type PublishContentArgs struct {
	Title   string `json:"title" jsonschema:"内容标题（掘金限制：最多20个中文字或英文单词）"`
	Content string `json:"content" jsonschema:"正文内容 限制为（Markdown 格式）"`
}

// 登录掘金
func loginTool(ctx context.Context, _req *goMcp.CallToolRequest, _ any) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	b := browser.New()
	defer b.Close()

	p := b.MustPage()
	defer p.Close()

	if err := juejin.Login(p, ctx); err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func publishTool(ctx context.Context, _req *goMcp.CallToolRequest, args PublishContentArgs) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	b := browser.New()
	defer b.Close()

	p := b.MustPage()
	defer p.Close()

	if err := juejin.Publish(p, ctx, juejin.PublishContent{
		Title:   args.Title,
		Content: args.Content,
	}); err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
