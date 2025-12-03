package mcp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/configs"
)

func initMCP() *goMcp.Server {
	server := goMcp.NewServer(&goMcp.Implementation{
		Name:    configs.MCPName,
		Version: configs.MCPVersion,
	}, nil)

	// 添加登录工具
	goMcp.AddTool(server, &goMcp.Tool{
		Name:        "login",
		Description: "登录掘金",
	}, loginTool)

	return server
}

func MCP() fiber.Handler {
	mcp := initMCP()
	mcpHandler := goMcp.NewStreamableHTTPHandler(func(r *http.Request) *goMcp.Server {
		return mcp
	}, configs.MCPStreamableHTTPOptions)
	return adaptor.HTTPHandler(mcpHandler)
}
