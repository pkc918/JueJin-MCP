package app

import (
	"github.com/gofiber/fiber/v2"
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/juejin"
	"github.com/unomcp/JueJin-MCP/mcp"
)

type App struct {
	JueJin   *juejin.JueJin
	MCP      *goMcp.Server
	FiberApp *fiber.App
}

func NewApp(app *fiber.App) *App {
	// 先初始化 JueJin
	jueJinInstance := juejin.NewJueJin()

	// 将 JueJin 实例注入到 MCP 中
	mcpServer := mcp.InitMCP(jueJinInstance)

	return &App{
		JueJin:   jueJinInstance,
		MCP:      mcpServer,
		FiberApp: app,
	}
}
