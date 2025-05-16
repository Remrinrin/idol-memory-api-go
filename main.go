package main

import (
	_ "github.com/Remrinrin/idol-memory-api-go/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/Remrinrin/idol-memory-api-go/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
