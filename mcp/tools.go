package mcp

import (
	"context"

	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/browser"
	"github.com/unomcp/JueJin-MCP/juejin"
)

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
